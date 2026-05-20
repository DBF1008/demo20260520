package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)


type SysOperLogDao struct {
	table   string
	group   string
	columns SysOperLogColumns
}


type SysOperLogColumns struct {
	OperId        string
	Title         string
	BusinessType  string
	Method        string
	RequestMethod string
	OperatorType  string
	OperName      string
	DeptName      string
	OperUrl       string
	OperIp        string
	OperLocation  string
	OperParam     string
	ErrorMsg      string
	OperTime      string
}


var sysOperLogColumns = SysOperLogColumns{
	OperId:        "oper_id",
	Title:         "title",
	BusinessType:  "business_type",
	Method:        "method",
	RequestMethod: "request_method",
	OperatorType:  "operator_type",
	OperName:      "oper_name",
	DeptName:      "dept_name",
	OperUrl:       "oper_url",
	OperIp:        "oper_ip",
	OperLocation:  "oper_location",
	OperParam:     "oper_param",
	ErrorMsg:      "error_msg",
	OperTime:      "oper_time",
}


func NewSysOperLogDao() *SysOperLogDao {
	return &SysOperLogDao{
		group:   "default",
		table:   "sys_oper_log",
		columns: sysOperLogColumns,
	}
}


func (dao *SysOperLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}


func (dao *SysOperLogDao) Table() string {
	return dao.table
}


func (dao *SysOperLogDao) Columns() SysOperLogColumns {
	return dao.columns
}


func (dao *SysOperLogDao) Group() string {
	return dao.group
}


func (dao *SysOperLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}







func (dao *SysOperLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
