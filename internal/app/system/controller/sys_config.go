package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	commonService "github.com/tiger1103/gfast/v3/internal/app/common/service"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

var Config = configController{}

type configController struct {
	BaseController
}


func (c *configController) List(ctx context.Context, req *system.ConfigSearchReq) (res *system.ConfigSearchRes, err error) {
	res, err = commonService.SysConfig().List(ctx, req)
	return
}


func (c *configController) Add(ctx context.Context, req *system.ConfigAddReq) (res *system.ConfigAddRes, err error) {
	err = commonService.SysConfig().Add(ctx, req, service.Context().GetUserId(ctx))
	return
}


func (c *configController) Get(ctx context.Context, req *system.ConfigGetReq) (res *system.ConfigGetRes, err error) {
	res, err = commonService.SysConfig().Get(ctx, req.Id)
	return
}


func (c *configController) Edit(ctx context.Context, req *system.ConfigEditReq) (res *system.ConfigEditRes, err error) {
	err = commonService.SysConfig().Edit(ctx, req, service.Context().GetUserId(ctx))
	return
}


func (c *configController) Delete(ctx context.Context, req *system.ConfigDeleteReq) (res *system.ConfigDeleteRes, err error) {
	err = commonService.SysConfig().Delete(ctx, req.Ids)
	return
}
