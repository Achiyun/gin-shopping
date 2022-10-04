package admin

import (
	"encoding/json"
	"net/http"

	"github.com/Achiyun/gin-shopping/server/intetnal/app/global"
	models "github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin"
	"github.com/Achiyun/gin-shopping/server/intetnal/app/model/common/response"
	"github.com/Achiyun/gin-shopping/server/pkg/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type MainController struct{}

func (con MainController) Index(c *gin.Context) {
	//获取userinfo 对应的session
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")
	//类型断言 来判断 userinfo是不是一个string
	userinfoStr, ok := userinfo.(string)
	if ok {
		//1、获取用户信息
		var managerinfo models.Manager
		err := json.Unmarshal([]byte(userinfoStr), &managerinfo)
		if err != nil {
			response.FailWithMessage("managerinfo转化为json失败", c)
			return
		}

		//2、获取所有的权限
		accessList := []models.Access{}
		roleAccess := []models.RoleAccess{}

		accessList, roleAccess, err = mainService.Index(&managerinfo, accessList, roleAccess)

		if err != nil {
			response.FailWithMessage("权限表查询失败", c)
			return
		}
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
		c.HTML(http.StatusOK, "admin/main/index.html", gin.H{
			"username":   managerinfo.Username,
			"accessList": accessList,
			"isSuper":    managerinfo.IsSuper,
		})
	} else {
		c.Redirect(302, "/admin/user/login")
	}

}

func (con MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}

// 公共修改状态的方法
func (con MainController) ChangeStatus(c *gin.Context) {
	id, err := utils.Int(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "传入的参数错误",
		})
		return
	}

	table := c.Query("table")
	field := c.Query("field")

	// status = ABS(0-1)   1

	// status = ABS(1-1)  0

	err1 := global.GVA_DB.Exec("update "+table+" set "+field+"=ABS("+field+"-1) where id=?", id).Error
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改失败 请重试",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "修改成功",
	})
}

// 公共修改状态的方法
func (con MainController) ChangeNum(c *gin.Context) {
	id, err := utils.Int(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "传入的参数错误",
		})
		return
	}

	table := c.Query("table")
	field := c.Query("field")
	num := c.Query("num")

	err1 := global.GVA_DB.Exec("update "+table+" set "+field+"="+num+" where id=?", id).Error
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改数据失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "修改成功",
		})
	}

}
