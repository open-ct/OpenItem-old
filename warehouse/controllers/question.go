package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
	"warehouse/logger"
	"warehouse/models"
)

type QuestionController struct {
	beego.Controller
}

// @Title CreateNewQuestion
// @Description 创建新的题目（临时题目）
// @Param   json body models.NewQuestionRequest true "新题目信息"
// @Success 200 {string}
// @Failure 400 "invalid body"
// @router / [post]
func (q *QuestionController) CreateNewQuestion() {
	var request models.NewQuestionRequest
	err := unmarshalBody(q.Ctx.Input.RequestBody, &request)
	if err != nil {
		logger.Recorder.Warning("[question] parse body error: " + err.Error())
		q.Ctx.Output.SetStatus(http.StatusBadRequest)
		q.Data["json"] = map[string]string{
			"msg": err.Error(),
		}
		q.ServeJSON()
		return
	}
	message := "ok"
	resp, err := models.CreateNewTempQuestion(&request)
	if err != nil {
		message = err.Error()
	}
	q.Ctx.Output.SetStatus(http.StatusOK)
	q.Data["json"] = map[string]interface{}{
		"msg":  message,
		"data": resp,
	}
	q.ServeJSON()
	return
}

// @Title UpdateQuestion
// @Description 更新新题目(创建一个新的分支)
// @Param   json body models.UpdateQuestionRequest true "更新的题目信息"
// @Success 200 {string}
// @Failure 400 "invalid body"
// @router / [put]
func (q *QuestionController) UpdateQuestion() {
	var request models.UpdateQuestionRequest
	err := unmarshalBody(q.Ctx.Input.RequestBody, &request)
	if err != nil {
		logger.Recorder.Warning("[question] parse body error: " + err.Error())
		q.Ctx.Output.SetStatus(http.StatusBadRequest)
		q.Data["json"] = map[string]string{
			"msg": err.Error(),
		}
		q.ServeJSON()
		return
	}
	message := "ok"
	resp, err := models.UpdateQuestion(&request)
	if err != nil {
		message = err.Error()
	}
	q.Ctx.Output.SetStatus(http.StatusOK)
	q.Data["json"] = map[string]interface{}{
		"msg":  message,
		"data": resp,
	}
	q.ServeJSON()
	return
}

// @Title AddQuestionComment
// @Description 添加一条题目的评价内容
// @Param   json body models.AddQuestionCommentRequest true "题目评价"
// @Success 200 {string}
// @Failure 400 "invalid body"
// @router /comment [post]
func (q *QuestionController) AddQuestionComment() {
	var request models.AddQuestionCommentRequest
	err := unmarshalBody(q.Ctx.Input.RequestBody, &request)
	if err != nil {
		logger.Recorder.Warning("[question] parse body error: " + err.Error())
		q.Ctx.Output.SetStatus(http.StatusBadRequest)
		q.Data["json"] = map[string]string{
			"msg": err.Error(),
		}
		q.ServeJSON()
		return
	}
	message := "ok"
	err = models.AddQuestionComment(&request)
	if err != nil {
		message = err.Error()
	}
	q.Ctx.Output.SetStatus(http.StatusOK)
	q.Data["json"] = map[string]interface{}{
		"msg": message,
	}
	q.ServeJSON()
	return
}

// @Title TraceQuestionVersion
// @Description 查询历史版本
// @Param   qid path string true "question id"
// @Success 200 {[]models.TempQuestion}
// @Failure 400 "invalid qid"
// @router /trace/:qid [get]
func (q *QuestionController) TraceQuestionVersion() {
	qid := q.GetString(":qid")
	if qid == "" {
		logger.Recorder.Warning("[question] get question id error")
		q.Ctx.Output.SetStatus(http.StatusBadRequest)
		q.Data["json"] = map[string]string{
			"msg": "invalid question id",
		}
		q.ServeJSON()
		return
	}
	message := "ok"
	resp, err := models.TraceQuestionVersion(qid)
	if err != nil {
		message = err.Error()
	}
	q.Ctx.Output.SetStatus(http.StatusOK)
	q.Data["json"] = map[string]interface{}{
		"msg":  message,
		"data": resp,
	}
	q.ServeJSON()
	return
}

// @Title FinishTempQuestion
// @Description 最终确定题目 (转移到final数据库下)
// @Param   qid path string true "question id"
// @Success 200 {string}
// @Failure 400 "invalid qid"
// @router /finish/:qid [get]
func (q *QuestionController) FinishTempQuestion() {
	qid := q.GetString(":qid")
	if qid == "" {
		logger.Recorder.Warning("[question] get question id error")
		q.Ctx.Output.SetStatus(http.StatusBadRequest)
		q.Data["json"] = map[string]string{
			"msg": "invalid question id",
		}
		q.ServeJSON()
		return
	}
	message := "ok"
	resp, err := models.FinishTempQuestion(qid)
	if err != nil {
		message = err.Error()
	}
	q.Ctx.Output.SetStatus(http.StatusOK)
	q.Data["json"] = map[string]interface{}{
		"msg":  message,
		"data": resp,
	}
	q.ServeJSON()
	return
}

// @Title GetUserTempQuestions
// @Description 获取用户创建的题目(temp数据库下)
// @Param   uid path string true "user id"
// @Success 200 {[]models.TempQuestion}
// @Failure 400 "invalid qid"
// @router /user_t/:uid [get]
func (q *QuestionController) GetUserTempQuestions() {
	uid := q.GetString(":uid")
	if uid == "" {
		logger.Recorder.Warning("[question] get user id error")
		q.Ctx.Output.SetStatus(http.StatusBadRequest)
		q.Data["json"] = map[string]string{
			"msg": "invalid user id",
		}
		q.ServeJSON()
		return
	}
	message := "ok"
	resp, err := models.GetUserTempQuestions(uid)
	if err != nil {
		message = err.Error()
	}
	q.Ctx.Output.SetStatus(http.StatusOK)
	q.Data["json"] = map[string]interface{}{
		"msg":  message,
		"data": resp,
	}
	q.ServeJSON()
	return
}

// @Title GetUserFinalQuestions
// @Description 获取用户创建的题目(final数据库下)
// @Param   uid path string true "user id"
// @Success 200 {[]models.FinalQuestion}
// @Failure 400 "invalid qid"
// @router /user_f/:uid [get]
func (q *QuestionController) GetUserFinalQuestions() {
	uid := q.GetString(":uid")
	if uid == "" {
		logger.Recorder.Warning("[question] get user id error")
		q.Ctx.Output.SetStatus(http.StatusBadRequest)
		q.Data["json"] = map[string]string{
			"msg": "invalid user id",
		}
		q.ServeJSON()
		return
	}
	message := "ok"
	resp, err := models.GetUserFinalQuestions(uid)
	if err != nil {
		message = err.Error()
	}
	q.Ctx.Output.SetStatus(http.StatusOK)
	q.Data["json"] = map[string]interface{}{
		"msg":  message,
		"data": resp,
	}
	q.ServeJSON()
	return
}

// @Title GetProjectTempQuestions
// @Description 获取项目下的题目(temp数据库下)
// @Param   uid path string true "project id"
// @Success 200 {[]models.TempQuestion}
// @Failure 400 "invalid qid"
// @router /proj_t/:pid [get]
func (q *QuestionController) GetProjectTempQuestions() {
	pid := q.GetString(":pid")
	if pid == "" {
		logger.Recorder.Warning("[question] get project id error")
		q.Ctx.Output.SetStatus(http.StatusBadRequest)
		q.Data["json"] = map[string]string{
			"msg": "invalid project id",
		}
		q.ServeJSON()
		return
	}
	message := "ok"
	resp, err := models.GetProjectTempQuestions(pid)
	if err != nil {
		message = err.Error()
	}
	q.Ctx.Output.SetStatus(http.StatusOK)
	q.Data["json"] = map[string]interface{}{
		"msg":  message,
		"data": resp,
	}
	q.ServeJSON()
	return
}

// @Title GetProjectFinalQuestions
// @Description 获取项目下的题目(final数据库下)
// @Param   uid path string true "project id"
// @Success 200 {[]models.FinalQuestion}
// @Failure 400 "invalid qid"
// @router /proj_f/:pid [get]
func (q *QuestionController) GetProjectFinalQuestions() {
	pid := q.GetString(":pid")
	if pid == "" {
		logger.Recorder.Warning("[question] get project id error")
		q.Ctx.Output.SetStatus(http.StatusBadRequest)
		q.Data["json"] = map[string]string{
			"msg": "invalid project id",
		}
		q.ServeJSON()
		return
	}
	message := "ok"
	resp, err := models.GetProjectFinalQuestions(pid)
	if err != nil {
		message = err.Error()
	}
	q.Ctx.Output.SetStatus(http.StatusOK)
	q.Data["json"] = map[string]interface{}{
		"msg":  message,
		"data": resp,
	}
	q.ServeJSON()
	return
}
