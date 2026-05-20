package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)


type SysDictData struct {
	g.Meta    `orm:"table:sys_dict_data, do:true"`
	DictCode  interface{}
	DictSort  interface{}
	DictLabel interface{}
	DictValue interface{}
	DictType  interface{}
	CssClass  interface{}
	ListClass interface{}
	IsDefault interface{}
	Status    interface{}
	CreateBy  interface{}
	UpdateBy  interface{}
	Remark    interface{}
	CreatedAt *gtime.Time
	UpdatedAt *gtime.Time
}
