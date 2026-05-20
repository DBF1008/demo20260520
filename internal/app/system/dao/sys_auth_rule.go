package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/system/dao/internal"
)



type sysAuthRuleDao struct {
	*internal.SysAuthRuleDao
}

var (

	SysAuthRule = sysAuthRuleDao{
		internal.NewSysAuthRuleDao(),
	}
)
