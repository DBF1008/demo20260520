package service

import (
	"context"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/tiger1103/gfast/v3/internal/app/common/dao"
	"github.com/tiger1103/gfast/v3/internal/app/common/model/entity"
)


type adapterCasbin struct {
	ctx context.Context
}


var (
	once        sync.Once
	enforcer    *casbin.SyncedEnforcer
	enforcerErr error
)


func CasbinEnforcer() (*casbin.SyncedEnforcer, error) {

	once.Do(func() {
		enforcer, enforcerErr = initEnforcerSingleton()
	})
	return enforcer, enforcerErr
}


func initEnforcerSingleton() (*casbin.SyncedEnforcer, error) {

	adapterCtx := gctx.GetInitCtx()
	a := &adapterCasbin{
		ctx: adapterCtx,
	}


	modelFile := g.Cfg().MustGet(a.ctx, "casbin.modelFile").String()
	e, err := casbin.NewSyncedEnforcer(modelFile, a)
	if err != nil {
		return nil, err
	}

	if !g.Cfg().MustGet(a.ctx, "casbin.cluster").Bool() {
		return e, nil
	}


	w, err := NewGfWatcher("casbin_policy_channel")
	if err != nil {
		g.Log().Error(a.ctx, "Casbin Watcher 初始化失败，集群同步功能不可用:", err)

	} else {

		_ = e.SetWatcher(w)


		_ = w.SetUpdateCallback(func(msg string) {
			g.Log().Info(a.ctx, "Casbin 规则变更，正在重新加载策略...")
			if err := e.LoadPolicy(); err != nil {
				g.Log().Error(a.ctx, "Casbin 重载失败:", err)
			} else {
				g.Log().Info(a.ctx, "Casbin 重载成功")
			}
		})
	}
	return e, nil
}




func (a *adapterCasbin) SavePolicy(model model.Model) (err error) {
	if err = a.dropTable(); err != nil {
		return err
	}

	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			_, err := dao.CasbinRule.Ctx(a.ctx).Data(line).Insert()
			if err != nil {
				return err
			}
		}
	}

	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			_, err := dao.CasbinRule.Ctx(a.ctx).Data(line).Insert()
			if err != nil {
				return err
			}
		}
	}
	return
}

func (a *adapterCasbin) dropTable() (err error) {
	return
}

func (a *adapterCasbin) createTable() (err error) {
	return
}


func (a *adapterCasbin) LoadPolicy(model model.Model) error {
	var lines []*entity.CasbinRule
	if err := dao.CasbinRule.Ctx(a.ctx).Scan(&lines); err != nil {
		return err
	}
	for _, line := range lines {
		loadPolicyLine(line, model)
	}
	return nil
}


func (a *adapterCasbin) AddPolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	_, err := dao.CasbinRule.Ctx(a.ctx).Data(line).Insert()
	return err
}


func (a *adapterCasbin) RemovePolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	err := rawDelete(a, line)
	return err
}

func (a *adapterCasbin) AddPolicies(sec string, ptype string, rules [][]string) error {
	lines := make([]*entity.CasbinRule, len(rules))
	for k, rule := range rules {
		lines[k] = savePolicyLine(ptype, rule)
	}
	_, err := dao.CasbinRule.Ctx(a.ctx).Data(lines).Insert()
	return err
}

func (a *adapterCasbin) RemovePolicies(sec string, ptype string, rules [][]string) error {
	for _, rule := range rules {
		err := a.RemovePolicy(sec, ptype, rule)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *adapterCasbin) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	line := &entity.CasbinRule{}
	line.Ptype = ptype
	if fieldIndex <= 0 && 0 < fieldIndex+len(fieldValues) {
		line.V0 = fieldValues[0-fieldIndex]
	}
	if fieldIndex <= 1 && 1 < fieldIndex+len(fieldValues) {
		line.V1 = fieldValues[1-fieldIndex]
	}
	if fieldIndex <= 2 && 2 < fieldIndex+len(fieldValues) {
		line.V2 = fieldValues[2-fieldIndex]
	}
	if fieldIndex <= 3 && 3 < fieldIndex+len(fieldValues) {
		line.V3 = fieldValues[3-fieldIndex]
	}
	if fieldIndex <= 4 && 4 < fieldIndex+len(fieldValues) {
		line.V4 = fieldValues[4-fieldIndex]
	}
	if fieldIndex <= 5 && 5 < fieldIndex+len(fieldValues) {
		line.V5 = fieldValues[5-fieldIndex]
	}
	err := rawDelete(a, line)
	return err
}


func loadPolicyLine(line *entity.CasbinRule, model model.Model) {
	lineText := line.Ptype
	if line.V0 != "" {
		lineText += ", " + line.V0
	}
	if line.V1 != "" {
		lineText += ", " + line.V1
	}
	if line.V2 != "" {
		lineText += ", " + line.V2
	}
	if line.V3 != "" {
		lineText += ", " + line.V3
	}
	if line.V4 != "" {
		lineText += ", " + line.V4
	}
	if line.V5 != "" {
		lineText += ", " + line.V5
	}
	persist.LoadPolicyLine(lineText, model)
}

func savePolicyLine(ptype string, rule []string) *entity.CasbinRule {
	line := &entity.CasbinRule{}
	line.Ptype = ptype
	if len(rule) > 0 {
		line.V0 = rule[0]
	}
	if len(rule) > 1 {
		line.V1 = rule[1]
	}
	if len(rule) > 2 {
		line.V2 = rule[2]
	}
	if len(rule) > 3 {
		line.V3 = rule[3]
	}
	if len(rule) > 4 {
		line.V4 = rule[4]
	}
	if len(rule) > 5 {
		line.V5 = rule[5]
	}
	return line
}

func rawDelete(a *adapterCasbin, line *entity.CasbinRule) error {
	db := dao.CasbinRule.Ctx(a.ctx).Where("ptype = ?", line.Ptype)
	if line.V0 != "" {
		db = db.Where("v0 = ?", line.V0)
	}
	if line.V1 != "" {
		db = db.Where("v1 = ?", line.V1)
	}
	if line.V2 != "" {
		db = db.Where("v2 = ?", line.V2)
	}
	if line.V3 != "" {
		db = db.Where("v3 = ?", line.V3)
	}
	if line.V4 != "" {
		db = db.Where("v4 = ?", line.V4)
	}
	if line.V5 != "" {
		db = db.Where("v5 = ?", line.V5)
	}
	_, err := db.Delete()
	return err
}
