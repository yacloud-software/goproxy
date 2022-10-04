package proto

import (
	"archive/zip"
	"context"
	"flag"
	"fmt"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/goproxy"
	pr "golang.conradwood.net/apis/protorenderer"
	"golang.conradwood.net/go-easyops/errors"
	"io"
	"strconv"
	"strings"
)

var (
	debug                    = flag.Bool("debug_proto", false, "debug proto stuff")
	WELL_KNOWN_PATH_PREFIXES = []string{
		"golang.conradwood.net/apis",
		"golang.singingcat.net/apis",
		"golang.yacloud.eu/apis",
	}
)

type protoHandler struct {
	pack    *pr.Package
	version *pr.Version
	path    string
}

func HandlerByPath(ctx context.Context, path string) (*protoHandler, error) {
	pn := &pr.PackageName{PackageName: path}
	pack, err := pr.GetProtoRendererClient().GetPackageByName(ctx, pn)
	if err != nil {
		if *debug {
			fmt.Printf("No proto package for path \"%s\": %s\n", path, err)
		}
		if isWellKnownProtoPath(path) {
			// we know user means us (path matches a proto path), but it's wrong
			return &protoHandler{}, nil
		}
		return nil, nil
	}
	p := &protoHandler{pack: pack, path: path}
	p.version, err = pr.GetProtoRendererClient().GetVersion(ctx, &common.Void{})
	return p, nil
}
func (ph *protoHandler) ModuleInfo() *pb.ModuleInfo {
	res := &pb.ModuleInfo{ModuleType: pb.MODULETYPE_PROTO}
	if ph.pack != nil {
		res.Exists = true
	}
	return res
}
func (ph *protoHandler) ListVersions(ctx context.Context) ([]*pb.VersionInfo, error) {
	if ph.pack == nil {
		return nil, nil
	}
	vi := &pb.VersionInfo{Version: ph.version.Version}
	return []*pb.VersionInfo{vi}, nil
}
func isWellKnownProtoPath(path string) bool {
	for _, s := range WELL_KNOWN_PATH_PREFIXES {
		if strings.HasPrefix(path, s) {
			return true
		}
	}
	return false
}
func (ph *protoHandler) GetVersion(ctx context.Context, v string) (*pb.VersionInfo, error) {
	vid, err := parseIDFromString(v)
	if err != nil {
		return nil, err
	}
	vi := &pb.VersionInfo{Version: vid}
	return vi, nil
}

func parseIDFromString(v string) (uint64, error) {
	sx := strings.Split(v, ".")
	if len(sx) != 3 {
		return 0, fmt.Errorf("invalid string \"%s\" (%d)\n", v, len(sx))
	}
	res, err := strconv.ParseUint(sx[1], 10, 64)
	if err != nil {
		return 0, err
	}
	return res, nil
}
func (ph *protoHandler) GetZip(ctx context.Context, w io.Writer, v string) error {
	vid, err := parseIDFromString(v)
	if err != nil {
		return err
	}
	fmt.Printf("Getting proto zip for package \"%s\", version \"%d\"\n", ph.pack.Name, ph.version.Version)
	fmt.Printf("Requested version: \"%d\", current version: \"%d\"\n", ph.version.Version, vid)
	if ph.version.Version != vid {
		s := fmt.Sprintf("proto in version %d is not available (current version is %d)", vid, ph.version.Version)
		return errors.InvalidArgs(ctx, s, s)
	}

	pn := &pr.PackageName{PackageName: ph.pack.Name}
	srv, err := pr.GetProtoRendererClient().GetFilesGoByPackageName(ctx, pn)
	if err != nil {
		return err
	}
	zw := zip.NewWriter(w)
	lastfile := ""
	var curwriter io.Writer
	for {
		zf, err := srv.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if lastfile == "" && zf.Filename != "" {
			lastfile = zf.Filename
			curwriter, err = zw.Create(zf.Filename)
			if err != nil {
				return err
			}
		}
		if len(zf.Payload) > 0 {
			_, err = curwriter.Write(zf.Payload)
			if err != nil {
				return err
			}
		}
	}
	err = zw.Close()
	if err != nil {
		return err
	}

	return nil
}
func (ph *protoHandler) GetMod(ctx context.Context, version string) ([]byte, error) {
	res := "module " + ph.path
	return []byte(res), nil
}
