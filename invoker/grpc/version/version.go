// This code is genetated by protoc-gen-gokit-micro.

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

// PLUGIN: protoc-gen-gokit-micro
//		VERSION : v0.0.1-alpha
//		GIT TAG : v0.0.1-alpha
//		GIT HASH : 3db443bb4aaf6b40ed5408edfa9db4a1cfa3a014
//		BUILDER VERSION : go version go1.16.5 linux/amd64
//		BUILD TIME: : 2022-08-11 17:40:17

// create time : 2022-08-11 18:01:35.615827369 +0800 CST m=+0.009031601

package version

import (
	common "/invoker/grpc/common"
	version1 "/invoker/protocol/version"
	version "/pb/golang/pkg/version"
	context "context"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	reflect "reflect"
)

//// THIS IS FOR CLIENT INVKE SERVICE REQUEST ENCODE
//// SERVICE.REQUEST - > PB.REQUEST
func EncodeVersionServiceGetRequest(ctx context.Context, request interface{}) (interface{}, error) {
	return &version.VersionRequest{}, nil
}

////

//// THIS IS FOR CLIENT DECODE RESPONSE FROM SERVICE
//// PB.RESPONSE -> SERVICE.RESPONSE
func DecodeVersionServiceGetResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	if response, ok := resp.(*version.VersionResponse); ok {
		r := version1.VersionResponse{}
		/// TODO
		r.Hash = response.Hash
		r.Tag = response.Tag
		r.Version = response.Version
		r.BuilderVersion = response.BuilderVersion
		r.Uptime = response.Uptime
		r.RunningTimes = response.RunningTimes
		r.BuildTime = response.BuildTime
		r.HostName = response.HostName
		r.Platform = response.Platform
		return &r, nil
	}
	return nil, fmt.Errorf("Error convert interface : %v -> *%v", reflect.TypeOf(resp), "version.VersionResponse")
}

////

func MakeVersionServiceGetClientEndpoint(serverName string) endpoint.Endpoint {
	var reply version.VersionResponse
	enc := EncodeVersionServiceGetRequest
	dec := DecodeVersionServiceGetResponse
	serviceName := "version.VersionService"
	methodName := "Get"
	return common.MakeClientEndpoint(serverName, serviceName, methodName, enc, dec, &reply)
}
