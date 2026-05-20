package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)


type SysLoginLogDao struct {
	table   string
	group   string
	columns SysLoginLogColumns
}


type SysLoginLogColumns struct {
	InfoId        string
	LoginName     string
	Ipaddr        string
	LoginLocation string
	Browser       string
	Os            string
	Status        string
	Msg           string
	LoginTime     string
	Module        string
}


var sysLoginLogColumns = SysLoginLogColumns{
	InfoId:        "info_id",
	LoginName:     "login_name",
	Ipaddr:        "ipaddr",
	LoginLocation: "login_location",
	Browser:       "browser",
	Os:            "os",
	Status:        "status",
	Msg:           "msg",
	LoginTime:     "login_time",
	Module:        "module",
}


func NewSysLoginLogDao() *SysLoginLogDao {
	return &SysLoginLogDao{
		group:   "default",
		table:   "sys_login_log",
		columns: sysLoginLogColumns,
	}
}


func (dao *SysLoginLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}


func (dao *SysLoginLogDao) Table() string {
	return dao.table
}


func (dao *SysLoginLogDao) Columns() SysLoginLogColumns {
	return dao.columns
}


func (dao *SysLoginLogDao) Group() string {
	return dao.group
}


func (dao *SysLoginLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}







func (dao *SysLoginLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
