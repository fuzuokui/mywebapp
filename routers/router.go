package routers

import (
	"github.com/astaxie/beego"
	"mywebapp/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/addBar", &controllers.TestController{}, "get:Add")
	beego.Router("/query", &controllers.TestController{}, "get:QueryAllBar")
}
