package controllers

import (
	"net/http"
	"review/logger"
	"review/models"
	"review/request"
	"review/response"
)

// a part of project controller

// @Title GetOneSubmit
// @Description 获取一个submit的信息
// @Param   token header string true "user token get at login"
// @Param   submitId path string true "要获取的submit uuid"
// @Success 200 {object} response.Default
// @Failure 400 "invalid submit id"
// @router /submit/:submitId [get]
func (p *ProjectController) GetOneSubmit() {
	submitId := p.GetString(":submitId")
	if submitId == "" {
		p.respondJson(http.StatusBadRequest, response.FAIL, "invalid id")
		return
	}
	resp, code := models.GetOneSubmit(submitId)
	p.respondJson(http.StatusOK, code, "", resp)
	return
}

// @Title GetSubmitInStep
// @Description 获取一个step下的所有submit
// @Param   token header string true "user token get at login"
// @Param   stepId path string true "step的uuid"
// @Success 200 {object} response.Default
// @Failure 400 "invalid step id"
// @router /submits/:stepId [get]
func (p *ProjectController) GetSubmitsInStep() {
	stepId := p.GetString(":stepId")
	if stepId == "" {
		p.respondJson(http.StatusBadRequest, response.FAIL, "invalid id")
		return
	}
	resp, code := models.GetStepSubmits(stepId)
	p.respondJson(http.StatusOK, code, "", resp)
	return
}

// @Title GetUserSubmitInStep
// @Description 获取某个用户在指定step下的submit
// @Param   token header string true "user token get at login"
// @Param   json body request.GetUserSubmitsInStep true "用户&step信息"
// @Success 200 {object} response.Default
// @Failure 400 "invalid json"
// @router /submits/user [post]
func (p *ProjectController) GetUserSubmitInStep() {
	var req request.GetUserSubmitsInStep
	err := unmarshalBody(p.Ctx.Input.RequestBody, &req)
	if err != nil {
		p.respondJson(http.StatusBadRequest, response.FAIL, "parse body failed")
		return
	}
	resp, code := models.GetUserSubmitsInStep(&req)
	p.respondJson(http.StatusOK, code, "", resp)
	return
}

// @Title MakeOneSubmit
// @Description 创建一个新的submit
// @Param   token header string true "user token get at login"
// @Param   json body request.CreateSubmit true "新submit信息"
// @Success 200 {object} response.Default
// @Failure 400 "invalid json"
// @router /submit [post]
func (p *ProjectController) MakeOneSubmit() {
	var req request.CreateSubmit
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
	req.UserId = user
	resp, code := models.MakeOneSubmit(&req)
	p.respondJson(http.StatusOK, code, "", resp)
	return
}

// @Title AppendContentInStep
// @Description 在一个step中的content下追加新的材料(即用户在上传材料审核的历史记录)
// @Param   token header string true "user token get at login"
// @Param   json body request.AppendContentInSubmit true "上传的材料信息"
// @Success 200 {object} response.Default
// @Failure 400 "invalid json"
// @router /submit/content [post]
func (p *ProjectController) AppendContentInStep() {
	var req request.AppendContentInSubmit
	err := unmarshalBody(p.Ctx.Input.RequestBody, &req)
	if err != nil {
		p.respondJson(http.StatusBadRequest, response.FAIL, "parse body failed")
		return
	}
	resp, code := models.AppendContent(&req)
	p.respondJson(http.StatusOK, code, "", resp)
	return
}

// @Title WithdrawContentInStep
// @Description 用户撤回某次提交的材料审核
// @Param   token header string true "user token get at login"
// @Param   json body request.WithdrawContentInSubmit true "撤回的信息"
// @Success 200 {object} response.Default
// @Failure 400 "invalid json"
// @router /submit/content [delete]
func (p *ProjectController) WithdrawContentInStep() {
	var req request.WithdrawContentInSubmit
	err := unmarshalBody(p.Ctx.Input.RequestBody, &req)
	if err != nil {
		p.respondJson(http.StatusBadRequest, response.FAIL, "parse body failed")
		return
	}
	resp, code := models.WithdrawContent(&req)
	p.respondJson(http.StatusOK, code, "", resp)
	return
}

// @Title SetSubmitStatus
// @Description 更改提交的状态(即管理员最终审核某次提交是否最终通过)
// @Param   token header string true "user token get at login"
// @Param   json body request.SetSubmitStatus true "设定的状态"
// @Success 200 {object} response.Default
// @Failure 400 "invalid json"
// @router /submit [put]
func (p *ProjectController) SetSubmitStatus() {
	var req request.SetSubmitStatus
	err := unmarshalBody(p.Ctx.Input.RequestBody, &req)
	if err != nil {
		p.respondJson(http.StatusBadRequest, response.FAIL, "parse body failed")
		return
	}
	code := models.SetSubmitStatus(&req)
	p.respondJson(http.StatusOK, code, "")
	return
}

// @Title DeleteSubmit
// @Description 删除一次submit
// @Param   token header string true "user token get at login"
// @Param   submitId path string true "要删除的submit的uuid"
// @Success 200 {object} response.Default
// @Failure 400 "invalid submit id"
// @router /submit/:submitId [delete]
func (p *ProjectController) DeleteSubmit() {
	submitId := p.GetString(":submitId")
	if submitId == "" {
		p.respondJson(http.StatusBadRequest, response.FAIL, "invalid id")
		return
	}
	code := models.DeleteSubmit(submitId)
	p.respondJson(http.StatusOK, code, "")
	return
}
