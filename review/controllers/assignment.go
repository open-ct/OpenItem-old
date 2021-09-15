package controllers

import (
	"net/http"
	"review/logger"
	"review/models"
	"review/request"
	"review/response"
)

// extension of ProjectController

// @Title MakeOneAssignment
// @Description create a user assignment for a project (创建一个人员分配记录: 项目-用户-角色) json字段说明: operator-进行角色分配的管理员id(根据token自行填充), project_id-该项分配记录对应的项目id, user_id-要进行分配的用户id, role-要分配的角色(1-项目管理员, 2-专家, 3-学科助理, 4-命题教师, 5-外审人员)
// @Param   token header string true "user token get at login"
// @Param   json body request.MakeOneAssignment true "assignment information"
// @Success 200 {object} response.Default
// @Failure 400 "invalid token(body)"
// @router /assign [post]
func (p *ProjectController) MakeOneAssignment() {
	var newAssign request.MakeOneAssignment
	err := unmarshalBody(p.Ctx.Input.RequestBody, &newAssign)
	if err != nil {
		p.respondJson(http.StatusBadRequest, response.FAIL, "parse body failed")
		return
	}
	creator, err := parseUserToken(p.Ctx.Request.Header["Token"][0])
	if err != nil {
		logger.Recorder.Warning("[user token] parse user token error: " + err.Error())
		p.respondJson(http.StatusBadRequest, response.FAIL, "invalid token")
		return
	}
	newAssign.Operator = creator
	resp, code := models.MakeOneAssignment(&newAssign)
	p.respondJson(http.StatusOK, code, "", resp)
	return
}

// @Title MakeAssignmentGroup
// @Description 创建一个项目的人员分配(同时分配多种角色); 字段说明: admins-要分配为管理员的人员id数组, experts-分配为专家的id数组, assistants-分配为学科助理的id, teachers-分配为命题教师的id, out_experts-分配为外审人员的id, project_id-所属项目id, operator-进行分配到管理员id
// @Param   token header string true "user token get at login"
// @Param   json body request.MakeAssignmentGroup true "需要每种角色的user id"
// @Success 200 {object} response.Default
// @Failure 400 "invalid token(body)"
// @router /group [post]
func (p *ProjectController) MakeAssignmentGroup() {
	var assignGroup request.MakeAssignmentGroup
	err := unmarshalBody(p.Ctx.Input.RequestBody, &assignGroup)
	if err != nil {
		p.respondJson(http.StatusBadRequest, response.FAIL, "parse body failed")
		return
	}
	creator, err := parseUserToken(p.Ctx.Request.Header["Token"][0])
	if err != nil {
		logger.Recorder.Warning("[user token] parse user token error: " + err.Error())
		p.respondJson(http.StatusBadRequest, response.FAIL, "invalid token")
		return
	}
	assignGroup.Operator = creator
	code := models.MakeAssignments(&assignGroup)
	p.respondJson(http.StatusOK, code, "")
	return
}

// @Title GetUserAssignment
// @Description 获取某一个用户的所有项目参与情况(即参与各个项目的角色分配情况)
// @Param   token header string true "user token get at login"
// @Param   uid path string true "user id"
// @Success 200 {object} response.Default
// @Failure 400 "invalid user id"
// @router /user/:uid [get]
func (p *ProjectController) GetUserAssignments() {
	uid := p.GetString(":uid")
	if uid == "" {
		p.respondJson(http.StatusBadRequest, response.FAIL, "invalid user id")
		return
	}
	assigns, code := models.GetUserAssignments(uid)
	p.respondJson(http.StatusOK, code, "", assigns)
	return
}

// @Title GetProjectAssignment
// @Description 获取一个项目的所有人员分配情况
// @Param   token header string true "user token get at login"
// @Param   pid path string true "project id"
// @Success 200 {object} response.Default
// @Failure 400 "invalid project id"
// @router /assign/:pid [get]
func (p *ProjectController) GetProjectAssignments() {
	pid := p.GetString(":pid")
	if pid == "" {
		p.respondJson(http.StatusBadRequest, response.FAIL, "invalid project id")
		return
	}
	assigns, code := models.GetProjectAssignment(pid)
	p.respondJson(http.StatusOK, code, "", assigns)
	return
}

// @Title ChangeAssignment
// @Description 更改一个角色分配; 字段说明: operator-进行更改的管理员id(根据token解析), assignment_id-更改的assign id, new_role-新的角色分配,
// @Param   token header string true "user token get at login"
// @Param   json body request.ChangeAssignment true "new role to change"
// @Success 200 {object} response.Default
// @Failure 400 "invalid bodu"
// @router /assign [patch]
func (p *ProjectController) ChangeAssignment() {
	var update request.ChangeAssignment
	err := unmarshalBody(p.Ctx.Input.RequestBody, &update)
	if err != nil {
		p.respondJson(http.StatusBadRequest, response.FAIL, "parse body failed")
		return
	}
	creator, err := parseUserToken(p.Ctx.Request.Header["Token"][0])
	if err != nil {
		logger.Recorder.Warning("[user token] parse user token error: " + err.Error())
		p.respondJson(http.StatusBadRequest, response.FAIL, "invalid token")
		return
	}
	update.Operator = creator
	code := models.ChangeAssignment(&update)
	p.respondJson(http.StatusOK, code, "")
	return
}

// @Title ConfirmAssignment
// @Description 用户端确认角色分配申请
// @Param   token header string true "user token get at login"
// @Param   aid path string true “要确认的addignment 的 uuid"
// @Success 200 {object} response.Default
// @Failure 400 "invalid uuid"
// @router /assign/confirm/:aid [get]
func (p *ProjectController) ConfirmAssignment() {
	aid := p.GetString(":aid")
	if aid == "" {
		p.respondJson(http.StatusBadRequest, response.FAIL, "invalid assignment id")
		return
	}
	code := models.ConfirmAssignment(aid)
	p.respondJson(http.StatusOK, code, "")
	return
}

// @Title DeleteAssignment
// @Description 删除一条角色分配
// @Param   token header string true "user token get at login"
// @Param   aid path string true "uuid of assignment to delete"
// @Success 200 {object} response.Default
// @Failure 400 "invalid uuid"
// @router /assign/:aid [delete]
func (p *ProjectController) DeleteAssignment() {
	aid := p.GetString(":aid")
	if aid == "" {
		p.respondJson(http.StatusBadRequest, response.FAIL, "invalid assignment id")
		return
	}
	code := models.RemoveAssignment(aid)
	p.respondJson(http.StatusOK, code, "")
	return
}
