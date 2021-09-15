package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "warehouse/logger"
	_ "warehouse/models"
	_ "warehouse/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
