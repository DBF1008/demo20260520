package context

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

func init() {
	service.RegisterContext(New())
}

type sContext struct{}

func New() *sContext {
	return &sContext{}
}


func (s *sContext) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.CtxKey, customCtx)
}


func (s *sContext) Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.CtxKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}


func (s *sContext) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}


func (s *sContext) GetLoginUser(ctx context.Context) *model.ContextUser {
	context := s.Get(ctx)
	if context == nil {
		return nil
	}
	return context.User
}


func (s *sContext) GetUserId(ctx context.Context) uint64 {
	user := s.GetLoginUser(ctx)
	if user != nil {
		return user.Id
	}
	return 0
}
