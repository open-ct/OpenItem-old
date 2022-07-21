package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
	"warehouse/logger"
	"warehouse/models"
)

type QueryController struct {
	beego.Controller
}

// @Title GetTempQuestionList
// @Description 根据id列表获取temp-question信息
// @Param   json body models.QueryListRequest true "要获取的id 列表"
// @Success 200 {object} []models.TempQuestion
// @Failure 400 "parse id list error"
// @router /t_question [post]
func (q *QueryController) GetTempQuestionList() {
	queryRequset := new(models.QueryListRequest)
	err := unmarshalBody(q.Ctx.Input.RequestBody, queryRequset)
	if err != nil {
		logger.Recorder.Warning("[query] parse id list error: " + err.Error())
		q.Ctx.Output.SetStatus(http.StatusBadRequest)
		q.Data["json"] = map[string]string{
			"msg": err.Error(),
		}
		q.ServeJSON()
		return
	}
	resp := models.QueryTempQuestions(queryRequset.IdList)
	q.Ctx.Output.SetStatus(http.StatusOK)
	q.Data["json"] = map[string]interface{}{
		"msg":  "ok",
		"data": resp,
	}
	q.ServeJSON()
	return
}

// @Title GetFinalQuestionList
// @Description 根据id列表获取temp-question信息
// @Param   json body models.QueryListRequest true "要获取的id 列表"
// @Success 200 {object} []models.FinalQuestion
// @Failure 400 "parse id list error"
// @router /f_question [post]
func (q *QueryController) GetFinalQuestionList() {
	queryRequset := new(models.QueryListRequest)
	err := unmarshalBody(q.Ctx.Input.RequestBody, queryRequset)
	if err != nil {
		logger.Recorder.Warning("[query] parse id list error: " + err.Error())
		q.Ctx.Output.SetStatus(http.StatusBadRequest)
		q.Data["json"] = map[string]string{
			"msg": err.Error(),
		}
		q.ServeJSON()
		return
	}
	resp := models.QueryFinalQuestions(queryRequset.IdList)
	q.Ctx.Output.SetStatus(http.StatusOK)
	q.Data["json"] = map[string]interface{}{
		"msg":  "ok",
		"data": resp,
	}
	q.ServeJSON()
	return
}

// @Title GetTempTestPaperList
// @Description 根据id列表获取temp-test-paper信息
// @Param   json body models.QueryListRequest true "要获取的id 列表"
// @Success 200 {object} []models.TempTestpaper
// @Failure 400 "parse id list error"
// @router /t_testpaper [post]
func (q *QueryController) GetTempTestPaperList() {
	queryRequset := new(models.QueryListRequest)
	err := unmarshalBody(q.Ctx.Input.RequestBody, queryRequset)
	if err != nil {
		logger.Recorder.Warning("[query] parse id list error: " + err.Error())
		q.Ctx.Output.SetStatus(http.StatusBadRequest)
		q.Data["json"] = map[string]string{
			"msg": err.Error(),
		}
		q.ServeJSON()
		return
	}
	resp := models.QueryTempTestpaper(queryRequset.IdList)
	q.Ctx.Output.SetStatus(http.StatusOK)
	q.Data["json"] = map[string]interface{}{
		"msg":  "ok",
		"data": resp,
	}
	q.ServeJSON()
	return
}

// @Title GetFinalTestPaperList
// @Description 根据id列表获取fianl-test-paper信息
// @Param   json body models.QueryListRequest true "要获取的id 列表"
// @Success 200 {object} []models.FinalTestpaper
// @Failure 400 "parse id list error"
// @router /f_testpaper [post]
func (q *QueryController) GetFinalTestPaperList() {
	queryRequset := new(models.QueryListRequest)
	err := unmarshalBody(q.Ctx.Input.RequestBody, queryRequset)
	if err != nil {
		logger.Recorder.Warning("[query] parse id list error: " + err.Error())
		q.Ctx.Output.SetStatus(http.StatusBadRequest)
		q.Data["json"] = map[string]string{
			"msg": err.Error(),
		}
		q.ServeJSON()
		return
	}
	resp := models.QueryFinalTestpaper(queryRequset.IdList)
	q.Ctx.Output.SetStatus(http.StatusOK)
	q.Data["json"] = map[string]interface{}{
		"msg":  "ok",
		"data": resp,
	}
	q.ServeJSON()
	return
}
