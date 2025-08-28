package main

import (
	"context"
	"flag"
	"fmt"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/goproxy"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/prometheus"
	"golang.conradwood.net/go-easyops/server"
	"golang.conradwood.net/go-easyops/utils"
	"google.golang.org/grpc"
	"os"
	"sync"
	"time"
)

var (
	testrunlock  sync.Mutex
	trigger      = flag.Bool("trigger", false, "if true, don't run,but trigger one in dc")
	delay        = flag.Duration("delay", time.Duration(30)*time.Minute, "delay between tests")
	enable       = flag.Bool("enable", true, "if false, do not run tests")
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
			Help: "V=1 UNIT=s DESC=Summmary for observed sections of tests",
		},
		[]string{"test", "section"},
	)
	timtotalsummary = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "goproxy_testrunner_total_timing",
			Help: "V=1 UNIT=s DESC=Summmary for observed tests",
		},
		[]string{"test", "result"},
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
   server.SetHealth(common.Health_STARTING)
	if *trigger {
		Trigger()
		os.Exit(0)
	}
	fmt.Printf("Starting GoProxy Testrunner Server...\n")
	prometheus.SetExpiry(time.Duration(24) * time.Hour)
	prometheus.MustRegister(timtotalsummary, timsummary, failCounter, totalCounter)
	go testrunner()
	sd := server.NewServerDef()
	sd.SetPort(*port)
sd.SetOnStartupCallback(startup)
	sd.SetRegister(server.Register(
		func(server *grpc.Server) error {
			e := new(echoServer)
			pb.RegisterGoProxyTestRunnerServer(server, e)
			return nil
		},
	))
	err = server.ServerStartup(sd)
	utils.Bail("Unable to start server", err)
	os.Exit(0)
}
func startup() {
	server.SetHealth(common.Health_READY)
}


func testrunner() {
	t := time.Duration(2) * time.Second

	for {
		time.Sleep(t)
		t = *delay
		if !*enable {
			continue
		}
		do_run()
	}
}
func do_run() {
	testrunlock.Lock()
	defer testrunlock.Unlock()
	var testruns []testrun
	testruns = append(testruns, NewModTestRun("test_mod1", "mod1"))
	testruns = append(testruns, NewModTestRun("test_mod2", "mod2"))

	fmt.Println("-------- starting run ---------")
	for _, test := range testruns {
		err := utils.RecreateSafely(godir())
		if err != nil {
			fmt.Printf("Failed to recreate godir (%s): %s\n", godir(), err)
			continue
		}
		test_started := time.Now()
		for section := 0; section < test.Sections(); section++ {
			l := prometheus.Labels{"test": test.Name(), "section": fmt.Sprintf("%d", section)}
			totalCounter.With(l).Inc()
			test.Printf("Starting Section %d, test %s...\n", section, test.Name())
			started := time.Now()
			err = test.Run(section)
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
		s := "ok"
		if err != nil {
			s = "failed"
		}
		test.Printf("Test completed, result: %s <---- *** TEST RESULT *** \n", s)
		l := prometheus.Labels{"test": test.Name(), "result": s}
		dur := time.Since(test_started).Seconds()
		timtotalsummary.With(l).Observe(dur)
	}
	fmt.Println("-------- finished run ---------")

}

type testrun interface {
	Name() string
	Sections() int
	Run(section int) error
	Printf(format string, args ...interface{})
}

func (e *echoServer) Trigger(ctx context.Context, req *common.Void) (*common.Void, error) {
	go do_run()
	return &common.Void{}, nil
}
func Trigger() {
	fmt.Printf("Triggering run...\n")
	ctx := authremote.Context()
	_, err := pb.GetGoProxyTestRunnerClient().Trigger(ctx, &common.Void{})
	utils.Bail("failed to trigger", err)
	fmt.Printf("Triggered\n")
}







