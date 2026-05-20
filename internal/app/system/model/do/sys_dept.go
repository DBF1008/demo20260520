package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)


type SysDept struct {
	g.Meta    `orm:"table:sys_dept, do:true"`
	DeptId    interface{}
	ParentId  interface{}
	Ancestors interface{}
	DeptName  interface{}
	OrderNum  interface{}
	Leader    interface{}
	Phone     interface{}
	Email     interface{}
	Status    interface{}
	CreatedBy interface{}
	UpdatedBy interface{}
	CreatedAt *gtime.Time
	UpdatedAt *gtime.Time
	DeletedAt *gtime.Time
}
