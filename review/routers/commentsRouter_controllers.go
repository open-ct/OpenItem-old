package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["review/controllers:FileController"] = append(beego.GlobalControllerRouter["review/controllers:FileController"],
		beego.ControllerComments{
			Method:           "UploadFile",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:FileController"] = append(beego.GlobalControllerRouter["review/controllers:FileController"],
		beego.ControllerComments{
			Method:           "DownloadFile",
			Router:           "/:fid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:FileController"] = append(beego.GlobalControllerRouter["review/controllers:FileController"],
		beego.ControllerComments{
			Method:           "DeleteFile",
			Router:           "/:fid",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:FileController"] = append(beego.GlobalControllerRouter["review/controllers:FileController"],
		beego.ControllerComments{
			Method:           "GetFileInfo",
			Router:           "/info/:fid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:FileController"] = append(beego.GlobalControllerRouter["review/controllers:FileController"],
		beego.ControllerComments{
			Method:           "SearchFiles",
			Router:           "/search",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "CreateEmptyProject",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "UpdateProjectInfo",
			Router:           "/",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "MakeOneAssignment",
			Router:           "/assign",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "ChangeAssignment",
			Router:           "/assign",
			AllowHTTPMethods: []string{"patch"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "DeleteAssignment",
			Router:           "/assign/:aid",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "GetProjectAssignments",
			Router:           "/assign/:pid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "ConfirmAssignment",
			Router:           "/assign/confirm/:aid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "GetBasicInfo",
			Router:           "/basic/:pid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "MakeAssignmentGroup",
			Router:           "/group",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "GetUserAssignments",
			Router:           "/user/:uid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:UserController"] = append(beego.GlobalControllerRouter["review/controllers:UserController"],
		beego.ControllerComments{
			Method:           "UserRegister",
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
			Method:           "UpdateUserInfo",
			Router:           "/",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:UserController"] = append(beego.GlobalControllerRouter["review/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetUser",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:UserController"] = append(beego.GlobalControllerRouter["review/controllers:UserController"],
		beego.ControllerComments{
			Method:           "UserDelete",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:UserController"] = append(beego.GlobalControllerRouter["review/controllers:UserController"],
		beego.ControllerComments{
			Method:           "UserLogin",
			Router:           "/login",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:UserController"] = append(beego.GlobalControllerRouter["review/controllers:UserController"],
		beego.ControllerComments{
			Method:           "UserLogout",
			Router:           "/logout",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:UserController"] = append(beego.GlobalControllerRouter["review/controllers:UserController"],
		beego.ControllerComments{
			Method:           "UpdateUserPassword",
			Router:           "/password",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
