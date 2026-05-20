package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)


type SysPost struct {
	g.Meta    `orm:"table:sys_post, do:true"`
	PostId    interface{}
	PostCode  interface{}
	PostName  interface{}
	PostSort  interface{}
	Status    interface{}
	Remark    interface{}
	CreatedBy interface{}
	UpdatedBy interface{}
	CreatedAt *gtime.Time
	UpdatedAt *gtime.Time
	DeletedAt *gtime.Time
}
