package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)


type SysLoginLog struct {
	g.Meta        `orm:"table:sys_login_log, do:true"`
	InfoId        interface{}
	LoginName     interface{}
	Ipaddr        interface{}
	LoginLocation interface{}
	Browser       interface{}
	Os            interface{}
	Status        interface{}
	Msg           interface{}
	LoginTime     *gtime.Time
	Module        interface{}
}
