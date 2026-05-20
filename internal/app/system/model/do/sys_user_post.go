package do

import (
	"github.com/gogf/gf/v2/frame/g"
)


type SysUserPost struct {
	g.Meta `orm:"table:sys_user_post, do:true"`
	UserId interface{}
	PostId interface{}
}
