package adminRouter

import (
	con "github.com/Achiyun/gin-shopping/server/intetnal/app/html/controllers"
	"github.com/gin-gonic/gin"
)

type ManagerRouters struct{}

func (s *UserRouters) InitManagerRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	managerRouter := adminRouter.Group("manager")
	managerCon := con.ConGroupApp.AdminConGroup.ManagerController
	{
		managerRouter.GET("/", managerCon.List)
		managerRouter.GET("add", managerCon.Add)
		managerRouter.POST("doAdd", managerCon.DoAdd)
		managerRouter.GET("edit", managerCon.Edit)
		managerRouter.POST("doEdit", managerCon.DoEdit)
		managerRouter.GET("delete", managerCon.Delete)

	}
	return managerRouter
}
