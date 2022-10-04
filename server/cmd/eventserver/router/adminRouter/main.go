// 对后台管理员的操作
package adminRouter

import (
	con "github.com/Achiyun/gin-shopping/server/intetnal/app/html/controllers"
	"github.com/gin-gonic/gin"
)

type MainRouters struct{}

func (s *UserRouters) InitMainRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	mainRouter := adminRouter.Group("")
	mainCon := con.ConGroupApp.AdminConGroup.MainController
	{
		mainRouter.GET("/", mainCon.Index)
		mainRouter.GET("welcome", mainCon.Welcome)
		mainRouter.GET("changeStatus", mainCon.ChangeStatus)
		mainRouter.GET("changeNum", mainCon.ChangeNum)

	}
	return mainRouter
}
