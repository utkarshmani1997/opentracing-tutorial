package main

import (
	"context"
	"fmt"
	"os"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	tracing "github.com/yurishkuro/opentracing-tutorial/go/lib/tracing"
)

func main() {

	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}

	tracer, closer := tracing.Init("hello-world")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)
	span := tracer.StartSpan("say-hello")
	defer span.Finish()

	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)

	helloTo := os.Args[1]
	helloStr := foramatString(ctx, helloTo)

	printHello(ctx, helloStr)
}

func foramatString(ctx context.Context, helloTo string) string {
	span, _ := opentracing.StartSpanFromContext(ctx, "formatString")
	helloStr := fmt.Sprintf("Hello, %s!", helloTo)
	defer span.Finish()
	// Using Tags
	//	span.SetTag("hello-To", helloTo)

	// Using Logs
	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
	)
	return helloStr
}

func printHello(ctx context.Context, helloStr string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "printHello")
	defer span.Finish()

	fmt.Println(helloStr)
	span.LogKV("event", "println")
}
