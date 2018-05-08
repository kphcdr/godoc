package controllers

import (
	"github.com/astaxie/beego"
	"showdoc/consts"
	"showdoc/models"
)

// 用户模块
type UserController struct {
	beego.Controller
}

// @Title 获取当前用户信息
// @Description get user info
// @router /info [post]
func (u *UserController) Info() {
	json := consts.Json{}
	uid := u.GetSession(consts.SESSION_UID)
	if uid == nil {
		json.Set(10000, "用户未登录")
		u.Data["json"] = json.VendorError()
		u.ServeJSON()
	} else {

		ret, user := models.GetOneUser(uid.(int64))
		if ret == true {
			json.SetData(user)
			u.Data["json"] = json.VendorOk()
			u.ServeJSON()
		} else {
			json.Set(10000, "用户未登录")
			u.Data["json"] = json.VendorError()
			u.ServeJSON()
		}
	}

	return

}

// @Title CreateUser
// @Description create users
// @Param	username	formData 	string	true		"用户名，这里是邮箱"
// @Param	password	formData 	string	true		"密码"
// @Param	confirm_password	formData 	string	true		"重复输入密码"
// @Success 200 {int} models.User.Id
// @router /register [post]
func (u *UserController) Register() {
	username := u.GetString("username")
	password := u.GetString("password")
	confirmPassword := u.GetString("confirm_password")
	json := consts.Json{}
	if password != confirmPassword {
		json.Set(10000, "两次密码不一致")
		u.Data["json"] = json.VendorError()
		u.ServeJSON()
		return
	}

	var err error

	var user models.User
	user.Email = username
	user.Password = models.CryptPassword(password)
	user.Id, err = user.Create()
	if err != nil {
		json.Set(500, err.Error())
		u.Data["json"] = json.VendorError()
		u.ServeJSON()
		return
	}

	u.loginSuccess(user)

	json.Set(0, "操作成功")
	u.Data["json"] = json.VendorOk()
	u.ServeJSON()

}

func (u *UserController) loginSuccess(user models.User) {
	u.SetSession(consts.SESSION_UID, user.Id)
}

// @Title Login
// @Description user login
// @Param	username		formData 	string	true		"用户名，这里是邮箱"
// @Param	password		formData 	string	true		"用户密码"
// @Success 200 {int} models.User.Id
// @router /login [post]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")

	password = models.CryptPassword(password)
	ret, user := models.Login(username, password)
	json := consts.Json{}

	if ret == true {
		json.Set(0, user.Email+"登录成功")
		u.loginSuccess(user)
		u.Data["json"] = json.VendorOk()
	} else {
		json.Set(10206, "用户名或密码错误")
		beego.Notice("用户登录失败" + username)
		u.Data["json"] = json.VendorError()
	}

	u.ServeJSON()
}

// @Title Login
// @Description user login
// @Param	username		formData 	string	true		"用户名，这里是邮箱"
// @Param	password		formData 	string	true		"用户密码"
// @Success 200 {int} models.User.Id
// @router /logout [get]
func (this *UserController) Logout() {
	json := consts.Json{}
	this.DestroySession()

	this.Data["json"] = json.VendorOk()
	this.ServeJSON()
}
