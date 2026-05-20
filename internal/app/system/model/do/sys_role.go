package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)


type SysRole struct {
	g.Meta    `orm:"table:sys_role, do:true"`
	Id        interface{}
	Status    interface{}
	ListOrder interface{}
	Name      interface{}
	Remark    interface{}
	DataScope interface{}
	CreatedAt *gtime.Time
	UpdatedAt *gtime.Time
}
