package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	commonController "github.com/tiger1103/gfast/v3/internal/app/common/controller"
)

type BaseController struct {
	commonController.BaseController
}


func (c *BaseController) Init(r *ghttp.Request) {
	c.BaseController.Init(r)
}
