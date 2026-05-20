package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/system/dao/internal"
)


type internalToolsGenTableDao = *internal.ToolsGenTableDao



type toolsGenTableDao struct {
	internalToolsGenTableDao
}

var (

	ToolsGenTable = toolsGenTableDao{
		internal.NewToolsGenTableDao(),
	}
)
