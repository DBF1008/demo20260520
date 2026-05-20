package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/system/dao/internal"
)


type internalSysUserPostDao = *internal.SysUserPostDao



type sysUserPostDao struct {
	internalSysUserPostDao
}

var (

	SysUserPost = sysUserPostDao{
		internal.NewSysUserPostDao(),
	}
)
