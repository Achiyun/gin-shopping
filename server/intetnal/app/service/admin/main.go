package admin

import (
	"github.com/Achiyun/gin-shopping/server/intetnal/app/global"
	models "github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin"
	"gorm.io/gorm"
)

type MainService struct {
}

func (mainService *MainService) Index(m *models.Manager, a []models.Access, r []models.RoleAccess) (
	accessInter []models.Access, roleAccess []models.RoleAccess, err error) {

	err = global.GVA_DB.Where("module_id=?", 0).Preload("AccessItem", func(db *gorm.DB) *gorm.DB {
		return db.Order("access.sort DESC")
	}).Order("sort DESC").Find(&a).Error

	if err != nil {
		return a, r, err
	}
	err = global.GVA_DB.Where("role_id=?", m.RoleId).Find(&roleAccess).Error

	if err != nil {
		return a, r, err
	}

	return a, r, err

}
