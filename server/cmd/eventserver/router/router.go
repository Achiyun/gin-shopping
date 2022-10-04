package router

import (
	"github.com/Achiyun/gin-shopping/server/cmd/eventserver/router/adminRouter"
	"github.com/Achiyun/gin-shopping/server/cmd/eventserver/router/baseRouter"
	_ "github.com/Achiyun/gin-shopping/server/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(r *gin.Engine) {

	// AdminRouters := adminRouter.AdminGroupApp.UserRouters
	BaseRouters := baseRouter.BaseGroupApp.BaseRouters
	// ApiRouters := apiRouter.ApiGroupApp.V1Group.UserRouters

	adminRouter.Init(r)
	// AdminRouters.InitUserRouter(PublicGroup)
	BaseRouters.InitBaseRouter(r.Group(""))

	// ApiRouters.InitUserRouter(r.Group(""))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
