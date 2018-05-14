// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"showdoc/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.IndexController{}, "GET:Index")


	ns := beego.NewNamespace("/v1/api",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/item",
			beego.NSInclude(
				&controllers.ItemController{},
			),
		),
		beego.NSNamespace("/catalog",
			beego.NSInclude(
				&controllers.CatalogController{},
			),
		),
		beego.NSNamespace("/page",
			beego.NSInclude(
				&controllers.PageController{},
			),
		),
		beego.NSNamespace("/template",
			beego.NSInclude(
				&controllers.TemplateController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
