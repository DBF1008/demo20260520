package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/system/dao/internal"
)



type sysJobDao struct {
	*internal.SysJobDao
}

var (

	SysJob = sysJobDao{
		internal.NewSysJobDao(),
	}
)
