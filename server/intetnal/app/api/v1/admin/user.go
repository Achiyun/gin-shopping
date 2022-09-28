package admin

import (
	models "github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin"
	"github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin/request"
	"github.com/Achiyun/gin-shopping/server/intetnal/app/model/common/response"
	"github.com/Achiyun/gin-shopping/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (b *BaseApi) Login(c *gin.Context) {
	var m request.Login
	err := c.ShouldBindJSON(&m) // 获取参数
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(m, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if flag := utils.VerifyCaptcha(m.CaptchaId, m.VerifyValue); flag {
		m := &models.Manager{Username: m.Username, Password: m.Password}
		manager, err := userService.Login(m)

		if err != nil {
			response.FailWithMessage("用户名不存在或者密码错误", c)
			return
		}
		b.TokenNext(c, *manager)
		return
	}
	response.FailWithMessage("验证码错误", c)
}
func (b *BaseApi) TokenNext(c *gin.Context, user models.Manager) {
}
