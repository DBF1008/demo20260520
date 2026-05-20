package sysAuthRule

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	commonService "github.com/tiger1103/gfast/v3/internal/app/common/service"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/internal/app/system/dao"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterSysAuthRule(New())
}

func New() *sSysAuthRule {
	return &sSysAuthRule{}
}

type sSysAuthRule struct {
}

func (s *sSysAuthRule) GetMenuListSearch(ctx context.Context, req *system.RuleSearchReq) (res []*model.SysAuthRuleInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysAuthRule.Ctx(ctx)
		if req.Title != "" {
			m = m.Where("title like ?", "%"+req.Title+"%")
		}
		if req.Component != "" {
			m = m.Where("component like ?", "%"+req.Component+"%")
		}
		err = m.Fields(model.SysAuthRuleInfoRes{}).Order("weigh desc,id asc").Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取菜单失败")
	})
	return
}


func (s *sSysAuthRule) GetIsMenuList(ctx context.Context) ([]*model.SysAuthRuleInfoRes, error) {
	list, err := s.GetMenuList(ctx)
	if err != nil {
		return nil, err
	}
	var gList = make([]*model.SysAuthRuleInfoRes, 0, len(list))
	for _, v := range list {
		if v.MenuType == 0 || v.MenuType == 1 {
			gList = append(gList, v)
		}
	}
	return gList, nil
}


func (s *sSysAuthRule) GetMenuList(ctx context.Context) (list []*model.SysAuthRuleInfoRes, err error) {
	cache := commonService.Cache()

	iList := cache.GetOrSetFuncLock(ctx, consts.CacheSysAuthMenu, s.getMenuListFromDb, 0, consts.CacheSysAuthTag)
	if !iList.IsEmpty() {
		err = gconv.Struct(iList, &list)
		liberr.ErrIsNil(ctx, err)
	}
	return
}


func (s *sSysAuthRule) getMenuListFromDb(ctx context.Context) (value interface{}, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var v []*model.SysAuthRuleInfoRes

		err = dao.SysAuthRule.Ctx(ctx).
			Fields(model.SysAuthRuleInfoRes{}).Order("weigh desc,id asc").Scan(&v)
		liberr.ErrIsNil(ctx, err, "获取菜单数据失败")
		value = v
	})
	return
}


func (s *sSysAuthRule) GetIsButtonList(ctx context.Context) ([]*model.SysAuthRuleInfoRes, error) {
	list, err := s.GetMenuList(ctx)
	if err != nil {
		return nil, err
	}
	var gList = make([]*model.SysAuthRuleInfoRes, 0, len(list))
	for _, v := range list {
		if v.MenuType == 2 {
			gList = append(gList, v)
		}
	}
	return gList, nil
}


func (s *sSysAuthRule) Add(ctx context.Context, req *system.RuleAddReq) (err error) {
	if s.menuNameExists(ctx, req.Name, 0) {
		err = gerror.New("接口规则已经存在")
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {

			data := do.SysAuthRule{
				Pid:       req.Pid,
				Name:      req.Name,
				Title:     req.Title,
				Icon:      req.Icon,
				Condition: req.Condition,
				Remark:    req.Remark,
				MenuType:  req.MenuType,
				Weigh:     req.Weigh,
				IsHide:    req.IsHide,
				Path:      req.Path,
				Component: req.Component,
				IsLink:    req.IsLink,
				IsIframe:  req.IsIframe,
				IsCached:  req.IsCached,
				Redirect:  req.Redirect,
				IsAffix:   req.IsAffix,
				LinkUrl:   req.LinkUrl,
			}
			ruleId, e := dao.SysAuthRule.Ctx(ctx).TX(tx).InsertAndGetId(data)
			liberr.ErrIsNil(ctx, e, "添加菜单失败")
			e = s.BindRoleRule(ctx, ruleId, req.Roles)
			liberr.ErrIsNil(ctx, e, "添加菜单失败")
		})
		return err
	})
	if err == nil {

		commonService.Cache().Remove(ctx, consts.CacheSysAuthMenu)
	}
	return
}


func (s *sSysAuthRule) menuNameExists(ctx context.Context, name string, id uint) bool {
	m := dao.SysAuthRule.Ctx(ctx).Where("name=?", name)
	if id != 0 {
		m = m.Where("id!=?", id)
	}
	c, err := m.Fields(dao.SysAuthRule.Columns().Id).Limit(1).One()
	if err != nil {
		g.Log().Error(ctx, err)
		return false
	}
	return !c.IsEmpty()
}


func (s *sSysAuthRule) BindRoleRule(ctx context.Context, ruleId interface{}, roleIds []uint) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := commonService.CasbinEnforcer()
		liberr.ErrIsNil(ctx, e)
		for _, roleId := range roleIds {
			_, err = enforcer.AddPolicy(fmt.Sprintf("%d", roleId), fmt.Sprintf("%d", ruleId), "All")
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

func (s *sSysAuthRule) Get(ctx context.Context, id uint) (rule *entity.SysAuthRule, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysAuthRule.Ctx(ctx).WherePri(id).Scan(&rule)
		liberr.ErrIsNil(ctx, err, "获取菜单失败")
	})
	return
}

func (s *sSysAuthRule) GetMenuRoles(ctx context.Context, id uint) (roleIds []uint, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := commonService.CasbinEnforcer()
		liberr.ErrIsNil(ctx, e)
		policies := enforcer.GetFilteredNamedPolicy("p", 1, gconv.String(id))
		for _, policy := range policies {
			roleIds = append(roleIds, gconv.Uint(policy[0]))
		}
	})
	return
}

func (s *sSysAuthRule) Update(ctx context.Context, req *system.RuleUpdateReq) (err error) {
	if s.menuNameExists(ctx, req.Name, req.Id) {
		err = gerror.New("接口规则已经存在")
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {

			data := do.SysAuthRule{
				Pid:       req.Pid,
				Name:      req.Name,
				Title:     req.Title,
				Icon:      req.Icon,
				Condition: req.Condition,
				Remark:    req.Remark,
				MenuType:  req.MenuType,
				Weigh:     req.Weigh,
				IsHide:    req.IsHide,
				Path:      req.Path,
				Component: req.Component,
				IsLink:    req.IsLink,
				IsIframe:  req.IsIframe,
				IsCached:  req.IsCached,
				Redirect:  req.Redirect,
				IsAffix:   req.IsAffix,
				LinkUrl:   req.LinkUrl,
			}
			_, e := dao.SysAuthRule.Ctx(ctx).TX(tx).WherePri(req.Id).Update(data)
			liberr.ErrIsNil(ctx, e, "添加菜单失败")
			e = s.UpdateRoleRule(ctx, req.Id, req.Roles)
			liberr.ErrIsNil(ctx, e, "添加菜单失败")
		})
		return err
	})
	if err == nil {

		commonService.Cache().Remove(ctx, consts.CacheSysAuthMenu)
	}
	return
}

func (s *sSysAuthRule) UpdateRoleRule(ctx context.Context, ruleId uint, roleIds []uint) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := commonService.CasbinEnforcer()
		liberr.ErrIsNil(ctx, e)

		_, e = enforcer.RemoveFilteredPolicy(1, gconv.String(ruleId))
		liberr.ErrIsNil(ctx, e)

		roleIdsStrArr := gconv.Strings(roleIds)
		for _, v := range roleIdsStrArr {
			_, e = enforcer.AddPolicy(v, gconv.String(ruleId), "All")
			liberr.ErrIsNil(ctx, e)
		}
	})
	return
}

func (s *sSysAuthRule) GetMenuListTree(pid uint, list []*model.SysAuthRuleInfoRes) []*model.SysAuthRuleTreeRes {
	tree := make([]*model.SysAuthRuleTreeRes, 0, len(list))
	for _, menu := range list {
		if menu.Pid == pid {
			t := &model.SysAuthRuleTreeRes{
				SysAuthRuleInfoRes: menu,
			}
			child := s.GetMenuListTree(menu.Id, list)
			if child != nil {
				t.Children = child
			}
			tree = append(tree, t)
		}
	}
	return tree
}


func (s *sSysAuthRule) DeleteMenuByIds(ctx context.Context, ids []int) (err error) {
	var list []*model.SysAuthRuleInfoRes
	list, err = s.GetMenuList(ctx)
	if err != nil {
		return
	}
	childrenIds := make([]int, 0, len(list))
	for _, id := range ids {
		rules := s.FindSonByParentId(list, gconv.Uint(id))
		for _, child := range rules {
			childrenIds = append(childrenIds, gconv.Int(child.Id))
		}
	}
	ids = append(ids, childrenIds...)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		return g.Try(ctx, func(ctx context.Context) {
			_, err = dao.SysAuthRule.Ctx(ctx).Where("id in (?)", ids).Delete()
			liberr.ErrIsNil(ctx, err, "删除失败")

			enforcer, err := commonService.CasbinEnforcer()
			liberr.ErrIsNil(ctx, err)
			for _, v := range ids {
				_, err = enforcer.RemoveFilteredPolicy(1, gconv.String(v))
				liberr.ErrIsNil(ctx, err)
			}

			commonService.Cache().Remove(ctx, consts.CacheSysAuthMenu)
		})
	})
	return
}

func (s *sSysAuthRule) FindSonByParentId(list []*model.SysAuthRuleInfoRes, pid uint) []*model.SysAuthRuleInfoRes {
	children := make([]*model.SysAuthRuleInfoRes, 0, len(list))
	for _, v := range list {
		if v.Pid == pid {
			children = append(children, v)
			fChildren := s.FindSonByParentId(list, v.Id)
			children = append(children, fChildren...)
		}
	}
	return children
}
