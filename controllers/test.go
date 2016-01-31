package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"mywebapp/models"
)

type TestController struct {
	beego.Controller
}

var bar models.Bar

func (c *TestController) Add() {
	barId := 100
	barName := "my bar"
	barDesc := "this is my first go bar"
	id, err := bar.Add(barId, barName, barDesc)
	fmt.Println(id, err)
	c.Data["id"] = id
	c.Data["error"] = err
	c.TplName = "bar/add.tpl"
}

func (c *TestController) QueryAllBar() {
	bars, err := bar.QueryAllBar()
	c.Data["bars"] = bars
	c.Data["error"] = err
	c.TplName = "bar/barList.tpl"
}
