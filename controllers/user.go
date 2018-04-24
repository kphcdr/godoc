package controllers

import (
	"github.com/astaxie/beego"
	"showdoc/consts"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /info [get]
func (u *UserController) Info() {

	u.Data["json"] = "hahaha";
	u.ServeJSON()
}


// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /register [post]
func (u *UserController) Register() {
	json := consts.Json{}
	json.Set(10206,"验证码不正确")
	u.Data["json"] =  json.VendorOk()
	u.ServeJSON()
}
