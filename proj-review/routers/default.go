package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"proj-review/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// 开发测试使用
	beego.Router("/api/doc", &controllers.MainController{}, "get:Doc")
	beego.Router("/api/demo", &controllers.MainController{}, "get:Demo")
}
