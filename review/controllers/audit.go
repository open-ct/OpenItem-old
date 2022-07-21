package controllers

import (
	"net/http"
	"review/logger"
	"review/models"
	"review/request"
	"review/response"
)

// a part of project controller

// @Title GetOneAudit
// @Description 获取项目管理员/负责人等对一个submit下的content做出的审核
// @Param   token header string true "user token get at login"
// @Param   auditId path string true "audit id"
// @Success 200 {object} response.Default
// @Failure 400 "invalid audit id"
// @router /audit/:auditId [get]
func (p *ProjectController) GetOneAudit() {
	auditId := p.GetString(":auditId")
	if auditId == "" {
		p.respondJson(http.StatusBadRequest, response.FAIL, "invalid id")
		return
	}
	resp, code := models.GetOneAudit(auditId)
	p.respondJson(http.StatusOK, code, "", resp)
	return
}

// @Title GetAuditsInSubmit
// @Description 获取一个submit下的所有audit
// @Param   token header string true "user token get at login"
// @Param   submitId path string true "submit id"
// @Success 200 {object} response.Default
// @Failure 400 "invalid submit uuid"
// @router /audits/:submitId [get]
func (p *ProjectController) GetAuditsInSubmit() {
	submitId := p.GetString(":submitId")
	if submitId == "" {
		p.respondJson(http.StatusBadRequest, response.FAIL, "invalid id")
		return
	}
	resp, code := models.GetSubmitAudits(submitId)
	p.respondJson(http.StatusOK, code, "", resp)
	return
}

// @Title VreateOneAudit
// @Description 项目负责人创建一个材料审核记录
// @Param   token header string true "user token get at login"
// @Param   json body request.MakeOneAudit true "审核信息"
// @Success 200 {object} response.Default
// @Failure 400 "invalid json"
// @router /audit [post]
func (p *ProjectController) CreateOneAudit() {
	var req request.MakeOneAudit
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
	resp, code := models.MakeOneAudit(&req)
	p.respondJson(http.StatusOK, code, "", resp)
	return
}

// @Title CorrectAudit
// @Description 负责人修改审核信息
// @Param   token header string true "user token get at login"
// @Param   json body request.UpdateAudit true "修改后的审核信息"
// @Success 200 {object} response.Default
// @Failure 400 "invalid json"
// @router /audit [put]
func (p *ProjectController) CorrectAudit() {
	var req request.UpdateAudit
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
	req.NewAuditor = user
	resp, code := models.CorrectAudit(&req)
	p.respondJson(http.StatusOK, code, "", resp)
	return
}

// @Title DeleteAudit
// @Description 删除一条审核记录
// @Param   token header string true "user token get at login"
// @Param   auditId path string true "要深处的审核记录Id"
// @Success 200 {object} response.Default
// @Failure 400 "invalid audit id"
// @router /audit [delete]
func (p *ProjectController) DeleteAudit() {
	auditId := p.GetString(":auditId")
	if auditId == "" {
		p.respondJson(http.StatusBadRequest, response.FAIL, "invalid id")
		return
	}
	code := models.DeleteAudit(auditId)
	p.respondJson(http.StatusOK, code, "")
	return
}
