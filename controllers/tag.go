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

func (c *TagController) Tag() {
	username := c.GetSession("username")
	c.Data["user"] = username

	tag := models.Tag{}
	tags := []models.Tag{}
	query := tag.Query()
	query.All(&tags)
	beego.Info(tags)

	if c.Ctx.Request.Method == "POST" {
		// 获取标签信息
		tag_name := strings.TrimSpace(c.GetString("tag_name"))
		if tag_name == "" {
			c.Data["errMsg"] = "请输入标签名称再添加"
		} else {

			beego.Info("标签名称：", tag_name)

			// 检查是否存在，不存在则添加
			tag := models.Tag{}

			tag.TagName = tag_name
			err := tag.Read("tag_name")
			if err == nil {
				c.Data["errMsg"] = "该标签已存在,请勿重复添加"
			} else {

				// 添加到数据库
				err = tag.Insert()
				if err != nil {
					beego.Info("插入数据库失败，", err)
					c.Redirect("/admin/tag/", http.StatusFound)
				}
				beego.Info(tag_name, "添加成功")
				// 添加成功需要重新刷新页面显示
				c.Redirect("/admin/tag/", http.StatusFound)
			}
		}
	}

	c.Data["tag"] = tags
	c.Layout = "admin/layout.tpl"
	c.TplName = "admin/tag.tpl"
}
