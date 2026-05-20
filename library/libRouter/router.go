package libRouter

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gregex"
	"reflect"
)



func RouterAutoBindBefore(ctx context.Context, R interface{}, group *ghttp.RouterGroup) (err error) {
	return bind(ctx, R, group, "before")
}



func RouterAutoBind(ctx context.Context, R interface{}, group *ghttp.RouterGroup) (err error) {
	return bind(ctx, R, group)
}

func bind(ctx context.Context, R interface{}, group *ghttp.RouterGroup, option ...string) (err error) {
	var rule string
	if len(option) > 0 && option[0] == "before" {
		rule = `^BeforeBind(.+)Controller$`
	} else {
		rule = `^Bind(.+)Controller$`
	}

	typ := reflect.TypeOf(R)

	val := reflect.ValueOf(R)
	if val.Elem().Kind() != reflect.Struct {
		err = gerror.New("expect struct but a " + val.Elem().Kind().String())
		return
	}
	for i := 0; i < typ.NumMethod(); i++ {
		if match := gregex.IsMatchString(rule, typ.Method(i).Name); match {

			val.Method(i).Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(group)})
		}
	}
	return
}
