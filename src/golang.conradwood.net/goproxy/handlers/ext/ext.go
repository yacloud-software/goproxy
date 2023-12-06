package ext

/*
this handles modules that are neither artefact nor proto, basically the "stuff from the net"

*/
import (
	"context"
	"flag"
	"fmt"
	pb "golang.conradwood.net/apis/goproxy"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/http"
	"golang.conradwood.net/goproxy/cacher"
	"golang.conradwood.net/goproxy/config"
	"io"
	"strings"
	"time"
)

var (
	debug         = flag.Bool("debug_exthandler", false, "if true debug exthandler")
	cache_ext     = flag.Bool("cache_ext_files", false, "if true cache external files")
	use_urlcacher = flag.Bool("use_urlcacher", true, "if true, use url cacher for external urls")
)

type exthandler struct {
	path string
}

func useInternalCache() bool {
	if *use_urlcacher {
		return false
	}
	return *cache_ext
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
	if strings.HasPrefix(e.path, "golang.conradwood.net") {
		return nil, errors.NotFound(ctx, "not implemented, must handle local paths")
	}
	h := getHTTP(ctx)
	h.SetTimeout(time.Duration(20) * time.Second)
	url := e.buildurl(e.path + "/@v/list")
	e.Printf("(list) Getting url \"%s\"...\n", url)
	started := time.Now()
	hr := h.Get(url)
	e.Printf("Got response for url \"%s\" in %0.2fs\n", url, time.Since(started).Seconds())
	err := hr.Error()
	if err != nil {
		if hr.HTTPCode() == 404 {
			e.Printf("(list) not found: \"%s\"\n", url)
		} else {
			e.Printf("(list) Failed to access \"%s\": %s (%d)\n", url, err, hr.HTTPCode())
		}
		//e.Printf("body:\n%s\n---endbody--\n", string(hr.Body()))
		return nil, err
	}
	e.Printf("(list) Reading response body for url \"%s\"\n", url)
	started = time.Now()
	body := hr.Body()
	e.Printf("(list) Read body for url %s in %0.2fs\n", url, time.Since(started).Seconds())
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
	u := e.buildurl(e.path + "/@v/latest")
	hr := h.Get(u)
	e.Printf("External getting url \"%s\"\n", u)
	err := hr.Error()
	if err != nil {
		return nil, err
	}
	b := hr.Body()
	e.Printf("Response (latest): %s\n", string(b))
	return nil, errors.NotFound(ctx, "'latest' not found", "'latest' not found")
}
func (e *exthandler) GetZip(ctx context.Context, c *cacher.Cache, w io.Writer, version string) error {
	h := getHTTP(ctx)
	h.SetTimeout(time.Duration(20) * time.Second)
	url := e.buildurl(e.path + "/@v/" + version + ".zip")
	e.Printf("External getting url \"%s\"\n", url)
	hr := h.Get(url)
	err := hr.Error()
	if err != nil {
		return err
	}
	b := hr.Body()
	_, err = w.Write(b)
	if err != nil {
		return err
	}
	if useInternalCache() {
		err = c.PutBytes(ctx, b)
		if err != nil {
			return err
		}
	}
	return nil
}
func (e *exthandler) GetMod(ctx context.Context, c *cacher.Cache, version string) ([]byte, error) {
	h := getHTTP(ctx)
	h.SetTimeout(time.Duration(20) * time.Second)
	url := e.buildurl(e.path + "/@v/" + version + ".mod")
	e.Printf("Getting url \"%s\"\n", url)
	hr := h.Get(url)
	err := hr.Error()
	if err != nil {
		return nil, err
	}
	b := hr.Body()
	if useInternalCache() {
		err = c.PutBytes(ctx, b)
		if err != nil {
			return nil, err
		}
	}
	return b, nil
}
func (e *exthandler) buildurl(path string) string {
	//return "https://proxy.golang.org/" + path
	// return "http://127.0.0.1:8080/" + path
	s := config.GetConfig().GoGetProxy
	s = strings.TrimSuffix(s, "/")
	return s + "/" + path
	// return "http://172.29.1.11:14231/" + path // the go-get-proxy from https://github.com/goproxy/goproxy
}
func (e *exthandler) CacheEnabled() bool {
	if !*cache_ext {
		return false
	}
	if *use_urlcacher {
		return false
	}
	return true
}
func (e *exthandler) Printf(format string, args ...interface{}) {
	if !*debug {
		return
	}
	s := "[exthandler] "
	sn := fmt.Sprintf(format, args...)
	fmt.Print(s + sn)
}

