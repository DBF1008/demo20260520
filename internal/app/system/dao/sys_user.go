package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/system/dao/internal"
)



type sysUserDao struct {
	*internal.SysUserDao
}

var (

	SysUser = sysUserDao{
		internal.NewSysUserDao(),
	}
)
