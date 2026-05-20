package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)


type SysAuthRuleDao struct {
	table   string
	group   string
	columns SysAuthRuleColumns
}


type SysAuthRuleColumns struct {
	Id         string
	Pid        string
	Name       string
	Title      string
	Icon       string
	Condition  string
	Remark     string
	MenuType   string
	Weigh      string
	IsHide     string
	Path       string
	Component  string
	IsLink     string
	ModuleType string
	ModelId    string
	IsIframe   string
	IsCached   string
	Redirect   string
	IsAffix    string
	LinkUrl    string
	CreatedAt  string
	UpdatedAt  string
}


var sysAuthRuleColumns = SysAuthRuleColumns{
	Id:         "id",
	Pid:        "pid",
	Name:       "name",
	Title:      "title",
	Icon:       "icon",
	Condition:  "condition",
	Remark:     "remark",
	MenuType:   "menu_type",
	Weigh:      "weigh",
	IsHide:     "is_hide",
	Path:       "path",
	Component:  "component",
	IsLink:     "is_link",
	ModuleType: "module_type",
	ModelId:    "model_id",
	IsIframe:   "is_iframe",
	IsCached:   "is_cached",
	Redirect:   "redirect",
	IsAffix:    "is_affix",
	LinkUrl:    "link_url",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}


func NewSysAuthRuleDao() *SysAuthRuleDao {
	return &SysAuthRuleDao{
		group:   "default",
		table:   "sys_auth_rule",
		columns: sysAuthRuleColumns,
	}
}


func (dao *SysAuthRuleDao) DB() gdb.DB {
	return g.DB(dao.group)
}


func (dao *SysAuthRuleDao) Table() string {
	return dao.table
}


func (dao *SysAuthRuleDao) Columns() SysAuthRuleColumns {
	return dao.columns
}


func (dao *SysAuthRuleDao) Group() string {
	return dao.group
}


func (dao *SysAuthRuleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}







func (dao *SysAuthRuleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
