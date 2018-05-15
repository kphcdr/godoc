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
			Method: "ChildCatList",
			Router: `/childCatList`,
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
			Method: "GetDefaultCat",
			Router: `/getDefaultCat`,
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

	beego.GlobalControllerRouter["showdoc/controllers:CatalogController"] = append(beego.GlobalControllerRouter["showdoc/controllers:CatalogController"],
		beego.ControllerComments{
			Method: "SecondCatList",
			Router: `/secondCatList`,
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
			Method: "Delete",
			Router: `/delete`,
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

	beego.GlobalControllerRouter["showdoc/controllers:PageController"] = append(beego.GlobalControllerRouter["showdoc/controllers:PageController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["showdoc/controllers:PageController"] = append(beego.GlobalControllerRouter["showdoc/controllers:PageController"],
		beego.ControllerComments{
			Method: "Info",
			Router: `/info`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["showdoc/controllers:PageController"] = append(beego.GlobalControllerRouter["showdoc/controllers:PageController"],
		beego.ControllerComments{
			Method: "Save",
			Router: `/save`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["showdoc/controllers:PageController"] = append(beego.GlobalControllerRouter["showdoc/controllers:PageController"],
		beego.ControllerComments{
			Method: "UploadImg",
			Router: `/uploadImg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["showdoc/controllers:TemplateController"] = append(beego.GlobalControllerRouter["showdoc/controllers:TemplateController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["showdoc/controllers:TemplateController"] = append(beego.GlobalControllerRouter["showdoc/controllers:TemplateController"],
		beego.ControllerComments{
			Method: "GetList",
			Router: `/getList`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["showdoc/controllers:TemplateController"] = append(beego.GlobalControllerRouter["showdoc/controllers:TemplateController"],
		beego.ControllerComments{
			Method: "Save",
			Router: `/save`,
			AllowHTTPMethods: []string{"post"},
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
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["showdoc/controllers:UserController"] = append(beego.GlobalControllerRouter["showdoc/controllers:UserController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["showdoc/controllers:UserController"] = append(beego.GlobalControllerRouter["showdoc/controllers:UserController"],
		beego.ControllerComments{
			Method: "ResetPassword",
			Router: `/resetPassword`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
