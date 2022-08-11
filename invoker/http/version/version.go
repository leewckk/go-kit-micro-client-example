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

// create time : 2022-08-11 18:01:35.615997963 +0800 CST m=+0.009202181

package version

import (
	common "/invoker/http/common"
	version "/invoker/protocol/version"
	gin "/middlewares/transport/http/gin"
	common1 "/transport/gin/common"
	context "context"
	json "encoding/json"
	endpoint "github.com/go-kit/kit/endpoint"
	http "net/http"
)

//// THIS IS FOR CLIENT INVKE SERVICE REQUEST ENCODE
//// CLIENT.REQUEST - > SERVICE.REQUEST
func EncodeVersionServiceGetRequest(ctx context.Context, req *http.Request, request interface{}) error {
	/// TODO

	pattern := req.URL.Path
	pattern = gin.MarshalPattern(pattern, request)
	req.URL.Path = pattern
	return common1.EncodeJSONRequest(ctx, req, request)
}

////

//// THIS IS FOR CLIENT DECODE RESPONSE FROM SERVICE
//// SERVICE.RESPONSE -> CLIENT.RESPONSE
func DecodeVersionServiceGetResponse(ctx context.Context, resp *http.Response) (interface{}, error) {
	//// TODO
	var r version.VersionResponse

	if err := json.NewDecoder(resp.Body).Decode(&r); nil != err {
		return nil, err
	}
	return &r, nil
}

////

func MakeVersionServiceGetClientEndpoint(serverName string) endpoint.Endpoint {
	enc := EncodeVersionServiceGetRequest
	dec := DecodeVersionServiceGetResponse
	api := "/v1/version"
	httpMethod := "GET"
	return common.MakeClientEndpoint(serverName, httpMethod, api, enc, dec)
}
