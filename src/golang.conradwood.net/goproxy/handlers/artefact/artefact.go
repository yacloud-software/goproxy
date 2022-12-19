package artefact

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	artefact "golang.conradwood.net/apis/artefact"
	git "golang.conradwood.net/apis/gitserver"
	pb "golang.conradwood.net/apis/goproxy"
	"golang.conradwood.net/go-easyops/cache"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/goproxy/cacher"
	"golang.conradwood.net/goproxy/hosts"
	"io"
	"strconv"
	"strings"
	"time"
)

const (
	MODULEDIR = "dist/gomod/"
)

var (
	debug     = flag.Bool("debug_artefact", false, "if true debug artefact handler")
	afcache   = cache.New("artefact_version_cache", time.Duration(5)*time.Second, 1000)
	afidcache = cache.New("artefact_id_for_url", time.Duration(24)*time.Hour, 1000)
)

type afidcache_entry struct {
	err     error
	afid    *artefact.ArtefactID
	created time.Time
}
type afhandler struct {
	afresolved *afresolved
	path       string
	modpath    string
	cacheentry *afcacheentry
}
type afcacheentry struct {
	versioninfolist []*pb.VersionInfo
}
type afresolved struct {
	afid      *artefact.ArtefactID
	PathMatch bool // if asked for, say, golang.conradwood.net/go-easyops/utils this will be false (because path is too long/too specific)
}

func (af *afhandler) url2artefactid(ctx context.Context, url string) (*afresolved, error) {
	afide := afidcache.Get(url)
	if afide != nil {
		afid := afide.(*afidcache_entry)
		if afid.err != nil {
			// last time it errored
			if time.Since(afid.created) < time.Duration(15)*time.Second { // negative (error) caching time
				return nil, afid.err
			}
		} else {
			// last time it did not error
			return &afresolved{afid: afid.afid, PathMatch: true}, nil
		}
	}
	bur := &git.ByURLRequest{URL: url}
	rr, err := git.GetGIT2Client().FindRepoByURL(ctx, bur)
	if err != nil {
		af.Printf("failed to find git repo by url \"%s\": %s\n", url, err)
		return nil, err
	}
	if !rr.Found {
		af.Printf("no git repo by url \"%s\"\n", url)
		return nil, nil
	}
	r := rr.Repository
	// we got a repo
	nc := &afidcache_entry{created: time.Now()}
	repoid := r.ID
	afid, err := artefact.GetArtefactClient().GetArtefactIDForRepo(ctx, &artefact.ID{ID: repoid})
	if err != nil {
		nc.err = err
		afidcache.Put(url, nc)
		return nil, err
	}
	nc.afid = afid
	afidcache.Put(url, nc)
	return &afresolved{afid: afid, PathMatch: true}, nil

}
func (af *afhandler) path2artefactid(ctx context.Context, path string) (*afresolved, string, error) {
	af.Printf("resolving \"%s\" as artefact\n", path)
	// we got to do some messy conversions, e.g. from
	// "golang.singingcat.net/scgolib" to "git.singingcat.net/git/scgolib.git"
	gu := []string{path}
	ghost := "git." + strings.TrimPrefix(path, "golang.")
	gpath := ""
	idx := strings.Index(ghost, "/")
	if idx != -1 {
		gpath = strings.TrimPrefix(ghost[idx:], "/")
		ghost = strings.TrimPrefix(ghost[:idx], "/")
	}
	gu = append(gu, "https://"+ghost+"/git/"+gpath+".git")

	for _, g := range gu {
		af, err := af.url2artefactid(ctx, g)
		if err != nil {
			return nil, "", err
		}
		if af != nil {
			return af, path, nil
		}
	}

	if path == "golang.conradwood.net/go-easyops" {
		afid := &artefact.ArtefactID{ID: 24, Domain: "conradwood.net", Name: "go-easysops"}
		return &afresolved{afid: afid, PathMatch: true}, path, nil
	}
	if strings.Contains(path, "golang.conradwood.net/go-easyops") {
		afid := &artefact.ArtefactID{ID: 24, Domain: "conradwood.net", Name: "go-easysops"}
		return &afresolved{afid: afid, PathMatch: false}, path, nil
	}
	af.Printf("Not an artefact: \"%s\"\n", path)
	return nil, "", nil
}

// this is where the resolver comes in.
// return nil means we are not responsible
func HandlerByPath(ctx context.Context, path string) (*afhandler, error) {
	af := &afhandler{path: path}
	b, err := hosts.IsOneOfUs(ctx, path)
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, nil
	}
	afid, modpath, err := af.path2artefactid(ctx, path)
	if err != nil {
		return nil, err
	}

	if afid == nil {
		return &afhandler{}, nil
	}

	af.Printf("artefact path: %s\n", path)
	ac := afcache.Get(modpath)
	var ace *afcacheentry
	if ac != nil {
		ace = ac.(*afcacheentry)
	} else {
		ace = &afcacheentry{}
		afcache.Put(modpath, ace)
	}

	af.afresolved = afid
	af.modpath = modpath
	af.cacheentry = ace
	return af, nil
}
func (af *afhandler) ModuleInfo() *pb.ModuleInfo {
	res := &pb.ModuleInfo{ModuleType: pb.MODULETYPE_ARTEFACT}
	if af.afresolved != nil && af.afresolved.PathMatch {
		res.Exists = true
	}
	return res
}

func (af *afhandler) ListVersions(ctx context.Context) ([]*pb.VersionInfo, error) {
	if af.cacheentry.versioninfolist != nil {
		return af.cacheentry.versioninfolist, nil
	}
	if !af.afresolved.PathMatch {
		return nil, errors.NotFound(ctx, "path not found (incomplete pathmatch)")
	}
	bl, err := artefact.GetArtefactClient().GetArtefactBuilds(ctx, af.afresolved.afid)
	if err != nil {
		af.Printf("unable to get build for artefact: %s\n", err)
		return nil, err
	}
	var res []*pb.VersionInfo
	for _, b := range bl.Builds {
		dir := MODULEDIR + af.modpath
		hasmod, err := af.check_if_has_file(ctx, b, af.afresolved.afid.ID, dir)

		if err != nil {
			af.Printf("Build #%d does not have a module in \"%s\" (%s)\n", b, dir, utils.ErrorString(err))
			continue
		}
		if !hasmod {
			af.Printf("Build #%d does not have the required files for modules in \"%s\"", b, dir)
			continue
		}
		af.Printf("Build #%d added as module in \"%s\"\n", b, dir)
		v := &pb.VersionInfo{VersionName: fmt.Sprintf("v0.1.%d", b)}
		res = append(res, v)
	}
	af.cacheentry.versioninfolist = res
	return res, nil

}

// return a versioninfo from a go string (e.g. "v0.120.0")
func (af *afhandler) GetLatestVersion(ctx context.Context) (*pb.VersionInfo, error) {
	if !af.afresolved.PathMatch {
		return nil, errors.NotFound(ctx, "path not found (incomplete pathmatch)")
	}
	versions, err := af.ListVersions(ctx)
	if err != nil {
		return nil, err
	}
	if versions == nil || len(versions) == 0 {
		return nil, errors.NotFound(ctx, "no versioninfo found for %s", af.modpath)
	}
	return versions[0], nil
}

// get the zip file for a version
func (af *afhandler) GetZip(ctx context.Context, c *cacher.Cache, w io.Writer, version string) error {
	if !af.afresolved.PathMatch {
		return errors.NotFound(ctx, "path not found (incomplete pathmatch)")
	}
	buildid, err := parseIDFromString(version)
	if err != nil {
		return err
	}
	fr := &artefact.FileRequest{
		ArtefactID: af.afresolved.afid.ID,
		Build:      buildid,
		Filename:   MODULEDIR + af.modpath + "/mod.zip",
	}
	fi, err := artefact.GetArtefactClient().DoesFileExist(ctx, fr)
	if err != nil {
		af.Printf("Failed to check if file exists: %s\n", utils.ErrorString(err))
		return err
	}
	if !fi.Exists {
		return errors.NotFound(ctx, "mod.zip not found", "mod.zip not found")
	}
	stream, err := artefact.GetArtefactClient().GetFileStream(ctx, fr)
	if err != nil {
		af.Printf("failed to get mod.zip (%s)\n", utils.ErrorString(err))
		return err
	}
	for {
		pl, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			af.Printf("recv() error: %s\n", err)
			return err
		}
		_, err = w.Write(pl.Payload)
		if err != nil {
			af.Printf("write error: %s\n", err)
			return err
		}
	}
	return nil
}

// get the go.mod (url like .../@v/[version].mod
func (af *afhandler) GetMod(ctx context.Context, c *cacher.Cache, version string) ([]byte, error) {
	if !af.afresolved.PathMatch {
		return nil, errors.NotFound(ctx, "path not found (incomplete pathmatch)")
	}
	buildid, err := parseIDFromString(version)
	if err != nil {
		return nil, err
	}
	fr := &artefact.FileRequest{
		ArtefactID: af.afresolved.afid.ID,
		Build:      buildid,
		Filename:   MODULEDIR + af.modpath + "/go.mod",
	}
	fi, err := artefact.GetArtefactClient().DoesFileExist(ctx, fr)
	if err != nil {
		af.Printf("Failed to check if file exists: %s\n", utils.ErrorString(err))
		return nil, err
	}
	if !fi.Exists {
		af.Printf("Returning standard go.mod file\n")
		return []byte("module " + af.modpath), nil
	}
	stream, err := artefact.GetArtefactClient().GetFileStream(ctx, fr)
	if err != nil {
		af.Printf("failed to get go.mod (%s)\n", utils.ErrorString(err))
		return nil, err
	}
	var buf bytes.Buffer
	for {
		pl, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			af.Printf("recv() error: %s\n", err)
			return nil, err
		}
		_, err = buf.Write(pl.Payload)
		if err != nil {
			af.Printf("write error: %s\n", err)
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func parseIDFromString(v string) (uint64, error) {
	sx := strings.Split(v, ".")
	if len(sx) != 3 {
		return 0, fmt.Errorf("invalid string \"%s\" (%d)\n", v, len(sx))
	}
	// new ones are 1.0.x
	res, err := strconv.ParseUint(sx[2], 10, 64)
	if err != nil {
		return 0, err
	}
	if res != 0 {
		return res, nil
	}
	// there are some old protos with version 0.x.0
	res, err = strconv.ParseUint(sx[1], 10, 64)
	if err != nil {
		return 0, err
	}
	return res, nil
}
func (af *afhandler) CacheEnabled() bool {
	return true
}
func (af *afhandler) Printf(format string, args ...interface{}) {
	if !*debug {
		return
	}
	s := "[artefacthandler] "
	sn := fmt.Sprintf(format, args...)
	fmt.Print(s + sn)
}
