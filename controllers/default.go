package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (main *MainController) HelloSitepoint() {
	main.Data["Website"] = "my website"
	main.Data["Email"] = "fzk-2008@163.com"
	main.Data["EmailName"] = "fuzuokui"
	main.Data["Id"] = main.Ctx.Input.Param(":id")
	main.TplName = "default/hello-sitepoint.tpl"

}
