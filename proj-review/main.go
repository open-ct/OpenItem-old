package main

import (
	_ "proj-review/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

