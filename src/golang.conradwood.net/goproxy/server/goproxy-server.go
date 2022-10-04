package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/goproxy/goproxy"
	pb "golang.conradwood.net/apis/goproxy"
	hg "golang.conradwood.net/apis/h2gproxy"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/server"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/goproxy/cacher"
	"golang.conradwood.net/goproxy/handlers"
	"google.golang.org/grpc"
	"os"
	"sort"
	"strings"
)

var (
	port      = flag.Int("port", 4100, "The grpc server port")
	http_port = flag.Int("http_port", 4108, "The http server port")
	gopr      *goproxy.Goproxy
	debug     = flag.Bool("debug", false, "debug mode")
)

type echoServer struct {
}

func main() {
	var err error
	flag.Parse()
	fmt.Printf("Starting GoProxyServer...\n")
	/*
		gopr = &goproxy.Goproxy{}
		gopr.Cacher = &cacher.GoCacher{}
		gopr.GoBinEnv = goenv()
		//	gopr.Cacher = nil
		go func() {
			adr := fmt.Sprintf(":%d", *http_port)
			err := http.ListenAndServe(adr, gopr)
			utils.Bail("failed to start http server", err)
		}()
		sd := server.NewHTMLServerDef("goproxy.GoProxy")
		sd.Port = *http_port
		server.AddRegistry(sd)
	*/
	sd := server.NewServerDef()
	sd.Port = *port
	sd.Register = server.Register(
		func(server *grpc.Server) error {
			e := new(echoServer)
			pb.RegisterGoProxyServer(server, e)
			return nil
		},
	)
	err = server.ServerStartup(sd)
	utils.Bail("Unable to start server", err)
	os.Exit(0)
}

func goenv() []string {
	res := []string{
		"GOPATH=/tmp/src",
		"PATH=/opt/yacloud/ctools/dev/bin:/opt/yacloud/ctools/dev/go/current/protoc:/opt/yacloud/ctools/dev/go/current/go/bin:/sbin:/usr/sbin:/bin:/sbin:/usr/sbin:/usr/bin",
	}
	return res
}

/************************************
* grpc functions
************************************/

func (e *echoServer) AnalyseURL(ctx context.Context, req *pb.ModuleInfoRequest) (*pb.ModuleInfo, error) {
	return nil, nil
}
func (e *echoServer) StreamHTTP(req *hg.StreamRequest, srv pb.GoProxy_StreamHTTPServer) error {
	ctx := srv.Context()
	u := auth.GetUser(ctx)
	if u == nil {
		if *debug {
			fmt.Printf("Unauthenticated access to \"https://%s://%s\"\n", req.Host, req.Path)
		}
		return errors.Unauthenticated(ctx, "login required")
	}
	path := strings.TrimPrefix(req.Path, "/")
	fmt.Printf("Access to path \"%s\" by %s\n", path, auth.Description(u))

	idx := strings.Index(path, "/@v/")
	find_path := path
	if idx != -1 {
		find_path = path[:idx]
	}
	handler, err := handlers.HandlerByPath(ctx, find_path)
	if err != nil {
		return err
	}
	if handler == nil {
		fmt.Printf("No handler found for \"%s\".\n", path)
		return errors.NotFound(ctx, "no handler \"%s\" found", path)
	}
	mi := handler.ModuleInfo()
	if err != nil {
		return err
	}
	if *debug {
		print := false
		if mi.ModuleType == pb.MODULETYPE_PROTO {
			print = true
		}
		if print {
			fmt.Printf("ModuleInfo for \"%s\"\n", find_path)
			if mi == nil {
				fmt.Printf("None.\n")
				return errors.NotFound(ctx, "no module \"%s\" found", path)
			}
			fmt.Printf("Type      : %s\n", mi.ModuleType)
			fmt.Printf("Exists    : %v\n", mi.Exists)
		}
	}
	if mi.ModuleType == pb.MODULETYPE_UNKNOWN {
		return errors.InvalidArgs(ctx, "module \"%s\" resolved to unknown", path)
	}
	if !mi.Exists {
		return errors.NotFound(ctx, "no module \"%s\" found", path)
	}

	if strings.HasSuffix(path, "/@v/list") {
		err = serveList(handler, req, srv)
	} else if strings.HasSuffix(path, "/@v/latest") {
		err = serveLatest(handler, req, srv)
	} else if strings.HasSuffix(path, ".info") {
		err = serveInfo(handler, req, srv)
	} else if strings.HasSuffix(path, ".zip") {
		err = serveZip(handler, req, srv)
	} else if strings.HasSuffix(path, ".mod") {
		err = serveMod(handler, req, srv)
	} else {
		err = errors.InvalidArgs(ctx, "invalid path", "invalid path \"%s\"", req.Path)
	}
	if err != nil {
		fmt.Printf("Error for \"%s\": %s\n", req.Path, err)
		return err
	}
	return nil

}

func versionFromPath(ctx context.Context, path string) (string, error) {
	idx := strings.Index(path, "/@v/")
	if idx == -1 {
		return "", errors.InvalidArgs(ctx, "invalid path", "invalid path")
	}
	version_string := path[idx+4:]
	idx = strings.LastIndex(version_string, ".")
	if idx == -1 {
		return "", errors.InvalidArgs(ctx, "invalid path", "invalid path")
	}

	version_string = version_string[:idx]
	return version_string, nil
}

// serve requests suffixed by /@v/[version].mod
func serveMod(handler handlers.Handler, req *hg.StreamRequest, srv pb.GoProxy_StreamHTTPServer) error {
	ctx := srv.Context()
	version_string, err := versionFromPath(ctx, req.Path)
	if err != nil {
		return err
	}
	fmt.Printf("Version: \"%s\"\n", version_string)
	nc, err := cacher.NewCacher(ctx, req.Path, version_string, "mod")
	if err != nil {
		return err
	}
	if nc.IsAvailable(ctx) {
		panic("cached already")
	}
	b, err := handler.GetMod(ctx, version_string)
	if err != nil {
		return err
	}
	return sendBytes(srv, b)
}

// serve requests suffixed by /@v/[version].zip
func serveZip(handler handlers.Handler, req *hg.StreamRequest, srv pb.GoProxy_StreamHTTPServer) error {
	ctx := srv.Context()
	version_string, err := versionFromPath(ctx, req.Path)
	if err != nil {
		return err
	}
	nc, err := cacher.NewCacher(ctx, req.Path, version_string, "zip")
	if err != nil {
		return err
	}
	if nc.IsAvailable(ctx) {
		panic("cached already")
	}
	fmt.Printf("Version: \"%s\"\n", version_string)
	sw := NewStreamWriter(srv)
	err = handler.GetZip(ctx, sw, version_string)
	if err != nil {
		return err
	}
	return nil
}

// serve requests suffixed by /@v/[version].info
func serveInfo(handler handlers.Handler, req *hg.StreamRequest, srv pb.GoProxy_StreamHTTPServer) error {
	fmt.Printf("Serving version for \"%s\"\n", req.Path)
	ctx := srv.Context()
	version_string, err := versionFromPath(ctx, req.Path)
	if err != nil {
		return err
	}
	fmt.Printf("Version: \"%s\"\n", version_string)
	res := fmt.Sprintf(`{"Version":"%s"}`, version_string)
	return sendBytes(srv, []byte(res))
}

// serve requests suffixed by /@v/latest
func serveLatest(handler handlers.Handler, req *hg.StreamRequest, srv pb.GoProxy_StreamHTTPServer) error {
	fmt.Printf("Serving latest for \"%s\"\n", req.Path)
	panic("not implemented")
}

// serve requests suffixed by /@v/list
func serveList(handler handlers.Handler, req *hg.StreamRequest, srv pb.GoProxy_StreamHTTPServer) error {
	fmt.Printf("Serving list for \"%s\"\n", req.Path)
	ctx := srv.Context()
	vls, err := handler.ListVersions(ctx)
	if err != nil {
		return err
	}
	sort.Slice(vls, func(i, j int) bool {
		return vls[i].Version < vls[j].Version
	})
	res := ""
	for _, v := range vls {
		res = versionToString(v.Version) + "\n" + res
	}
	return sendBytes(srv, []byte(res))
}
func versionToString(v uint64) string {
	return fmt.Sprintf("v1.1.%d", v)
}
func sendBytes(srv pb.GoProxy_StreamHTTPServer, b []byte) error {
	if *debug {
		if len(b) < 800 {
			fmt.Printf("Response:\n")
			fmt.Printf("%s\n", string(b))
		}
	}
	err := srv.Send(&hg.StreamDataResponse{Response: &hg.StreamResponse{
		Filename: "foo",
		Size:     uint64(len(b)),
		MimeType: "application/octet-stream",
	}})
	if err != nil {
		return err
	}
	total := 0
	for {
		size := 8192
		if total+size >= len(b) {
			size = len(b) - total
		}
		eo := total + size
		err = srv.Send(&hg.StreamDataResponse{Data: b[total:eo]})
		if err != nil {
			return err
		}
		total = total + size
		if total >= len(b) {
			break
		}
	}
	return nil
}
