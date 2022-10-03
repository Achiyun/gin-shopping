// 对后台管理员的操作
package v1

import (
	v1 "github.com/Achiyun/gin-shopping/server/intetnal/app/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRouters struct{}

func (s *UserRouters) InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	userRouter := Router.Group("v1")
	userApi := v1.ApiGroupApp.AdminApiGroup.UserApi
	{
		userRouter.POST("login", userApi.Login)
		userRouter.POST("register", userApi.Register)
		userRouter.GET("list", userApi.ManagerList)
		userRouter.POST("captcha", userApi.Captcha)
	}
	return userRouter
}
