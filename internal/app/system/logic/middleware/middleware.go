package middleware

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	commonService "github.com/tiger1103/gfast/v3/internal/app/common/service"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/libResponse"
)

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}

type sMiddleware struct{}


func (s *sMiddleware) Ctx(r *ghttp.Request) {
	ctx := r.GetCtx()

	data, err := service.GfToken().ParseToken(r)
	if err != nil {
		r.Middleware.Next()
	}
	if data != nil {
		context := new(model.Context)
		err = gconv.Struct(data.Data, &context.User)
		if err != nil {
			g.Log().Error(ctx, err)
			r.Middleware.Next()
		}
		service.Context().Init(r, context)
	}

	r.Middleware.Next()
}


func (s *sMiddleware) Auth(r *ghttp.Request) {
	ctx := r.GetCtx()

	adminId := service.Context().GetUserId(ctx)
	accessParams := r.Get("accessParams").Strings()
	accessParamsStr := ""
	if len(accessParams) > 0 && accessParams[0] != "undefined" {
		accessParamsStr = "?" + gstr.Join(accessParams, "&")
	}
	url := gstr.TrimLeft(r.Request.URL.Path, "/") + accessParamsStr


	tagSuperAdmin := false
	service.SysUser().NotCheckAuthAdminIds(ctx).Iterator(func(v interface{}) bool {
		if gconv.Uint64(v) == adminId {
			tagSuperAdmin = true
			return false
		}
		return true
	})
	if tagSuperAdmin {
		r.Middleware.Next()

		return
	}

	menuList, err := service.SysAuthRule().GetMenuList(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
		libResponse.FailJson(true, r, "请求数据失败")
	}
	var menu *model.SysAuthRuleInfoRes
	for _, m := range menuList {
		ms := gstr.SubStr(m.Name, 0, gstr.Pos(m.Name, "?"))
		if m.Name == url || ms == url {
			menu = m
			break
		}
	}

	if menu != nil {

		excludePaths := g.Cfg().MustGet(ctx, "gfToken.excludePaths").Strings()
		for _, p := range excludePaths {
			if gstr.Equal(menu.Name, gstr.TrimLeft(p, "/")) {
				r.Middleware.Next()
				return
			}
		}

		if gstr.Equal(menu.Condition, "nocheck") {
			r.Middleware.Next()
			return
		}
		menuId := menu.Id

		if menuId != 0 {

			enforcer, err := commonService.CasbinEnforcer()
			if err != nil {
				g.Log().Error(ctx, err)
				libResponse.FailJson(true, r, "获取权限失败")
			}
			hasAccess := false
			hasAccess, err = enforcer.Enforce(fmt.Sprintf("%s%d", service.SysUser().GetCasBinUserPrefix(), adminId), gconv.String(menuId), "All")
			if err != nil {
				g.Log().Error(ctx, err)
				libResponse.FailJson(true, r, "判断权限失败")
			}
			if !hasAccess {
				libResponse.FailJson(true, r, "没有访问权限")
			}
		}
	} else if menu == nil && accessParamsStr != "" {
		libResponse.FailJson(true, r, "没有访问权限")
	}
	r.Middleware.Next()
}
