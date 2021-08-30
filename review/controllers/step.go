package controllers

import (
	"net/http"
	"review/logger"
	"review/models"
	"review/request"
	"review/response"
)

// continue with projectController

// @Title CreateOneStep
// @Description 在指定项目下创建一个流程
// @Param   token header string true "user token get at login"
// @Param   json body request.CreateStep true "新建步骤的基本信息结构"
// @Success 200 {object} response.Default
// @Failure 400 "invalid step json body"
// @router /step [post]
func (p *ProjectController) CreateOneStep() {
	var req request.CreateStep
	err := unmarshalBody(p.Ctx.Input.RequestBody, &req)
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
	req.Creator = creator
	resp, code := models.CreateOneStep(&req)
	p.respondJson(http.StatusOK, code, "", resp)
	return
}

// @Title GetOneStepInfo
// @Description 获取某一个指定id的step信息
// @Param   token header string true "user token get at login"
// @Param   stepId path string true "对应step的uuid"
// @Success 200 {object} response.Default
// @Failure 400 "invalid stepID"
// @router /step/:stepId [get]
func (p *ProjectController) GetOneStepInfo() {
	stepId := p.GetString(":stepId")
	if stepId == "" {
		p.respondJson(http.StatusBadRequest, response.FAIL, "invalid id")
		return
	}
	resp, code := models.GetStepInfo(stepId)
	p.respondJson(http.StatusOK, code, "", resp)
	return
}

// @Title GetStepsInProject
// @Description 获取一个项目下的所有step信息
// @Param   token header string true "user token get at login"
// @Param   pid path string true "指定的项目uuid"
// @Success 200 {object} response.Default
// @Failure 400 "invalid project id"
// @router /steps/:pid [get]
func (p *ProjectController) GetStepsInProject() {
	pId := p.GetString(":pid")
	if pId == "" {
		p.respondJson(http.StatusBadRequest, response.FAIL, "invalid id")
		return
	}
	resp, code := models.GetAllStepsInProject(pId)
	p.respondJson(http.StatusOK, code, "", resp)
	return
}

// @Title UploadStepAttachment
// @Description 绑定某个step的附件(这里并不上传文件, 需要调用uploadFile相关接口上传后绑定文件uuid到指定step)
// @Param   token header string true "user token get at login"
// @Param   json body request.AddStepAttachment true "附件信息"
// @Success 200 {object} response.Default
// @Failure 400 "invalid attachment json"
// @router /step/attachment [post]
func (p *ProjectController) UploadStepAttachment() {
	var req request.AddStepAttachment
	err := unmarshalBody(p.Ctx.Input.RequestBody, &req)
	if err != nil {
		p.respondJson(http.StatusBadRequest, response.FAIL, "parse body failed")
		return
	}
	user, err := parseUserToken(p.Ctx.Request.Header["Token"][0])
	if err != nil {
		logger.Recorder.Warning("[user token] parse user token error: " + err.Error())
		p.respondJson(http.StatusBadRequest, response.FAIL, "invalid token")
		return
	}
	req.Uploader = user
	code := models.UploadStepAttachments(&req)
	p.respondJson(http.StatusOK, code, "")
	return
}

// @Title UpdateStepInfo
// @Description 更新某个step的信息
// @Param   token header string true "user token get at login"
// @Param   json body request.UpdateStepInfo true "要更新的信息"
// @Success 200 {object} response.Default
// @Failure 400 "invalid json"
// @router /step [put]
func (p *ProjectController) UpdateStepInfo() {
	var req request.UpdateStepInfo
	err := unmarshalBody(p.Ctx.Input.RequestBody, &req)
	if err != nil {
		p.respondJson(http.StatusBadRequest, response.FAIL, "parse body failed")
		return
	}
	code := models.UpdateStepInfo(&req)
	p.respondJson(http.StatusOK, code, "")
	return
}

// @Title SetStepStatus
// @Description 更改流程的进度状态
// @Param   token header string true "user token get at login"
// @Param   json body request.SetStepStatus true "新的状态信息"
// @Success 200 {object} response.Default
// @Failure 400 "invalid json body"
// @router /step/status [put]
func (p *ProjectController) SetStepStatus() {
	var req request.SetStepStatus
	err := unmarshalBody(p.Ctx.Input.RequestBody, &req)
	if err != nil {
		p.respondJson(http.StatusBadRequest, response.FAIL, "parse body failed")
		return
	}
	code := models.SetStepStatus(&req)
	p.respondJson(http.StatusOK, code, "")
	return
}

// @Title SetStepTimePoint
// @Description 为step设置时间点
// @Param   token header string true "user token get at login"
// @Param   json body request.SetStepTimePoint true "时间点信息"
// @Success 200 {object} response.Default
// @Failure 400 "invalid json"
// @router /step/timepoint [put]
func (p *ProjectController) SetStepTimePoint() {
	var req request.SetStepTimePoint
	err := unmarshalBody(p.Ctx.Input.RequestBody, &req)
	if err != nil {
		p.respondJson(http.StatusBadRequest, response.FAIL, "parse body failed")
		return
	}
	resp, code := models.SetStepTimePoint(&req)
	p.respondJson(http.StatusOK, code, "", resp)
	return
}

// @Title DeleteStepTimePoint
// @Description 删除step下的某个时间点
// @Param   token header string true "user token get at login"
// @Param   json body request.DeleteStepTimePoint true "要删除的时间点信息"
// @Success 200 {object} response.Default
// @Failure 400 "invalid jsob"
// @router /step/timepoint [delete]
func (p *ProjectController) DeleteStepTimePoint() {
	var req request.DeleteStepTimePoint
	err := unmarshalBody(p.Ctx.Input.RequestBody, &req)
	if err != nil {
		p.respondJson(http.StatusBadRequest, response.FAIL, "parse body failed")
		return
	}
	code := models.DeleteStepTimePoint(&req)
	p.respondJson(http.StatusOK, code, "")
	return
}

// @Title DeleteStep
// @Description 删除step
// @Param   token header string true "user token get at login"
// @Param   stepId path string true "要删除的的step的uuid"
// @Success 200 {object} response.Default
// @Failure 400 "invalid step id"
// @router /step/:stepId [delete]
func (p *ProjectController) DeleteStep() {
	stepId := p.GetString(":stepId")
	if stepId == "" {
		p.respondJson(http.StatusBadRequest, response.FAIL, "invalid id")
		return
	}
	code := models.DeleteStep(stepId)
	p.respondJson(http.StatusOK, code, "")
	return
}
