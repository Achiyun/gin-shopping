package v1

import "github.com/Achiyun/gin-shopping/server/intetnal/app/api/v1/admin"

type ApiGroup struct {
	AdminApiGroup admin.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
