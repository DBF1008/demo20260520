package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)


type SysOperLog struct {
	g.Meta        `orm:"table:sys_oper_log, do:true"`
	OperId        interface{}
	Title         interface{}
	BusinessType  interface{}
	Method        interface{}
	RequestMethod interface{}
	OperatorType  interface{}
	OperName      interface{}
	DeptName      interface{}
	OperUrl       interface{}
	OperIp        interface{}
	OperLocation  interface{}
	OperParam     interface{}
	ErrorMsg      interface{}
	OperTime      *gtime.Time
}
