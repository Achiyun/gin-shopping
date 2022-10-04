package adminRouter

import (
	con "github.com/Achiyun/gin-shopping/server/intetnal/app/html/controllers"
	"github.com/gin-gonic/gin"
)

type FocusRouters struct{}

func (s *RoleRouters) InitFocusRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	focusRouter := adminRouter.Group("focus")
	focusCon := con.ConGroupApp.AdminConGroup.FocusController
	{

		focusRouter.GET("/", focusCon.Index)
		focusRouter.GET("add", focusCon.Add)
		focusRouter.POST("doAdd", focusCon.DoAdd)
		focusRouter.GET("edit", focusCon.Edit)
		focusRouter.POST("doEdit", focusCon.DoEdit)
		focusRouter.GET("delete", focusCon.Delete)

	}
	return focusRouter
}
