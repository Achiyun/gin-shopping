package grouprouter

import (
	admin "github.com/Achiyun/gin-shopping/server/intetnal/app/service/admin"
	"github.com/gin-gonic/gin"
)

type AdminRouters struct{}

func (s *AdminRouters) InitAdminRouter(r *gin.Engine) {
	adminRouters := r.Group("/")
	{
		adminRouters.GET("/", admin.LoginService{}.Index)
		adminRouters.POST("/DoLogin", admin.LoginService{}.Login)
	}
}
