package controllers

import (
	"github.com/astaxie/beego"
	"showdoc/consts"
	"showdoc/models"
)

// 分类模块
type CatalogController struct {
	beego.Controller
}

// @Title  删除分类
// @Description 分类列表
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.catalogs
// @router /delete [post]
func (this *CatalogController) Delete() {
	json := consts.Json{}
	id,_ := this.GetInt64("cat_id")
	uid := this.GetSession(consts.SESSION_UID)
	if uid == nil {
		this.Abort("403")
	} else {
		ret,catalogs := models.GetOneCataLogs(id)
		if ret {
			 err := catalogs.Delete()
			 if err == nil {
			 	json.Set(0,"删除成功")
			 } else {
			 	json.Set(404,"删除失败")
			 }
		} else {
			json.Set(404,"数据不存在")
		}
		this.Data["json"] = json.VendorOk()

		this.ServeJSON()
	}
}

// @Title  分类列表
// @Description 分类列表
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.catalogs
// @router /catListGroup [post]
func (this *CatalogController) CatListGroup() {
	/**
	{"error_code":0,"data":[{"cat_id":"3","cat_name":"tt","item_id":"3","s_number":"99","addtime":"2018-05-04 13:43:34","parent_cat_id":"0","level":"2","sub":[]}]}
	 */
	json := consts.Json{}
	id,_ := this.GetInt64("item_id")
	uid := this.GetSession(consts.SESSION_UID)
	if uid == nil {
		this.Abort("403")
	} else {
		catalogs := models.GetCatalogsByItemId(id)
		json.SetData(catalogs)
		this.Data["json"] = json.VendorOk()
		this.ServeJSON()
	}
}

// @Title  分类列表
// @Description 分类列表
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.catalogs
// @router /save [post]
func (this *CatalogController) Save() {
//item_id: 2
//cat_id: undefined
//parent_cat_id: undefined
//cat_name: test
//s_number: 123
	json := consts.Json{}

	uid := this.GetSession(consts.SESSION_UID)
	if uid == nil {
		this.Abort("403")
	} else {
		item_id,_ := this.GetInt("item_id")
		cat_name := this.GetString("cat_name")
		s_number,_ := this.GetInt("s_number")
		cat_id,_  := this.GetInt64("cat_id")
		parent_cat_id,_ := this.GetInt("parent_cat_id")
		var catalog models.Catalogs

		catalog.Name = cat_name
		catalog.SNumber = s_number
		catalog.ItemId = item_id
		catalog.Id = cat_id

		var parent_id int

		if parent_cat_id == 0 {
			parent_id = 0
		} else {
			parent_id = 1
		}
		catalog.ParentCatId = parent_cat_id
		catalog.Level = parent_id
		catalog.Save()

		json.SetData(catalog)
		this.Data["json"] = json.VendorOk()
		this.ServeJSON()
	}
}