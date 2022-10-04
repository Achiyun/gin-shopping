package adminRouter

import (
	con "github.com/Achiyun/gin-shopping/server/intetnal/app/html/controllers"
	"github.com/gin-gonic/gin"
)

type GoodsRouters struct{}

func (s *RoleRouters) InitGoodsRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	goodsRouter := adminRouter.Group("goods")
	goodsCon := con.ConGroupApp.AdminConGroup.GoodsController
	{

		goodsRouter.GET("/", goodsCon.Index)
		goodsRouter.GET("add", goodsCon.Add)
		goodsRouter.POST("doAdd", goodsCon.DoAdd)
		goodsRouter.GET("edit", goodsCon.Edit)
		goodsRouter.POST("doEdit", goodsCon.DoEdit)
		goodsRouter.GET("delete", goodsCon.Delete)

	}
	return goodsRouter
}
