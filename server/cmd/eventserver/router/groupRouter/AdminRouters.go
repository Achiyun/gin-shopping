package grouprouter

import (
	admin "github.com/Achiyun/gin-shopping/server/intetnal/app/service/admin"
	"github.com/gin-gonic/gin"
)

func AdminRouters(r *gin.Engine) {
	adminRouters := r.Group("/")
	{
		adminRouters.GET("/", admin.LoginService{}.Index)
		adminRouters.GET("/captcha", admin.LoginService{}.Captcha)
	}
}
