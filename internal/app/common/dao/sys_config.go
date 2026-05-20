package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/common/dao/internal"
)



type sysConfigDao struct {
	*internal.SysConfigDao
}

var (

	SysConfig = sysConfigDao{
		internal.NewSysConfigDao(),
	}
)
