package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

var Dept = sysDeptController{}

type sysDeptController struct {
	BaseController
}


func (c *sysDeptController) List(ctx context.Context, req *system.DeptSearchReq) (res *system.DeptSearchRes, err error) {
	res = new(system.DeptSearchRes)
	res.DeptList, err = service.SysDept().GetList(ctx, req)
	return
}


func (c *sysDeptController) Add(ctx context.Context, req *system.DeptAddReq) (res *system.DeptAddRes, err error) {
	err = service.SysDept().Add(ctx, req)
	return
}


func (c *sysDeptController) Edit(ctx context.Context, req *system.DeptEditReq) (res *system.DeptEditRes, err error) {
	err = service.SysDept().Edit(ctx, req)
	return
}


func (c *sysDeptController) Delete(ctx context.Context, req *system.DeptDeleteReq) (res *system.DeptDeleteRes, err error) {
	err = service.SysDept().Delete(ctx, req.Id)
	return
}


func (c *sysDeptController) TreeSelect(ctx context.Context, req *system.DeptTreeSelectReq) (res *system.DeptTreeSelectRes, err error) {
	var deptList []*entity.SysDept
	deptList, err = service.SysDept().GetList(ctx, &system.DeptSearchReq{
		Status: "1",
	})
	if err != nil {
		return
	}
	res = new(system.DeptTreeSelectRes)
	res.Deps = service.SysDept().GetListTree(0, deptList)
	return
}
