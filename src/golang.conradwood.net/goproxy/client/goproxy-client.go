package main

import (
	"bytes"
	"flag"
	"fmt"
	pb "golang.conradwood.net/apis/goproxy"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/utils"
	"io"
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

	empty := &pb.GetPathRequest{Path: flag.Args()[0]}
	srv, err := echoClient.GetPath(ctx, empty)
	utils.Bail("Failed to ping server", err)
	var buf bytes.Buffer
	for {
		o, err := srv.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			utils.Bail("failed to recv", err)
		}
		_, err = buf.Write(o.Data)
		utils.Bail("failed to write", err)
	}
	fname := "/tmp/rec.bin"
	bd := buf.Bytes()
	utils.Bail("failed to write", utils.WriteFile(fname, bd))
	fmt.Printf("%d bytes written to %s\n", len(bd), fname)
	fmt.Printf("Done.\n")
	os.Exit(0)
}
