package router

import (
	"fmt"

	grouter "github.com/Achiyun/gin-shopping/server/cmd/eventserver/router/groupRouter"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	fmt.Println("Initialization router")

	grouter.AdminRouters(r)
}
