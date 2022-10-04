package main

import (
	"flag"
	"fmt"
	pb "golang.conradwood.net/apis/goproxy"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/utils"
	"os"
)

var (
	echoClient pb.GoProxyClient
)

func main() {
	flag.Parse()

	echoClient = pb.GetGoProxyClient()

	// a context with authentication
	ctx := authremote.Context()

	empty := &pb.ModuleInfoRequest{}
	response, err := echoClient.AnalyseURL(ctx, empty)
	utils.Bail("Failed to ping server", err)
	fmt.Printf("Response to ping: %v\n", response)

	fmt.Printf("Done.\n")
	os.Exit(0)
}
