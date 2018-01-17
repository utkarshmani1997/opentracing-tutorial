package main

import (
	"fmt"
	"os"

	"github.com/opentracing/opentracing-go/log"
	tracing "github.com/yurishkuro/opentracing-tutorial/go/lib/tracing"
)

func main() {

	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}

	tracer, closer := tracing.Init("hello-world")
	defer closer.Close()

	span := tracer.StartSpan("say-hello")

	helloTo := os.Args[1]
	helloStr := fmt.Sprintf("Hello, %s!", helloTo)

	// Using Tags
	//	span.SetTag("hello-To", helloTo)
	//	println(helloStr)

	// Using Logs
	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
	)

	println(helloStr)
	span.LogKV("event", "println")
	span.Finish()
}
