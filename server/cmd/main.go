package main

import (
	"fmt"

	eventserver "github.com/Achiyun/gin-shopping/server/cmd/eventserver"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("starting eventserver")

	eventserver.GinEngine = gin.Default()

	eventserver.Init()
}
