package model

import (
	"net/url"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)


type SysOperLogAdd struct {
	User         *ContextUser
	Menu         *SysAuthRuleInfoRes
	Url          *url.URL
	Params       g.Map
	Method       string
	ClientIp     string
	OperatorType int
}


type SysOperLogInfoRes struct {
	gmeta.Meta    `orm:"table:sys_oper_log"`
	OperId        uint64      `orm:"oper_id,primary" json:"operId"`
	Title         string      `orm:"title" json:"title"`
	BusinessType  int         `orm:"business_type" json:"businessType"`
	Method        string      `orm:"method" json:"method"`
	RequestMethod string      `orm:"request_method" json:"requestMethod"`
	OperatorType  int         `orm:"operator_type" json:"operatorType"`
	OperName      string      `orm:"oper_name" json:"operName"`
	DeptName      string      `orm:"dept_name" json:"deptName"`
	OperUrl       string      `orm:"oper_url" json:"operUrl"`
	OperIp        string      `orm:"oper_ip" json:"operIp"`
	OperLocation  string      `orm:"oper_location" json:"operLocation"`
	OperParam     string      `orm:"oper_param" json:"operParam"`
	ErrorMsg      string      `orm:"error_msg" json:"errorMsg"`
	OperTime      *gtime.Time `orm:"oper_time" json:"operTime"`
}

type SysOperLogListRes struct {
	OperId        uint64      `json:"operId"`
	Title         string      `json:"title"`
	RequestMethod string      `json:"requestMethod"`
	OperName      string      `json:"operName"`
	DeptName      string      `json:"deptName"`
	OperUrl       string      `json:"operUrl"`
	OperIp        string      `json:"operIp"`
	OperLocation  string      `json:"operLocation"`
	OperParam     string      `json:"operParam"`
	OperTime      *gtime.Time `json:"operTime"`
}
