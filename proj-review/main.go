package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "proj-review/database"
	_ "proj-review/routers"
)

func main() {
	beego.Run()
}
