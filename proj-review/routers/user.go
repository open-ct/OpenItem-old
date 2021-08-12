package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"proj-review/controllers"
)

func init() {
	// 用户登录 & 注册路由
	beego.Router("/user/register", &controllers.UserController{}, "post:UserRegister")
	beego.Router("/user/login", &controllers.UserController{}, "post:UserLogin")
	// todo: 用户退出登陆操作
	beego.Router("/user/logout", &controllers.UserController{}, "get:UserLogout")
	beego.Router("/user/delete", &controllers.UserController{}, "get:DeleteUser")
	// todp: search user
	beego.Router("/user/search", &controllers.UserController{}, "post:SearchUser")
	beego.Router("/user/getInfo", &controllers.UserController{}, "get:GetUserInfo")
	beego.Router("/user/updateInfo", &controllers.UserController{}, "post:UpdateUserInfo")
	beego.Router("/user/updatePwd", &controllers.UserController{}, "post:UpdateUserPassword")
}
