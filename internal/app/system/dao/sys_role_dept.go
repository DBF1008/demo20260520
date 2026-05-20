package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/system/dao/internal"
)


type internalSysRoleDeptDao = *internal.SysRoleDeptDao



type sysRoleDeptDao struct {
	internalSysRoleDeptDao
}

var (

	SysRoleDept = sysRoleDeptDao{
		internal.NewSysRoleDeptDao(),
	}
)
