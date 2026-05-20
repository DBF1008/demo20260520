package model

import "github.com/gogf/gf/v2/os/gtime"

type SysDictTypeInfoRes struct {
	DictId    uint64      `orm:"dict_id,primary"  json:"dictId"`
	DictName  string      `orm:"dict_name"        json:"dictName"`
	DictType  string      `orm:"dict_type,unique" json:"dictType"`
	Status    uint        `orm:"status"           json:"status"`
	Remark    string      `orm:"remark"           json:"remark"`
	CreatedAt *gtime.Time `orm:"created_at"       json:"createdAt"`
}
