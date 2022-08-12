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
//		VERSION : v0.0.1-alpha-1-gb9cf26b
//		GIT TAG : v0.0.1-alpha-1-gb9cf26b
//		GIT HASH : b9cf26b42671c569051d633fd8c790a87dd2caaf
//		BUILDER VERSION : go version go1.16.5 linux/amd64
//		BUILD TIME: : 2022-08-12 10:13:03

// create time : 2022-08-12 10:13:09.874656916 +0800 CST m=+0.009850083

package calculate

import (
	context "context"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	consul "github.com/go-kit/kit/sd/consul"
	grpc "github.com/go-kit/kit/transport/grpc"
	grpc1 "github.com/leewckk/go-kit-micro-service/middlewares/endpoint/grpc"
	calculate "micro-client-sample/invoker/protocol/calculate"
	calculate1 "micro-client-sample/pb/golang/pkg/calculate"
	reflect "reflect"
)

//// THIS IS FOR CLIENT INVKE SERVICE REQUEST ENCODE
//// SERVICE.REQUEST - > PB.REQUEST
func EncodeCalculateAddRequest(ctx context.Context, request interface{}) (interface{}, error) {
	if r, ok := request.(*calculate.AddRequest); ok {
		req := calculate1.AddRequest{}
		/// TODO
		req.Element0 = r.Element0
		req.Element1 = r.Element1
		if nil != r.Msg {
			var s0 calculate1.CalMessage
			req.Msg = &s0
			req.Msg.Evalint = r.Msg.Evalint
			req.Msg.Evalfloat = r.Msg.Evalfloat
			req.Msg.Evaldouble = r.Msg.Evaldouble
			req.Msg.Evalbool = r.Msg.Evalbool
			req.Msg.Evalsint32 = r.Msg.Evalsint32
			req.Msg.Evalunit32 = r.Msg.Evalunit32
			req.Msg.Evalint64 = r.Msg.Evalint64
			req.Msg.Evalsint64 = r.Msg.Evalsint64
			req.Msg.Evaluint64 = r.Msg.Evaluint64
			req.Msg.Evalsfix32 = r.Msg.Evalsfix32
			req.Msg.Evalfix32 = r.Msg.Evalfix32
			req.Msg.Evalfixed64 = r.Msg.Evalfixed64
			req.Msg.Evalstring = r.Msg.Evalstring
			req.Msg.Evalbytes = r.Msg.Evalbytes
			req.Msg.Msg = r.Msg.Msg
		}
		for _, sub := range r.Externals {
			var s0 calculate1.ExternalInfo
			s0.Evalint = sub.Evalint
			s0.Evalfloat = sub.Evalfloat
			s0.Evaldouble = sub.Evaldouble
			s0.Evalbool = sub.Evalbool
			s0.Evalsint32 = sub.Evalsint32
			s0.Evalunit32 = sub.Evalunit32
			s0.Evalint64 = sub.Evalint64
			s0.Evalsint64 = sub.Evalsint64
			s0.Evaluint64 = sub.Evaluint64
			s0.Evalsfix32 = sub.Evalsfix32
			s0.Evalfix32 = sub.Evalfix32
			s0.Evalfixed64 = sub.Evalfixed64
			s0.Evalstring = sub.Evalstring
			s0.Evalbytes = sub.Evalbytes
			s0.CallMsgs = make(map[string]*calculate1.CalMessage)
			for key, sub := range sub.CallMsgs {
				var s1 calculate1.CalMessage
				s1.Evalint = sub.Evalint
				s1.Evalfloat = sub.Evalfloat
				s1.Evaldouble = sub.Evaldouble
				s1.Evalbool = sub.Evalbool
				s1.Evalsint32 = sub.Evalsint32
				s1.Evalunit32 = sub.Evalunit32
				s1.Evalint64 = sub.Evalint64
				s1.Evalsint64 = sub.Evalsint64
				s1.Evaluint64 = sub.Evaluint64
				s1.Evalsfix32 = sub.Evalsfix32
				s1.Evalfix32 = sub.Evalfix32
				s1.Evalfixed64 = sub.Evalfixed64
				s1.Evalstring = sub.Evalstring
				s1.Evalbytes = sub.Evalbytes
				s1.Msg = sub.Msg
				s0.CallMsgs[key] = &s1
			}
			req.Externals = append(req.Externals, &s0)
		}
		req.Datelist = r.Datelist
		req.Emapstr = make(map[string]*calculate1.ExternalInfo)
		for key, sub := range r.Emapstr {
			var s0 calculate1.ExternalInfo
			s0.Evalint = sub.Evalint
			s0.Evalfloat = sub.Evalfloat
			s0.Evaldouble = sub.Evaldouble
			s0.Evalbool = sub.Evalbool
			s0.Evalsint32 = sub.Evalsint32
			s0.Evalunit32 = sub.Evalunit32
			s0.Evalint64 = sub.Evalint64
			s0.Evalsint64 = sub.Evalsint64
			s0.Evaluint64 = sub.Evaluint64
			s0.Evalsfix32 = sub.Evalsfix32
			s0.Evalfix32 = sub.Evalfix32
			s0.Evalfixed64 = sub.Evalfixed64
			s0.Evalstring = sub.Evalstring
			s0.Evalbytes = sub.Evalbytes
			s0.CallMsgs = make(map[string]*calculate1.CalMessage)
			for key, sub := range sub.CallMsgs {
				var s1 calculate1.CalMessage
				s1.Evalint = sub.Evalint
				s1.Evalfloat = sub.Evalfloat
				s1.Evaldouble = sub.Evaldouble
				s1.Evalbool = sub.Evalbool
				s1.Evalsint32 = sub.Evalsint32
				s1.Evalunit32 = sub.Evalunit32
				s1.Evalint64 = sub.Evalint64
				s1.Evalsint64 = sub.Evalsint64
				s1.Evaluint64 = sub.Evaluint64
				s1.Evalsfix32 = sub.Evalsfix32
				s1.Evalfix32 = sub.Evalfix32
				s1.Evalfixed64 = sub.Evalfixed64
				s1.Evalstring = sub.Evalstring
				s1.Evalbytes = sub.Evalbytes
				s1.Msg = sub.Msg
				s0.CallMsgs[key] = &s1
			}
			req.Emapstr[key] = &s0
		}
		req.Emapint = make(map[int32]*calculate1.ExternalInfo)
		for key, sub := range r.Emapint {
			var s0 calculate1.ExternalInfo
			s0.Evalint = sub.Evalint
			s0.Evalfloat = sub.Evalfloat
			s0.Evaldouble = sub.Evaldouble
			s0.Evalbool = sub.Evalbool
			s0.Evalsint32 = sub.Evalsint32
			s0.Evalunit32 = sub.Evalunit32
			s0.Evalint64 = sub.Evalint64
			s0.Evalsint64 = sub.Evalsint64
			s0.Evaluint64 = sub.Evaluint64
			s0.Evalsfix32 = sub.Evalsfix32
			s0.Evalfix32 = sub.Evalfix32
			s0.Evalfixed64 = sub.Evalfixed64
			s0.Evalstring = sub.Evalstring
			s0.Evalbytes = sub.Evalbytes
			s0.CallMsgs = make(map[string]*calculate1.CalMessage)
			for key, sub := range sub.CallMsgs {
				var s1 calculate1.CalMessage
				s1.Evalint = sub.Evalint
				s1.Evalfloat = sub.Evalfloat
				s1.Evaldouble = sub.Evaldouble
				s1.Evalbool = sub.Evalbool
				s1.Evalsint32 = sub.Evalsint32
				s1.Evalunit32 = sub.Evalunit32
				s1.Evalint64 = sub.Evalint64
				s1.Evalsint64 = sub.Evalsint64
				s1.Evaluint64 = sub.Evaluint64
				s1.Evalsfix32 = sub.Evalsfix32
				s1.Evalfix32 = sub.Evalfix32
				s1.Evalfixed64 = sub.Evalfixed64
				s1.Evalstring = sub.Evalstring
				s1.Evalbytes = sub.Evalbytes
				s1.Msg = sub.Msg
				s0.CallMsgs[key] = &s1
			}
			req.Emapint[key] = &s0
		}
		req.Emadouble = make(map[int32]float64)
		for key, sub := range r.Emadouble {
			req.Emadouble[key] = sub
		}
		return &req, nil
	}
	return nil, fmt.Errorf("Error convert interface : %v -> *%v", reflect.TypeOf(request), "calculate.AddRequest")
}

//// THIS IS FOR CLIENT INVKE SERVICE REQUEST ENCODE
//// SERVICE.REQUEST - > PB.REQUEST
func EncodeCalculateAdd2Request(ctx context.Context, request interface{}) (interface{}, error) {
	if r, ok := request.(*calculate.AddRequest2); ok {
		req := calculate1.AddRequest2{}
		/// TODO
		req.Requests = make(map[string]*calculate1.AddRequest)
		for key, sub := range r.Requests {
			var s0 calculate1.AddRequest
			s0.Element0 = sub.Element0
			s0.Element1 = sub.Element1
			if nil != sub.Msg {
				var s1 calculate1.CalMessage
				s0.Msg = &s1
				s0.Msg.Evalint = sub.Msg.Evalint
				s0.Msg.Evalfloat = sub.Msg.Evalfloat
				s0.Msg.Evaldouble = sub.Msg.Evaldouble
				s0.Msg.Evalbool = sub.Msg.Evalbool
				s0.Msg.Evalsint32 = sub.Msg.Evalsint32
				s0.Msg.Evalunit32 = sub.Msg.Evalunit32
				s0.Msg.Evalint64 = sub.Msg.Evalint64
				s0.Msg.Evalsint64 = sub.Msg.Evalsint64
				s0.Msg.Evaluint64 = sub.Msg.Evaluint64
				s0.Msg.Evalsfix32 = sub.Msg.Evalsfix32
				s0.Msg.Evalfix32 = sub.Msg.Evalfix32
				s0.Msg.Evalfixed64 = sub.Msg.Evalfixed64
				s0.Msg.Evalstring = sub.Msg.Evalstring
				s0.Msg.Evalbytes = sub.Msg.Evalbytes
				s0.Msg.Msg = sub.Msg.Msg
			}
			for _, sub := range sub.Externals {
				var s1 calculate1.ExternalInfo
				s1.Evalint = sub.Evalint
				s1.Evalfloat = sub.Evalfloat
				s1.Evaldouble = sub.Evaldouble
				s1.Evalbool = sub.Evalbool
				s1.Evalsint32 = sub.Evalsint32
				s1.Evalunit32 = sub.Evalunit32
				s1.Evalint64 = sub.Evalint64
				s1.Evalsint64 = sub.Evalsint64
				s1.Evaluint64 = sub.Evaluint64
				s1.Evalsfix32 = sub.Evalsfix32
				s1.Evalfix32 = sub.Evalfix32
				s1.Evalfixed64 = sub.Evalfixed64
				s1.Evalstring = sub.Evalstring
				s1.Evalbytes = sub.Evalbytes
				s1.CallMsgs = make(map[string]*calculate1.CalMessage)
				for key, sub := range sub.CallMsgs {
					var s2 calculate1.CalMessage
					s2.Evalint = sub.Evalint
					s2.Evalfloat = sub.Evalfloat
					s2.Evaldouble = sub.Evaldouble
					s2.Evalbool = sub.Evalbool
					s2.Evalsint32 = sub.Evalsint32
					s2.Evalunit32 = sub.Evalunit32
					s2.Evalint64 = sub.Evalint64
					s2.Evalsint64 = sub.Evalsint64
					s2.Evaluint64 = sub.Evaluint64
					s2.Evalsfix32 = sub.Evalsfix32
					s2.Evalfix32 = sub.Evalfix32
					s2.Evalfixed64 = sub.Evalfixed64
					s2.Evalstring = sub.Evalstring
					s2.Evalbytes = sub.Evalbytes
					s2.Msg = sub.Msg
					s1.CallMsgs[key] = &s2
				}
				s0.Externals = append(s0.Externals, &s1)
			}
			s0.Datelist = sub.Datelist
			s0.Emapstr = make(map[string]*calculate1.ExternalInfo)
			for key, sub := range sub.Emapstr {
				var s1 calculate1.ExternalInfo
				s1.Evalint = sub.Evalint
				s1.Evalfloat = sub.Evalfloat
				s1.Evaldouble = sub.Evaldouble
				s1.Evalbool = sub.Evalbool
				s1.Evalsint32 = sub.Evalsint32
				s1.Evalunit32 = sub.Evalunit32
				s1.Evalint64 = sub.Evalint64
				s1.Evalsint64 = sub.Evalsint64
				s1.Evaluint64 = sub.Evaluint64
				s1.Evalsfix32 = sub.Evalsfix32
				s1.Evalfix32 = sub.Evalfix32
				s1.Evalfixed64 = sub.Evalfixed64
				s1.Evalstring = sub.Evalstring
				s1.Evalbytes = sub.Evalbytes
				s1.CallMsgs = make(map[string]*calculate1.CalMessage)
				for key, sub := range sub.CallMsgs {
					var s2 calculate1.CalMessage
					s2.Evalint = sub.Evalint
					s2.Evalfloat = sub.Evalfloat
					s2.Evaldouble = sub.Evaldouble
					s2.Evalbool = sub.Evalbool
					s2.Evalsint32 = sub.Evalsint32
					s2.Evalunit32 = sub.Evalunit32
					s2.Evalint64 = sub.Evalint64
					s2.Evalsint64 = sub.Evalsint64
					s2.Evaluint64 = sub.Evaluint64
					s2.Evalsfix32 = sub.Evalsfix32
					s2.Evalfix32 = sub.Evalfix32
					s2.Evalfixed64 = sub.Evalfixed64
					s2.Evalstring = sub.Evalstring
					s2.Evalbytes = sub.Evalbytes
					s2.Msg = sub.Msg
					s1.CallMsgs[key] = &s2
				}
				s0.Emapstr[key] = &s1
			}
			s0.Emapint = make(map[int32]*calculate1.ExternalInfo)
			for key, sub := range sub.Emapint {
				var s1 calculate1.ExternalInfo
				s1.Evalint = sub.Evalint
				s1.Evalfloat = sub.Evalfloat
				s1.Evaldouble = sub.Evaldouble
				s1.Evalbool = sub.Evalbool
				s1.Evalsint32 = sub.Evalsint32
				s1.Evalunit32 = sub.Evalunit32
				s1.Evalint64 = sub.Evalint64
				s1.Evalsint64 = sub.Evalsint64
				s1.Evaluint64 = sub.Evaluint64
				s1.Evalsfix32 = sub.Evalsfix32
				s1.Evalfix32 = sub.Evalfix32
				s1.Evalfixed64 = sub.Evalfixed64
				s1.Evalstring = sub.Evalstring
				s1.Evalbytes = sub.Evalbytes
				s1.CallMsgs = make(map[string]*calculate1.CalMessage)
				for key, sub := range sub.CallMsgs {
					var s2 calculate1.CalMessage
					s2.Evalint = sub.Evalint
					s2.Evalfloat = sub.Evalfloat
					s2.Evaldouble = sub.Evaldouble
					s2.Evalbool = sub.Evalbool
					s2.Evalsint32 = sub.Evalsint32
					s2.Evalunit32 = sub.Evalunit32
					s2.Evalint64 = sub.Evalint64
					s2.Evalsint64 = sub.Evalsint64
					s2.Evaluint64 = sub.Evaluint64
					s2.Evalsfix32 = sub.Evalsfix32
					s2.Evalfix32 = sub.Evalfix32
					s2.Evalfixed64 = sub.Evalfixed64
					s2.Evalstring = sub.Evalstring
					s2.Evalbytes = sub.Evalbytes
					s2.Msg = sub.Msg
					s1.CallMsgs[key] = &s2
				}
				s0.Emapint[key] = &s1
			}
			s0.Emadouble = make(map[int32]float64)
			for key, sub := range sub.Emadouble {
				s0.Emadouble[key] = sub
			}
			req.Requests[key] = &s0
		}
		return &req, nil
	}
	return nil, fmt.Errorf("Error convert interface : %v -> *%v", reflect.TypeOf(request), "calculate.AddRequest2")
}

////

//// THIS IS FOR CLIENT DECODE RESPONSE FROM SERVICE
//// PB.RESPONSE -> SERVICE.RESPONSE
func DecodeCalculateAddResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	if response, ok := resp.(*calculate1.CalculateResult); ok {
		r := calculate.CalculateResult{}
		/// TODO
		r.Status = response.Status
		r.Msg = response.Msg
		r.Result = response.Result
		return &r, nil
	}
	return nil, fmt.Errorf("Error convert interface : %v -> *%v", reflect.TypeOf(resp), "calculate1.CalculateResult")
}

//// THIS IS FOR CLIENT DECODE RESPONSE FROM SERVICE
//// PB.RESPONSE -> SERVICE.RESPONSE
func DecodeCalculateAdd2Response(ctx context.Context, resp interface{}) (interface{}, error) {
	if response, ok := resp.(*calculate1.CalculateResult); ok {
		r := calculate.CalculateResult{}
		/// TODO
		r.Status = response.Status
		r.Msg = response.Msg
		r.Result = response.Result
		return &r, nil
	}
	return nil, fmt.Errorf("Error convert interface : %v -> *%v", reflect.TypeOf(resp), "calculate1.CalculateResult")
}

////

//// THIS IS FOR CLIENT INVKE SERVICE REQUEST ENCODE
//// SERVICE.REQUEST - > PB.REQUEST
func EncodeMessagingGetMessageRequest(ctx context.Context, request interface{}) (interface{}, error) {
	if r, ok := request.(*calculate.GetMessageRequest); ok {
		req := calculate1.GetMessageRequest{}
		/// TODO
		req.MessageId = r.MessageId
		req.UserId = r.UserId
		return &req, nil
	}
	return nil, fmt.Errorf("Error convert interface : %v -> *%v", reflect.TypeOf(request), "calculate.GetMessageRequest")
}

////

//// THIS IS FOR CLIENT DECODE RESPONSE FROM SERVICE
//// PB.RESPONSE -> SERVICE.RESPONSE
func DecodeMessagingGetMessageResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	if response, ok := resp.(*calculate1.GetMessageResponse); ok {
		r := calculate.GetMessageResponse{}
		/// TODO
		r.MessageId = response.MessageId
		r.UserId = response.UserId
		r.Text = response.Text
		return &r, nil
	}
	return nil, fmt.Errorf("Error convert interface : %v -> *%v", reflect.TypeOf(resp), "calculate1.GetMessageResponse")
}

////

func MakeCalculateAddClientEndpoint(serverName string, client consul.Client, opts ...grpc.ClientOption) endpoint.Endpoint {
	var reply calculate1.CalculateResult
	enc := EncodeCalculateAddRequest
	dec := DecodeCalculateAddResponse
	serviceName := "calculate.Calculate"
	methodName := "Add"
	return grpc1.MakeClientEndpoint(client, serverName, serviceName, methodName, enc, dec, &reply, opts...)
}
func MakeCalculateAdd2ClientEndpoint(serverName string, client consul.Client, opts ...grpc.ClientOption) endpoint.Endpoint {
	var reply calculate1.CalculateResult
	enc := EncodeCalculateAdd2Request
	dec := DecodeCalculateAdd2Response
	serviceName := "calculate.Calculate"
	methodName := "Add2"
	return grpc1.MakeClientEndpoint(client, serverName, serviceName, methodName, enc, dec, &reply, opts...)
}

func MakeMessagingGetMessageClientEndpoint(serverName string, client consul.Client, opts ...grpc.ClientOption) endpoint.Endpoint {
	var reply calculate1.GetMessageResponse
	enc := EncodeMessagingGetMessageRequest
	dec := DecodeMessagingGetMessageResponse
	serviceName := "calculate.Messaging"
	methodName := "GetMessage"
	return grpc1.MakeClientEndpoint(client, serverName, serviceName, methodName, enc, dec, &reply, opts...)
}
