package adminRouter

import (
	con "github.com/Achiyun/gin-shopping/server/intetnal/app/html/controllers"
	"github.com/gin-gonic/gin"
)

type NavRouters struct{}

func (s *RoleRouters) InitNavRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	navRouter := adminRouter.Group("nav")
	navCon := con.ConGroupApp.AdminConGroup.NavController
	{

		navRouter.GET("/", navCon.Index)
		navRouter.GET("add", navCon.Add)
		navRouter.POST("doAdd", navCon.DoAdd)
		navRouter.GET("edit", navCon.Edit)
		navRouter.POST("doEdit", navCon.DoEdit)
		navRouter.GET("delete", navCon.Delete)

	}
	return navRouter
}
