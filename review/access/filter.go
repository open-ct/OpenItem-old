package access

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web/context"
	"log"
	"review/logger"
	"review/utils"
)

type FilterResponse struct {
	Status string      `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func denyRequest(ctx *context.Context, message string) {
	w := ctx.ResponseWriter
	w.WriteHeader(403)
	// resp := &controllers.Response{Status: "error", Msg: "Unauthorized operation"}
	resp := &FilterResponse{Status: "error", Msg: message}
	_, err := w.Write([]byte(utils.StructToJson(resp)))
	if err != nil {
		logger.Recorder.Error("filter error:" + err.Error())
	}
}

func TokenFilter(ctx *context.Context) {
	method := ctx.Request.Method
	if method == "POST" {
		var requestBody map[string]interface{}
		json.Unmarshal(ctx.Input.RequestBody, &requestBody)
		logger.Recorder.Info(fmt.Sprintf("%s", requestBody))
		if requestBody["token"] == nil {
			log.Println("deny")
			denyRequest(ctx, "no token")
		}
	}
}
