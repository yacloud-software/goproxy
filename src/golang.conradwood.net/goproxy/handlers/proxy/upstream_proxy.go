package proxy

/*
this handler forwards requests to an upstream proxy
*/
import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	pb "golang.conradwood.net/apis/goproxy"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/http"
	"golang.conradwood.net/goproxy/cacher"
	"golang.conradwood.net/goproxy/config"
	"golang.conradwood.net/goproxy/datatypes"
	"io"
	"regexp"
	"strings"
)

var (
	debug_upstream = flag.Bool("debug_upstream_proxy", false, "if true, debug upstream proxy code")
)

type upstream_proxy struct {
	path     string
	fullpath string
	matched  *pb.UpStreamProxy
}

func HandlerByPath(ctx context.Context, path string, fullpath string) (*upstream_proxy, error) {
	for _, gp := range config.GetConfig().GoProxies {
		r := get_regex(gp.Matcher)
		matches := r.MatchString(path)
		if *debug_upstream {
			fmt.Printf("matching \"%s\" with \"%s\"? %v\n", path, gp.Matcher, matches)
		}
		if !matches {
			continue
		}
		if *debug_upstream {
			fmt.Printf("best match is \"%s\" from %s\n", path, gp.Proxy)
		}

		return &upstream_proxy{matched: gp, path: path, fullpath: fullpath}, nil

	}
	return nil, nil
}

// TODO: add cache
func get_regex(r string) *regexp.Regexp {
	return regexp.MustCompile(r)
}
func (up *upstream_proxy) ModuleInfo() *pb.ModuleInfo {
	mi := &pb.ModuleInfo{ModuleType: pb.MODULETYPE_UPSTREAMPROXY, Exists: true}
	return mi
}
func (up *upstream_proxy) ListVersions(ctx context.Context) ([]*pb.VersionInfo, error) {
	body, err := up.download(ctx, nil)
	if err != nil {
		return nil, err
	}
	//	fmt.Printf("Response:\n%s\n", string(body))
	var res []*pb.VersionInfo
	for _, line := range strings.Split(string(body), "\n") {
		v := &pb.VersionInfo{VersionName: line}
		res = append(res, v)
	}
	return res, nil
}

// return the latest versioninfo from a go string (e.g. "v0.120.0")
func (up *upstream_proxy) GetLatestVersion(ctx context.Context) (*pb.VersionInfo, error) {
	b, err := up.download(ctx, nil)
	if err != nil {
		return nil, err
	}
	lf := &datatypes.LatestVersion{}
	err = json.Unmarshal(b, lf)
	if err != nil {
		fmt.Printf("Response:\n%s\n", string(b))
		fmt.Printf("Failed to parse: %s\n", err)
		return nil, err
	}
	vi := &pb.VersionInfo{VersionName: lf.Version}
	return vi, nil
}

// get the zip file for a version
func (up *upstream_proxy) GetZip(ctx context.Context, c *cacher.Cache, w io.Writer, version string) error {
	b, err := up.download(ctx, c)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	if err != nil {
		return err
	}
	return nil

}

// get the go.mod (url like .../@v/[version].mod
func (up *upstream_proxy) GetMod(ctx context.Context, c *cacher.Cache, version string) ([]byte, error) {
	b, err := up.download(ctx, c)
	return b, err
}

func (up *upstream_proxy) download(ctx context.Context, c *cacher.Cache) ([]byte, error) {
	ht := up.getHttp()
	url := strings.TrimSuffix(up.matched.Proxy, "/") + "/" + up.fullpath
	var b []byte

	tb := &bytes.Buffer{}
	var xerr error
	if c != nil && c.IsAvailable(ctx) {
		xerr = c.Get(ctx, func(b []byte) error {
			_, err := tb.Write(b)
			return err
		})
	}
	b = tb.Bytes()
	if xerr != nil || len(b) == 0 {
		up.Debugf("downloading %s", url)
		hr := ht.Get(url)
		err := hr.Error()
		if err != nil {
			up.Debugf("failed to get url %s: %s", url, err)
			return nil, err
		}
		b = hr.Body()
		code := hr.HTTPCode()
		up.Debugf("retrieved %s, %d bytes, code=%d", url, len(b), code)
		if code == 404 {
			return nil, errors.NotFound(ctx, "%s returned with code %d", url, code)
		}
		if code < 200 || code >= 300 {
			return nil, fmt.Errorf("%s returned with code %d", url, code)
		}
	}
	if c != nil {
		c.PutBytes(ctx, b)
	}
	return b, nil
}
func (up *upstream_proxy) getHttp() http.HTTPIF {
	res := http.NewCachingClient(authremote.Context())
	if up.matched.Username != "" || up.matched.Password != "" || up.matched.Token != "" {
		res = http.NewDirectClient()
	}
	if up.matched.Username != "" || up.matched.Password != "" {
		res.SetCreds(up.matched.Username, up.matched.Password)
	}
	if up.matched.Token != "" {
		res.SetHeader("Authorization", "Bearer "+up.matched.Token)
	}
	return res
}

// only serve from cache if this is enabled
func (up *upstream_proxy) CacheEnabled() bool {
	return true
}
func (up *upstream_proxy) Debugf(format string, args ...interface{}) {
	if !*debug_upstream {
		return
	}
	s1 := fmt.Sprintf("[UPSTREAMPROXY path=%s] ", up.path)
	s2 := fmt.Sprintf(format, args...)
	fmt.Println(s1 + s2)
}




