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
			Method:           "CreateOneAudit",
			Router:           "/audit",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "CorrectAudit",
			Router:           "/audit",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "DeleteAudit",
			Router:           "/audit",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "GetOneAudit",
			Router:           "/audit/:auditId",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "GetAuditsInSubmit",
			Router:           "/audits/:submitId",
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
			Method:           "GetDetailedInfo",
			Router:           "/detailed/:pid",
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
			Method:           "UpdateStepInfo",
			Router:           "/step",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "CreateOneStep",
			Router:           "/step",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "GetOneStepInfo",
			Router:           "/step/:stepId",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "DeleteStep",
			Router:           "/step/:stepId",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "UploadStepAttachment",
			Router:           "/step/attachment",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "GetStepStatisticData",
			Router:           "/step/stat/:stepId",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "SetStepStatus",
			Router:           "/step/status",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "SetStepTimePoint",
			Router:           "/step/timepoint",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "DeleteStepTimePoint",
			Router:           "/step/timepoint",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "GetStepsInProject",
			Router:           "/steps/:pid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "SetSubmitStatus",
			Router:           "/submit",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "MakeOneSubmit",
			Router:           "/submit",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "GetOneSubmit",
			Router:           "/submit/:submitId",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "DeleteSubmit",
			Router:           "/submit/:submitId",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "AppendContentInStep",
			Router:           "/submit/content",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "WithdrawContentInStep",
			Router:           "/submit/content",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "GetSubmitsInStep",
			Router:           "/submits/:stepId",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "GetUserSubmitInStep",
			Router:           "/submits/user",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:ProjectController"] = append(beego.GlobalControllerRouter["review/controllers:ProjectController"],
		beego.ControllerComments{
			Method:           "CreatTemplateProject",
			Router:           "/template",
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

	beego.GlobalControllerRouter["review/controllers:QueryController"] = append(beego.GlobalControllerRouter["review/controllers:QueryController"],
		beego.ControllerComments{
			Method:           "GetAssignmentList",
			Router:           "/assign",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:QueryController"] = append(beego.GlobalControllerRouter["review/controllers:QueryController"],
		beego.ControllerComments{
			Method:           "GetAuditList",
			Router:           "/audit",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:QueryController"] = append(beego.GlobalControllerRouter["review/controllers:QueryController"],
		beego.ControllerComments{
			Method:           "GetFilesList",
			Router:           "/file",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:QueryController"] = append(beego.GlobalControllerRouter["review/controllers:QueryController"],
		beego.ControllerComments{
			Method:           "GetProjectList",
			Router:           "/proj",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:QueryController"] = append(beego.GlobalControllerRouter["review/controllers:QueryController"],
		beego.ControllerComments{
			Method:           "GetStepList",
			Router:           "/step",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:QueryController"] = append(beego.GlobalControllerRouter["review/controllers:QueryController"],
		beego.ControllerComments{
			Method:           "GetSubmitList",
			Router:           "/submit",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["review/controllers:QueryController"] = append(beego.GlobalControllerRouter["review/controllers:QueryController"],
		beego.ControllerComments{
			Method:           "GetUserList",
			Router:           "/user",
			AllowHTTPMethods: []string{"post"},
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
