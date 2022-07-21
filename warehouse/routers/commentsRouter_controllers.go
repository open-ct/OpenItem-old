package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["warehouse/controllers:QueryController"] = append(beego.GlobalControllerRouter["warehouse/controllers:QueryController"],
		beego.ControllerComments{
			Method:           "GetFinalQuestionList",
			Router:           "/f_question",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:QueryController"] = append(beego.GlobalControllerRouter["warehouse/controllers:QueryController"],
		beego.ControllerComments{
			Method:           "GetFinalTestPaperList",
			Router:           "/f_testpaper",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:QueryController"] = append(beego.GlobalControllerRouter["warehouse/controllers:QueryController"],
		beego.ControllerComments{
			Method:           "GetTempQuestionList",
			Router:           "/t_question",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:QueryController"] = append(beego.GlobalControllerRouter["warehouse/controllers:QueryController"],
		beego.ControllerComments{
			Method:           "GetTempTestPaperList",
			Router:           "/t_testpaper",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:QuestionController"] = append(beego.GlobalControllerRouter["warehouse/controllers:QuestionController"],
		beego.ControllerComments{
			Method:           "CreateNewQuestion",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:QuestionController"] = append(beego.GlobalControllerRouter["warehouse/controllers:QuestionController"],
		beego.ControllerComments{
			Method:           "UpdateQuestion",
			Router:           "/",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:QuestionController"] = append(beego.GlobalControllerRouter["warehouse/controllers:QuestionController"],
		beego.ControllerComments{
			Method:           "AddQuestionComment",
			Router:           "/comment",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:QuestionController"] = append(beego.GlobalControllerRouter["warehouse/controllers:QuestionController"],
		beego.ControllerComments{
			Method:           "FinishTempQuestion",
			Router:           "/finish/:qid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:QuestionController"] = append(beego.GlobalControllerRouter["warehouse/controllers:QuestionController"],
		beego.ControllerComments{
			Method:           "GetProjectFinalQuestions",
			Router:           "/proj_f/:pid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:QuestionController"] = append(beego.GlobalControllerRouter["warehouse/controllers:QuestionController"],
		beego.ControllerComments{
			Method:           "GetProjectTempQuestions",
			Router:           "/proj_t/:pid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:QuestionController"] = append(beego.GlobalControllerRouter["warehouse/controllers:QuestionController"],
		beego.ControllerComments{
			Method:           "TraceQuestionVersion",
			Router:           "/trace/:qid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:QuestionController"] = append(beego.GlobalControllerRouter["warehouse/controllers:QuestionController"],
		beego.ControllerComments{
			Method:           "GetUserFinalQuestions",
			Router:           "/user_f/:uid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:QuestionController"] = append(beego.GlobalControllerRouter["warehouse/controllers:QuestionController"],
		beego.ControllerComments{
			Method:           "GetUserTempQuestions",
			Router:           "/user_t/:uid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:TestpaperController"] = append(beego.GlobalControllerRouter["warehouse/controllers:TestpaperController"],
		beego.ControllerComments{
			Method:           "CreateNewTestpaper",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:TestpaperController"] = append(beego.GlobalControllerRouter["warehouse/controllers:TestpaperController"],
		beego.ControllerComments{
			Method:           "UpdateTestpaper",
			Router:           "/",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:TestpaperController"] = append(beego.GlobalControllerRouter["warehouse/controllers:TestpaperController"],
		beego.ControllerComments{
			Method:           "AddTestpaperComment",
			Router:           "/comment",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:TestpaperController"] = append(beego.GlobalControllerRouter["warehouse/controllers:TestpaperController"],
		beego.ControllerComments{
			Method:           "FinishTempTestpaper",
			Router:           "/finish/:qid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:TestpaperController"] = append(beego.GlobalControllerRouter["warehouse/controllers:TestpaperController"],
		beego.ControllerComments{
			Method:           "GetProjectFinalTestpaper",
			Router:           "/proj_f/:pid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:TestpaperController"] = append(beego.GlobalControllerRouter["warehouse/controllers:TestpaperController"],
		beego.ControllerComments{
			Method:           "GetProjectTempTestpaper",
			Router:           "/proj_t/:pid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:TestpaperController"] = append(beego.GlobalControllerRouter["warehouse/controllers:TestpaperController"],
		beego.ControllerComments{
			Method:           "TraceTestpaperVersion",
			Router:           "/trace/:qid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:TestpaperController"] = append(beego.GlobalControllerRouter["warehouse/controllers:TestpaperController"],
		beego.ControllerComments{
			Method:           "GetUserFinalTestpaper",
			Router:           "/user_f/:uid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:TestpaperController"] = append(beego.GlobalControllerRouter["warehouse/controllers:TestpaperController"],
		beego.ControllerComments{
			Method:           "GetUserTempTestpaper",
			Router:           "/user_t/:uid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:UserController"] = append(beego.GlobalControllerRouter["warehouse/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:UserController"] = append(beego.GlobalControllerRouter["warehouse/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:UserController"] = append(beego.GlobalControllerRouter["warehouse/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:UserController"] = append(beego.GlobalControllerRouter["warehouse/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:UserController"] = append(beego.GlobalControllerRouter["warehouse/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:UserController"] = append(beego.GlobalControllerRouter["warehouse/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           "/login",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["warehouse/controllers:UserController"] = append(beego.GlobalControllerRouter["warehouse/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Logout",
			Router:           "/logout",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
