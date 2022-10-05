package artefact

import (
	"bytes"
	"context"
	"fmt"
	artefact "golang.conradwood.net/apis/artefact"
	git "golang.conradwood.net/apis/gitserver"
	pb "golang.conradwood.net/apis/goproxy"
	"golang.conradwood.net/go-easyops/cache"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/utils"
	"io"
	"strconv"
	"strings"
	"time"
)

const (
	MODULEDIR = "dist/gomod/"
)

var (
	afcache = cache.New("artefact_version_cache", time.Duration(5)*time.Second, 1000)
)

type afhandler struct {
	artefactid *artefact.ArtefactID
	path       string
	modpath    string
	cacheentry *afcacheentry
}
type afcacheentry struct {
	versioninfolist []*pb.VersionInfo
}

func url2artefactid(ctx context.Context, url string) (*artefact.ArtefactID, error) {
	bur := &git.ByURLRequest{URL: url}
	r, err := git.GetGIT2Client().RepoByURL(ctx, bur)
	if err != nil {
		fmt.Printf("no git repo by url \"%s\": %s\n", url, err)
		return nil, nil
	}
	// we got a repo
	repoid := r.ID
	afid, err := artefact.GetArtefactClient().GetArtefactIDForRepo(ctx, &artefact.ID{ID: repoid})
	if err != nil {
		return nil, err
	}
	return afid, nil

}
func path2artefactid(ctx context.Context, path string) (*artefact.ArtefactID, string, error) {
	fmt.Printf("resolving \"%s\" as artefact\n", path)
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
		af, err := url2artefactid(ctx, g)
		if err != nil {
			return nil, "", err
		}
		if af != nil {
			return af, path, nil
		}
	}

	if path == "golang.conradwood.net/go-easyops" {
		return &artefact.ArtefactID{ID: 24, Domain: "conradwood.net", Name: "go-easysops"}, path, nil
	}
	fmt.Printf("Not an artefact: \"%s\"\n", path)
	return nil, "", nil
}

func HandlerByPath(ctx context.Context, path string) (*afhandler, error) {
	afid, modpath, err := path2artefactid(ctx, path)
	if err != nil {
		return nil, err
	}
	if afid == nil {
		return nil, nil
	}

	fmt.Printf("artefact path: %s\n", path)
	ac := afcache.Get(modpath)
	var ace *afcacheentry
	if ac != nil {
		ace = ac.(*afcacheentry)
	} else {
		ace = &afcacheentry{}
		afcache.Put(modpath, ace)
	}

	res := &afhandler{artefactid: afid, modpath: modpath, path: path, cacheentry: ace}

	return res, nil
}
func (af *afhandler) ModuleInfo() *pb.ModuleInfo {
	res := &pb.ModuleInfo{ModuleType: pb.MODULETYPE_LOCALMODULE}
	if af.artefactid != nil {
		res.Exists = true
	}
	return res
}

func (af *afhandler) ListVersions(ctx context.Context) ([]*pb.VersionInfo, error) {
	if af.cacheentry.versioninfolist != nil {
		return af.cacheentry.versioninfolist, nil
	}
	bl, err := artefact.GetArtefactClient().GetArtefactBuilds(ctx, af.artefactid)
	if err != nil {
		fmt.Printf("unable to get build for artefact: %s\n", err)
		return nil, err
	}
	var res []*pb.VersionInfo
	for _, b := range bl.Builds {
		//TODO: fix gitbuilder to create zip files with v1.1.%d instead of 0.1.%d
		v := &pb.VersionInfo{VersionName: fmt.Sprintf("v0.1.%d", b)}

		dlr := &artefact.DirListRequest{
			Build:      b,
			ArtefactID: af.artefactid.ID,
			Dir:        MODULEDIR + af.modpath,
		}
		ds, err := artefact.GetArtefactClient().GetDirListing(ctx, dlr)

		if err != nil {
			fmt.Printf("Build #%d does not have a module in \"%s\" (%s)\n", b, dlr.Dir, utils.ErrorString(err))
			continue
		}
		if len(ds.Files) == 0 {
			fmt.Printf("Build #%d does not have the required files for modules in \"%s\"", b, dlr.Dir)
			continue
		}
		fmt.Printf("Build #%d added as module in \"%s\"\n", b, dlr.Dir)
		res = append(res, v)
	}
	af.cacheentry.versioninfolist = res
	return res, nil

}

// return a versioninfo from a go string (e.g. "v0.120.0")
func (af *afhandler) GetLatestVersion(ctx context.Context) (*pb.VersionInfo, error) {
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
func (af *afhandler) GetZip(ctx context.Context, w io.Writer, version string) error {
	buildid, err := parseIDFromString(version)
	if err != nil {
		return err
	}
	fr := &artefact.FileRequest{
		ArtefactID: af.artefactid.ID,
		Build:      buildid,
		Filename:   MODULEDIR + af.modpath + "/mod.zip",
	}
	fi, err := artefact.GetArtefactClient().DoesFileExist(ctx, fr)
	if err != nil {
		fmt.Printf("Failed to check if file exists: %s\n", utils.ErrorString(err))
		return err
	}
	if !fi.Exists {
		return errors.NotFound(ctx, "mod.zip not found", "mod.zip not found")
	}
	stream, err := artefact.GetArtefactClient().GetFileStream(ctx, fr)
	if err != nil {
		fmt.Printf("failed to get mod.zip (%s)\n", utils.ErrorString(err))
		return err
	}
	for {
		pl, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("recv() error: %s\n", err)
			return err
		}
		_, err = w.Write(pl.Payload)
		if err != nil {
			fmt.Printf("write error: %s\n", err)
			return err
		}
	}
	return nil
}

// get the go.mod (url like .../@v/[version].mod
func (af *afhandler) GetMod(ctx context.Context, version string) ([]byte, error) {
	buildid, err := parseIDFromString(version)
	if err != nil {
		return nil, err
	}
	fr := &artefact.FileRequest{
		ArtefactID: af.artefactid.ID,
		Build:      buildid,
		Filename:   MODULEDIR + af.modpath + "/go.mod",
	}
	fi, err := artefact.GetArtefactClient().DoesFileExist(ctx, fr)
	if err != nil {
		fmt.Printf("Failed to check if file exists: %s\n", utils.ErrorString(err))
		return nil, err
	}
	if !fi.Exists {
		fmt.Printf("Returning standard go.mod file\n")
		return []byte("module " + af.modpath), nil
	}
	stream, err := artefact.GetArtefactClient().GetFileStream(ctx, fr)
	if err != nil {
		fmt.Printf("failed to get go.mod (%s)\n", utils.ErrorString(err))
		return nil, err
	}
	var buf bytes.Buffer
	for {
		pl, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("recv() error: %s\n", err)
			return nil, err
		}
		_, err = buf.Write(pl.Payload)
		if err != nil {
			fmt.Printf("write error: %s\n", err)
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
