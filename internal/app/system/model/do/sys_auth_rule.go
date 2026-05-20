package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)


type SysAuthRule struct {
	g.Meta     `orm:"table:sys_auth_rule, do:true"`
	Id         interface{}
	Pid        interface{}
	Name       interface{}
	Title      interface{}
	Icon       interface{}
	Condition  interface{}
	Remark     interface{}
	MenuType   interface{}
	Weigh      interface{}
	IsHide     interface{}
	Path       interface{}
	Component  interface{}
	IsLink     interface{}
	ModuleType interface{}
	ModelId    interface{}
	IsIframe   interface{}
	IsCached   interface{}
	Redirect   interface{}
	IsAffix    interface{}
	LinkUrl    interface{}
	CreatedAt  *gtime.Time
	UpdatedAt  *gtime.Time
}
