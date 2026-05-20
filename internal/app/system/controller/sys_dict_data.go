package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	commonService "github.com/tiger1103/gfast/v3/internal/app/common/service"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

var DictData = dictDataController{}

type dictDataController struct {
}


func (c *dictDataController) GetDictData(ctx context.Context, req *system.GetDictReq) (res *system.GetDictRes, err error) {
	res, err = commonService.SysDictData().GetDictWithDataByType(ctx, req.DictType, req.DefaultValue)
	return
}


func (c *dictDataController) List(ctx context.Context, req *system.DictDataSearchReq) (res *system.DictDataSearchRes, err error) {
	res, err = commonService.SysDictData().List(ctx, req)
	return
}


func (c *dictDataController) Add(ctx context.Context, req *system.DictDataAddReq) (res *system.DictDataAddRes, err error) {
	err = commonService.SysDictData().Add(ctx, req, service.Context().GetUserId(ctx))
	return
}


func (c *dictDataController) Get(ctx context.Context, req *system.DictDataGetReq) (res *system.DictDataGetRes, err error) {
	res, err = commonService.SysDictData().Get(ctx, req.DictCode)
	return
}


func (c *dictDataController) Edit(ctx context.Context, req *system.DictDataEditReq) (res *system.DictDataEditRes, err error) {
	err = commonService.SysDictData().Edit(ctx, req, service.Context().GetUserId(ctx))
	return
}

func (c *dictDataController) Delete(ctx context.Context, req *system.DictDataDeleteReq) (res *system.DictDataDeleteRes, err error) {
	err = commonService.SysDictData().Delete(ctx, req.Ids)
	return
}
