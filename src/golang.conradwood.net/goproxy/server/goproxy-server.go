package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/goproxy/goproxy"
	pb "golang.conradwood.net/apis/goproxy"
	h2g "golang.conradwood.net/apis/h2gproxy"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/prometheus"
	"golang.conradwood.net/go-easyops/server"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/goproxy/cacher"
	"golang.conradwood.net/goproxy/handlers"
	"google.golang.org/grpc"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	answer_all_with_404 = flag.Bool("answer_all_with_404", false, "if true, just answer 404 on all")
	totalCounter        = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "goproxy_total_requests",
			Help: "V=1 UNIT=none DESC=incremented each time a request is received",
		},
		[]string{"reqtype"},
	)
	failCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "goproxy_failed_requests",
			Help: "V=1 UNIT=none DESC=incremented each time a request failed",
		},
		[]string{"reqtype"},
	)
	timsummary = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "goproxy_req_timing",
			Help: "V=1 UNIT=s DESC=Summmary for observed requests",
		},
		[]string{"handler", "reqtype"},
	)

	singleton_lock sync.Mutex
	singleton      = flag.Bool("singleton", false, "if true only allow one access at a time (for debugging)")
	port           = flag.Int("port", 4100, "The grpc server port")
	http_port      = flag.Int("http_port", 4108, "The http server port")
	gopr           *goproxy.Goproxy
	debug          = flag.Bool("debug", false, "debug mode")
	index          int
)

type echoServer struct {
}
type SingleRequest struct {
	Prefix  string // for printing
	Started time.Time
	mi      *pb.ModuleInfo
}

func (sr *SingleRequest) promLabels() prometheus.Labels {
	r := "undef"
	if sr.mi != nil {
		r = fmt.Sprintf("%v", sr.mi.ModuleType)
	}
	l := prometheus.Labels{"reqtype": r}
	return l
}
func (sr *SingleRequest) Printf(format string, args ...interface{}) {
	mis := ""
	if sr.mi != nil {
		mis = fmt.Sprintf("%v ", sr.mi.ModuleType)
	}
	pre := fmt.Sprintf("[%s %s%0.2fs] ", sr.Prefix, mis, time.Since(sr.Started).Seconds())
	sn := fmt.Sprintf(format, args...)
	fmt.Print(pre + sn)
}

func main() {
	var err error
	flag.Parse()
	fmt.Printf("Starting GoProxyServer...\n")
	prometheus.MustRegister(timsummary, failCounter, totalCounter)
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

type streamer interface {
	Send(*h2g.StreamDataResponse) error
	Context() context.Context
}

func (e *echoServer) StreamHTTP(req *h2g.StreamRequest, srv pb.GoProxy_StreamHTTPServer) error {
	return e.streamHTTP(req, srv)
}
func (e *echoServer) streamHTTP(req *h2g.StreamRequest, srv streamer) error {
	if *singleton {
		singleton_lock.Lock()
		defer singleton_lock.Unlock()
	}
	if *answer_all_with_404 {
		fmt.Printf("Returning 404 because of -answer_all_with_404 to %s\n", req.Path)
		return errors.NotFound(srv.Context(), "proxy serves all with 404 atm")
	}
	index++
	sr := SingleRequest{
		Prefix:  fmt.Sprintf("%d", index),
		Started: time.Now(),
	}
	sr.Printf("-------------------\nStarted...\n")
	ctx := srv.Context()
	u := auth.GetUser(ctx)
	if u == nil {
		sr.Printf("unauthenticated access failure\n")
		if *debug {
			fmt.Printf("Unauthenticated access to \"https://%s://%s\"\n", req.Host, req.Path)
		}
		return errors.Unauthenticated(ctx, "login required")
	}
	path := strings.TrimPrefix(req.Path, "/")
	sr.Printf("Access to path \"%s\" by %s\n", path, auth.Description(u))

	idx := strings.Index(path, "/@v/")
	find_path := path
	if idx != -1 {
		find_path = path[:idx]
	}
	handler, err := handlers.HandlerByPath(ctx, find_path)
	if handler != nil {
		sr.mi = handler.ModuleInfo()
	}
	totalCounter.With(sr.promLabels()).Inc()
	if err != nil {
		sr.Printf("failed\n")
		failCounter.With(sr.promLabels()).Inc()
		return err
	}
	if handler == nil {
		failCounter.With(sr.promLabels()).Inc()
		sr.Printf("No handler found for \"%s\".\n", path)
		return errors.NotFound(ctx, "no handler \"%s\" found", path)
	}
	mi := handler.ModuleInfo()
	if err != nil {
		sr.Printf("no moduleinfo error\n")
		failCounter.With(sr.promLabels()).Inc()
		return err
	}
	if *debug {
		print := false
		if mi.ModuleType == pb.MODULETYPE_PROTO {
			print = true
		}
		if print {
			sr.Printf("ModuleInfo for \"%s\"\n", find_path)
			if mi == nil {
				fmt.Printf("None.\n")
				return errors.NotFound(ctx, "no module \"%s\" found", path)
			}
			sr.Printf("Type      : %s\n", mi.ModuleType)
			sr.Printf("Exists    : %v\n", mi.Exists)
		}
	}
	if mi.ModuleType == pb.MODULETYPE_UNKNOWN {
		failCounter.With(sr.promLabels()).Inc()
		sr.Printf("unknown module serving this path\n")
		return errors.InvalidArgs(ctx, "module \"%s\" resolved to unknown", path)
	}
	if !mi.Exists {
		failCounter.With(sr.promLabels()).Inc()
		sr.Printf("module serving this path does not exist\n")
		return errors.NotFound(ctx, "no module \"%s\" found", path)
	}
	reqtype := ""
	if strings.HasSuffix(path, "/@v/list") {
		reqtype = "list"
		err = sr.serveList(handler, req, srv)
	} else if strings.HasSuffix(path, "/@v/latest") {
		reqtype = "latest"
		err = fmt.Errorf("/@v/latest not supported")
	} else if strings.HasSuffix(path, "@latest") {
		reqtype = "latest"
		err = sr.serveLatest(handler, req, srv)
	} else if strings.HasSuffix(path, ".info") {
		reqtype = "info"
		err = sr.serveInfo(handler, req, srv)
	} else if strings.HasSuffix(path, ".zip") {
		reqtype = "zip"
		err = sr.serveZip(handler, req, srv)
	} else if strings.HasSuffix(path, ".mod") {
		reqtype = "mod"
		err = sr.serveMod(handler, req, srv)
	} else {
		// must be notfound so that go tries to download alternative paths
		err = errors.NotFound(ctx, "invalid path \"%s\"", req.Path)
	}
	if err != nil {
		failCounter.With(sr.promLabels()).Inc()
		if errors.ToHTTPCode(err).ErrorCode == 404 {
			sr.Printf("[from %s] not found: \"%s\"\n", mi.ModuleType, path)
			sendError(srv, 404)
		} else {
			sr.Printf("Error for \"%s\" in %v: %s\n", path, mi.ModuleType, utils.ErrorString(err))
			return err
		}
	}
	req_timing(fmt.Sprintf("%v", mi.ModuleType), reqtype, time.Since(sr.Started))
	sr.Printf("Completed in %0.2fs.\n", time.Since(sr.Started).Seconds())
	return nil
}
func sendError(srv streamer, code uint32) {
	e := &h2g.StreamDataResponse{
		Response: &h2g.StreamResponse{
			StatusCode: code,
		},
	}
	err := srv.Send(e)
	if err != nil {
		fmt.Printf("Failed to send error!! (%s)\n", err)
	}
}
func req_timing(modtype, reqtype string, dur time.Duration) {
	l := prometheus.Labels{"handler": modtype, "reqtype": reqtype}
	timsummary.With(l).Observe(dur.Seconds())
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
func (sr *SingleRequest) serveMod(handler handlers.Handler, req *h2g.StreamRequest, srv streamer) error {
	ctx := srv.Context()
	version_string, err := versionFromPath(ctx, req.Path)
	if err != nil {
		return err
	}
	sr.Printf("Version: \"%s\"\n", version_string)
	nc, err := cacher.NewCacher(ctx, req.Path, version_string, "mod")
	if err != nil {
		return err
	}
	if handler.CacheEnabled() {
		if nc.IsAvailable(ctx) {
			sr.Printf("Serving from cache...\n")
			return nc.Get(ctx, func(data []byte) error {
				return srv.Send(&h2g.StreamDataResponse{Data: data})
			})
		}
	}
	b, err := handler.GetMod(ctx, nc, version_string)
	if err != nil {
		return err
	}
	return sendBytes(srv, b)
}

// serve requests suffixed by /@v/[version].zip
func (sr *SingleRequest) serveZip(handler handlers.Handler, req *h2g.StreamRequest, srv streamer) error {
	ctx := srv.Context()
	version_string, err := versionFromPath(ctx, req.Path)
	if err != nil {
		return err
	}
	nc, err := cacher.NewCacher(ctx, req.Path, version_string, "zip")
	if err != nil {
		return err
	}
	if handler.CacheEnabled() {
		if nc.IsAvailable(ctx) {
			sr.Printf("Serving from cache...\n")
			return nc.Get(ctx, func(data []byte) error {
				return srv.Send(&h2g.StreamDataResponse{Data: data})
			})
		}
	}
	sr.Printf("Version: \"%s\"\n", version_string)
	sw := NewStreamWriter(srv)
	err = handler.GetZip(ctx, nc, sw, version_string)
	if err != nil {
		return err
	}
	return nil
}

// serve requests suffixed by /@v/[version].info
func (sr *SingleRequest) serveInfo(handler handlers.Handler, req *h2g.StreamRequest, srv streamer) error {
	sr.Printf("Serving version for \"%s\"\n", req.Path)
	ctx := srv.Context()
	version_string, err := versionFromPath(ctx, req.Path)
	if err != nil {
		return err
	}
	sr.Printf("Version: \"%s\"\n", version_string)
	res := fmt.Sprintf(`{"Version":"%s"}`, version_string)
	return sendBytes(srv, []byte(res))
}

// serve requests suffixed by /@v/latest
func (sr *SingleRequest) serveLatest(handler handlers.Handler, req *h2g.StreamRequest, srv streamer) error {
	sr.Printf("Serving latest for \"%s\"\n", req.Path)
	vi, err := handler.GetLatestVersion(srv.Context())
	if err != nil {
		return err
	}
	version_string := versionToString(vi)
	sr.Printf("Version: \"%s\"\n", version_string)
	res := fmt.Sprintf(`{"Version":"%s"}`, version_string)
	return sendBytes(srv, []byte(res))
}

// serve requests suffixed by /@v/list
func (sr *SingleRequest) serveList(handler handlers.Handler, req *h2g.StreamRequest, srv streamer) error {
	sr.Printf("Serving list for \"%s\" from %v\n", req.Path, handler.ModuleInfo().ModuleType)
	started := time.Now()
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
		res = versionToString(v) + "\n" + res
	}
	sr.Printf("Created list for %s in %0.2fs\n", req.Path, time.Since(started).Seconds())
	return sendBytes(srv, []byte(res))
}
func versionToString(v *pb.VersionInfo) string {
	if v.Version != 0 {
		return fmt.Sprintf("v1.1.%d", v.Version)
	}
	return v.VersionName
}
func sendBytes(srv streamer, b []byte) error {
	if *debug {
		if len(b) < 800 {
			fmt.Printf("Response:\n")
			fmt.Printf("%s\n", string(b))
		}
	}
	err := srv.Send(&h2g.StreamDataResponse{Response: &h2g.StreamResponse{
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
		err = srv.Send(&h2g.StreamDataResponse{Data: b[total:eo]})
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
