package controllers

import (
	"net/http"
	"review/models"
	"review/request"
	"review/response"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// respondJson: do response for user operations
func (u *UserController) respondJson(httpCode int, opCode int, message string, data ...interface{}) {
	u.Ctx.Output.SetStatus(httpCode)
	var d interface{}
	if len(data) == 1 {
		d = data[0]
	} else {
		d = data
	}
	resp := response.GenResponse(opCode, message, d)
	u.Data["json"] = resp
	u.ServeJSON()
}

// @Title UserRegister
// @Description 用户注册部分, 邮箱和手机号不能和已注册用户重复
// @Param   json body request.UserRegister true "要注册的用户信息"
// @Success 200 {object} response.Default
// @Failure 400 "parse body failed"
// @router / [post]
func (u *UserController) UserRegister() {
	var user request.UserRegister
	err := unmarshalBody(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		u.respondJson(http.StatusBadRequest, response.FAIL, "Parse Body Failed")
		return
	}
	uid, code := models.AddUser(&user)
	u.respondJson(http.StatusOK, code, "", uid)
	return
}

// @Title GetAll
// @Description get all Users (已废弃)
// @Success 200 {array} []models.User
// @router / [get]
func (u *UserController) GetAll() {
	users, code := models.GetProjectUsers()
	u.respondJson(http.StatusOK, code, "", users)
}

// @Title GetUser
// @Description get a user info (profile): 根据用户id获取用户基本信息 (用户密码不会返回)
// @Param   token header string true "user token recived at login"
// @Param   uid path string true "user 的 uuid"
// @Success 200 {object} response.Default
// @Failure 400 "invalid user id"
// @router /:uid [get]
func (u *UserController) GetUser() {
	uid := u.GetString(":uid")
	if uid == "" {
		u.respondJson(http.StatusBadRequest, response.FAIL, "invalid user ID")
		return
	}
	user, code := models.GetUser(uid)
	u.respondJson(http.StatusOK, code, "", user)
}

// @Title UpdateUserInfo
// @Description 更新用户信息 (主要提交post时必须所有选项都要填写, 没发生更改的应当填充原来的信息数据)
// @Param   token header string true "user token get at login"
// @Param   json body request.UserUpdateInfo true "要更新的用户信息数据"
// @Success 200 {object} response.Default
// @Failure 400 "parse body error"
// @router / [put]
func (u *UserController) UpdateUserInfo() {
	newInfo := request.UserUpdateInfo{}
	err := unmarshalBody(u.Ctx.Input.RequestBody, &newInfo)
	if err != nil {
		u.respondJson(http.StatusBadRequest, response.FAIL, "Parse body failed")
		return
	}
	res, code := models.UpdateUserInfo(&newInfo)
	u.respondJson(http.StatusOK, code, "", res)
	return
}

// @Title UserDelete
// @Description delete a user by user id: 根据id删除用户记录
// @Param   token header string true "user token get at login"
// @Param   uid path string true "user uuid"
// @Success 200 {object} response.Default
// @Failure 400 "invalid user uuid"
// @router /:uid [delete]
func (u *UserController) UserDelete() {
	uid := u.GetString(":uid")
	if uid == "" {
		u.respondJson(http.StatusBadRequest, response.FAIL, "invalid user id")
		return
	}
	code := models.DeleteUser(uid)
	u.respondJson(http.StatusOK, code, "")
	return
}

// @Title UserLogin
// @Description user login: 用户登录(邮箱 & 手机号 二选一即可登录)
// @Param   json body request.UserLogin true "user email / phone & password"
// @Success 200 {object} response.Default
// @Failure 400 "parse body failed"
// @router /login [post]
func (u *UserController) UserLogin() {
	var login request.UserLogin
	err := unmarshalBody(u.Ctx.Input.RequestBody, &login)
	if err != nil {
		u.respondJson(http.StatusBadRequest, response.FAIL, "parse body failed")
		return
	}
	res, token, code := models.UserLogin(&login)
	u.respondJson(http.StatusOK, code, token, res)
	return
}

// @Title UpdateUserPassword
// @Description update user's password: 更新用户密码(需要重新认证原密码) todo: 更新完成后退出当前登录 (注销 token)
// @Param   token header string true "user token get at login"
// @Param   json body request.UserUpdatePassword true "需要填写旧密码和新密码"
// @Success 200 {object} response.Default
// @Failure 400 "invalid body data"
// @router /password [put]
func (u *UserController) UpdateUserPassword() {
	var req request.UserUpdatePassword
	err := unmarshalBody(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		u.respondJson(http.StatusBadRequest, response.FAIL, "parse body failed")
		return
	}
	res, code := models.UpdateUserPassword(&req)
	u.respondJson(http.StatusOK, code, "", res)
	return
}

// @Title UserLogout
// @Description user logout: 用户退出, 注销token的有效期
// @Param   token header string true "user token get at login"
// @Success 200 {object} response.Default
// @Failure 400 "invalid ??"
// @router /logout [get]
func (u *UserController) UserLogout() {
	uid, err := parseUserToken(u.Ctx.Request.Header["Token"][0])
	if err != nil {
		u.respondJson(http.StatusBadRequest, response.FAIL, "need token", uid)
		return
	}
	// todo
	u.respondJson(http.StatusOK, response.SUCCESS, "", uid)
	return
}
