package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/system/dao/internal"
)


type internalToolsGenTableColumnDao = *internal.ToolsGenTableColumnDao



type toolsGenTableColumnDao struct {
	internalToolsGenTableColumnDao
}

var (

	ToolsGenTableColumn = toolsGenTableColumnDao{
		internal.NewToolsGenTableColumnDao(),
	}
)
