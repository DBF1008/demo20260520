package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)


type SysUser struct {
	g.Meta        `orm:"table:sys_user, do:true"`
	Id            interface{}
	UserName      interface{}
	Mobile        interface{}
	UserNickname  interface{}
	Birthday      interface{}
	UserPassword  interface{}
	UserSalt      interface{}
	UserStatus    interface{}
	UserEmail     interface{}
	Sex           interface{}
	Avatar        interface{}
	DeptId        interface{}
	Remark        interface{}
	IsAdmin       interface{}
	Address       interface{}
	Describe      interface{}
	LastLoginIp   interface{}
	LastLoginTime *gtime.Time
	CreatedAt     *gtime.Time
	UpdatedAt     *gtime.Time
	DeletedAt     *gtime.Time
}
