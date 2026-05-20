package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/system/dao/internal"
)



type sysRoleDao struct {
	*internal.SysRoleDao
}

var (

	SysRole = sysRoleDao{
		internal.NewSysRoleDao(),
	}
)
