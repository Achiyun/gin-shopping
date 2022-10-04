package admin

import (
	"github.com/Achiyun/gin-shopping/server/intetnal/app/global"
	models "github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin"
	"github.com/Achiyun/gin-shopping/server/pkg/utils"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type GoodsTypeAttributeController struct {
	BaseController
}

func (con GoodsTypeAttributeController) Index(c *gin.Context) {

	cateId, err := utils.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入的参数不正确", "/admin/goodsType")
		return
	}
	//获取商品类型属性
	goodsTypeAttributeList := []models.GoodsTypeAttribute{}
	global.GVA_DB.Where("cate_id=?", cateId).Find(&goodsTypeAttributeList)
	//获取商品类型属性对应的类型

	goodsType := models.GoodsType{}
	global.GVA_DB.Where("id=?", cateId).Find(&goodsType)

	c.HTML(http.StatusOK, "admin/goodsTypeAttribute/index.html", gin.H{
		"cateId":                 cateId,
		"goodsTypeAttributeList": goodsTypeAttributeList,
		"goodsType":              goodsType,
	})

}
func (con GoodsTypeAttributeController) Add(c *gin.Context) {
	//获取当前商品类型属性对应的类型id

	cateId, err := utils.Int(c.Query("cate_id"))
	if err != nil {
		con.Error(c, "传入的参数不正确", "/admin/goodsType")
		return
	}

	//获取所有的商品类型
	goodsTypeList := []models.GoodsType{}
	global.GVA_DB.Find(&goodsTypeList)
	c.HTML(http.StatusOK, "admin/goodsTypeAttribute/add.html", gin.H{
		"goodsTypeList": goodsTypeList,
		"cateId":        cateId,
	})
}

func (con GoodsTypeAttributeController) DoAdd(c *gin.Context) {

	title := strings.Trim(c.PostForm("title"), " ")
	cateId, err1 := utils.Int(c.PostForm("cate_id"))
	attrType, err2 := utils.Int(c.PostForm("attr_type"))
	attrValue := c.PostForm("attr_value")
	sort, err3 := utils.Int(c.PostForm("sort"))

	if err1 != nil || err2 != nil {
		con.Error(c, "非法请求", "/admin/goodsType")
		return
	}
	if title == "" {
		con.Error(c, "商品类型属性名称不能为空", "/admin/goodsTypeAttribute/add?cate_id="+utils.String(cateId))
		return
	}

	if err3 != nil {
		con.Error(c, "排序值不对", "/admin/goodsTypeAttribute/add?cate_id="+utils.String(cateId))
		return
	}

	goodsTypeAttr := models.GoodsTypeAttribute{
		Title:     title,
		CateId:    cateId,
		AttrType:  attrType,
		AttrValue: attrValue,
		Status:    1,
		Sort:      sort,
		AddTime:   int(utils.GetUnix()),
	}
	err := global.GVA_DB.Create(&goodsTypeAttr).Error
	if err != nil {
		con.Error(c, "增加商品类型属性失败 请重试", "/admin/goodsTypeAttribute/add?cate_id="+utils.String(cateId))
	} else {
		con.Success(c, "增加商品类型属性成功", "/admin/goodsTypeAttribute?id="+utils.String(cateId))
	}

}

func (con GoodsTypeAttributeController) Edit(c *gin.Context) {

	//获取当前要修改数据的id
	id, err := utils.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入的参数不正确", "/admin/goodsType")
		return
	}
	//获取当前id对应的商品类型属性
	goodsTypeAttribute := models.GoodsTypeAttribute{Id: id}
	global.GVA_DB.Find(&goodsTypeAttribute)

	//获取所有的商品类型
	goodsTypeList := []models.GoodsType{}
	global.GVA_DB.Find(&goodsTypeList)

	c.HTML(http.StatusOK, "admin/goodsTypeAttribute/edit.html", gin.H{
		"goodsTypeAttribute": goodsTypeAttribute,
		"goodsTypeList":      goodsTypeList,
	})
}

func (con GoodsTypeAttributeController) DoEdit(c *gin.Context) {
	id, err1 := utils.Int(c.PostForm("id"))
	title := strings.Trim(c.PostForm("title"), " ")
	cateId, err2 := utils.Int(c.PostForm("cate_id"))
	attrType, err3 := utils.Int(c.PostForm("attr_type"))
	attrValue := c.PostForm("attr_value")
	sort, err4 := utils.Int(c.PostForm("sort"))

	if err1 != nil || err2 != nil || err3 != nil {
		con.Error(c, "非法请求", "/admin/goodsType")
		return
	}
	if title == "" {
		con.Error(c, "商品类型属性名称不能为空", "/admin/goodsTypeAttribute/edit?id="+utils.String(id))
		return
	}
	if err4 != nil {
		con.Error(c, "排序值不对", "/admin/goodsTypeAttribute/edit?id="+utils.String(id))
		return
	}

	goodsTypeAttr := models.GoodsTypeAttribute{Id: id}
	global.GVA_DB.Find(&goodsTypeAttr)
	goodsTypeAttr.Title = title
	goodsTypeAttr.CateId = cateId
	goodsTypeAttr.AttrType = attrType
	goodsTypeAttr.AttrValue = attrValue
	goodsTypeAttr.Sort = sort
	err := global.GVA_DB.Save(&goodsTypeAttr).Error
	if err != nil {
		con.Error(c, "修改数据失败", "/admin/goodsTypeAttribute/edit?id="+utils.String(id))
		return
	}
	con.Success(c, "需改数据成功", "/admin/goodsTypeAttribute?id="+utils.String(cateId))
}

func (con GoodsTypeAttributeController) Delete(c *gin.Context) {
	id, err1 := utils.Int(c.Query("id"))
	cateId, err2 := utils.Int(c.Query("cate_id"))
	if err1 != nil || err2 != nil {
		con.Error(c, "传入参数错误", "/admin/goodsType")
	} else {
		goodsTypeAttr := models.GoodsTypeAttribute{Id: id}
		global.GVA_DB.Delete(&goodsTypeAttr)
		con.Success(c, "删除数据成功", "/admin/goodsTypeAttribute?id="+utils.String(cateId))
	}
}
