package adminRouter

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	UserRouters
	MainRouters
	ManagerRouters
}

var (
	adminRouter   *gin.RouterGroup
	AdminGroupApp = new(RouterGroup)
)

func Init(r *gin.Engine) {
	// PrivateGroup
	adminRouter = r.Group("admin")
	PublicGroup := r.Group("")
	{
		AdminGroupApp.InitUserRouter(PublicGroup)
		AdminGroupApp.InitMainRouter(PublicGroup)
		AdminGroupApp.InitManagerRouter(PublicGroup)
	}

}
