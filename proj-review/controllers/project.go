package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
	"proj-review/constant"
	"proj-review/log"
	"proj-review/models"
	"proj-review/request"
	"proj-review/response"
)

type ProjectController struct {
	beego.Controller
}

// respondJson
func (p *ProjectController) respondJson(httpCode int, opCode int, message string, data ...interface{}) {
	p.Ctx.Output.SetStatus(httpCode)
	var d interface{}
	if len(data) == 1 {
		d = data[0]
	} else {
		d = data
	}
	resp := response.GenResponse(
		opCode,
		message,
		d,
	)
	p.Data["json"] = resp
	p.ServeJSON()
}

/*
	Projects Basic operations
*/
// CreateNewProject
func (p *ProjectController) CreateNewProject() {
	createReq := new(request.CreateProject)
	err := unmarshalBody(p.Ctx.Input.RequestBody, createReq)
	if err != nil {
		log.Logger.Warn("[Login] " + err.Error())
		p.respondJson(
			http.StatusOK,
			constant.FAIL,
			"params fail", // todo
			createReq,
		)
		return
	}
	// get creator id:
	creator, err := parseUserToken(p.Ctx.Request.Header["Token"][0])
	if err != nil {
		log.Logger.Warn("[Project] " + err.Error())
		p.respondJson(
			http.StatusOK,
			constant.FAIL,
			"unknown creator, need valid token...", // todo: unknown upload user ... operation denied
		)
		return
	}
	createReq.Creator = creator
	creatResp, ok := models.DoCreateNewProject(createReq)
	if !ok {
		p.respondJson(
			http.StatusOK,
			constant.FAIL,
			"create fail", // todo
			creatResp,
		)
	} else {
		p.respondJson(
			http.StatusOK,
			constant.SUCCESS,
			"create successfully", // todo
			creatResp,
		)
	}
	return
}

// CreateTemplateProject

// GetProjectInfo
func (p *ProjectController) GetProjectInfo() {
	pid := p.GetString("proj_id")
	if pid == "" {
		log.Logger.Warn("[Get Info] invalid project id")
		p.respondJson(
			http.StatusOK,
			constant.FAIL,
			constant.BasicMsg.GetInfoFail,
			pid,
		)
		return
	}
	projInfo, ok := models.DoGetProject(pid)
	if !ok {
		p.respondJson(
			http.StatusOK,
			constant.FAIL,
			constant.BasicMsg.GetInfoFail,
			pid,
		)
	} else {
		p.respondJson(
			http.StatusOK,
			constant.SUCCESS,
			constant.BasicMsg.GetInfoSuccess,
			projInfo,
		)
	}
	return
}

/*
	Assignments
*/
// GetUserAssignments
func (p *ProjectController) GetUserAssignments() {
	uid := p.GetString("user_id")
	if uid == "" {
		if uid == "" {
			log.Logger.Warn("[Get Info] invalid user id")
			p.respondJson(http.StatusOK, constant.FAIL, constant.BasicMsg.GetInfoFail, uid)
			return
		}
	}
	assigns, ok := models.DoGetUserAssignments(uid)
	if !ok {
		p.respondJson(http.StatusOK, constant.FAIL, constant.BasicMsg.GetInfoFail, assigns)
	} else {
		p.respondJson(http.StatusOK, constant.SUCCESS, constant.BasicMsg.GetInfoSuccess, assigns)
	}
	return
}

// GetProjectAssignments
func (p *ProjectController) GetProjectAssignments() {
	pid := p.GetString("proj_id")
	if pid == "" {
		if pid == "" {
			log.Logger.Warn("[Get Info] invalid user id")
			p.respondJson(http.StatusOK, constant.FAIL, constant.BasicMsg.GetInfoFail, pid)
			return
		}
	}
	assigns, ok := models.DoGetProjectAssignments(pid)
	if !ok {
		p.respondJson(http.StatusOK, constant.FAIL, constant.BasicMsg.GetInfoFail, assigns)
	} else {
		p.respondJson(http.StatusOK, constant.SUCCESS, constant.BasicMsg.GetInfoSuccess, assigns)
	}
	return
}
