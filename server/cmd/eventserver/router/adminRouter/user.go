// 对后台管理员的操作
package adminRouter

import (
	con "github.com/Achiyun/gin-shopping/server/intetnal/app/html/controllers"
	"github.com/gin-gonic/gin"
)

type UserRouters struct{}

func (s *UserRouters) InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	userRouter := adminRouter.Group("user")
	userCon := con.ConGroupApp.AdminConGroup.UserController

	{
		userRouter.GET("login", userCon.Index)
		userRouter.POST("doLogin", userCon.DoLogin)

	}
	return userRouter
}
