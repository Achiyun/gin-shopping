package eventserver

import (
	Router "github.com/Achiyun/gin-shopping/server/cmd/eventserver/router"
	"github.com/Achiyun/gin-shopping/server/cmd/eventserver/templates"
	"github.com/Achiyun/gin-shopping/server/intetnal/app/global"
	"github.com/Achiyun/gin-shopping/server/intetnal/app/initialize"
	"github.com/Achiyun/gin-shopping/server/intetnal/pkg/core"
	"github.com/gin-gonic/gin"
)

var GinEngine *gin.Engine

func Init() {
	GinEngine = gin.Default()
	global.GVA_VP = core.Viper()      // 初始化Viper, 读取.yaml配置文件
	global.GVA_DB = initialize.Gorm() // gorm连接数据库

	templates.Init(GinEngine)
	// 初始化路由
	Router.Router(GinEngine)
	// todo: 读取配置文件设置端口
	GinEngine.Run(":8082")
}
