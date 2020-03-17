package controllers

import "github.com/astaxie/beego"

type BlogCoontroller struct {
	beego.Controller
}

func (c *BlogCoontroller) Get() {
	c.TplName = "admin/list.tpl"
}
