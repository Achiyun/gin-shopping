package main

import (
	"fmt"

	eventserver "github.com/Achiyun/gin-shopping/server/cmd/eventserver"
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

	eventserver.GinEngine = gin.Default()

	eventserver.Init()

}
