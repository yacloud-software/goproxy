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
	version    = flag.String("version", "", "if set get for that version, otherwise get 'latest'")
)

func main() {
	flag.Parse()

	echoClient = pb.GetGoProxyClient()

	// a context with authentication
	pname := flag.Args()[0]
	v := *version
	if v == "" {
		// get version
		bd, err := get(pname + "/@v/latest")
		utils.Bail("failed to get latest", err)
		fmt.Printf("Latest: %s\n", string(bd))
		bd, err = get(pname + "/@v/list")
		utils.Bail("failed to get list", err)
		fmt.Printf("Other Versions: %s\n", string(bd))
		os.Exit(0)
	}

	bd, err := get(pname + "/@v/" + v + ".info")
	utils.Bail("failed to get info", err)
	fmt.Printf("Info for version \"%s\": %s\n", v, string(bd))

	bd, err = get(pname + "/@v/" + v + ".mod")
	utils.Bail("failed to get mod", err)
	fmt.Printf("Mod for version \"%s\":\n%s\n", v, string(bd))

	bd, err = get(pname + "/@v/" + v + ".zip")
	utils.Bail("failed to get", err)
	fname := "/tmp/rec.bin"
	utils.Bail("failed to write", utils.WriteFile(fname, bd))
	fmt.Printf("%d bytes written to %s\n", len(bd), fname)
	fmt.Printf("Done.\n")
	os.Exit(0)
}

func get(path string) ([]byte, error) {
	ctx := authremote.Context()
	empty := &pb.GetPathRequest{Path: path}
	srv, err := echoClient.GetPath(ctx, empty)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	for {
		o, err := srv.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		_, err = buf.Write(o.Data)
		if err != nil {
			return nil, err
		}
	}
	bd := buf.Bytes()
	return bd, nil
}
