package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)


type SysDictTypeDao struct {
	table   string
	group   string
	columns SysDictTypeColumns
}


type SysDictTypeColumns struct {
	DictId    string
	DictName  string
	DictType  string
	Status    string
	CreateBy  string
	UpdateBy  string
	Remark    string
	CreatedAt string
	UpdatedAt string
}


var sysDictTypeColumns = SysDictTypeColumns{
	DictId:    "dict_id",
	DictName:  "dict_name",
	DictType:  "dict_type",
	Status:    "status",
	CreateBy:  "create_by",
	UpdateBy:  "update_by",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}


func NewSysDictTypeDao() *SysDictTypeDao {
	return &SysDictTypeDao{
		group:   "default",
		table:   "sys_dict_type",
		columns: sysDictTypeColumns,
	}
}


func (dao *SysDictTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}


func (dao *SysDictTypeDao) Table() string {
	return dao.table
}


func (dao *SysDictTypeDao) Columns() SysDictTypeColumns {
	return dao.columns
}


func (dao *SysDictTypeDao) Group() string {
	return dao.group
}


func (dao *SysDictTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}







func (dao *SysDictTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
