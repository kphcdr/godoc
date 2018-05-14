package controllers

import (
	"github.com/astaxie/beego"
	"showdoc/consts"
	"showdoc/models"
)

// 用户模块
type TemplateController struct {
	beego.Controller
}

// @Title 保存模板
// @Description user login
// @Param	username		formData 	string	true		"用户名，这里是邮箱"
// @Param	password		formData 	string	true		"用户密码"
// @Success 200 {int} models.User.Id
// @router /save [post]
func (this *TemplateController) Save() {
	uid,err := consts.IsLogin(this.Controller)
	if err != nil {
		this.Abort("403")
	} else {

		title :=this.GetString("template_title")
		content :=this.GetString("template_content")

		var template models.Template
		template.Title = title
		template.Content = content
		template.UserId = uid
		println(title,content)
		template.Save()

		json := consts.Json{}
		json.SetData(template)
		this.Data["json"] = json.VendorOk()
		this.ServeJSON()
	}
}



// @Title 获取列表
// @Description user login
// @Param	username		formData 	string	true		"用户名，这里是邮箱"
// @Param	password		formData 	string	true		"用户密码"
// @Success 200 {int} models.User.Id
// @router /getList [post]
func (this *TemplateController) GetList() {
	uid,err := consts.IsLogin(this.Controller)
	if err != nil {
		this.Abort("403")
	} else {


		template := models.GetTemplateByUid(uid)

		json := consts.Json{}
		json.SetData(template)
		this.Data["json"] = json.VendorOk()
		this.ServeJSON()
	}
}

// @Title 删除模板
// @Description user login
// @Param	username		formData 	string	true		"用户名，这里是邮箱"
// @Param	password		formData 	string	true		"用户密码"
// @Success 200 {int} models.User.Id
// @router /delete [post]
	func (this *TemplateController) Delete() {
		json := consts.Json{}
		id,_ := this.GetInt("id")
		_,err := consts.IsLogin(this.Controller)
		if err != nil {
			this.Abort("403")
		} else {

			ret,template := models.GetOneTemplate(id)
			if ret {
				err := template.Delete()
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