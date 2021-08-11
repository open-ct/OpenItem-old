package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"proj-review/controllers"
)

func init() {
	beego.Router("/user/login", &controllers.UserController{}, "post:UserRegister")

}
