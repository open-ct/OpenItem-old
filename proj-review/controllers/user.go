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
		log.Logger.Warn("[Register] " + err.Error())
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

// UserLogin
func (u *UserController) UserLogin() {
	loginReq := new(request.UserLogin)
	err := unmarshalBody(u.Ctx.Input.RequestBody, loginReq)
	if err != nil {
		log.Logger.Warn("[Login] " + err.Error())
		u.respondJson(
			http.StatusOK,
			constant.FAIL,
			constant.BasicMsg.LoginFail,
			loginReq,
		)
		return
	}
	loginResp, ok := models.DoUserLogin(loginReq)
	if !ok {
		u.respondJson(
			http.StatusOK,
			constant.FAIL,
			constant.BasicMsg.LoginFail,
			loginResp,
		)
	} else {
		u.respondJson(
			http.StatusOK,
			constant.SUCCESS,
			constant.BasicMsg.LoginSuccess,
			loginResp,
		)
	}
	return
}

// todo: 取消 token 有效期
func (u *UserController) UserLogout() {

}

func (u *UserController) GetUserInfo() {
	uid := u.GetString("user_id")
	if uid == "" {
		log.Logger.Warn("[Get Info] invalid user id")
		u.respondJson(
			http.StatusOK,
			constant.FAIL,
			constant.BasicMsg.GetInfoFail,
			uid,
		)
		return
	}
	//
	getInfoResp, ok := models.DoGetUserInfo(uid)
	if !ok {
		u.respondJson(
			http.StatusOK,
			constant.FAIL,
			constant.BasicMsg.GetInfoFail,
			uid,
		)
	} else {
		u.respondJson(
			http.StatusOK,
			constant.SUCCESS,
			constant.BasicMsg.GetInfoSuccess,
			getInfoResp,
		)
	}
	return
}

// UpdateUserPassword
func (u *UserController) UpdateUserPassword() {
	updatePwdReq := new(request.UserUpdatePassword)
	err := unmarshalBody(u.Ctx.Input.RequestBody, updatePwdReq)
	if err != nil {
		log.Logger.Warn("[Update] " + err.Error())
		u.respondJson(
			http.StatusOK,
			constant.FAIL,
			constant.BasicMsg.UpdateInfoFail,
			updatePwdReq,
		)
		return
	}
	// 处理
	updateResponse, ok := models.DoUpdatePassword(updatePwdReq)
	if !ok {
		u.respondJson(
			http.StatusOK,
			constant.FAIL,
			constant.BasicMsg.UpdateInfoFail,
			updateResponse,
		)
	} else {
		u.respondJson(
			http.StatusOK,
			constant.SUCCESS,
			constant.BasicMsg.UpdateInfoSuccess,
			updateResponse,
		)
	}
	return
}

func (u *UserController) UpdateUserInfo() {
	updateInfoReq := new(request.UserUpdateInfo)
	err := unmarshalBody(u.Ctx.Input.RequestBody, updateInfoReq)
	if err != nil {
		log.Logger.Warn("[Update] " + err.Error())
		u.respondJson(
			http.StatusOK,
			constant.FAIL,
			constant.BasicMsg.UpdateInfoFail,
			updateInfoReq,
		)
		return
	}
	// do update
	updateResponse, ok := models.DoUpdateUserInfo(updateInfoReq)
	if !ok {
		u.respondJson(
			http.StatusOK,
			constant.FAIL,
			constant.BasicMsg.UpdateInfoFail,
			updateResponse,
		)
	} else {
		u.respondJson(
			http.StatusOK,
			constant.SUCCESS,
			constant.BasicMsg.UpdateInfoSuccess,
			updateResponse,
		)
	}
	return
}

// SearchUser todo:
func (u *UserController) SearchUser() {

}

// DeleteUser
func (u *UserController) DeleteUser() {
	uid := u.GetString("user_id")
	if uid == "" {
		log.Logger.Warn("[Delete] invalid user id")
		u.respondJson(
			http.StatusOK,
			constant.FAIL,
			constant.BasicMsg.DeleteUserFail,
			uid,
		)
		return
	}
	// do delete
	resp, ok := models.DoDeleteUser(uid)
	if !ok {
		u.respondJson(
			http.StatusOK,
			constant.FAIL,
			constant.BasicMsg.DeleteUserFail,
			resp,
		)
	} else {
		u.respondJson(
			http.StatusOK,
			constant.SUCCESS,
			constant.BasicMsg.DeleteUserSuccess,
			resp,
		)
	}
	return
}
