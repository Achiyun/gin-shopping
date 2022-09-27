package service

import (
	"fmt"
	"net/http"

	service "github.com/Achiyun/gin-shopping/server/intetnal/app/service"
	captcha "github.com/Achiyun/gin-shopping/server/intetnal/pkg/utils"

	"github.com/gin-gonic/gin"
)

type LoginService struct {
	service.BaseController
}

func (con LoginService) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "张三",
	})
}
func (con LoginService) Captcha(c *gin.Context) {
	id, b64s, err := captcha.MakeCaptcha()

	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
	})
}
