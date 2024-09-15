package proto

import (
	"archive/zip"
	"context"
	"flag"
	"fmt"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/goproxy"
	"golang.conradwood.net/apis/protorenderer"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/goproxy/cacher"
	hh "golang.conradwood.net/goproxy/handlerhelpers"
	"golang.yacloud.eu/apis/protomanager"
	//	"golang.conradwood.net/goproxy/hosts"
	"io"
)

var (
	use_protomanager = flag.Bool("use_protomanager", true, "use protomanager instead of protorenderer")
	cache_mod        = flag.Bool("cache_go_mod_proto", false, "if true also cache the go.mod file from protos. Should not be necessary because it is autogenerated")
	//	use_gomodule_to_serve = flag.Bool("use_gomodule_to_serve_proto_zips", false, "if true, we serve all proto zip files from gomodule")
	debug = flag.Bool("debug_proto", false, "debug proto stuff")
)

type protoHandler struct {
	//	pack       *protorenderer.Package
	package_id string // "" == not one of our packages
	version    uint64 //*protorenderer.Version
	path       string
	modname    string
}

func HandlerByPath(ctx context.Context, path string, isoneofus func(context.Context, string) (bool, error)) (*protoHandler, error) {
	/*
		idx := strings.Index(path, "/")
		if idx == -1 {
			return nil, nil
		}
		if !strings.HasPrefix(path[idx:], "/apis/") {
			return nil, nil
		}
	*/
	b, err := isoneofus(ctx, path)
	if err != nil {
		return nil, err
	}
	if !b {
		if *debug {
			fmt.Printf("[protos] not a host we serve: %s\n", path)
		}
		return nil, nil
	}
	package_id := ""
	if *use_protomanager {
		pn := &protomanager.PackagesByNameRequest{Name: path}
		pnl, err := protomanager.GetProtoManagerClient().FindPackagesByName(ctx, pn)
		if err != nil {
			return nil, err
		}
		if len(pnl.Packages) == 0 {
			if *debug {
				fmt.Printf("[protos] protomanager said it is not a proto: %s\n", path)
			}
			return nil, nil
		}
		if len(pnl.Packages) > 1 {
			if *debug {
				fmt.Printf("[protos] protomanager found %d packages for proto: %s\n", len(pnl.Packages), path)
				for _, p := range pnl.Packages {
					fmt.Printf("[protos] protomanager found: %s\n", p.Name)
				}
			}
			return nil, nil
		}
		package_id = pnl.Packages[0].ID
	} else {
		pn := &protorenderer.PackageName{PackageName: path}
		fr, err := protorenderer.GetProtoRendererClient().FindPackageByName(ctx, pn)
		if err != nil {
			return nil, err
		}
		if !fr.Exists {
			if *debug {
				fmt.Printf("[protos] protorenderer said it is not a proto: %s\n", path)
			}
			return nil, nil
		}
		package_id = fr.Package.ID
	}
	//	pack := fr.Package
	ph := &protoHandler{package_id: package_id, path: path}

	if *use_protomanager {
		pn := &protomanager.VersionRequest{Package: path}
		pid, err := protomanager.GetProtoManagerClient().GetCurrentVersion(ctx, pn)
		if err != nil {
			return nil, err
		}
		ph.version = pid.VersionNumber
	} else {
		pid, err := protorenderer.GetProtoRendererClient().GetVersion(ctx, &common.Void{})
		if err != nil {
			return nil, err
		}
		ph.version = pid.Version
	}

	ph.modname = path
	if *debug {
		ph.Printf("path %s is a valid proto module\n", path)
	}
	return ph, nil
}
func (ph *protoHandler) ModuleInfo() *pb.ModuleInfo {
	res := &pb.ModuleInfo{ModuleType: pb.MODULETYPE_PROTO}
	if ph.package_id != "" {
		res.Exists = true
	}
	return res
}
func (ph *protoHandler) ListVersions(ctx context.Context) ([]*pb.VersionInfo, error) {
	if ph.package_id == "" {
		return nil, nil
	}
	a_err := ph.checkAccess(ctx)
	if a_err != nil {
		return nil, a_err
	}
	vi := &pb.VersionInfo{Version: ph.version}
	return []*pb.VersionInfo{vi}, nil
}

func (ph *protoHandler) GetLatestVersion(ctx context.Context) (*pb.VersionInfo, error) {
	a_err := ph.checkAccess(ctx)
	if a_err != nil {
		return nil, a_err
	}
	vi := &pb.VersionInfo{Version: ph.version}
	return vi, nil
}

func (ph *protoHandler) GetZip(ctx context.Context, c *cacher.Cache, w io.Writer, version string) error {
	a_err := ph.checkAccess(ctx)
	if a_err != nil {
		return a_err
	}

	vid, err := hh.ParseIDFromString(version)
	if err != nil {
		return err
	}

	//	u := auth.GetUser(ctx)

	//	ph.Printf("Getting proto zip for package \"%s\", version \"%d\"\n", ph.pack.Name, ph.version.Version)
	ph.Printf("Requested version: \"%d\", current version: \"%d\"\n", vid, ph.version)

	// don't enforce version. instead we rely on caches (otherwise we are unable to build the cache)
	/*
		if ph.version.Version != vid {
			s := fmt.Sprintf("proto in version %d is not available (current version is %d)", vid, ph.version.Version)
			return errors.InvalidArgs(ctx, s, s)
		}
	*/
	zw := zip.NewWriter(w)
	receiver := &changeable_file_receiver{}
	if *use_protomanager {
		pn := &protomanager.FilesExtensionRequest{
			PackageID:  ph.package_id,
			Extensions: []string{".go"},
		}
		srv, err := protomanager.GetProtoManagerClient().GetFilesForPackageByExtension(ctx, pn)
		if err != nil {
			return err
		}
		receiver.pm = srv
	} else {
		pn := &protorenderer.PackageName{PackageName: ph.modname}
		srv, err := protorenderer.GetProtoRendererClient().GetFilesGoByPackageName(ctx, pn)
		if err != nil {
			return err
		}
		receiver.pr = srv
	}

	// now add the files
	lastfile := ""
	var curwriter io.Writer
	for {
		zf, err := receiver.Recv()
		if zf != nil {
			//		ph.Printf("Gofile received: %s (%d bytes)\n", zf.Filename, len(zf.Payload))
			if lastfile != zf.Filename && zf.Filename != "" {
				lastfile = zf.Filename
				curwriter, err = zw.Create(hh.Filename2ZipFilename(ph.modname, version, zf.Filename))
				if err != nil {
					return err
				}
			}
			if len(zf.Data) > 0 {
				_, err = curwriter.Write(zf.Data)
				if err != nil {
					return err
				}
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			ph.Printf("receive Err:%s\n", utils.ErrorString(err))
			return err
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

type changeable_file_receiver struct {
	pm protomanager.ProtoManager_GetFilesForPackageByExtensionClient
	pr protorenderer.ProtoRendererService_GetFilesGoByPackageNameClient
}
type changeable_file_receiver_data struct {
	Filename string
	Data     []byte
}

func (cfr *changeable_file_receiver) Recv() (*changeable_file_receiver_data, error) {
	var res *changeable_file_receiver_data
	if cfr.pm == nil {
		b, err := cfr.pr.Recv()
		if b != nil {
			res = &changeable_file_receiver_data{Filename: b.Filename, Data: b.Payload}
		}
		return res, err
	}
	b, err := cfr.pm.Recv()
	if b != nil {
		res = &changeable_file_receiver_data{Filename: b.Filename, Data: b.Data}
	}
	return res, err

}
