package controllers

import "github.com/astaxie/beego"

type TagController struct {
	beego.Controller
}

func (c *TagController) Get() {
	c.TplName = "admin/tag.tpl"
}
