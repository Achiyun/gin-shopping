package admin

import (
	models "github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin"
	"github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin/request"
	AdminRes "github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin/response"
	"github.com/Achiyun/gin-shopping/server/intetnal/app/model/common/response"
	"github.com/Achiyun/gin-shopping/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (b *UserApi) Login(c *gin.Context) {
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
func (b *UserApi) TokenNext(c *gin.Context, user models.Manager) {
}

// 先做注册, 加入uuid, 用uuid生成jwt
func (b *UserApi) Register(c *gin.Context) {
	var r request.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.RegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	manager := &models.Manager{Username: r.Username, Password: r.Password}
	managerReturn, err := userService.Register(*manager)

	if err != nil {
		response.FailWithDetailed(AdminRes.AdminUserResponse{Manager: managerReturn}, "注册失败", c)
		return
	}
	response.OkWithDetailed(AdminRes.AdminUserResponse{Manager: managerReturn}, "注册成功", c)
}
