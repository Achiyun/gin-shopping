package grouprouter

import (
	service "github.com/Achiyun/gin-shopping/server/intetnal/app/service"
	"github.com/gin-gonic/gin"
)

func AdminRouters(r *gin.Engine) {
	adminRouters := r.Group("/")
	{
		adminRouters.GET("/", service.LoginService{}.Index)
	}
}
