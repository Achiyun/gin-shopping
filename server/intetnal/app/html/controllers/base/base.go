package base

import (
	"github.com/Achiyun/gin-shopping/server/intetnal/app/model/common/response"
	"github.com/Achiyun/gin-shopping/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

type Base struct {
}

func (b *Base) Captcha(c *gin.Context) {
	id, b64s, err := utils.MakeCaptcha()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.OkWithDetailed(gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
	}, "验证码获取成功", c)

}
