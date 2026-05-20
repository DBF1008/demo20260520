package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)


type SysUserOnlineDao struct {
	table   string
	group   string
	columns SysUserOnlineColumns
}


type SysUserOnlineColumns struct {
	Id         string
	Uuid       string
	Token      string
	CreateTime string
	UserName   string
	Ip         string
	Explorer   string
	Os         string
}


var sysUserOnlineColumns = SysUserOnlineColumns{
	Id:         "id",
	Uuid:       "uuid",
	Token:      "token",
	CreateTime: "create_time",
	UserName:   "user_name",
	Ip:         "ip",
	Explorer:   "explorer",
	Os:         "os",
}


func NewSysUserOnlineDao() *SysUserOnlineDao {
	return &SysUserOnlineDao{
		group:   "default",
		table:   "sys_user_online",
		columns: sysUserOnlineColumns,
	}
}


func (dao *SysUserOnlineDao) DB() gdb.DB {
	return g.DB(dao.group)
}


func (dao *SysUserOnlineDao) Table() string {
	return dao.table
}


func (dao *SysUserOnlineDao) Columns() SysUserOnlineColumns {
	return dao.columns
}


func (dao *SysUserOnlineDao) Group() string {
	return dao.group
}


func (dao *SysUserOnlineDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}







func (dao *SysUserOnlineDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
