package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
	"warehouse/logger"
	"warehouse/models"
)

type TestpaperController struct {
	beego.Controller
}

// @Title CreateNewTestpaper
// @Description 创建新的题目（临时题目）
// @Param   json body models.NewTestpaperRequest true "新试卷信息"
// @Success 200 {string}
// @Failure 400 "invalid body"
// @router / [post]
func (t *TestpaperController) CreateNewTestpaper() {
	var request models.NewTestpaperRequest
	err := unmarshalBody(t.Ctx.Input.RequestBody, &request)
	if err != nil {
		logger.Recorder.Warning("[test-paper] parse body error: " + err.Error())
		t.Ctx.Output.SetStatus(http.StatusBadRequest)
		t.Data["json"] = map[string]string{
			"msg": err.Error(),
		}
		t.ServeJSON()
		return
	}
	message := "ok"
	resp, err := models.CreateNewTestpaper(&request)
	if err != nil {
		message = err.Error()
	}
	t.Ctx.Output.SetStatus(http.StatusOK)
	t.Data["json"] = map[string]interface{}{
		"msg":  message,
		"data": resp,
	}
	t.ServeJSON()
	return
}

// @Title UpdateTestpaper
// @Description 更新新题目(创建一个新的分支)
// @Param   json body models.UpdateTestpaperRequest true "更新的试卷信息or内容"
// @Success 200 {string}
// @Failure 400 "invalid body"
// @router / [put]
func (t *TestpaperController) UpdateTestpaper() {
	var request models.UpdateTestpaperRequest
	err := unmarshalBody(t.Ctx.Input.RequestBody, &request)
	if err != nil {
		logger.Recorder.Warning("[test-paper] parse body error: " + err.Error())
		t.Ctx.Output.SetStatus(http.StatusBadRequest)
		t.Data["json"] = map[string]string{
			"msg": err.Error(),
		}
		t.ServeJSON()
		return
	}
	message := "ok"
	resp, err := models.UpdateTestpaper(&request)
	if err != nil {
		message = err.Error()
	}
	t.Ctx.Output.SetStatus(http.StatusOK)
	t.Data["json"] = map[string]interface{}{
		"msg":  message,
		"data": resp,
	}
	t.ServeJSON()
	return
}

// @Title AddTestpaperComment
// @Description 添加一条题目的评价内容
// @Param   json body models.AddTestpaperCommentRequest true "新建一个试卷评估记录"
// @Success 200 {string}
// @Failure 400 "invalid body"
// @router /comment [post]
func (t *TestpaperController) AddTestpaperComment() {
	var request models.AddTestpaperCommentRequest
	err := unmarshalBody(t.Ctx.Input.RequestBody, &request)
	if err != nil {
		logger.Recorder.Warning("[test-paper] parse body error: " + err.Error())
		t.Ctx.Output.SetStatus(http.StatusBadRequest)
		t.Data["json"] = map[string]string{
			"msg": err.Error(),
		}
		t.ServeJSON()
		return
	}
	message := "ok"
	err = models.AddTestpaperComment(&request)
	if err != nil {
		message = err.Error()
	}
	t.Ctx.Output.SetStatus(http.StatusOK)
	t.Data["json"] = map[string]interface{}{
		"msg": message,
	}
	t.ServeJSON()
	return
}

// @Title TraceTestpaperVersion
// @Description 查询试卷的历史版本(向前查询)
// @Param   qid path string true "temp test-paper id"
// @Success 200 {[]models.TempTestpaper}
// @Failure 400 "invalid qid"
// @router /trace/:qid [get]
func (t *TestpaperController) TraceTestpaperVersion() {
	tid := t.GetString(":tid")
	if tid == "" {
		logger.Recorder.Warning("[test-paper] get test-paper id error")
		t.Ctx.Output.SetStatus(http.StatusBadRequest)
		t.Data["json"] = map[string]string{
			"msg": "invalid test-paper id",
		}
		t.ServeJSON()
		return
	}
	message := "ok"
	resp, err := models.TraceTestpaperVersion(tid)
	if err != nil {
		message = err.Error()
	}
	t.Ctx.Output.SetStatus(http.StatusOK)
	t.Data["json"] = map[string]interface{}{
		"msg":  message,
		"data": resp,
	}
	t.ServeJSON()
	return
}

// @Title FinishTempTestpaper
// @Description 最终确定题目 (试卷完成评审, 转移到final数据库下)
// @Param   qid path string true "test-paper id"
// @Success 200 {string}
// @Failure 400 "invalid qid"
// @router /finish/:qid [get]
func (t *TestpaperController) FinishTempTestpaper() {
	tid := t.GetString(":tid")
	if tid == "" {
		logger.Recorder.Warning("[test-paper] get test-paper id error")
		t.Ctx.Output.SetStatus(http.StatusBadRequest)
		t.Data["json"] = map[string]string{
			"msg": "invalid test-paper id",
		}
		t.ServeJSON()
		return
	}
	message := "ok"
	resp, err := models.FinishTempTestpaper(tid)
	if err != nil {
		message = err.Error()
	}
	t.Ctx.Output.SetStatus(http.StatusOK)
	t.Data["json"] = map[string]interface{}{
		"msg":  message,
		"data": resp,
	}
	t.ServeJSON()
	return
}

// @Title GetUserTempTestpaper
// @Description 获取用户创建的试卷(temp数据库下)
// @Param   uid path string true "user id"
// @Success 200 {[]models.TempTestpaper}
// @Failure 400 "invalid qid"
// @router /user_t/:uid [get]
func (t *TestpaperController) GetUserTempTestpaper() {
	uid := t.GetString(":uid")
	if uid == "" {
		logger.Recorder.Warning("[test-paper] get user id error")
		t.Ctx.Output.SetStatus(http.StatusBadRequest)
		t.Data["json"] = map[string]string{
			"msg": "invalid user id",
		}
		t.ServeJSON()
		return
	}
	message := "ok"
	resp, err := models.GetUserTempTestpaper(uid)
	if err != nil {
		message = err.Error()
	}
	t.Ctx.Output.SetStatus(http.StatusOK)
	t.Data["json"] = map[string]interface{}{
		"msg":  message,
		"data": resp,
	}
	t.ServeJSON()
	return
}

// @Title GetUserFinalTestpaper
// @Description 获取用户创建的试卷(final数据库下)
// @Param   uid path string true "user id"
// @Success 200 {[]models.FinalTestpaper}
// @Failure 400 "invalid qid"
// @router /user_f/:uid [get]
func (t *TestpaperController) GetUserFinalTestpaper() {
	uid := t.GetString(":uid")
	if uid == "" {
		logger.Recorder.Warning("[test-paper] get user id error")
		t.Ctx.Output.SetStatus(http.StatusBadRequest)
		t.Data["json"] = map[string]string{
			"msg": "invalid user id",
		}
		t.ServeJSON()
		return
	}
	message := "ok"
	resp, err := models.GetUserFinalTestpaper(uid)
	if err != nil {
		message = err.Error()
	}
	t.Ctx.Output.SetStatus(http.StatusOK)
	t.Data["json"] = map[string]interface{}{
		"msg":  message,
		"data": resp,
	}
	t.ServeJSON()
	return
}

// @Title GetProjectTempTestpaper
// @Description 获取项目下的试卷(temp数据库下)
// @Param   uid path string true "project id"
// @Success 200 {[]models.TempTestpaper}
// @Failure 400 "invalid qid"
// @router /proj_t/:pid [get]
func (t *TestpaperController) GetProjectTempTestpaper() {
	pid := t.GetString(":pid")
	if pid == "" {
		logger.Recorder.Warning("[question] get project id error")
		t.Ctx.Output.SetStatus(http.StatusBadRequest)
		t.Data["json"] = map[string]string{
			"msg": "invalid project id",
		}
		t.ServeJSON()
		return
	}
	message := "ok"
	resp, err := models.GetProjectTempTestpaper(pid)
	if err != nil {
		message = err.Error()
	}
	t.Ctx.Output.SetStatus(http.StatusOK)
	t.Data["json"] = map[string]interface{}{
		"msg":  message,
		"data": resp,
	}
	t.ServeJSON()
	return
}

// @Title GetProjectFinalTestpaper
// @Description 获取项目下的试卷(final数据库下)
// @Param   uid path string true "project id"
// @Success 200 {[]models.FinalTestpaper}
// @Failure 400 "invalid qid"
// @router /proj_f/:pid [get]
func (t *TestpaperController) GetProjectFinalTestpaper() {
	pid := t.GetString(":pid")
	if pid == "" {
		logger.Recorder.Warning("[question] get project id error")
		t.Ctx.Output.SetStatus(http.StatusBadRequest)
		t.Data["json"] = map[string]string{
			"msg": "invalid project id",
		}
		t.ServeJSON()
		return
	}
	message := "ok"
	resp, err := models.GetProjecgtFinalTestpaper(pid)
	if err != nil {
		message = err.Error()
	}
	t.Ctx.Output.SetStatus(http.StatusOK)
	t.Data["json"] = map[string]interface{}{
		"msg":  message,
		"data": resp,
	}
	t.ServeJSON()
	return
}
