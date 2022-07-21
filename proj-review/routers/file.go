package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"proj-review/controllers"
)

func init() {
	beego.Router("/file/getInfo", &controllers.FileController{}, "get:GetFileInfo")
	beego.Router("/file/download", &controllers.FileController{}, "get:DownloadFile")
	beego.Router("/file/upload", &controllers.FileController{}, "post:UploadFile")
	beego.Router("/file/search", &controllers.FileController{}, "post:SearchFile")
	beego.Router("/file/delete", &controllers.FileController{}, "get:DeleteFile")
}
