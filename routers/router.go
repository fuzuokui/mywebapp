package routers

import (
	"github.com/astaxie/beego"
	"mywebapp/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello-world/:id([0-9]+)", &controllers.MainController{}, "get:HelloSitepoint")
	beego.Router("/addBar", &controllers.TestController{}, "get:Add")
	beego.Router("/query", &controllers.TestController{}, "get:QueryAllBar")
}
