package baseRouter

import (
	"github.com/Achiyun/gin-shopping/server/intetnal/app/html/controllers/base"
	"github.com/gin-gonic/gin"
)

type BaseRouters struct{}

func (s *BaseRouters) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	baseRouter := Router.Group("base")
	baseCon := base.BaseGroupApp.Base
	{
		baseRouter.POST("captcha", baseCon.Captcha)
	}
	return baseRouter
}
