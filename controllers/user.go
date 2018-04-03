package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
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
// @router /captch [get]
func (u *UserController) Captch() {
	store := cache.NewMemoryCache()
	urlPrefix := "/captcha/"
	cpt := captcha.NewWithFilter(urlPrefix, store)


	str,_ := cpt.CreateCaptcha()
	img := "http://localhost:8080" + urlPrefix +str + ".png"
	json := map[string]string{"img": img}

	u.Data["json"] = json
	u.ServeJSON()
	 //= urlPrefix + cpt_result + ".png"


}
