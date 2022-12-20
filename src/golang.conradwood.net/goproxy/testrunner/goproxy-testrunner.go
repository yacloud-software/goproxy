package main

import (
	"flag"
	"fmt"
	pb "golang.conradwood.net/apis/goproxy"
	"golang.conradwood.net/go-easyops/prometheus"
	"golang.conradwood.net/go-easyops/server"
	"golang.conradwood.net/go-easyops/utils"
	"google.golang.org/grpc"
	"os"
	"time"
)

var (
	enable       = flag.Bool("enable", false, "if false, do not run tests")
	totalCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "goproxy_testrunner_total_requests",
			Help: "V=1 UNIT=none DESC=incremented each time a request is received",
		},
		[]string{"test", "section"},
	)
	failCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "goproxy_testrunner_failed_requests",
			Help: "V=1 UNIT=none DESC=incremented each time a request failed",
		},
		[]string{"test", "section"},
	)
	timsummary = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "goproxy_testrunner_req_timing",
			Help: "V=1 UNIT=s DESC=Summmary for observed requests",
		},
		[]string{"test", "section"},
	)

	port      = flag.Int("port", 4100, "The grpc server port")
	http_port = flag.Int("http_port", 4108, "The http server port")
	debug     = flag.Bool("debug", false, "debug mode")
)

type echoServer struct {
}

func main() {
	var err error
	flag.Parse()
	fmt.Printf("Starting GoProxy Testrunner Server...\n")
	prometheus.MustRegister(timsummary, failCounter, totalCounter)
	go testrunner()
	sd := server.NewServerDef()
	sd.Port = *port
	sd.Register = server.Register(
		func(server *grpc.Server) error {
			e := new(echoServer)
			pb.RegisterGoProxyTestRunnerServer(server, e)
			return nil
		},
	)
	err = server.ServerStartup(sd)
	utils.Bail("Unable to start server", err)
	os.Exit(0)
}

func testrunner() {
	t := time.Duration(2) * time.Second

	for {
		time.Sleep(t)
		t = time.Duration(120) * time.Second
		if !*enable {
			continue
		}
		var testruns []testrun
		testruns = append(testruns, &testrun1{})
		for _, test := range testruns {
			err := utils.RecreateSafely(godir())
			if err != nil {
				fmt.Printf("Failed to recreate godir (%s): %s\n", godir(), err)
				continue
			}
			for section := 0; section < test.Sections(); section++ {
				l := prometheus.Labels{"test": "test1", "section": fmt.Sprintf("%d", section)}
				totalCounter.With(l).Inc()
				test.Printf("Starting Section %d...\n", section)
				started := time.Now()
				err := test.Run(section)
				dur := time.Since(started).Seconds()
				if err != nil {
					failCounter.With(l).Inc()
					test.Printf("TestRun section %d failed: %s\n", section, utils.ErrorString(err))
					break
				} else {
					test.Printf("section %d completed after %0.2fs\n", section, dur)
					timsummary.With(l).Observe(dur)
				}
			}
		}
	}
}

type testrun interface {
	Sections() int
	Run(section int) error
	Printf(format string, args ...interface{})
}
