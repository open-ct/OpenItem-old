package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["review/controllers:ObjectController"] = append(beego.GlobalControllerRouter["review/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ObjectController"] = append(beego.GlobalControllerRouter["review/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ObjectController"] = append(beego.GlobalControllerRouter["review/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/:objectId",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ObjectController"] = append(beego.GlobalControllerRouter["review/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:objectId",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ObjectController"] = append(beego.GlobalControllerRouter["review/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:objectId",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:UserController"] = append(beego.GlobalControllerRouter["review/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:UserController"] = append(beego.GlobalControllerRouter["review/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:UserController"] = append(beego.GlobalControllerRouter["review/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:UserController"] = append(beego.GlobalControllerRouter["review/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:UserController"] = append(beego.GlobalControllerRouter["review/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:UserController"] = append(beego.GlobalControllerRouter["review/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           "/login",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:UserController"] = append(beego.GlobalControllerRouter["review/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Logout",
			Router:           "/logout",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
