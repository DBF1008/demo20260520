package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)


type SysOperLogSearchReq struct {
	g.Meta        `path:"/operLog/list" tags:"操作日志" method:"get" summary:"操作日志列表"`
	Title         string `p:"title"`
	RequestMethod string `p:"requestMethod"`
	OperName      string `p:"operName"`
	commonApi.PageReq
	commonApi.Author
}


type SysOperLogSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*model.SysOperLogListRes `json:"list"`
}


type SysOperLogGetReq struct {
	g.Meta `path:"/operLog/get" tags:"操作日志" method:"get" summary:"获取操作日志信息"`
	commonApi.Author
	OperId uint64 `p:"operId" v:"required#主键必须"`
}


type SysOperLogGetRes struct {
	g.Meta `mime:"application/json"`
	*model.SysOperLogInfoRes
}


type SysOperLogDeleteReq struct {
	g.Meta `path:"/operLog/delete" tags:"操作日志" method:"delete" summary:"删除操作日志"`
	commonApi.Author
	OperIds []uint64 `p:"operIds" v:"required#主键必须"`
}


type SysOperLogDeleteRes struct {
	commonApi.EmptyRes
}

type SysOperLogClearReq struct {
	g.Meta `path:"/operLog/clear" tags:"操作日志" method:"delete" summary:"清除日志"`
	commonApi.Author
}

type SysOperLogClearRes struct {
	commonApi.EmptyRes
}
