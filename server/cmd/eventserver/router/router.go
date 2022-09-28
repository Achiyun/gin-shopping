package router

import (
	"fmt"

	grouter "github.com/Achiyun/gin-shopping/server/cmd/eventserver/router/groupRouter"
	_ "github.com/Achiyun/gin-shopping/server/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(r *gin.Engine) {
	fmt.Println("Initialization router")

	br := grouter.RouterGroupApp.AdminRouters
	br1 := grouter.RouterGroupApp.BaseRouter
	PublicGroup := r.Group("")

	br1.InitBaseRouter(PublicGroup)
	br.InitAdminRouter(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
