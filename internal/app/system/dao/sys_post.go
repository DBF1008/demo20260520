package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/system/dao/internal"
)



type sysPostDao struct {
	*internal.SysPostDao
}

var (

	SysPost = sysPostDao{
		internal.NewSysPostDao(),
	}
)
