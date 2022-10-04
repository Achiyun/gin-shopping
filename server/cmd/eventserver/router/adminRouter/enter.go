package adminRouter

import (
	"github.com/Achiyun/gin-shopping/server/intetnal/app/middlewares"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	UserRouters
	MainRouters
	ManagerRouters
	RoleRouters
	AccessRouters
	FocusRouters
	GoodsRouters
	GoodsCateRouters
	GoodsTypeRouters
	GoodsTypeAttributeRouters
	NavRouters
	SettingRouters
}

var (
	adminRouter   *gin.RouterGroup
	AdminGroupApp = new(RouterGroup)
)

func Init(r *gin.Engine) {
	// PrivateGroup
	adminRouter = r.Group("admin", middlewares.InitAdminAuthMiddleware)
	PublicGroup := r.Group("")
	{
		AdminGroupApp.InitUserRouter(PublicGroup)
		AdminGroupApp.InitMainRouter(PublicGroup)
		AdminGroupApp.InitManagerRouter(PublicGroup)
		AdminGroupApp.InitRoleRouter(PublicGroup)
		AdminGroupApp.InitAccessRouter(PublicGroup)
		AdminGroupApp.InitFocusRouter(PublicGroup)
		AdminGroupApp.InitGoodsRouter(PublicGroup)
		AdminGroupApp.InitGoodsCateRouter(PublicGroup)
		AdminGroupApp.InitGoodsTypeRouter(PublicGroup)
		AdminGroupApp.InitGoodsTypeAttributeRouter(PublicGroup)
		AdminGroupApp.InitNavRouter(PublicGroup)
		AdminGroupApp.InitSettingRouter(PublicGroup)

	}

}
