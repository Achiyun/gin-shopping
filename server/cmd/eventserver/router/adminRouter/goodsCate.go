package adminRouter

import (
	con "github.com/Achiyun/gin-shopping/server/intetnal/app/html/controllers"
	"github.com/gin-gonic/gin"
)

type GoodsCateRouters struct{}

func (s *RoleRouters) InitGoodsCateRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	goodsCateRouter := adminRouter.Group("goodsCate")
	goodsCateCon := con.ConGroupApp.AdminConGroup.GoodsCateController
	{

		goodsCateRouter.GET("/", goodsCateCon.Index)
		goodsCateRouter.GET("add", goodsCateCon.Add)
		goodsCateRouter.POST("doAdd", goodsCateCon.DoAdd)
		goodsCateRouter.GET("edit", goodsCateCon.Edit)
		goodsCateRouter.POST("doEdit", goodsCateCon.DoEdit)
		goodsCateRouter.GET("delete", goodsCateCon.Delete)

	}
	return goodsCateRouter
}
