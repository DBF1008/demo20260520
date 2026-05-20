package do

import (
	"github.com/gogf/gf/v2/frame/g"
)


type SysRoleDept struct {
	g.Meta `orm:"table:sys_role_dept, do:true"`
	RoleId interface{}
	DeptId interface{}
}
