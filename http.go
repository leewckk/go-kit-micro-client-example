package main

import (
	"context"
	"encoding/json"
	"fmt"
	"micro-client-sample/invoker/http/calculate"
	"micro-client-sample/invoker/http/version"
	proto "micro-client-sample/invoker/protocol/calculate"

	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/transport/http"
	"github.com/leewckk/go-kit-micro-service/middlewares/tracing/report"
	"github.com/openzipkin/zipkin-go/reporter"
)

func InvokeGetVersionHTTP(server string, sd consul.Client, tracing reporter.Reporter) {

	opts := make([]http.ClientOption, 0, 0)
	opts = append(opts, report.NewHttpClientTracer(tracing))

	ep := version.MakeVersionServiceGetClientEndpoint(server, sd, opts...)
	resp, err := ep(context.Background(), nil)
	if nil != err {
		panic(err)
	}
	js, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println("resp from HTTP: ", string(js))
}

func InvokeCalculateAddHTTP(server string, sd consul.Client, tracing reporter.Reporter) {
	opts := make([]http.ClientOption, 0, 0)
	opts = append(opts, report.NewHttpClientTracer(tracing))

	ep := calculate.MakeCalculateAddClientEndpoint(server, sd, opts...)

	request := proto.AddRequest{
		Element0: 123,
		Element1: 556,
	}
	resp, err := ep(context.Background(), &request)
	if nil != err {
		panic(err)
	}
	js, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println("resp from HTTP: ", string(js))

}
