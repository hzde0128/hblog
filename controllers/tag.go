package controllers

import (
	"github.com/astaxie/beego"
	"github.com/hzde0128/hblog/models"
	"net/http"
	"strings"
)

type TagController struct {
	beego.Controller
}

func (c *TagController) Get() {
	tag := models.Tag{}
	tags := []models.Tag{}
	query := tag.Query()
	query.All(&tags)
	beego.Info(tags)

	c.Data["tag"] = tags
	c.TplName = "admin/tag.tpl"
}

func (c *TagController) Post() {
	// 获取标签信息
	tag_name := strings.TrimSpace(c.GetString("tag_name"))

	if tag_name != "" {

		beego.Info("标签名称：", tag_name)

		// 检查是否存在，不存在则添加
		tag := models.Tag{}

		tag.TagName = tag_name
		err := tag.Read("tag_name")
		if err == nil {
			beego.Info("该标签已存在,请勿重复添加", err)
			c.Redirect("/admin/tag/", http.StatusFound)
		} else {

			// 添加到数据库
			err = tag.Insert()
			if err != nil {
				beego.Info("插入数据库失败，", err)
				c.Redirect("/admin/tag/", http.StatusFound)
			}
			beego.Info(tag_name, "添加成功")
		}
	}

	c.Redirect("/admin/tag/", http.StatusFound)
}