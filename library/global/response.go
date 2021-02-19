package global

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"net/http"
)

const (
	FormatXml   = "xml"
	FormatJson  = "json"
	FormatJsonp = "jsonp"
	SuccessCode = 0
	FailCode    = 1
)

type Response struct {
	Request *ghttp.Request
	Format  string
}

type Payload struct {
	Status  int         `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Data map[string]interface{}

// 新建响应
func NewResponse(request *ghttp.Request, format string) *Response {
	return &Response{
		Request: request,
		Format:  format,
	}
}

// 附带请求
func (l *Response) WithRequest(request *ghttp.Request) *Response {
	l.Request = request

	return l
}

// 附带HEADER头信息
func (l *Response) WithHeader(key string, value string) *Response {
	l.Request.Response.Header().Add(key, value)

	return l
}

// 响应JSON
func (l *Response) Json(status int, code int, message string, data ...interface{}) {
	l.Request.Response.WriteHeader(status)

	err := l.Request.Response.WriteJsonExit(Payload{
		Status:  status,
		Code:    code,
		Message: message,
		Data:    l.handleData(data...),
	})

	l.handleError(err)
}

// 响应JSONP
func (l *Response) Jsonp(status int, code int, message string, data ...interface{}) {
	l.Request.Response.WriteHeader(status)

	err := l.Request.Response.WriteJsonPExit(Payload{
		Status:  status,
		Code:    code,
		Message: message,
		Data:    l.handleData(data...),
	})

	l.handleError(err)
}

// 响应XML
func (l *Response) Xml(status int, code int, message string, data ...interface{}) {
	l.Request.Response.WriteHeader(status)

	err := l.Request.Response.WriteXmlExit(Payload{
		Status:  status,
		Code:    code,
		Message: message,
		Data:    l.handleData(data...),
	})

	l.handleError(err)
}

// 响应数据
func (l *Response) Response(status int, code int, message string, data ...interface{}) {
	switch l.Format {
	case FormatXml:
		l.Xml(status, code, message, data...)
	case FormatJson:
		l.Json(status, code, message, data...)
	case FormatJsonp:
		l.Jsonp(status, code, message, data...)
	default:
		l.Json(status, code, message, data...)
	}
}

// 成功响应
func (l *Response) Success(message string, data ...interface{}) {
	l.Response(http.StatusOK, SuccessCode, message, data...)
}

// 失败响应
func (l *Response) Fail(status int, message string, data ...interface{}) {
	l.Response(status, FailCode, message, data...)
}

// 处理响应数据
func (l *Response) handleData(data ...interface{}) interface{} {
	if data == nil {
		return g.Map{}
	} else {
		return data[0]
	}
}

// 处理错误
func (l *Response) handleError(err error) {
	glog.Error(err.Error())
}
