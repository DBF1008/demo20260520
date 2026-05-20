package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	commonRouter "github.com/tiger1103/gfast/v3/internal/app/common/router"
	commonService "github.com/tiger1103/gfast/v3/internal/app/common/service"
	systemRouter "github.com/tiger1103/gfast/v3/internal/app/system/router"
	"github.com/tiger1103/gfast/v3/library/libRouter"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/api/v1", func(group *ghttp.RouterGroup) {

		group.Middleware(commonService.Middleware().MiddlewareCORS)
		group.Middleware(ghttp.MiddlewareHandlerResponse)

		systemRouter.R.BindController(ctx, group)

		commonRouter.R.BindController(ctx, group)

		if err := libRouter.RouterAutoBind(ctx, router, group); err != nil {
			panic(err)
		}
	})
}
