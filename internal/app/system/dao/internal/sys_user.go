package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)


type SysUserDao struct {
	table   string
	group   string
	columns SysUserColumns
}


type SysUserColumns struct {
	Id            string
	UserName      string
	Mobile        string
	UserNickname  string
	Birthday      string
	UserPassword  string
	UserSalt      string
	UserStatus    string
	UserEmail     string
	Sex           string
	Avatar        string
	DeptId        string
	Remark        string
	IsAdmin       string
	Address       string
	Describe      string
	LastLoginIp   string
	LastLoginTime string
	CreatedAt     string
	UpdatedAt     string
	DeletedAt     string
}


var sysUserColumns = SysUserColumns{
	Id:            "id",
	UserName:      "user_name",
	Mobile:        "mobile",
	UserNickname:  "user_nickname",
	Birthday:      "birthday",
	UserPassword:  "user_password",
	UserSalt:      "user_salt",
	UserStatus:    "user_status",
	UserEmail:     "user_email",
	Sex:           "sex",
	Avatar:        "avatar",
	DeptId:        "dept_id",
	Remark:        "remark",
	IsAdmin:       "is_admin",
	Address:       "address",
	Describe:      "describe",
	LastLoginIp:   "last_login_ip",
	LastLoginTime: "last_login_time",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
}


func NewSysUserDao() *SysUserDao {
	return &SysUserDao{
		group:   "default",
		table:   "sys_user",
		columns: sysUserColumns,
	}
}


func (dao *SysUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}


func (dao *SysUserDao) Table() string {
	return dao.table
}


func (dao *SysUserDao) Columns() SysUserColumns {
	return dao.columns
}


func (dao *SysUserDao) Group() string {
	return dao.group
}


func (dao *SysUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}







func (dao *SysUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
