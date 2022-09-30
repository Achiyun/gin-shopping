// 对后台管理员的操作
package grouprouter

import (
	v1 "github.com/Achiyun/gin-shopping/server/intetnal/app/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRouters struct{}

func (s *UserRouters) InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	userRouter := Router.Group("user")
	userApi := v1.ApiGroupApp.AdminApiGroup.UserApi
	{
		userRouter.POST("login", userApi.Login)
		userRouter.POST("register", userApi.Register)
		userRouter.POST("captcha", userApi.Captcha)
	}
	return userRouter
}
