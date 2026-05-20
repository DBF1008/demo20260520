package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)


type SysDictDataDao struct {
	table   string
	group   string
	columns SysDictDataColumns
}


type SysDictDataColumns struct {
	DictCode  string
	DictSort  string
	DictLabel string
	DictValue string
	DictType  string
	CssClass  string
	ListClass string
	IsDefault string
	Status    string
	CreateBy  string
	UpdateBy  string
	Remark    string
	CreatedAt string
	UpdatedAt string
}


var sysDictDataColumns = SysDictDataColumns{
	DictCode:  "dict_code",
	DictSort:  "dict_sort",
	DictLabel: "dict_label",
	DictValue: "dict_value",
	DictType:  "dict_type",
	CssClass:  "css_class",
	ListClass: "list_class",
	IsDefault: "is_default",
	Status:    "status",
	CreateBy:  "create_by",
	UpdateBy:  "update_by",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}


func NewSysDictDataDao() *SysDictDataDao {
	return &SysDictDataDao{
		group:   "default",
		table:   "sys_dict_data",
		columns: sysDictDataColumns,
	}
}


func (dao *SysDictDataDao) DB() gdb.DB {
	return g.DB(dao.group)
}


func (dao *SysDictDataDao) Table() string {
	return dao.table
}


func (dao *SysDictDataDao) Columns() SysDictDataColumns {
	return dao.columns
}


func (dao *SysDictDataDao) Group() string {
	return dao.group
}


func (dao *SysDictDataDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}







func (dao *SysDictDataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
