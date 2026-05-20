package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)


type ToolsGenTableColumnDao struct {
	table   string
	group   string
	columns ToolsGenTableColumnColumns
}


type ToolsGenTableColumnColumns struct {
	ColumnId              string
	TableId               string
	ColumnName            string
	ColumnComment         string
	ColumnType            string
	GoType                string
	TsType                string
	GoField               string
	HtmlField             string
	IsPk                  string
	IsIncrement           string
	IsRequired            string
	IsInsert              string
	IsEdit                string
	IsList                string
	IsDetail              string
	IsQuery               string
	SortOrderEdit         string
	SortOrderList         string
	SortOrderDetail       string
	SortOrderQuery        string
	QueryType             string
	HtmlType              string
	DictType              string
	LinkTableName         string
	LinkTableClass        string
	LinkTableModuleName   string
	LinkTableBusinessName string
	LinkTablePackage      string
	LinkLabelId           string
	LinkLabelName         string
	ColSpan               string
	RowSpan               string
	IsRowStart            string
	MinWidth              string
	IsFixed               string
	IsOverflowTooltip     string
	IsCascade             string
	ParentColumnName      string
	CascadeColumnName     string
}


var toolsGenTableColumnColumns = ToolsGenTableColumnColumns{
	ColumnId:              "column_id",
	TableId:               "table_id",
	ColumnName:            "column_name",
	ColumnComment:         "column_comment",
	ColumnType:            "column_type",
	GoType:                "go_type",
	TsType:                "ts_type",
	GoField:               "go_field",
	HtmlField:             "html_field",
	IsPk:                  "is_pk",
	IsIncrement:           "is_increment",
	IsRequired:            "is_required",
	IsInsert:              "is_insert",
	IsEdit:                "is_edit",
	IsList:                "is_list",
	IsDetail:              "is_detail",
	IsQuery:               "is_query",
	SortOrderEdit:         "sort_order_edit",
	SortOrderList:         "sort_order_list",
	SortOrderDetail:       "sort_order_detail",
	SortOrderQuery:        "sort_order_query",
	QueryType:             "query_type",
	HtmlType:              "html_type",
	DictType:              "dict_type",
	LinkTableName:         "link_table_name",
	LinkTableClass:        "link_table_class",
	LinkTableModuleName:   "link_table_module_name",
	LinkTableBusinessName: "link_table_business_name",
	LinkTablePackage:      "link_table_package",
	LinkLabelId:           "link_label_id",
	LinkLabelName:         "link_label_name",
	ColSpan:               "col_span",
	RowSpan:               "row_span",
	IsRowStart:            "is_row_start",
	MinWidth:              "min_width",
	IsFixed:               "is_fixed",
	IsOverflowTooltip:     "is_overflow_tooltip",
	IsCascade:             "is_cascade",
	ParentColumnName:      "parent_column_name",
	CascadeColumnName:     "cascade_column_name",
}


func NewToolsGenTableColumnDao() *ToolsGenTableColumnDao {
	return &ToolsGenTableColumnDao{
		group:   "default",
		table:   "tools_gen_table_column",
		columns: toolsGenTableColumnColumns,
	}
}


func (dao *ToolsGenTableColumnDao) DB() gdb.DB {
	return g.DB(dao.group)
}


func (dao *ToolsGenTableColumnDao) Table() string {
	return dao.table
}


func (dao *ToolsGenTableColumnDao) Columns() ToolsGenTableColumnColumns {
	return dao.columns
}


func (dao *ToolsGenTableColumnDao) Group() string {
	return dao.group
}


func (dao *ToolsGenTableColumnDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}







func (dao *ToolsGenTableColumnDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
