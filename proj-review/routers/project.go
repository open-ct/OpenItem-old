package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"proj-review/controllers"
)

func init() {
	// projects
	beego.Router("/proj/createEmpty", &controllers.ProjectController{}, "post:CreateNewProject")
	beego.Router("/proj/getInfo", &controllers.ProjectController{}, "get:GetProjectInfo")
	// assignments
	beego.Router("/proj/userAssign", &controllers.ProjectController{}, "get:GetUserAssignments")
	beego.Router("/proj/projectAssign", &controllers.ProjectController{}, "get:GetProjectAssignments")
	beego.Router("/proj/makeAssign", &controllers.ProjectController{}, "post:MakeNewAssignment")
	beego.Router("/proj/changeAssign", &controllers.ProjectController{}, "post:ChangeAssignment")
	beego.Router("/proj/removeAssign", &controllers.ProjectController{}, "get:RemoveAssignment")

}
