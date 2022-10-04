package adminRouter

import (
	con "github.com/Achiyun/gin-shopping/server/intetnal/app/html/controllers"
	"github.com/gin-gonic/gin"
)

type AccessRouters struct{}

func (s *RoleRouters) InitAccessRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	accessRouter := adminRouter.Group("access")
	accessCon := con.ConGroupApp.AdminConGroup.AccessController
	{

		accessRouter.GET("/", accessCon.Index)
		accessRouter.GET("add", accessCon.Add)
		accessRouter.POST("doAdd", accessCon.DoAdd)
		accessRouter.GET("edit", accessCon.Edit)
		accessRouter.POST("doEdit", accessCon.DoEdit)
		accessRouter.GET("delete", accessCon.Delete)

	}
	return accessRouter
}
