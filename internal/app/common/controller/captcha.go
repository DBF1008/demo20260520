package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/common/service"
)

var Captcha = captchaController{}

type captchaController struct {
}


func (c *captchaController) Get(ctx context.Context, req *common.CaptchaReq) (res *common.CaptchaRes, err error) {
	var (
		idKeyC, base64stringC string
	)
	idKeyC, base64stringC, err = service.Captcha().GetVerifyImgString(ctx)
	res = &common.CaptchaRes{
		Key: idKeyC,
		Img: base64stringC,
	}
	return
}
