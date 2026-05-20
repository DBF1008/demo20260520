package model

import (
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)


type LoginUserRes struct {
	Id           uint64 `orm:"id,primary"       json:"id"`
	UserName     string `orm:"user_name,unique" json:"userName"`
	UserNickname string `orm:"user_nickname"    json:"userNickname"`
	UserPassword string `orm:"user_password"    json:"userPassword"`
	UserSalt     string `orm:"user_salt"        json:"userSalt"`
	UserStatus   uint   `orm:"user_status"      json:"userStatus"`
	IsAdmin      int    `orm:"is_admin"         json:"isAdmin"`
	Avatar       string `orm:"avatar" json:"avatar"`
	DeptId       uint64 `orm:"dept_id"       json:"deptId"`
}


type SysUserRoleDeptRes struct {
	*entity.SysUser
	Dept     *entity.SysDept       `json:"dept"`
	RoleInfo []*SysUserRoleInfoRes `json:"roleInfo"`
	Post     []*SysUserPostInfoRes `json:"post"`
}

type SysUserRoleInfoRes struct {
	RoleId uint   `json:"roleId"`
	Name   string `json:"name"`
}

type SysUserPostInfoRes struct {
	PostId   int64  `json:"postId"`
	PostName string `json:"postName"`
}

type SysUserSimpleRes struct {
	gmeta.Meta   `orm:"table:sys_user"`
	Id           uint64 `orm:"id"       json:"id"`
	Avatar       string `orm:"avatar" json:"avatar"`
	Sex          int    `orm:"sex" json:"sex"`
	UserName     string `orm:"user_name" json:"userName"`
	UserNickname string `orm:"user_nickname"    json:"userNickname"`
}
