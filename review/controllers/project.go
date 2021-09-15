package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
	"review/logger"
	"review/models"
	"review/request"
	"review/response"
)

type ProjectController struct {
	beego.Controller
}

// respondJson: do response for project operations
func (p *ProjectController) respondJson(httpCode int, opCode int, message string, data ...interface{}) {
	p.Ctx.Output.SetStatus(httpCode)
	var d interface{}
	if len(data) == 1 {
		d = data[0]
	} else {
		d = data
	}
	resp := response.GenResponse(opCode, message, d)
	p.Data["json"] = resp
	p.ServeJSON()
}

// @Title CreateEmptyProject
// @Description 创建一个空项目(不创建相关流程和任务)
// @Param   token header string true "user token get at login"
// @Param   json body request.CreateProject true "基本的项目信息, 创建人(creator)一项不需要填写,会根据token自动解析填充"
// @Success 200 {object} response.Default
// @Failure 400 "invalid project json body(token)"
// @router / [post]
func (p *ProjectController) CreateEmptyProject() {
	var req request.CreateProject
	err := unmarshalBody(p.Ctx.Input.RequestBody, &req)
	if err != nil {
		p.respondJson(http.StatusBadRequest, response.FAIL, "parse body (project json) failed")
		return
	}
	// get user id from token
	CreatorId, err := parseUserToken(p.Ctx.Request.Header["Token"][0])
	if err != nil {
		logger.Recorder.Warning("[request token] invalid request token: " + err.Error())
		p.respondJson(http.StatusBadRequest, response.FAIL, "invalid token")
		return
	}
	req.UserId = CreatorId
	projUuid, code := models.CreateEmptyProject(&req)
	p.respondJson(http.StatusOK, code, "", projUuid)
	return
}

// @Title CreatTemplateProject
// @Description 创建一个模板项目(创建标准的7个流程和任务)
// @Param   token header string true "user token get at login"
// @Param   json body request.CreateProject true "基本的项目信息, 创建人(creator)一项不需要填写,会根据token自动解析填充"
// @Success 200 {object} response.Default
// @Failure 400 "invalid project json body(token)"
// @router /template [post]
func (p *ProjectController) CreatTemplateProject() {
	var req request.CreateProject
	err := unmarshalBody(p.Ctx.Input.RequestBody, &req)
	if err != nil {
		p.respondJson(http.StatusBadRequest, response.FAIL, "parse body (project json) failed")
		return
	}
	// get user id from token
	CreatorId, err := parseUserToken(p.Ctx.Request.Header["Token"][0])
	if err != nil {
		logger.Recorder.Warning("[request token] invalid request token: " + err.Error())
		p.respondJson(http.StatusBadRequest, response.FAIL, "invalid token")
		return
	}
	req.UserId = CreatorId
	projectUuid, code := models.CreateTemplateProject(&req)
	p.respondJson(http.StatusOK, code, "", projectUuid)
	return
}

// @Title UpdateProjectInfo
// @Description 更新项目相关信息
// @Param   token header string true "user token get at login"
// @Param   json body request.UpdateProjectInfo true 要更新的项目信息数据
// @Success 200 {object} response.Default
// @Failure 400 "invalid body"
// @router / [put]
func (p *ProjectController) UpdateProjectInfo() {
	var req request.UpdateProjectInfo
	err := unmarshalBody(p.Ctx.Input.RequestBody, &req)
	if err != nil {
		p.respondJson(http.StatusBadRequest, response.FAIL, "parse body (json) failed")
		return
	}
	code := models.UpdateProjectInfo(&req)
	p.respondJson(http.StatusOK, code, "")
	return
}

// @Title GetBasicInfo
// @Description 获取项目的基本信息数据
// @Param   token header string true "user token get at login"
// @Param   pid path string true "项目的uuid"
// @Success 200 {object} response.Default
// @Failure 400 "invalid project id"
// @router /basic/:pid [get]
func (p *ProjectController) GetBasicInfo() {
	pid := p.GetString(":pid")
	if pid == "" {
		p.respondJson(http.StatusBadRequest, response.FAIL, "invalid project id")
		return
	}
	proj, code := models.GetProjectBasicInfo(pid)
	p.respondJson(http.StatusOK, code, "", proj)
	return
}

// @Title GetDetailedInfo
// @Description 获取项目的详细信息数据(basic info 基本信息, group-人员情况, steps-项目所有流程信息, materials-项目使用的参考材料)
// @Param   token header string true "user token get at login"
// @Param   pid path string true "项目的uuid"
// @Success 200 {object} response.Default
// @Failure 400 "invalid project id"
// @router /detailed/:pid [get]
func (p *ProjectController) GetDetailedInfo() {
	pid := p.GetString(":pid")
	if pid == "" {
		p.respondJson(http.StatusBadRequest, response.FAIL, "invalid project id")
		return
	}
	proj, code := models.GetProjectDetailedInfo(pid)
	p.respondJson(http.StatusOK, code, "", proj)
	return
}
