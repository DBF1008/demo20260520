package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)


type SysUserOnline struct {
	g.Meta     `orm:"table:sys_user_online, do:true"`
	Id         interface{}
	Uuid       interface{}
	Token      interface{}
	CreateTime *gtime.Time
	UserName   interface{}
	Ip         interface{}
	Explorer   interface{}
	Os         interface{}
}
