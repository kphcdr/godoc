package controllers

import (
	"github.com/astaxie/beego"
	"showdoc/consts"
	"showdoc/models"
	"strings"
	"beego_blog/util"
)

// 分类模块
type PageController struct {
	beego.Controller
}

// @Title  分类列表
// @Description 分类列表
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.page
// @router /uploadImg [post]
func (this *PageController) UploadImg() {

	var json consts.Json
	var ext string
	var filename string
	f, h, _ := this.GetFile("myfile")                  //获取上传的文件

	ext = h.Filename[strings.LastIndex( h.Filename, "."):]
	filename = util.UniqueId() + ext
	path := "./web/upload/" + filename  //文件目录
	f.Close()                                          //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	this.SaveToFile("myfile", path)                    //存文件
	println(path)
	this.Data["json"] = json.VendorOk()
	this.ServeJSON()

}

// @Title  分类列表
// @Description 分类列表
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.catalogs
// @router /save [post]
func (this *PageController) Save() {
	/**
	page_id=0&
	item_id=2&
	s_number=&
	page_title=tes&
	page_content=%E6%AC%A2%E8%BF%8E%E4%BD&
	cat_id=0

	 */
	uid := this.GetSession(consts.SESSION_UID)
	if uid == nil {
		this.Abort("403")
	} else {

		id,_ :=this.GetInt("page_id")
		item_id,_ :=this.GetInt("item_id")
		s_number,_ := this.GetInt("s_number")
		page_title := this.GetString("page_title")
		page_content := this.GetString("page_content")
		cat_id,_ := this.GetInt("cat_id")


		var page models.Page
		page.Id = id
		page.ItemId = item_id
		page.SNumber = s_number
		page.PageTitle = page_title
		page.PageContent = page_content
		page.CatId = cat_id

		page.SavePage()

		json := consts.Json{}
		json.SetData(page)
		this.Data["json"] = json.VendorOk()
		this.ServeJSON()
	}


}

// @Title  分类列表
// @Description 分类列表
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.catalogs
// @router /info [post]
func (this *PageController) Info() {
	json := consts.Json{}
	id,_ := this.GetInt("page_id")
	//uid := this.GetSession(consts.SESSION_UID)

	_,item := models.GetOnePage(id)

	json.SetData(item)
	this.Data["json"] = json.VendorOk()
	this.ServeJSON()

}
