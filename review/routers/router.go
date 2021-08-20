// @APIVersion 1.0.0
// @Title OpenCT PQBS - api document
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact kkoogqw@hotmail.com
// @TermsOfServiceUrl http://
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"review/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/review",
		beego.NSNamespace("/demo",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/file",
			beego.NSInclude(
				&controllers.FileController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
