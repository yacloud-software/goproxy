package ext

import (
	"context"
	"flag"
	"fmt"
	pb "golang.conradwood.net/apis/goproxy"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/http"
	"golang.conradwood.net/goproxy/cacher"
	"io"
	"strings"
	"time"
)

var (
	use_urlcacher = flag.Bool("use_urlcacher", false, "if true, use url cacher for external urls")
)

type exthandler struct {
	path string
}

func HandlerByPath(ctx context.Context, path string) (*exthandler, error) {
	if strings.Contains(path, "/apis/") {
		return nil, nil
	}
	res := &exthandler{path: path}
	return res, nil
}
func (e *exthandler) ModuleInfo() *pb.ModuleInfo {
	mi := &pb.ModuleInfo{ModuleType: pb.MODULETYPE_EXTERNALMODULE, Exists: true}
	return mi
}
func getHTTP(ctx context.Context) http.HTTPIF {
	if *use_urlcacher {
		return http.NewCachingClient(ctx)
	}
	return http.NewDirectClient()

}
func (e *exthandler) ListVersions(ctx context.Context) ([]*pb.VersionInfo, error) {
	h := getHTTP(ctx)
	h.SetTimeout(time.Duration(20) * time.Second)
	url := "https://proxy.golang.org/" + e.path + "/@v/list"
	hr := h.Get(url)
	err := hr.Error()
	if err != nil {
		fmt.Printf("Failed to access \"%s\": %s (%d)\n", url, err, hr.HTTPCode())
		fmt.Printf("body:\n%s\n---endbody--\n", string(hr.Body()))
		return nil, err
	}
	body := hr.Body()
	var res []*pb.VersionInfo
	for _, line := range strings.Split(string(body), "\n") {
		v := &pb.VersionInfo{VersionName: line}
		res = append(res, v)
	}
	return res, nil

}
func (e *exthandler) GetLatestVersion(ctx context.Context) (*pb.VersionInfo, error) {
	h := getHTTP(ctx)
	h.SetTimeout(time.Duration(20) * time.Second)
	u := "https://proxy.golang.org/" + e.path + "/@v/latest"
	hr := h.Get(u)
	fmt.Printf("External getting url \"%s\"\n", u)
	err := hr.Error()
	if err != nil {
		return nil, err
	}
	b := hr.Body()
	fmt.Printf("Response (latest): %s\n", string(b))
	return nil, errors.NotFound(ctx, "'latest' not found", "'latest' not found")
}
func (e *exthandler) GetZip(ctx context.Context, c *cacher.Cache, w io.Writer, version string) error {
	h := getHTTP(ctx)
	h.SetTimeout(time.Duration(20) * time.Second)
	hr := h.Get("https://proxy.golang.org/" + e.path + "/@v/" + version + ".zip")
	err := hr.Error()
	if err != nil {
		return err
	}
	b := hr.Body()
	_, err = w.Write(b)
	if err != nil {
		return err
	}

	err = c.PutBytes(ctx, b)
	if err != nil {
		return err
	}

	return nil
}
func (e *exthandler) GetMod(ctx context.Context, c *cacher.Cache, version string) ([]byte, error) {
	h := getHTTP(ctx)
	h.SetTimeout(time.Duration(20) * time.Second)
	hr := h.Get("https://proxy.golang.org/" + e.path + "/@v/" + version + ".mod")
	err := hr.Error()
	if err != nil {
		return nil, err
	}
	b := hr.Body()
	err = c.PutBytes(ctx, b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
