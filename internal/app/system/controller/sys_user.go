package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

var (
	User = userController{}
)

type userController struct {
	BaseController
}


func (c *userController) GetUserMenus(ctx context.Context, req *system.UserMenusReq) (res *system.UserMenusRes, err error) {
	var (
		permissions []string
		menuList    []*model.UserMenus
	)
	userId := service.Context().GetUserId(ctx)
	menuList, permissions, err = service.SysUser().GetAdminRules(ctx, userId)
	res = &system.UserMenusRes{
		MenuList:    menuList,
		Permissions: permissions,
	}
	return
}


func (c *userController) List(ctx context.Context, req *system.UserSearchReq) (res *system.UserSearchRes, err error) {
	var (
		total    interface{}
		userList []*entity.SysUser
	)
	res = new(system.UserSearchRes)
	total, userList, err = service.SysUser().List(ctx, req)
	if err != nil || total == 0 {
		return
	}
	res.Total = total
	res.UserList, err = service.SysUser().GetUsersRoleDept(ctx, userList)
	return
}


func (c *userController) GetParams(ctx context.Context, req *system.UserGetParamsReq) (res *system.UserGetParamsRes, err error) {
	res = new(system.UserGetParamsRes)
	res.RoleList, err = service.SysRole().GetRoleList(ctx)
	if err != nil {
		return
	}
	res.Posts, err = service.SysPost().GetUsedPost(ctx)
	return
}


func (c *userController) Add(ctx context.Context, req *system.UserAddReq) (res *system.UserAddRes, err error) {
	err = service.SysUser().Add(ctx, req)
	return
}


func (c *userController) GetEditUser(ctx context.Context, req *system.UserGetEditReq) (res *system.UserGetEditRes, err error) {
	res, err = service.SysUser().GetEditUser(ctx, req.Id)
	return
}


func (c *userController) Edit(ctx context.Context, req *system.UserEditReq) (res *system.UserEditRes, err error) {
	err = service.SysUser().Edit(ctx, req)
	return
}


func (c *userController) ResetPwd(ctx context.Context, req *system.UserResetPwdReq) (res *system.UserResetPwdRes, err error) {
	err = service.SysUser().ResetUserPwd(ctx, req)
	return
}


func (c *userController) SetStatus(ctx context.Context, req *system.UserStatusReq) (res *system.UserStatusRes, err error) {
	err = service.SysUser().ChangeUserStatus(ctx, req)
	return
}


func (c *userController) Delete(ctx context.Context, req *system.UserDeleteReq) (res *system.UserDeleteRes, err error) {
	err = service.SysUser().Delete(ctx, req.Ids)
	return
}


func (c *userController) GetUsers(ctx context.Context, req *system.UserGetByIdsReq) (res *system.UserGetByIdsRes, err error) {
	res = new(system.UserGetByIdsRes)
	res.List, err = service.SysUser().GetUsers(ctx, req.Ids)
	return
}
