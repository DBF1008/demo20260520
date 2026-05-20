package mounter

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

type MountHandler func(ctx context.Context, s *ghttp.Server)

var (
	funcOptions = make([]MountHandler, 0)
)

func Mount(handler MountHandler) {
	funcOptions = append(funcOptions, handler)
}

func DoMount(ctx context.Context, s *ghttp.Server) {
	for _, fn := range funcOptions {
		fn(ctx, s)
	}
}
