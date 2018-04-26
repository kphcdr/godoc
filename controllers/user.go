package controllers

import (
	"github.com/astaxie/beego"
	"showdoc/consts"
	"showdoc/models"
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
// @router /register [post]
func (u *UserController) Register() {
//username: kphcdr@163.com
//password: 3551488
//confirm_password: 3551488
	username := u.GetString("username")
	password := u.GetString("password")
	confirmPassword := u.GetString("confirm_password")
	println(username,password,confirmPassword)
	json := consts.Json{}
	if password != confirmPassword {
		json.Set(10000,"两次密码不一致")
		u.Data["json"] =  json.VendorError()
		u.ServeJSON()
		return
	}

	var err error

	user := new(models.User)
	user.Email = username
	user.Password = models.CryptPassword(password)
	user.Id,err = user.Create()
	if err != nil {
		json.Set(500,err.Error())
		u.Data["json"] =  json.VendorError()
		u.ServeJSON()
		return
	}

	u.SetSession("uid",user.Id)

	println("session",u.GetSession("uid"))
	json.Set(0,"操作成功")
	u.Data["json"] = json.VendorOk()
	u.ServeJSON()


}

func (u *UserController) Login() {
	json := consts.Json{}
	json.Set(10206,"验证码不正确")
	u.Data["json"] =  json.VendorOk()
	u.ServeJSON()
}
