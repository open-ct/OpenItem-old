package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
	"review/noticer"
	"review/response"
)

type NoticerController struct {
	beego.Controller
}

func (n *NoticerController) respondJson(httpCode int, opCode int, message string, data ...interface{}) {
	n.Ctx.Output.SetStatus(httpCode)
	var d interface{}
	if len(data) == 1 {
		d = data[0]
	} else {
		d = data
	}
	resp := response.GenResponse(opCode, message, d)
	n.Data["json"] = resp
	n.ServeJSON()
}

// @Title SendEmail
// @Description 发送邮件
// @Param   json body noticer.EmailPackage true "邮件内容"
// @Success 200 {object} response.Default
// @Failure 400 "invalid email-structure json body(token)"
// @router /email [post]
func (n *NoticerController) SendEmail() {
	var emailInfo noticer.EmailPackage
	err := unmarshalBody(n.Ctx.Input.RequestBody, &emailInfo)
	if err != nil {
		n.respondJson(http.StatusBadRequest, response.FAIL, "parse body (email information) failed")
		return
	}
	err = noticer.SendEmail(&emailInfo)
	if err != nil {
		n.respondJson(http.StatusOK, response.FAIL, "send email failed: "+err.Error())
		return
	}
	n.respondJson(http.StatusOK, response.SUCCESS, "send successfully")
	return
}
