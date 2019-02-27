package controllers

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}

func (c *HomeController)ShowHome()  {
	c.TplName="manager/home.html"
}