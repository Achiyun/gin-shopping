package adminRouter

import (
	con "github.com/Achiyun/gin-shopping/server/intetnal/app/html/controllers"
	"github.com/gin-gonic/gin"
)

type SettingRouters struct{}

func (s *RoleRouters) InitSettingRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	settingRouter := adminRouter.Group("setting")
	settingCon := con.ConGroupApp.AdminConGroup.SettingController
	{

		settingRouter.GET("/", settingCon.Index)
		settingRouter.POST("doEdit", settingCon.DoEdit)

	}
	return settingRouter
}
