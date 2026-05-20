package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)


type SysJobLogDao struct {
	table   string
	group   string
	columns SysJobLogColumns
}


type SysJobLogColumns struct {
	Id         string
	TargetName string
	CreatedAt  string
	Result     string
}


var sysJobLogColumns = SysJobLogColumns{
	Id:         "id",
	TargetName: "target_name",
	CreatedAt:  "created_at",
	Result:     "result",
}


func NewSysJobLogDao() *SysJobLogDao {
	return &SysJobLogDao{
		group:   "default",
		table:   "sys_job_log",
		columns: sysJobLogColumns,
	}
}


func (dao *SysJobLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}


func (dao *SysJobLogDao) Table() string {
	return dao.table
}


func (dao *SysJobLogDao) Columns() SysJobLogColumns {
	return dao.columns
}


func (dao *SysJobLogDao) Group() string {
	return dao.group
}


func (dao *SysJobLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}







func (dao *SysJobLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
