package controllers

import "github.com/astaxie/beego"

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	c.TplName = "admin/category.tpl"
}
