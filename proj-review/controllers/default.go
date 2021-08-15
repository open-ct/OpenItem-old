package controllers

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"proj-review/auth"
	"proj-review/log"
)

// Test controller, no used.
type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) List() {
	c.TplName = "api.html"
}

func (c *MainController) Demo() {
	c.Ctx.Output.Download("tests/api-test.http", "api.http")
}

// universal functions in controller

// unmarshalBody parse the response body (json type)
func unmarshalBody(body []byte, obj interface{}) error {
	// marshal json body
	if err := json.Unmarshal(body, obj); err != nil {
		log.Logger.Warn("[JSON] " + err.Error())
		return err
	}
	return nil
}

// parseUserToken get user id from token string
func parseUserToken(token string) (string, error) {
	//
	fmt.Println("recv token: ", token)
	//
	tokenData, err := auth.ParseToken(token)
	if err != nil {
		return "", err
	}
	//
	fmt.Println("recv user: ", tokenData.UserID)
	fmt.Println("recv token time from ", tokenData.IssuedAt, " to ", tokenData.ExpiresAt)
	//
	return tokenData.UserID, nil
}

func checkHttpBodyEmpty(body []byte) bool {
	if len(body) == 0 {
		return false
	}
	return true
}
