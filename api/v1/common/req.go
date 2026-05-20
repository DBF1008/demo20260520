package common

import "github.com/tiger1103/gfast/v3/internal/app/common/model"


type PageReq struct {
	model.PageReq
}

type Author struct {
	Authorization string `p:"Authorization" in:"header" dc:"Bearer {{token}}"`
}
