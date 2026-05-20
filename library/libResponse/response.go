package libResponse

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	SuccessCode int = 0
	ErrorCode   int = -1
)

type Response struct {

	Code int `json:"code" example:"200"`

	Data interface{} `json:"data"`

	Msg string `json:"message"`
}

var response = new(Response)

func JsonExit(r *ghttp.Request, code int, msg string, data ...interface{}) {
	response.JsonExit(r, code, msg, data...)
}

func RJson(r *ghttp.Request, code int, msg string, data ...interface{}) {
	response.RJson(r, code, msg, data...)
}

func SusJson(isExit bool, r *ghttp.Request, msg string, data ...interface{}) {
	response.SusJson(isExit, r, msg, data...)
}

func FailJson(isExit bool, r *ghttp.Request, msg string, data ...interface{}) {
	response.FailJson(isExit, r, msg, data...)
}

func WriteTpl(r *ghttp.Request, tpl string, view *gview.View, params ...gview.Params) error {
	return response.WriteTpl(r, tpl, view, params...)
}


func (res *Response) JsonExit(r *ghttp.Request, code int, msg string, data ...interface{}) {
	res.RJson(r, code, msg, data...)
	r.ExitAll()
}






func (res *Response) RJson(r *ghttp.Request, code int, msg string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(&Response{
		Code: code,
		Msg:  msg,
		Data: responseData,
	})
}


func (res *Response) SusJson(isExit bool, r *ghttp.Request, msg string, data ...interface{}) {
	if isExit {
		res.JsonExit(r, SuccessCode, msg, data...)
	}
	res.RJson(r, SuccessCode, msg, data...)
}


func (res *Response) FailJson(isExit bool, r *ghttp.Request, msg string, data ...interface{}) {
	if isExit {
		res.JsonExit(r, ErrorCode, msg, data...)
	}
	res.RJson(r, ErrorCode, msg, data...)
}


func (res *Response) WriteTpl(r *ghttp.Request, tpl string, view *gview.View, params ...gview.Params) error {

	view.BindFuncMap(g.Map{

		"subStr": func(str interface{}, i int) (s string) {
			s1 := gconv.String(str)
			if gstr.LenRune(s1) > i {
				s = gstr.SubStrRune(s1, 0, i) + "..."
				return s
			}
			return s1
		},
	})
	r.Response.Write(view.Parse(r.GetCtx(), tpl, params...))
	return nil
}

func (res *Response) Redirect(r *ghttp.Request, location string, code ...int) {
	r.Response.RedirectTo(location, code...)
}
