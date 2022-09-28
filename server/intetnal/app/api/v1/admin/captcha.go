package admin

import (
	"github.com/Achiyun/gin-shopping/server/intetnal/app/model/common/response"
	systemRes "github.com/Achiyun/gin-shopping/server/intetnal/app/model/system/response"
	"github.com/Achiyun/gin-shopping/server/intetnal/pkg/utils"
	"github.com/gin-gonic/gin"
)

type BaseApi struct{}

func (b *BaseApi) Captcha(c *gin.Context) {
	id, b64s, err := utils.MakeCaptcha()

	if err != nil {
		response.FailWithMessage("验证码获取失败", c)
		return
	}
	// c.JSON(http.StatusOK, gin.H{
	// 	"captchaId":    id,
	// 	"captchaImage": b64s,
	// })
	response.OkWithDetailed(systemRes.SysCaptchaResponse{
		CaptchaId:    id,
		CaptchaImage: b64s,
	}, "验证码获取成功", c)
}
