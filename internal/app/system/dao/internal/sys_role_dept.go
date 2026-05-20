package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)


type SysRoleDeptDao struct {
	table   string
	group   string
	columns SysRoleDeptColumns
}


type SysRoleDeptColumns struct {
	RoleId string
	DeptId string
}


var sysRoleDeptColumns = SysRoleDeptColumns{
	RoleId: "role_id",
	DeptId: "dept_id",
}


func NewSysRoleDeptDao() *SysRoleDeptDao {
	return &SysRoleDeptDao{
		group:   "default",
		table:   "sys_role_dept",
		columns: sysRoleDeptColumns,
	}
}


func (dao *SysRoleDeptDao) DB() gdb.DB {
	return g.DB(dao.group)
}


func (dao *SysRoleDeptDao) Table() string {
	return dao.table
}


func (dao *SysRoleDeptDao) Columns() SysRoleDeptColumns {
	return dao.columns
}


func (dao *SysRoleDeptDao) Group() string {
	return dao.group
}


func (dao *SysRoleDeptDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}







func (dao *SysRoleDeptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
