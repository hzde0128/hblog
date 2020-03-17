package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	cookie := c.Ctx.GetCookie("uname")
	fmt.Printf("%#v\n", cookie)

	c.TplName = "admin/index.tpl"
}

func (c *AdminController) Info() {
	c.TplName = "admin/info.tpl"
}
