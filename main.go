package main

import (
	_ "showdoc/routers"

	"github.com/astaxie/beego"
	"showdoc/models"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		beego.BConfig.WebConfig.StaticDir["/web"] = "web"
	}
	beego.Run()
}

func init() {
	//数据库
	models.MysqlInit()
	// session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider="file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "storage/session"


	//
	beego.SetLogger("file", `{"filename":"storage/logs/test.log","daily":true}`)
}
