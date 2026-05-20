package common

import "github.com/gogf/gf/v2/frame/g"


type EmptyRes struct {
	g.Meta `mime:"application/json"`
}


type ListRes struct {
	CurrentPage int         `json:"currentPage"`
	Total       interface{} `json:"total"`
}
