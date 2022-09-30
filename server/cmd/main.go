package main

import (
	"fmt"

	eventserver "github.com/Achiyun/gin-shopping/server/cmd/eventserver"
	"github.com/Achiyun/gin-shopping/server/intetnal/app/global"
	"github.com/Achiyun/gin-shopping/server/intetnal/pkg/core"
	"github.com/gin-gonic/gin"
)

// @title gin-shopping
// @version 1.0
// @description 购物网站
// @termsOfService http://swagger.io/terms/

// @contact.name chenchiyu

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 8082
// @BasePath /
func main() {
	fmt.Println("starting eventserver")
	global.GVA_VP = core.Viper() // 初始化Viper
	fmt.Printf("global.GVA_CONFIG.JWT.SigningKey: %v\n", global.GVA_CONFIG.JWT.SigningKey)

	eventserver.GinEngine = gin.Default()

	eventserver.Init()
}
