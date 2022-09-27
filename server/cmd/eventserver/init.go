package eventserver

import (
	Router "github.com/Achiyun/gin-shopping/server/cmd/eventserver/router"
	"github.com/gin-gonic/gin"
)

var GinEngine *gin.Engine

func Init() {
	// 初始化路由
	Router.Router(GinEngine)
	// todo: 读取配置文件设置端口
	GinEngine.Run(":8082")
}
