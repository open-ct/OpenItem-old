package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
	"review/logger"
	"review/query"
	"review/response"
)

// Operations about query apis
type QueryController struct {
	beego.Controller
}

// respondJson: do response for user operations
func (q *QueryController) respondJson(httpCode int, opCode int, message string, data ...interface{}) {
	q.Ctx.Output.SetStatus(httpCode)
	var d interface{}
	if len(data) == 1 {
		d = data[0]
	} else {
		d = data
	}
	resp := response.GenResponse(opCode, message, d)
	q.Data["json"] = resp
	q.ServeJSON()
}

type QueryRequest struct {
	IdList []string `json:"id_list"`
}

// router definitions:

// @Title GetUserList
// @Description 根据id列表获取user信息
// @Param   token header string false "user token get at login"
// @Param   json body QueryRequest true "要获取的id 列表"
// @Success 200 {object} response.Default
// @Failure 400 "parse id list error"
// @router /user [get]
func (q *QueryController) GetUserList() {
	queryRequset := new(QueryRequest)
	err := unmarshalBody(q.Ctx.Input.RequestBody, queryRequset)
	if err != nil {
		logger.Recorder.Warning("[query] parse id list error: " + err.Error())
		q.respondJson(http.StatusBadRequest, response.FAIL, "parse id list error")
		return
	}
	resp := query.QueryUsers(queryRequset.IdList)
	q.respondJson(http.StatusOK, response.SUCCESS, "", resp)
	return
}

// @Title GetUserList
// @Description 根据id列表获取user信息
// @Param   token header string false "user token get at login"
// @Param   json body QueryRequest true "要获取的id 列表"
// @Success 200 {object} response.Default
// @Failure 400 "parse id list error"
// @router /proj [get]
func (q *QueryController) GetProjectList() {
	queryRequset := new(QueryRequest)
	err := unmarshalBody(q.Ctx.Input.RequestBody, queryRequset)
	if err != nil {
		logger.Recorder.Warning("[query] parse id list error: " + err.Error())
		q.respondJson(http.StatusBadRequest, response.FAIL, "parse id list error")
		return
	}
	resp := query.QueryProjects(queryRequset.IdList)
	q.respondJson(http.StatusOK, response.SUCCESS, "", resp)
	return
}

// @Title GetAssignmentList
// @Description 根据id列表获取assignment信息
// @Param   token header string false "user token get at login"
// @Param   json body QueryRequest true "要获取的id 列表"
// @Success 200 {object} response.Default
// @Failure 400 "parse id list error"
// @router /assign [get]
func (q *QueryController) GetAssignmentList() {
	queryRequset := new(QueryRequest)
	err := unmarshalBody(q.Ctx.Input.RequestBody, queryRequset)
	if err != nil {
		logger.Recorder.Warning("[query] parse id list error: " + err.Error())
		q.respondJson(http.StatusBadRequest, response.FAIL, "parse id list error")
		return
	}
	resp := query.QueryAssignments(queryRequset.IdList)
	q.respondJson(http.StatusOK, response.SUCCESS, "", resp)
	return
}

// @Title GetFilesList
// @Description 根据id列表获取files信息
// @Param   token header string false "user token get at login"
// @Param   json body QueryRequest true "要获取的id 列表"
// @Success 200 {object} response.Default
// @Failure 400 "parse id list error"
// @router /file [get]
func (q *QueryController) GetFilesList() {
	queryRequset := new(QueryRequest)
	err := unmarshalBody(q.Ctx.Input.RequestBody, queryRequset)
	if err != nil {
		logger.Recorder.Warning("[query] parse id list error: " + err.Error())
		q.respondJson(http.StatusBadRequest, response.FAIL, "parse id list error")
		return
	}
	resp := query.QueryFiles(queryRequset.IdList)
	q.respondJson(http.StatusOK, response.SUCCESS, "", resp)
	return
}

// @Title GetStepList
// @Description 根据id列表获取step信息
// @Param   token header string false "user token get at login"
// @Param   json body QueryRequest true "要获取的id 列表"
// @Success 200 {object} response.Default
// @Failure 400 "parse id list error"
// @router /step [get]
func (q *QueryController) GetStepList() {
	queryRequset := new(QueryRequest)
	err := unmarshalBody(q.Ctx.Input.RequestBody, queryRequset)
	if err != nil {
		logger.Recorder.Warning("[query] parse id list error: " + err.Error())
		q.respondJson(http.StatusBadRequest, response.FAIL, "parse id list error")
		return
	}
	resp := query.QuerySteps(queryRequset.IdList)
	q.respondJson(http.StatusOK, response.SUCCESS, "", resp)
	return
}

// @Title GetSubmitList
// @Description 根据id列表获取submit信息
// @Param   token header string false "user token get at login"
// @Param   json body QueryRequest true "要获取的id 列表"
// @Success 200 {object} response.Default
// @Failure 400 "parse id list error"
// @router /submit [get]
func (q *QueryController) GetSubmitList() {
	queryRequset := new(QueryRequest)
	err := unmarshalBody(q.Ctx.Input.RequestBody, queryRequset)
	if err != nil {
		logger.Recorder.Warning("[query] parse id list error: " + err.Error())
		q.respondJson(http.StatusBadRequest, response.FAIL, "parse id list error")
		return
	}
	resp := query.QuerySubmits(queryRequset.IdList)
	q.respondJson(http.StatusOK, response.SUCCESS, "", resp)
	return
}

// @Title GetAuditList
// @Description 根据id列表获取audit信息
// @Param   token header string false "user token get at login"
// @Param   json body QueryRequest true "要获取的id 列表"
// @Success 200 {object} response.Default
// @Failure 400 "parse id list error"
// @router /audit [get]
func (q *QueryController) GetAuditList() {
	queryRequset := new(QueryRequest)
	err := unmarshalBody(q.Ctx.Input.RequestBody, queryRequset)
	if err != nil {
		logger.Recorder.Warning("[query] parse id list error: " + err.Error())
		q.respondJson(http.StatusBadRequest, response.FAIL, "parse id list error")
		return
	}
	resp := query.QueryAudits(queryRequset.IdList)
	q.respondJson(http.StatusOK, response.SUCCESS, "", resp)
	return
}
