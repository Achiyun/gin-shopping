package admin

import (
	"github.com/Achiyun/gin-shopping/server/intetnal/app/global"
	models "github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin"
	"github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin/request"
	AdminRes "github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin/response"
	"github.com/Achiyun/gin-shopping/server/intetnal/app/model/common/response"
	"github.com/Achiyun/gin-shopping/server/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (b *UserApi) Login(c *gin.Context) {
	var m request.Login
	err := c.ShouldBindJSON(&m) // 获取参数
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// err = utils.Verify(m, utils.LoginVerify)
	// if err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	m2 := &models.Manager{Username: m.Username, Password: m.Password}
	manager, _ := userService.Login(m2)
	b.TokenNext(c, *manager)

	// if flag := utils.VerifyCaptcha(m.CaptchaId, m.VerifyValue); flag {
	// 	m := &models.Manager{Username: m.Username, Password: m.Password}
	// 	manager, err := userService.Login(m)

	// 	if err != nil {
	// 		response.FailWithMessage("用户名不存在或者密码错误", c)
	// 		return
	// 	}
	// 	b.TokenNext(c, *manager)
	// 	return
	// }
	// response.FailWithMessage("验证码错误", c)
}

// TokenNext 登录以后签发jwt
func (b *UserApi) TokenNext(c *gin.Context, manager models.Manager) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(request.BaseClaims{
		UUID:     manager.UUID,
		ID:       manager.Id,
		Username: manager.Username,
	})
	token, err := j.CreateToken(claims)

	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint { // 这里禁止使用多点登陆
		response.OkWithDetailed(AdminRes.LoginResponse{
			Manager:   manager,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}
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
func (b *UserApi) ManagerList(c *gin.Context) {

	managerReturn, _ := userService.ManagerList()
	response.OkWithDetailed(AdminRes.AdminUserResponse{ManagerList: managerReturn}, "管理员列表", c)

}
