package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/system/dao/internal"
)


type internalSysJobLogDao = *internal.SysJobLogDao



type sysJobLogDao struct {
	internalSysJobLogDao
}

var (

	SysJobLog = sysJobLogDao{
		internal.NewSysJobLogDao(),
	}
)
