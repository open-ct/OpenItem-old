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
}
