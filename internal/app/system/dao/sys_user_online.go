package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/system/dao/internal"
)


type internalSysUserOnlineDao = *internal.SysUserOnlineDao



type sysUserOnlineDao struct {
	internalSysUserOnlineDao
}

var (

	SysUserOnline = sysUserOnlineDao{
		internal.NewSysUserOnlineDao(),
	}
)
