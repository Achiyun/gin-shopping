package admin

import (
	"fmt"

	"github.com/Achiyun/gin-shopping/server/intetnal/app/global"
	models "github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin"
	"github.com/Achiyun/gin-shopping/server/pkg/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

type FocusController struct {
	BaseController
}

func (con FocusController) Index(c *gin.Context) {
	focusList := []models.Focus{}
	global.GVA_DB.Find(&focusList)
	c.HTML(http.StatusOK, "admin/focus/index.html", gin.H{
		"focusList": focusList,
	})

}
func (con FocusController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}
func (con FocusController) DoAdd(c *gin.Context) {

	title := c.PostForm("title")
	focusType, err1 := utils.Int(c.PostForm("focus_type"))
	link := c.PostForm("link")
	sort, err2 := utils.Int(c.PostForm("sort"))
	status, err3 := utils.Int(c.PostForm("status"))

	if err1 != nil || err3 != nil {
		con.Error(c, "非法请求", "/admin/focus/add")
	}
	if err2 != nil {
		con.Error(c, "请输入正确的排序值", "/admin/focus/add")
	}
	//上传文件
	focusImgSrc, err4 := utils.UploadImg(c, "focus_img")
	if err4 != nil {
		fmt.Println(err4)
	}

	focus := models.Focus{
		Title:     title,
		FocusType: focusType,
		FocusImg:  focusImgSrc,
		Link:      link,
		Sort:      sort,
		Status:    status,
		AddTime:   int(utils.GetUnix()),
	}
	err5 := global.GVA_DB.Create(&focus).Error
	if err5 != nil {
		con.Error(c, "增加轮播图失败", "/admin/focus/add")
	} else {
		con.Success(c, "增加轮播图成功", "/admin/focus")
	}

}

func (con FocusController) Edit(c *gin.Context) {
	id, err1 := utils.Int(c.Query("id"))
	if err1 != nil {
		con.Error(c, "传入参数错误", "/admin/focus")
		return
	}
	focus := models.Focus{Id: id}
	global.GVA_DB.Find(&focus)
	c.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{
		"focus": focus,
	})
}
func (con FocusController) DoEdit(c *gin.Context) {
	id, err1 := utils.Int(c.PostForm("id"))
	title := c.PostForm("title")
	focusType, err2 := utils.Int(c.PostForm("focus_type"))
	link := c.PostForm("link")
	sort, err3 := utils.Int(c.PostForm("sort"))
	status, err4 := utils.Int(c.PostForm("status"))

	if err1 != nil || err2 != nil || err4 != nil {
		con.Error(c, "非法请求", "/admin/focus")
	}
	if err3 != nil {
		con.Error(c, "请输入正确的排序值", "/admin/focus/edit?id="+utils.String(id))
	}
	//上传文件
	focusImg, _ := utils.UploadImg(c, "focus_img")

	focus := models.Focus{Id: id}
	global.GVA_DB.Find(&focus)
	focus.Title = title
	focus.FocusType = focusType
	focus.Link = link
	focus.Sort = sort
	focus.Status = status
	if focusImg != "" {
		focus.FocusImg = focusImg
	}
	err5 := global.GVA_DB.Save(&focus).Error
	if err5 != nil {
		con.Error(c, "修改数据失败请重新尝试", "/admin/focus/edit?id="+utils.String(id))
	} else {
		con.Success(c, "增加轮播图成功", "/admin/focus")
	}
}

func (con FocusController) Delete(c *gin.Context) {
	id, err := utils.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/focus")
	} else {
		focus := models.Focus{Id: id}
		global.GVA_DB.Delete(&focus)
		//根据自己的需要 要不要删除图片
		// os.Remove("static/upload/20210915/1631694117.jpg")
		con.Success(c, "删除数据成功", "/admin/focus")
	}
}
