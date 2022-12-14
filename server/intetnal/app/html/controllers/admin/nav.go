package admin

import (
	"fmt"

	"github.com/Achiyun/gin-shopping/server/intetnal/app/global"
	models "github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin"
	"github.com/Achiyun/gin-shopping/server/pkg/utils"

	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NavController struct {
	BaseController
}

func (con NavController) Index(c *gin.Context) {
	//当前页
	page, _ := utils.Int(c.Query("page"))
	if page == 0 {
		page = 1
	}
	fmt.Println(page)
	//每页显示的数量
	pageSize := 8
	//获取数据
	navList := []models.Nav{}
	global.GVA_DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&navList)

	//获取总数量
	var count int64
	global.GVA_DB.Table("nav").Count(&count)
	c.HTML(http.StatusOK, "admin/nav/index.html", gin.H{
		"navList": navList,
		//注意float64类型
		"totalPages": math.Ceil(float64(count) / float64(pageSize)),
		"page":       page,
	})
}
func (con NavController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/nav/add.html", gin.H{})
}
func (con NavController) DoAdd(c *gin.Context) {
	title := c.PostForm("title")
	link := c.PostForm("link")
	position, _ := utils.Int(c.PostForm("position"))
	isOpennew, _ := utils.Int(c.PostForm("is_opennew"))
	relation := c.PostForm("relation")
	sort, _ := utils.Int(c.PostForm("sort"))
	status, _ := utils.Int(c.PostForm("status"))
	if title == "" {
		con.Error(c, "标题不能为空", "/admin/nav/add")
		return
	}

	nav := models.Nav{
		Title:     title,
		Link:      link,
		Position:  position,
		IsOpennew: isOpennew,
		Relation:  relation,
		Sort:      sort,
		Status:    status,
		AddTime:   int(utils.GetUnix()),
	}
	err := global.GVA_DB.Create(&nav).Error
	if err != nil {
		con.Error(c, "增加导航失败 请重试", "/admin/nav/add")
	} else {
		con.Success(c, "增加导航成功", "/admin/nav")
	}

}
func (con NavController) Edit(c *gin.Context) {
	id, err := utils.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/nav")
	} else {
		nav := models.Nav{Id: id}
		global.GVA_DB.Find(&nav)
		fmt.Println(nav)
		c.HTML(http.StatusOK, "admin/nav/edit.html", gin.H{
			"nav":      nav,
			"prevPage": c.Request.Referer(), //获取上一页的地址
		})
	}

}
func (con NavController) DoEdit(c *gin.Context) {

	id, err1 := utils.Int(c.PostForm("id"))
	if err1 != nil {
		con.Error(c, "传入数据错误", "/admin/nav")
		return
	}
	//获取上一页的地址
	prevPage := c.PostForm("prevPage")
	title := c.PostForm("title")
	link := c.PostForm("link")
	position, _ := utils.Int(c.PostForm("position"))
	isOpennew, _ := utils.Int(c.PostForm("is_opennew"))
	relation := c.PostForm("relation")
	sort, _ := utils.Int(c.PostForm("sort"))
	status, _ := utils.Int(c.PostForm("status"))
	if title == "" {
		con.Error(c, "标题不能为空", "/admin/nav/add")
		return
	}

	nav := models.Nav{Id: id}
	global.GVA_DB.Find(&nav)
	nav.Title = title
	nav.Link = link
	nav.Position = position
	nav.IsOpennew = isOpennew
	nav.Relation = relation
	nav.Sort = sort
	nav.Status = status
	err2 := global.GVA_DB.Save(&nav).Error
	if err2 != nil {
		con.Error(c, "修改数据失败", "/admin/nav/edit?id="+utils.String(id))
	} else {
		if len(prevPage) > 0 {
			con.Success(c, "修改数据成功", prevPage)
		} else {
			con.Success(c, "修改数据成功", "/admin/nav")
		}
	}

}
func (con NavController) Delete(c *gin.Context) {
	id, err := utils.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/nav")
	} else {
		nav := models.Nav{Id: id}
		global.GVA_DB.Delete(&nav)
		//获取上一页
		prevPage := c.Request.Referer()
		if len(prevPage) > 0 {
			con.Success(c, "删除数据成功", prevPage)
		} else {
			con.Success(c, "删除数据成功", "/admin/nav")
		}
	}
}
