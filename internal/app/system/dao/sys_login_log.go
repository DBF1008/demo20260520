package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/system/dao/internal"
)



type sysLoginLogDao struct {
	*internal.SysLoginLogDao
}

var (

	SysLoginLog = sysLoginLogDao{
		internal.NewSysLoginLogDao(),
	}
)
