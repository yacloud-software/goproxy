package proto

import (
	"archive/zip"
	"context"
	"flag"
	"fmt"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/goproxy"
	pr "golang.conradwood.net/apis/protorenderer"
	//	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/goproxy/cacher"
	hh "golang.conradwood.net/goproxy/handlerhelpers"
	"io"
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
	modname string
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
	p.modname = path
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
	vid, err := hh.ParseIDFromString(v)
	if err != nil {
		return nil, err
	}
	vi := &pb.VersionInfo{Version: vid}
	return vi, nil
}

func (ph *protoHandler) GetZip(ctx context.Context, w io.Writer, version string) error {
	vid, err := hh.ParseIDFromString(version)
	if err != nil {
		return err
	}
	//	fmt.Printf("Getting proto zip for package \"%s\", version \"%d\"\n", ph.pack.Name, ph.version.Version)
	fmt.Printf("Requested version: \"%d\", current version: \"%d\"\n", vid, ph.version.Version)

	// don't enforce version. instead we rely on caches (otherwise we are unable to build the cache)
	/*
		if ph.version.Version != vid {
			s := fmt.Sprintf("proto in version %d is not available (current version is %d)", vid, ph.version.Version)
			return errors.InvalidArgs(ctx, s, s)
		}
	*/

	pn := &pr.PackageName{PackageName: ph.modname}
	srv, err := pr.GetProtoRendererClient().GetFilesGoByPackageName(ctx, pn)
	if err != nil {
		return err
	}
	zw := zip.NewWriter(w)

	/*
		// now add go.mod
		modfile, err := ph.GetMod(ctx, version)
		if err != nil {
			return err
		}
		modwriter, err := zw.Create(hh.Filename2ZipFilename(ph.modname, version, "go.mod"))
		if err != nil {
			return err
		}
		_, err = modwriter.Write(modfile)
		if err != nil {
			return err
		}
	*/
	// now add the files
	lastfile := ""
	var curwriter io.Writer
	for {
		zf, err := srv.Recv()
		if err != nil {
			fmt.Printf("Err:%s\n", err)
			if err == io.EOF {
				break
			}
			return err
		}
		//		fmt.Printf("Gofile received: %s (%d bytes)\n", zf.Filename, len(zf.Payload))
		if lastfile != zf.Filename && zf.Filename != "" {
			lastfile = zf.Filename
			curwriter, err = zw.Create(hh.Filename2ZipFilename(ph.modname, version, zf.Filename))
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
	buf := []byte(res)
	c, err := cacher.NewCacher(ctx, ph.path, version, "mod")
	if err != nil {
		return nil, err
	}
	err = c.PutBytes(ctx, buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
