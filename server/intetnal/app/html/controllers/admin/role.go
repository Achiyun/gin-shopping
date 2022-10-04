package admin

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Achiyun/gin-shopping/server/intetnal/app/global"
	models "github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin"
	"github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin/request"
	"github.com/Achiyun/gin-shopping/server/intetnal/app/model/common/response"
	"github.com/Achiyun/gin-shopping/server/pkg/utils"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	BaseController
}

func (con RoleController) List(c *gin.Context) {

	roleList := []models.Role{}
	err := roleService.List(&roleList)
	if err != nil {
		con.ErrorPlus(c, err, "admin/")
	}
	c.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"roleList": roleList,
	})

}
func (con RoleController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}
func (con RoleController) DoAdd(c *gin.Context) {

	var m request.Role

	err := c.ShouldBind(&m)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 先去除空格再校验
	m.Title = strings.Trim(m.Title, " ")
	m.Description = strings.Trim(m.Description, " ")

	err = utils.Verify(m, utils.RoleDoAddVerify)
	if err != nil {
		con.Error(c, "角色的标题不能为空", "/admin/role/add")
		return
	}

	role := models.Role{
		Title:       m.Title,
		Description: m.Description,
		Status:      1,
		AddTime:     int(utils.GetUnix()),
	}

	err = roleService.Create(&role)
	if err != nil {
		con.Error(c, "增加角色失败 请重试", "/admin/role/add")
		return
	}
	con.Success(c, "增加角色成功", "/admin/role")

}
func (con RoleController) Edit(c *gin.Context) {

	var m request.Role

	err := c.ShouldBind(&m)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	id, err := utils.Int(m.Id)
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}

	role := models.Role{Id: id}
	err = roleService.Edit(&role)
	if err != nil {
		con.Error(c, "查找数据错误", "/admin/role")
		return
	}
	c.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
		"role": role,
	})

}
func (con RoleController) DoEdit(c *gin.Context) {

	var m request.Role

	err := c.ShouldBind(&m)
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}
	// 格式化数据
	m.Title = strings.Trim(m.Title, " ")
	m.Description = strings.Trim(m.Description, " ")

	err = utils.Verify(m, utils.RoleDoAddVerify)
	if err != nil {
		con.Error(c, "角色的标题不能为空", "/admin/role/add")
		return
	}

	id, err := utils.Int(c.PostForm("id"))
	if err != nil {
		con.Error(c, "int转string错误", "/admin/role")
		return
	}

	role := models.Role{
		Id:          id,
		Title:       m.Title,
		Description: m.Description,
	}

	err = global.GVA_DB.Save(&role).Error
	if err != nil {
		con.Error(c, "修改数据失败", "/admin/role/edit?id="+utils.String(id))
		return
	}
	con.Success(c, "修改数据成功", "/admin/role/edit?id="+utils.String(id))

}
func (con RoleController) Delete(c *gin.Context) {
	var m request.Role
	err := c.ShouldBind(&m)

	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}

	id, err := utils.Int(m.Id)
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}

	role := models.Role{Id: id}
	err = roleService.Delete(&role)

	if err != nil {
		con.Error(c, "删除数据失败", "/admin/role")
		return
	}
	con.Success(c, "删除数据成功", "/admin/role")
}

func (con RoleController) Auth(c *gin.Context) {
	//1、获取角色id
	var m request.Role
	err := c.ShouldBind(&m)

	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}

	roleId, err := utils.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}
	//2、获取所有的权限
	accessList := []models.Access{}

	
	global.GVA_DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)

	//3、获取当前角色拥有的权限 ，并把权限id放在一个map对象里面
	roleAccess := []models.RoleAccess{}
	global.GVA_DB.Where("role_id=?", roleId).Find(&roleAccess)
	roleAccessMap := make(map[int]int)
	for _, v := range roleAccess {
		roleAccessMap[v.AccessId] = v.AccessId
	}

	//4、循环遍历所有的权限数据，判断当前权限的id是否在角色权限的Map对象中,如果是的话给当前数据加入checked属性

	for i := 0; i < len(accessList); i++ {
		if _, ok := roleAccessMap[accessList[i].Id]; ok {
			accessList[i].Checked = true
		}
		for j := 0; j < len(accessList[i].AccessItem); j++ {
			if _, ok := roleAccessMap[accessList[i].AccessItem[j].Id]; ok {
				accessList[i].AccessItem[j].Checked = true
			}
		}
	}

	c.HTML(http.StatusOK, "admin/role/auth.html", gin.H{
		"roleId":     roleId,
		"accessList": accessList,
	})

}

func (con RoleController) DoAuth(c *gin.Context) {
	//获取角色id
	roleId, err1 := utils.Int(c.PostForm("role_id"))
	if err1 != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}
	//获取权限id  切片
	accessIds := c.PostFormArray("access_node[]")

	//删除当前角色对应的权限
	roleAccess := models.RoleAccess{}
	global.GVA_DB.Where("role_id=?", roleId).Delete(&roleAccess)

	//增加当前角色对应的权限
	for _, v := range accessIds {
		roleAccess.RoleId = roleId
		accessId, _ := utils.Int(v)
		roleAccess.AccessId = accessId
		global.GVA_DB.Create(&roleAccess)
	}
	fmt.Println(roleId)
	fmt.Println(accessIds)

	fmt.Println("/admin/role/auth?id=?" + utils.String(roleId))
	con.Success(c, "授权成功", "/admin/role/auth?id="+utils.String(roleId))
}
