package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/common/dao/internal"
)



type sysDictDataDao struct {
	*internal.SysDictDataDao
}

var (

	SysDictData = sysDictDataDao{
		internal.NewSysDictDataDao(),
	}
)
