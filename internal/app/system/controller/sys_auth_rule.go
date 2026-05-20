package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

var Menu = menuController{}

type menuController struct {
	BaseController
}

func (c *menuController) List(ctx context.Context, req *system.RuleSearchReq) (res *system.RuleListRes, err error) {
	var list []*model.SysAuthRuleInfoRes
	res = &system.RuleListRes{
		Rules: make([]*model.SysAuthRuleTreeRes, 0),
	}
	list, err = service.SysAuthRule().GetMenuListSearch(ctx, req)
	if req.Title != "" || req.Component != "" {
		for _, menu := range list {
			res.Rules = append(res.Rules, &model.SysAuthRuleTreeRes{
				SysAuthRuleInfoRes: menu,
			})
		}
	} else {
		res.Rules = service.SysAuthRule().GetMenuListTree(0, list)
	}
	return
}

func (c *menuController) Add(ctx context.Context, req *system.RuleAddReq) (res *system.RuleAddRes, err error) {
	err = service.SysAuthRule().Add(ctx, req)
	return
}


func (c *menuController) GetAddParams(ctx context.Context, req *system.RuleGetParamsReq) (res *system.RuleGetParamsRes, err error) {

	res = new(system.RuleGetParamsRes)
	res.Roles, err = service.SysRole().GetRoleList(ctx)
	if err != nil {
		return
	}
	res.Menus, err = service.SysAuthRule().GetIsMenuList(ctx)
	return
}


func (c *menuController) Get(ctx context.Context, req *system.RuleInfoReq) (res *system.RuleInfoRes, err error) {
	res = new(system.RuleInfoRes)
	res.Rule, err = service.SysAuthRule().Get(ctx, req.Id)
	if err != nil {
		return
	}
	res.RoleIds, err = service.SysAuthRule().GetMenuRoles(ctx, req.Id)
	return
}


func (c *menuController) Update(ctx context.Context, req *system.RuleUpdateReq) (res *system.RuleUpdateRes, err error) {
	err = service.SysAuthRule().Update(ctx, req)
	return
}


func (c *menuController) Delete(ctx context.Context, req *system.RuleDeleteReq) (res *system.RuleDeleteRes, err error) {
	err = service.SysAuthRule().DeleteMenuByIds(ctx, req.Ids)
	return
}
