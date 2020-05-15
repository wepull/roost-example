package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"

	_ "github.com/project-flogo/core/data/expression/script"
	"github.com/project-flogo/core/engine"
	//	"github.com/sirupsen/logrus"
	//"github.com/sirupsen/logrus"
	//"github.com/project-flogo/core/activity"
)

/*const (
	listenPort          = "5050"
	usdCurrency         = "USD"
	zbioServiceEndpoint = "zbio-service:50002"
	topicName           = "checkoutservice"
	zbioEnabled         = true
)

var (
	//log      *logrus.Logger
	zbclient *zb.Client
)*/

var (
	cpuProfile    = flag.String("cpuprofile", "", "Writes CPU profile to the specified file")
	memProfile    = flag.String("memprofile", "", "Writes memory profile to the specified file")
	cfgJson       string
	cfgEngine     string
	cfgCompressed bool
)

func main() {

	cpuProfiling := false
	/*if zbioEnabled {
		initZBIO()
	}*/
	flag.Parse()
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create CPU profiling file: %v\n", err)
			os.Exit(1)
		}
		if err = pprof.StartCPUProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to start CPU profiling: %v\n", err)
			os.Exit(1)
		}
		cpuProfiling = true
	}

	cfg, err := engine.LoadAppConfig(cfgJson, cfgCompressed)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create engine: %v\n", err)
		os.Exit(1)
	}

	e, err := engine.New(cfg, engine.ConfigOption(cfgEngine, cfgCompressed))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create engine: %v\n", err)
		os.Exit(1)
	}
	//fmt.Println(e)
	code := engine.RunEngine(e)
	//code1, err := activity.Activity.Eval(code)
	//fmt.Println(code1)

	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create memory profiling file: %v\n", err)
			os.Exit(1)
		}

		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write memory profiling data: %v", err)
			os.Exit(1)
		}
		_ = f.Close()
	}

	if cpuProfiling {
		pprof.StopCPUProfile()
	}

	/*code1, err := activity.Eval()
	if err != nil {
		errors.New("Error")
	}*/
	//zbclient.sendMessageToZBIO("hello")

	os.Exit(code)
	//os.Exit(code1)
}
