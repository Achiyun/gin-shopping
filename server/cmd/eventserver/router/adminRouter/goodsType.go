package adminRouter

import (
	con "github.com/Achiyun/gin-shopping/server/intetnal/app/html/controllers"
	"github.com/gin-gonic/gin"
)

type GoodsTypeRouters struct{}

func (s *RoleRouters) InitGoodsTypeRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	goodsTypeRouter := adminRouter.Group("goodsType")
	goodsTypeCon := con.ConGroupApp.AdminConGroup.GoodsTypeController
	{

		goodsTypeRouter.GET("/", goodsTypeCon.Index)
		goodsTypeRouter.GET("add", goodsTypeCon.Add)
		goodsTypeRouter.POST("doAdd", goodsTypeCon.DoAdd)
		goodsTypeRouter.GET("edit", goodsTypeCon.Edit)
		goodsTypeRouter.POST("doEdit", goodsTypeCon.DoEdit)
		goodsTypeRouter.GET("delete", goodsTypeCon.Delete)

	}
	return goodsTypeRouter
}
