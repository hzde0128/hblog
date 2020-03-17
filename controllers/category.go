package controllers

import (
	"github.com/astaxie/beego"
	"github.com/hzde0128/hblog/models"
	"net/http"
	"strings"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	cate := models.Category{}
	cates := []models.Category{}
	query := cate.Query()
	query.All(&cates)
	beego.Info(cates)

	c.Data["category"] = cates
	c.TplName = "admin/category.tpl"
}

func (c *CategoryController) Post() {
	cate_name := strings.TrimSpace(c.GetString("cate_name"))
	beego.Info("CategoryName: ", cate_name)

	if cate_name != "" {
		// 检查是否存在，不存在则添加
		cate := models.Category{}

		cate.CategoryName = cate_name
		err := cate.Read("category_name")
		if err == nil {
			beego.Info("该分类已存在,", err)
			c.Redirect("/admin/category/", http.StatusFound)
		} else {

			// 添加到数据库
			err = cate.Insert()
			if err != nil {
				beego.Info("插入数据库失败，", err)
				c.Redirect("/admin/category/", http.StatusFound)
			}
			beego.Info(cate_name, "添加成功")
		}
	}
	c.Redirect("/admin/category/", http.StatusFound)
}
