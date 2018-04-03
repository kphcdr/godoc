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
	models.MysqlInit()
}
