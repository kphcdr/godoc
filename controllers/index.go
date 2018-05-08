package controllers

import (
	"github.com/astaxie/beego"
)

// 用户模块
type IndexController struct {
	beego.Controller
}


func (this *IndexController) Index() {

	this.Redirect("/web",301)

	return
}
