package main

import (
	"context"
	"encoding/json"
	"fmt"
	"micro-client-sample/invoker/grpc/calculate"
	"micro-client-sample/invoker/grpc/version"
	proto "micro-client-sample/invoker/protocol/calculate"

	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/leewckk/go-kit-micro-service/middlewares/tracing/report"
	"github.com/openzipkin/zipkin-go/reporter"
)

func InvokeGetVersionGRPC(server string, sd consul.Client, tracing reporter.Reporter) {

	opts := make([]grpc.ClientOption, 0, 0)
	opts = append(opts, report.NewGrpcClientTracer(tracing))

	// zkTracer, err0 := opzipkin.NewTracer(tracing)
	// if nil != err0 {
	// 	panic(err0)
	// }

	// opts = append(opts, zipkin.GRPCClientTrace(zkTracer, zipkin.Tags(map[string]string{"name": "liwenchao"})))

	ep := version.MakeVersionServiceGetClientEndpoint(server, sd, opts...)
	resp, err := ep(context.Background(), nil)
	if nil != err {
		panic(err)
	}
	js, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println("resp from GRPC: ", string(js))
}

func InvokeCalculateAddGRPC(server string, sd consul.Client, tracing reporter.Reporter) {
	opts := make([]grpc.ClientOption, 0, 0)
	opts = append(opts, report.NewGrpcClientTracer(tracing))

	ep := calculate.MakeCalculateAddClientEndpoint(server, sd, opts...)

	request := proto.AddRequest{
		Element0: 100,
		Element1: 200,
	}
	resp, err := ep(context.Background(), &request)
	if nil != err {
		panic(err)
	}
	js, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println("resp from GRPC: ", string(js))

}
