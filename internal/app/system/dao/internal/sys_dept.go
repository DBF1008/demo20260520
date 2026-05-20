package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)


type SysDeptDao struct {
	table   string
	group   string
	columns SysDeptColumns
}


type SysDeptColumns struct {
	DeptId    string
	ParentId  string
	Ancestors string
	DeptName  string
	OrderNum  string
	Leader    string
	Phone     string
	Email     string
	Status    string
	CreatedBy string
	UpdatedBy string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}


var sysDeptColumns = SysDeptColumns{
	DeptId:    "dept_id",
	ParentId:  "parent_id",
	Ancestors: "ancestors",
	DeptName:  "dept_name",
	OrderNum:  "order_num",
	Leader:    "leader",
	Phone:     "phone",
	Email:     "email",
	Status:    "status",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}


func NewSysDeptDao() *SysDeptDao {
	return &SysDeptDao{
		group:   "default",
		table:   "sys_dept",
		columns: sysDeptColumns,
	}
}


func (dao *SysDeptDao) DB() gdb.DB {
	return g.DB(dao.group)
}


func (dao *SysDeptDao) Table() string {
	return dao.table
}


func (dao *SysDeptDao) Columns() SysDeptColumns {
	return dao.columns
}


func (dao *SysDeptDao) Group() string {
	return dao.group
}


func (dao *SysDeptDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}







func (dao *SysDeptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
