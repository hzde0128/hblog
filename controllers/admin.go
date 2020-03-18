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

	c.Data["title"] = "首页"
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
	c.Data["title"] = "系统设置"
	username := c.GetSession("username")
	c.Data["user"] = username

	c.Layout = "admin/layout.tpl"
	c.TplName = "admin/setting.tpl"
}
