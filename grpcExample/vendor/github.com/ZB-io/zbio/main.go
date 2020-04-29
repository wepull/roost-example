package util

import (
	"flag"
	"fmt"
	"github.com/ZB-io/zbio/config"
	"net/http"
	"net/http/httptrace"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"golang.org/x/net/context"
)

/*
func main() {

	// closing...
	defer close()

	// Set up logging
	initLogs()

	// parse from command line if any
	var configuration = parseCLI()

	fmt.Print(configuration)

	//start broker
	startBroker()
}
*/

var tracer opentracing.Tracer

func TestOpenTracing(ctx context.Context) error {
	// retrieve current Span from Context
	var parentCtx opentracing.SpanContext
	parentSpan := opentracing.SpanFromContext(ctx)
	if parentSpan != nil {
		parentCtx = parentSpan.Context()
	}

	// start a new Span to wrap HTTP request
	span := tracer.StartSpan(
		"",
		opentracing.ChildOf(parentCtx),
	)

	// make sure the Span is finished once we're done
	defer span.Finish()

	// make the Span current in the context
	ctx = opentracing.ContextWithSpan(ctx, span)

	req, err := http.NewRequest("GET", "http://zb.io", nil)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)

	return nil
}

func initLogs() {

}

func parseCLI() config.Config {
	configToReturn := config.Config{}

	flag.IntVar(&configToReturn.Port, "port", 10092, "The port to listen on")

	flag.Parse()

	return configToReturn
}

func startBroker() {

}

func close() {

}
