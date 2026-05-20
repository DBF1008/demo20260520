package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)


type SysDictType struct {
	g.Meta    `orm:"table:sys_dict_type, do:true"`
	DictId    interface{}
	DictName  interface{}
	DictType  interface{}
	Status    interface{}
	CreateBy  interface{}
	UpdateBy  interface{}
	Remark    interface{}
	CreatedAt *gtime.Time
	UpdatedAt *gtime.Time
}
