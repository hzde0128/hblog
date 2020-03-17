package controllers

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	c.Data["user"] = "admin"
	c.TplName = "admin/index.tpl"
}

func (c *AdminController) Info() {
	c.Data["user"] = "admin"
	c.TplName = "admin/info.tpl"
}
