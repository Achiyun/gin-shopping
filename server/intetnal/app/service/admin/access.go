package admin

import (
	"github.com/Achiyun/gin-shopping/server/intetnal/app/global"
	models "github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin"
)

type AccessService struct {
}

func (role *AccessService) List(m *[]models.Access) (err error) {

	err = global.GVA_DB.Where("module_id=?", 0).Preload("AccessItem").Find(&m).Error

	return err
}

func (role *AccessService) Create(m *models.Access) (err error) {

	err = global.GVA_DB.Create(&m).Error
	return err
}

func (role *AccessService) Edit(m *models.Access) (err error) {

	err = global.GVA_DB.Find(&m).Error
	return err
}
func (role *AccessService) Save(m *models.Access) (err error) {

	err = global.GVA_DB.Save(&m).Error

	return err
}
func (role *AccessService) Delete(m *models.Access) (err error) {

	err = global.GVA_DB.Delete(&m).Error

	return err
}
