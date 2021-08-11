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

type UserController struct {
	beego.Controller
}

// respondJson
func (u *UserController) respondJson(httpCode int, opCode int, message string, data ...interface{}) {
	u.Ctx.Output.SetStatus(httpCode)
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
	u.Data["json"] = resp
	u.ServeJSON()
}

// UserRegister
func (u *UserController) UserRegister() {
	registerReq := new(request.UserRegister)
	err := unmarshalBody(u.Ctx.Input.RequestBody, registerReq)
	if err != nil {
		log.Logger.Warn("[Register] %s" + err.Error())
		u.respondJson(
			http.StatusOK,
			constant.FAIL,
			constant.BasicMsg.RegisterFail,
			registerReq,
		)
		return
	}
	// 处理注册
	registerResponse, ok := models.DoUserRegister(registerReq)
	if !ok {
		u.respondJson(
			http.StatusOK,
			constant.FAIL,
			constant.BasicMsg.RegisterFail,
			registerResponse,
		)
	} else {
		u.respondJson(
			http.StatusOK,
			constant.SUCCESS,
			constant.BasicMsg.RegisterSuccess,
			registerResponse,
		)
	}
	return
}
