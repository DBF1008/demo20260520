package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	commonService "github.com/tiger1103/gfast/v3/internal/app/common/service"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

var DictType = &SysDictTypeController{}

type SysDictTypeController struct {
}


func (c *SysDictTypeController) List(ctx context.Context, req *system.DictTypeSearchReq) (res *system.DictTypeSearchRes, err error) {
	res, err = commonService.SysDictType().List(ctx, req)
	return
}


func (c *SysDictTypeController) Add(ctx context.Context, req *system.DictTypeAddReq) (res *system.DictTypeAddRes, err error) {
	err = commonService.SysDictType().Add(ctx, req, service.Context().GetUserId(ctx))
	return
}


func (c *SysDictTypeController) Get(ctx context.Context, req *system.DictTypeGetReq) (res *system.DictTypeGetRes, err error) {
	res = new(system.DictTypeGetRes)
	res.DictType, err = commonService.SysDictType().Get(ctx, req)
	return
}


func (c *SysDictTypeController) Edit(ctx context.Context, req *system.DictTypeEditReq) (res *system.DictTypeEditRes, err error) {
	err = commonService.SysDictType().Edit(ctx, req, service.Context().GetUserId(ctx))
	return
}

func (c *SysDictTypeController) Delete(ctx context.Context, req *system.DictTypeDeleteReq) (res *system.DictTypeDeleteRes, err error) {
	err = commonService.SysDictType().Delete(ctx, req.DictIds)
	return
}


func (c *SysDictTypeController) OptionSelect(ctx context.Context, req *system.DictTypeAllReq) (res *system.DictTYpeAllRes, err error) {
	res = new(system.DictTYpeAllRes)
	res.DictType, err = commonService.SysDictType().GetAllDictType(ctx)
	return
}
