// The MIT License (MIT)

// Copyright (c) 2022 leewckk@126.com

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/leewckk/go-kit-micro-service/discovery"
	"github.com/leewckk/go-kit-micro-service/middlewares/tracing/report"
)

func init() {
}

func main() {
	server := flag.String("server", "micro-service", "invoke server name")
	center := flag.String("center", "127.0.0.1:8500", "service discovery center API")
	tracing := flag.String("tracing", "http://localhost:9411/api/v2/spans", "service API tracing")
	flag.Parse()

	if *server == "" || *center == "" || *tracing == "" {
		flag.Usage()
		return
	}
	sd := discovery.NewClient(*center)
	fmt.Println(FullVersion())

	if "" != *tracing {
		reporter := report.NewZipkinReporter(*tracing)
		InvokeGetVersionGRPC(*server, sd, reporter)
		InvokeGetVersionHTTP(*server, sd, reporter)
		InvokeCalculateAddGRPC(*server, sd, reporter)
		InvokeCalculateAddHTTP(*server, sd, reporter)
	}

	select {
	case <-time.After(3 * time.Second):
		break
	}
	fmt.Println("exit 0")
}
