package response

import (
	models "github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin"
)

type AdminUserResponse struct {
	Manager     models.Manager   `json:"manager"`
	ManagerList []models.Manager `json:"managerList"`
}

type LoginResponse struct {
	Manager   models.Manager `json:"manager"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}
