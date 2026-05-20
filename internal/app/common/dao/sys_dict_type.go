package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/common/dao/internal"
)



type sysDictTypeDao struct {
	*internal.SysDictTypeDao
}

var (

	SysDictType = sysDictTypeDao{
		internal.NewSysDictTypeDao(),
	}
)
