package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/system/dao/internal"
)


type internalSysDeptDao = *internal.SysDeptDao



type sysDeptDao struct {
	internalSysDeptDao
}

var (

	SysDept = sysDeptDao{
		internal.NewSysDeptDao(),
	}
)
