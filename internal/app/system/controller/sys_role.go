package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

var Role = roleController{}

type roleController struct {
	BaseController
}


func (c *roleController) List(ctx context.Context, req *system.RoleListReq) (res *system.RoleListRes, err error) {
	res, err = service.SysRole().GetRoleListSearch(ctx, req)
	return
}


func (c *roleController) GetParams(ctx context.Context, req *system.RoleGetParamsReq) (res *system.RoleGetParamsRes, err error) {
	res = new(system.RoleGetParamsRes)
	res.Menu, err = service.SysAuthRule().GetMenuList(ctx)
	return
}


func (c *roleController) Add(ctx context.Context, req *system.RoleAddReq) (res *system.RoleAddRes, err error) {
	err = service.SysRole().AddRole(ctx, req)
	return
}


func (c *roleController) Get(ctx context.Context, req *system.RoleGetReq) (res *system.RoleGetRes, err error) {
	res = new(system.RoleGetRes)
	res.Role, err = service.SysRole().Get(ctx, req.Id)
	if err != nil {
		return
	}
	res.MenuIds, err = service.SysRole().GetFilteredNamedPolicy(ctx, req.Id)
	return
}


func (c *roleController) Edit(ctx context.Context, req *system.RoleEditReq) (res *system.RoleEditRes, err error) {
	err = service.SysRole().EditRole(ctx, req)
	return
}


func (c *roleController) Delete(ctx context.Context, req *system.RoleDeleteReq) (res *system.RoleDeleteRes, err error) {
	err = service.SysRole().DeleteByIds(ctx, req.Ids)
	return
}
