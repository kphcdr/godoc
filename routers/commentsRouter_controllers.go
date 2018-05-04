package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["showdoc/controllers:CatalogController"] = append(beego.GlobalControllerRouter["showdoc/controllers:CatalogController"],
		beego.ControllerComments{
			Method: "CatListGroup",
			Router: `/catListGroup`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["showdoc/controllers:CatalogController"] = append(beego.GlobalControllerRouter["showdoc/controllers:CatalogController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["showdoc/controllers:CatalogController"] = append(beego.GlobalControllerRouter["showdoc/controllers:CatalogController"],
		beego.ControllerComments{
			Method: "Save",
			Router: `/save`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["showdoc/controllers:ItemController"] = append(beego.GlobalControllerRouter["showdoc/controllers:ItemController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["showdoc/controllers:ItemController"] = append(beego.GlobalControllerRouter["showdoc/controllers:ItemController"],
		beego.ControllerComments{
			Method: "Info",
			Router: `/info`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["showdoc/controllers:ItemController"] = append(beego.GlobalControllerRouter["showdoc/controllers:ItemController"],
		beego.ControllerComments{
			Method: "MyList",
			Router: `/myList`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["showdoc/controllers:UserController"] = append(beego.GlobalControllerRouter["showdoc/controllers:UserController"],
		beego.ControllerComments{
			Method: "Info",
			Router: `/info`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["showdoc/controllers:UserController"] = append(beego.GlobalControllerRouter["showdoc/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["showdoc/controllers:UserController"] = append(beego.GlobalControllerRouter["showdoc/controllers:UserController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
