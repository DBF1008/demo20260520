package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)


type ToolsGenTableDao struct {
	table   string
	group   string
	columns ToolsGenTableColumns
}


type ToolsGenTableColumns struct {
	TableId        string
	TableName      string
	TableComment   string
	ClassName      string
	TplCategory    string
	PackageName    string
	ModuleName     string
	BusinessName   string
	FunctionName   string
	FunctionAuthor string
	Options        string
	CreateTime     string
	UpdateTime     string
	Remark         string
	Overwrite      string
	SortColumn     string
	SortType       string
	ShowDetail     string
}


var toolsGenTableColumns = ToolsGenTableColumns{
	TableId:        "table_id",
	TableName:      "table_name",
	TableComment:   "table_comment",
	ClassName:      "class_name",
	TplCategory:    "tpl_category",
	PackageName:    "package_name",
	ModuleName:     "module_name",
	BusinessName:   "business_name",
	FunctionName:   "function_name",
	FunctionAuthor: "function_author",
	Options:        "options",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
	Remark:         "remark",
	Overwrite:      "overwrite",
	SortColumn:     "sort_column",
	SortType:       "sort_type",
	ShowDetail:     "show_detail",
}


func NewToolsGenTableDao() *ToolsGenTableDao {
	return &ToolsGenTableDao{
		group:   "default",
		table:   "tools_gen_table",
		columns: toolsGenTableColumns,
	}
}


func (dao *ToolsGenTableDao) DB() gdb.DB {
	return g.DB(dao.group)
}


func (dao *ToolsGenTableDao) Table() string {
	return dao.table
}


func (dao *ToolsGenTableDao) Columns() ToolsGenTableColumns {
	return dao.columns
}


func (dao *ToolsGenTableDao) Group() string {
	return dao.group
}


func (dao *ToolsGenTableDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}







func (dao *ToolsGenTableDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
