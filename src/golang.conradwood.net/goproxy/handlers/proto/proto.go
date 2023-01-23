package proto

import (
	"archive/zip"
	"context"
	"flag"
	"fmt"
	"golang.conradwood.net/apis/common"
	gom "golang.conradwood.net/apis/gomodule"
	pb "golang.conradwood.net/apis/goproxy"
	pr "golang.conradwood.net/apis/protorenderer"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/goproxy/cacher"
	hh "golang.conradwood.net/goproxy/handlerhelpers"
	"golang.conradwood.net/goproxy/hosts"
	"io"
)

var (
	cache_mod             = flag.Bool("cache_go_mod_proto", false, "if true also cache the go.mod file from protos. Should not be necessary because it is autogenerated")
	use_gomodule_to_serve = flag.Bool("use_gomodule_to_serve_proto_zips", false, "if true, we serve all proto zip files from gomodule")
	debug                 = flag.Bool("debug_proto", false, "debug proto stuff")
)

type protoHandler struct {
	pack    *pr.Package
	version *pr.Version
	path    string
	modname string
}

func HandlerByPath(ctx context.Context, path string) (*protoHandler, error) {
	/*
		idx := strings.Index(path, "/")
		if idx == -1 {
			return nil, nil
		}
		if !strings.HasPrefix(path[idx:], "/apis/") {
			return nil, nil
		}
	*/
	b, err := hosts.IsOneOfUs(ctx, path)
	if err != nil {
		return nil, err
	}
	if !b {
		if *debug {
			fmt.Printf("[protos] not a host we serve: %s\n", path)
		}
		return nil, nil
	}
	pn := &pr.PackageName{PackageName: path}
	fr, err := pr.GetProtoRendererClient().FindPackageByName(ctx, pn)
	if err != nil {
		return nil, err
	}
	if !fr.Exists {
		if *debug {
			fmt.Printf("[protos] protorenderer said it is not a proto: %s\n", path)
		}
		return nil, nil
	}
	pack := fr.Package
	ph := &protoHandler{pack: pack, path: path}
	if err != nil {
		if *debug {
			ph.Printf("error getting proto package for path \"%s\": %s\n", path, err)
		}
		// we know user means us (path matches a proto path), but it's wrong
		return &protoHandler{}, nil

	}
	ph.version, err = pr.GetProtoRendererClient().GetVersion(ctx, &common.Void{})
	if err != nil {
		if *debug {
			ph.Printf("error getting proto version for path \"%s\": %s\n", path, err)
		}
		return nil, err
	}
	ph.modname = path
	if *debug {
		ph.Printf("path %s is a valid proto module\n", path)
	}
	return ph, nil
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
	a_err := ph.checkAccess(ctx)
	if a_err != nil {
		return nil, a_err
	}
	vi := &pb.VersionInfo{Version: ph.version.Version}
	return []*pb.VersionInfo{vi}, nil
}

func (ph *protoHandler) GetLatestVersion(ctx context.Context) (*pb.VersionInfo, error) {
	a_err := ph.checkAccess(ctx)
	if a_err != nil {
		return nil, a_err
	}
	vi := &pb.VersionInfo{Version: ph.version.Version}
	return vi, nil
}

func (ph *protoHandler) GetZip(ctx context.Context, c *cacher.Cache, w io.Writer, version string) error {
	a_err := ph.checkAccess(ctx)
	if a_err != nil {
		return a_err
	}

	use_go_mod := false
	if *use_gomodule_to_serve {
		use_go_mod = true
	}
	vid, err := hh.ParseIDFromString(version)
	if err != nil {
		return err
	}
	if vid < 2122 {
		use_go_mod = true
	}
	u := auth.GetUser(ctx)
	if use_go_mod {
		fmt.Printf("Calling gomodule as user %s\n", auth.UserIDString(u))
		pr := &gom.ProtoRequest{
			PackageName: ph.path,
			Version:     version,
		}
		srv, err := gom.GetGoModuleClient().GetProtoZip(ctx, pr)
		if err != nil {
			return err
		}
		for {
			r, err := srv.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				return err
			}
			_, err = w.Write(r.Data)
			if err != nil {
				return err
			}
		}
		return nil
	}

	//	ph.Printf("Getting proto zip for package \"%s\", version \"%d\"\n", ph.pack.Name, ph.version.Version)
	ph.Printf("Requested version: \"%d\", current version: \"%d\"\n", vid, ph.version.Version)

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
			ph.Printf("Err:%s\n", err)
			if err == io.EOF {
				break
			}
			return err
		}
		//		ph.Printf("Gofile received: %s (%d bytes)\n", zf.Filename, len(zf.Payload))
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
func (ph *protoHandler) GetMod(ctx context.Context, c *cacher.Cache, version string) ([]byte, error) {
	a_err := ph.checkAccess(ctx)
	if a_err != nil {
		return nil, a_err
	}
	res := fmt.Sprintf("module %s\n", ph.path)
	res = res + fmt.Sprintf("// version requested %s\ngo 1.18\n", version)
	buf := []byte(res)
	if *cache_mod {
		c, err := cacher.NewCacher(ctx, ph.path, version, "mod")
		if err != nil {
			return nil, err
		}
		err = c.PutBytes(ctx, buf)
		if err != nil {
			return nil, err
		}
	}
	return buf, nil
}
func (ph *protoHandler) CacheEnabled() bool {
	return true
}
func (ph *protoHandler) Printf(format string, args ...interface{}) {
	s := "[protohandler] "
	sn := fmt.Sprintf(format, args...)
	fmt.Print(s + sn)
}

func (ph *protoHandler) checkAccess(ctx context.Context) error {
	u := auth.GetUser(ctx)
	if u == nil {
		return errors.Unauthenticated(ctx, "need login for protos")
	}
	return nil
}
