package model

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TokenOptions struct {

	ServerName string `json:"serverName"`

	CacheKey string `json:"cacheKey"`

	Timeout int64 `json:"timeout"`



	MaxRefresh int64 `json:"maxRefresh"`

	MultiLogin bool `json:"multiLogin"`

	EncryptKey []byte `json:"encryptKey"`

	ExcludePaths g.SliceStr `json:"excludePaths"`
	CacheModel   string     `json:"cacheModel"`
}
