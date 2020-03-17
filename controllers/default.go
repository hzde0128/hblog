package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hzde0128/hblog/models"
)

type MainController struct {
	beego.Controller
}

// 首页
func (c *MainController) Get() {
	// 获取文章信息
	o := orm.NewOrm()
	articles := []models.Article{}
	o.QueryTable("article").RelatedSel("category").All(&articles)
	c.Data["article"] = articles
	c.TplName = "index.tpl"
}
