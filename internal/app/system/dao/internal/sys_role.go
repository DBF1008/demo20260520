package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)


type SysRoleDao struct {
	table   string
	group   string
	columns SysRoleColumns
}


type SysRoleColumns struct {
	Id        string
	Status    string
	ListOrder string
	Name      string
	Remark    string
	DataScope string
	CreatedAt string
	UpdatedAt string
}


var sysRoleColumns = SysRoleColumns{
	Id:        "id",
	Status:    "status",
	ListOrder: "list_order",
	Name:      "name",
	Remark:    "remark",
	DataScope: "data_scope",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}


func NewSysRoleDao() *SysRoleDao {
	return &SysRoleDao{
		group:   "default",
		table:   "sys_role",
		columns: sysRoleColumns,
	}
}


func (dao *SysRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}


func (dao *SysRoleDao) Table() string {
	return dao.table
}


func (dao *SysRoleDao) Columns() SysRoleColumns {
	return dao.columns
}


func (dao *SysRoleDao) Group() string {
	return dao.group
}


func (dao *SysRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}







func (dao *SysRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
