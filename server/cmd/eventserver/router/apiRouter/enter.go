package apiRouter

import v1 "github.com/Achiyun/gin-shopping/server/cmd/eventserver/router/apiRouter/v1"

type RouterGroup struct {
	V1Group v1.RouterV1Group
}

var ApiGroupApp = new(RouterGroup)
