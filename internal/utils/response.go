package utils

import (
	"encoding/json"
	"fmt"
	"gapi/pkg/conf"
	"gapi/pkg/consts"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (result *Result) String() string {
	by, _ := json.Marshal(result)
	return fmt.Sprintf("%#v", string(by))
}

func NewResponse(context *gin.Context) *Response {
	return &Response{context: context, result: &Result{Code: 0, Msg: "ok", Data: nil}}
}

type Response struct {
	result  *Result
	context *gin.Context
}

func (res *Response) WithCode(code int) *Response {
	res.result.Code = code
	return res
}

func (res *Response) WithMsg(msg interface{}) *Response {
	r := ""
	switch m := msg.(type) {
	case error:
		r = m.Error()
	case string:
		r = m
	}
	res.result.Msg = r
	return res
}

func (res *Response) WithData(data interface{}) *Response {
	res.result.Data = data
	return res
}

func (res *Response) debugPrint() {
	if conf.GetConfig().System.Env == consts.DebugMode {
		log.Println(res.result)
	}
}

func (res *Response) JsonOutput() {
	res.debugPrint()
	res.context.JSON(http.StatusOK, res.result)
}
