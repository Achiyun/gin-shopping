package admin

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin"
	"github.com/gin-contrib/sessions"

	"github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin/request"
	"github.com/Achiyun/gin-shopping/server/intetnal/app/model/common/response"
	"github.com/Achiyun/gin-shopping/server/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {

	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}

func (con *UserController) DoLogin(c *gin.Context) {

	var m request.Login

	err := c.ShouldBind(&m)
	fmt.Println(m)
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
		session := sessions.Default(c)
		//注意：session.Set没法直接保存结构体对应的切片 把结构体转换成json字符串
		userinfo, _ := json.Marshal(manager)
		session.Set("userinfo", string(userinfo))
		session.Save()
		con.Success(c, "登录成功", "/admin")
		return
	}

	response.FailWithMessage("验证码错误", c)
}
