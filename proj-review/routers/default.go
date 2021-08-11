package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"proj-review/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/doc", &controllers.MainController{}, "get:ApiDoc")
}
