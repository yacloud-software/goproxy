package artefact

import (
	"bytes"
	"context"
	"fmt"
	artefact "golang.conradwood.net/apis/artefact"
	pb "golang.conradwood.net/apis/goproxy"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/utils"
	"io"
	"strconv"
	"strings"
)

const (
	MODULEDIR = "dist/gomod/"
)

type afhandler struct {
	artefactid *artefact.ArtefactID
	path       string
	modpath    string
}

func path2artefactid(ctx context.Context, path string) (*artefact.ArtefactID, string, error) {
	if strings.Contains(path, "go-easyops") {
		return &artefact.ArtefactID{ID: 24, Domain: "conradwood.net", Name: "go-easysops"}, "golang.conradwood.net/go-easyops", nil
	}
	return nil, "", nil
}

func HandlerByPath(ctx context.Context, path string) (*afhandler, error) {
	fmt.Printf("artefact path: %s\n", path)
	afid, modpath, err := path2artefactid(ctx, path)
	if err != nil {
		return nil, err
	}
	if afid == nil {
		return nil, nil
	}
	res := &afhandler{artefactid: afid, modpath: modpath, path: path}
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
	bl, err := artefact.GetArtefactClient().GetArtefactBuilds(ctx, af.artefactid)
	if err != nil {
		fmt.Printf("unable to get build for artefact: %s\n", err)
		return nil, err
	}
	var res []*pb.VersionInfo
	for _, b := range bl.Builds {
		v := &pb.VersionInfo{Version: b}

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

	return res, nil

}

// return a versioninfo from a go string (e.g. "v0.120.0")
func (af *afhandler) GetVersion(ctx context.Context, version string) (*pb.VersionInfo, error) {
	return nil, nil
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
