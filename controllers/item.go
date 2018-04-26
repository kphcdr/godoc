package controllers

import (
	"github.com/astaxie/beego"
	"showdoc/consts"
)

// Operations about Users
type ItemController struct {
	beego.Controller
}

// @Title MyList
// @Description mylist
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /myList [get]
func (u *ItemController) MyList() {
	json := consts.Json{}
	var data [0]int
	json.Set(0,"获取成功")
	json.SetData(data)
	u.Data["json"] = json.VendorOk()
	u.ServeJSON()
}

// @Title add item
// @Description add item
// @Param   item_type     formData    int  false        "项目类型 1常规项目  2单页项目"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /add [post]
func (u *ItemController) Add() {
	/**
	item_type: 1
item_name:
password:
item_domain:
copy_item_id:
item_description:
	 */
	json := consts.Json{}
	var data [0]int
	json.Set(0,"获取成功")
	json.SetData(data)
	u.Data["json"] = json.VendorOk()
	u.ServeJSON()
}
