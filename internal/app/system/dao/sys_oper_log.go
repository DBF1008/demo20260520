package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/system/dao/internal"
)


type internalSysOperLogDao = *internal.SysOperLogDao



type sysOperLogDao struct {
	internalSysOperLogDao
}

var (

	SysOperLog = sysOperLogDao{
		internal.NewSysOperLogDao(),
	}
)
