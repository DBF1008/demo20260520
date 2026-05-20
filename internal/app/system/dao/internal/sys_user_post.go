package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)


type SysUserPostDao struct {
	table   string
	group   string
	columns SysUserPostColumns
}


type SysUserPostColumns struct {
	UserId string
	PostId string
}


var sysUserPostColumns = SysUserPostColumns{
	UserId: "user_id",
	PostId: "post_id",
}


func NewSysUserPostDao() *SysUserPostDao {
	return &SysUserPostDao{
		group:   "default",
		table:   "sys_user_post",
		columns: sysUserPostColumns,
	}
}


func (dao *SysUserPostDao) DB() gdb.DB {
	return g.DB(dao.group)
}


func (dao *SysUserPostDao) Table() string {
	return dao.table
}


func (dao *SysUserPostDao) Columns() SysUserPostColumns {
	return dao.columns
}


func (dao *SysUserPostDao) Group() string {
	return dao.group
}


func (dao *SysUserPostDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}







func (dao *SysUserPostDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
