package adminRouter

import (
	con "github.com/Achiyun/gin-shopping/server/intetnal/app/html/controllers"
	"github.com/gin-gonic/gin"
)

type GoodsTypeAttributeRouters struct{}

func (s *RoleRouters) InitGoodsTypeAttributeRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	goodsTypeAttributeRouter := adminRouter.Group("goodsTypeAttribute")
	goodsTypeAttributeCon := con.ConGroupApp.AdminConGroup.GoodsTypeAttributeController
	{

		goodsTypeAttributeRouter.GET("/", goodsTypeAttributeCon.Index)
		goodsTypeAttributeRouter.GET("add", goodsTypeAttributeCon.Add)
		goodsTypeAttributeRouter.POST("doAdd", goodsTypeAttributeCon.DoAdd)
		goodsTypeAttributeRouter.GET("edit", goodsTypeAttributeCon.Edit)
		goodsTypeAttributeRouter.POST("doEdit", goodsTypeAttributeCon.DoEdit)
		goodsTypeAttributeRouter.GET("delete", goodsTypeAttributeCon.Delete)

	}
	return goodsTypeAttributeRouter
}
