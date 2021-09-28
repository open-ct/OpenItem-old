// @APIVersion 1.0.0
// @Title OpenCT Question-warehouse API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact kkoogqw@hotmial.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"warehouse/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/qbank",
		beego.NSNamespace("/question",
			beego.NSInclude(
				&controllers.QuestionController{},
			),
		),
		beego.NSNamespace("/testpaper",
			beego.NSInclude(
				&controllers.TestpaperController{},
			),
		),
		beego.NSNamespace("/query",
			beego.NSInclude(
				&controllers.QueryController{},
			),
		),
	)
	beego.AddNamespace(ns)
}