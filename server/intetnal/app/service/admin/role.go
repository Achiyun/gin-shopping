package admin

import (
	"github.com/Achiyun/gin-shopping/server/intetnal/app/global"
	models "github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin"
)

type RoleService struct {
}

func (role *RoleService) List(m *[]models.Role) (err error) {

	err = global.GVA_DB.Find(&m).Error
	return err
}

func (role *RoleService) Create(m *models.Role) (err error) {

	err = global.GVA_DB.Create(&m).Error
	return err
}

func (role *RoleService) Edit(m *models.Role) (err error) {

	err = global.GVA_DB.Find(&m).Error
	return err
}
func (role *RoleService) Save(m *models.Role) (err error) {

	err = global.GVA_DB.Save(&m).Error

	return err
}
func (role *RoleService) Delete(m *models.Role) (err error) {

	err = global.GVA_DB.Delete(&m).Error

	return err
}
