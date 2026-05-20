package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)


type SysJobDao struct {
	table   string
	group   string
	columns SysJobColumns
}


type SysJobColumns struct {
	JobId          string
	JobName        string
	JobParams      string
	JobGroup       string
	InvokeTarget   string
	CronExpression string
	MisfirePolicy  string
	Concurrent     string
	Status         string
	CreatedBy      string
	UpdatedBy      string
	Remark         string
	CreatedAt      string
	UpdatedAt      string
}

var sysJobColumns = SysJobColumns{
	JobId:          "job_id",
	JobName:        "job_name",
	JobParams:      "job_params",
	JobGroup:       "job_group",
	InvokeTarget:   "invoke_target",
	CronExpression: "cron_expression",
	MisfirePolicy:  "misfire_policy",
	Concurrent:     "concurrent",
	Status:         "status",
	CreatedBy:      "created_by",
	UpdatedBy:      "updated_by",
	Remark:         "remark",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}


func NewSysJobDao() *SysJobDao {
	return &SysJobDao{
		group:   "default",
		table:   "sys_job",
		columns: sysJobColumns,
	}
}


func (dao *SysJobDao) DB() gdb.DB {
	return g.DB(dao.group)
}


func (dao *SysJobDao) Table() string {
	return dao.table
}


func (dao *SysJobDao) Columns() SysJobColumns {
	return dao.columns
}


func (dao *SysJobDao) Group() string {
	return dao.group
}


func (dao *SysJobDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}







func (dao *SysJobDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
