package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["showdoc/controllers:ObjectController"] = append(beego.GlobalControllerRouter["showdoc/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["showdoc/controllers:ObjectController"] = append(beego.GlobalControllerRouter["showdoc/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["showdoc/controllers:ObjectController"] = append(beego.GlobalControllerRouter["showdoc/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["showdoc/controllers:ObjectController"] = append(beego.GlobalControllerRouter["showdoc/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["showdoc/controllers:ObjectController"] = append(beego.GlobalControllerRouter["showdoc/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["showdoc/controllers:UserController"] = append(beego.GlobalControllerRouter["showdoc/controllers:UserController"],
		beego.ControllerComments{
			Method: "Info",
			Router: `/info`,
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

}
