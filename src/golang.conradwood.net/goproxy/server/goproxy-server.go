package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/goproxy/goproxy"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/goproxy"
	"golang.conradwood.net/go-easyops/server"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/goproxy/cacher"
	"google.golang.org/grpc"
	"net/http"
	"os"
)

var (
	port      = flag.Int("port", 4100, "The grpc server port")
	http_port = flag.Int("http_port", 4108, "The http server port")
	gopr      *goproxy.Goproxy
)

type echoServer struct {
}

func main() {
	var err error
	flag.Parse()
	fmt.Printf("Starting GoProxyServer...\n")
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

	sd = server.NewServerDef()
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

func (e *echoServer) Ping(ctx context.Context, req *common.Void) (*pb.PingResponse, error) {
	resp := &pb.PingResponse{Response: "pingresponse"}
	return resp, nil
}
