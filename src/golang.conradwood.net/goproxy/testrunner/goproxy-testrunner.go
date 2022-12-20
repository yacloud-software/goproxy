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
	totalCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "goproxy_testrunner_total_requests",
			Help: "V=1 UNIT=none DESC=incremented each time a request is received",
		},
		[]string{"test"},
	)
	failCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "goproxy_testrunner_failed_requests",
			Help: "V=1 UNIT=none DESC=incremented each time a request failed",
		},
		[]string{"test"},
	)
	timsummary = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "goproxy_testrunner_req_timing",
			Help: "V=1 UNIT=s DESC=Summmary for observed requests",
		},
		[]string{"test"},
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

		var test testrun

		l := prometheus.Labels{"test": "test1"}
		totalCounter.With(l).Inc()
		started := time.Now()
		test = &testrun1{}
		err := test.Run()
		if err != nil {
			failCounter.With(l).Inc()
			fmt.Printf("TestRun failed: %s\n", err)
		} else {
			timsummary.With(l).Observe(time.Since(started).Seconds())
		}
	}
}

type testrun interface {
	Run() error
}
