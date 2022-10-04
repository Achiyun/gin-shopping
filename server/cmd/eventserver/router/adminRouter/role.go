package adminRouter

import (
	con "github.com/Achiyun/gin-shopping/server/intetnal/app/html/controllers"
	"github.com/gin-gonic/gin"
)

type RoleRouters struct{}

func (s *RoleRouters) InitRoleRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	roleRouter := adminRouter.Group("role")
	roleCon := con.ConGroupApp.AdminConGroup.RoleController
	{

		roleRouter.GET("/", roleCon.List)
		roleRouter.GET("add", roleCon.Add)
		roleRouter.POST("doAdd", roleCon.DoAdd)
		roleRouter.GET("edit", roleCon.Edit)
		roleRouter.POST("doEdit", roleCon.DoEdit)
		roleRouter.GET("auth", roleCon.Auth)
		roleRouter.POST("doAuth", roleCon.DoAuth)
		roleRouter.GET("delete", roleCon.Delete)

	}
	return roleRouter
}
