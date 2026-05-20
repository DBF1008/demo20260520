package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/common/dao/internal"
)



type casbinRuleDao struct {
	*internal.CasbinRuleDao
}

var (

	CasbinRule = casbinRuleDao{
		internal.NewCasbinRuleDao(),
	}
)
