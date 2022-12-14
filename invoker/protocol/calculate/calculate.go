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

// create time : 2022-08-12 10:13:09.87625416 +0800 CST m=+0.011447342

package calculate

type CalMessage struct {
	Evalint     int32   `json:"evalint" uri:"evalint" form:"evalint"`
	Evalfloat   float32 `json:"evalfloat" uri:"evalfloat" form:"evalfloat"`
	Evaldouble  float64 `json:"evaldouble" uri:"evaldouble" form:"evaldouble"`
	Evalbool    bool    `json:"evalbool" uri:"evalbool" form:"evalbool"`
	Evalsint32  int32   `json:"evalsint32" uri:"evalsint32" form:"evalsint32"`
	Evalunit32  uint32  `json:"evalunit32" uri:"evalunit32" form:"evalunit32"`
	Evalint64   int64   `json:"evalint64" uri:"evalint64" form:"evalint64"`
	Evalsint64  int64   `json:"evalsint64" uri:"evalsint64" form:"evalsint64"`
	Evaluint64  uint64  `json:"evaluint64" uri:"evaluint64" form:"evaluint64"`
	Evalsfix32  int32   `json:"evalsfix32" uri:"evalsfix32" form:"evalsfix32"`
	Evalfix32   uint32  `json:"evalfix32" uri:"evalfix32" form:"evalfix32"`
	Evalfixed64 int64   `json:"evalfixed64" uri:"evalfixed64" form:"evalfixed64"`
	Evalstring  string  `json:"evalstring" uri:"evalstring" form:"evalstring"`
	Evalbytes   []byte  `json:"evalbytes" uri:"evalbytes" form:"evalbytes"`
	Msg         string  `json:"msg" uri:"msg" form:"msg"`
}

type ExternalInfo struct {
	Evalint     int32                  `json:"evalint" uri:"evalint" form:"evalint"`
	Evalfloat   float32                `json:"evalfloat" uri:"evalfloat" form:"evalfloat"`
	Evaldouble  float64                `json:"evaldouble" uri:"evaldouble" form:"evaldouble"`
	Evalbool    bool                   `json:"evalbool" uri:"evalbool" form:"evalbool"`
	Evalsint32  int32                  `json:"evalsint32" uri:"evalsint32" form:"evalsint32"`
	Evalunit32  uint32                 `json:"evalunit32" uri:"evalunit32" form:"evalunit32"`
	Evalint64   int64                  `json:"evalint64" uri:"evalint64" form:"evalint64"`
	Evalsint64  int64                  `json:"evalsint64" uri:"evalsint64" form:"evalsint64"`
	Evaluint64  uint64                 `json:"evaluint64" uri:"evaluint64" form:"evaluint64"`
	Evalsfix32  int32                  `json:"evalsfix32" uri:"evalsfix32" form:"evalsfix32"`
	Evalfix32   uint32                 `json:"evalfix32" uri:"evalfix32" form:"evalfix32"`
	Evalfixed64 int64                  `json:"evalfixed64" uri:"evalfixed64" form:"evalfixed64"`
	Evalstring  string                 `json:"evalstring" uri:"evalstring" form:"evalstring"`
	Evalbytes   []byte                 `json:"evalbytes" uri:"evalbytes" form:"evalbytes"`
	CallMsgs    map[string]*CalMessage `json:"callMsgs,omitempty" uri:"callMsgs" form:"callMsgs"`
}

type AddRequest struct {
	Element0  float32                  `json:"element0" uri:"element0" form:"element0"`              /// ??????1
	Element1  float32                  `json:"element1" uri:"element1" form:"element1"`              /// ??????2
	Msg       *CalMessage              `json:"msg,omitempty" uri:"msg" form:"msg"`                   /// ??????
	Externals []*ExternalInfo          `json:"Externals,omitempty" uri:"Externals" form:"Externals"` /// ????????????
	Datelist  []float32                `json:"datelist,omitempty" uri:"datelist" form:"datelist"`    /// ????????????
	Emapstr   map[string]*ExternalInfo `json:"emapstr,omitempty" uri:"emapstr" form:"emapstr"`       /// ????????????1
	Emapint   map[int32]*ExternalInfo  `json:"emapint,omitempty" uri:"emapint" form:"emapint"`       /// ????????????2
	Emadouble map[int32]float64        `json:"emadouble,omitempty" uri:"emadouble" form:"emadouble"` /// ????????????3
}

type AddRequest2 struct {
	Requests map[string]*AddRequest `json:"requests,omitempty" uri:"requests" form:"requests"`
}

type CalculateResult struct {
	Status int32   `json:"status" uri:"status" form:"status"`
	Msg    string  `json:"msg" uri:"msg" form:"msg"`
	Result float32 `json:"result" uri:"result" form:"result"`
}

type GetMessageRequest struct {
	MessageId string `json:"messageId" uri:"messageId" form:"messageId"` /// message id
	UserId    string `json:"userId" uri:"userId" form:"userId"`          /// user id
}

type GetMessageResponse struct {
	MessageId string `json:"messageId" uri:"messageId" form:"messageId"` /// ???????????????Id
	UserId    string `json:"userId" uri:"userId" form:"userId"`          /// ?????????userId
	Text      string `json:"text" uri:"text" form:"text"`                /// ????????????
}
