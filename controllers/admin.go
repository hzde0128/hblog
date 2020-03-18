package controllers

import (
	"github.com/astaxie/beego"
	"runtime"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	username := c.GetSession("username")

	c.Data["user"] = username
	c.Data["version"] = "v0.1.3"
	c.Data["goVer"] = runtime.Version()
	c.Data["os"] = runtime.GOOS
	c.Data["cpuNum"] = runtime.NumCPU()
	c.Data["arch"] = runtime.GOARCH

	c.Layout = "admin/layout.tpl"
	c.TplName = "admin/index.tpl"
}

func (c *AdminController) Setting() {
	username := c.GetSession("username")
	c.Data["user"] = username

	c.Layout = "admin/layout.tpl"
	c.TplName = "admin/setting.tpl"
}