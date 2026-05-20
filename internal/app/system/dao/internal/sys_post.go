package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)


type SysPostDao struct {
	table   string
	group   string
	columns SysPostColumns
}


type SysPostColumns struct {
	PostId    string
	PostCode  string
	PostName  string
	PostSort  string
	Status    string
	Remark    string
	CreatedBy string
	UpdatedBy string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}


var sysPostColumns = SysPostColumns{
	PostId:    "post_id",
	PostCode:  "post_code",
	PostName:  "post_name",
	PostSort:  "post_sort",
	Status:    "status",
	Remark:    "remark",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}


func NewSysPostDao() *SysPostDao {
	return &SysPostDao{
		group:   "default",
		table:   "sys_post",
		columns: sysPostColumns,
	}
}


func (dao *SysPostDao) DB() gdb.DB {
	return g.DB(dao.group)
}


func (dao *SysPostDao) Table() string {
	return dao.table
}


func (dao *SysPostDao) Columns() SysPostColumns {
	return dao.columns
}


func (dao *SysPostDao) Group() string {
	return dao.group
}


func (dao *SysPostDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}







func (dao *SysPostDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
