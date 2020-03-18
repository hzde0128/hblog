package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hzde0128/hblog/models"
	"net/http"
	"strconv"
	"strings"
)

type BlogCoontroller struct {
	beego.Controller
}

// 文章列表页面
func (c *BlogCoontroller) Get() {
	username := c.GetSession("username")
	c.Data["user"] = username

	c.Data["title"] = "文章列表"
	// 获取分类
	//cate := models.Category{}
	//cates := []models.Category{}
	//query := cate.Query()
	//query.All(&cates)
	//beego.Info(cates)

	// 获取全部文章
	o := orm.NewOrm()
	articles := []models.Article{}
	_, err := o.QueryTable("article").OrderBy("-article__id","-article_create_at").RelatedSel("category").All(&articles)
	if err != nil {
		beego.Info("查询失败", err)
	}

	// 查询文章对应的标签
	//o.QueryM2M()
	c.Data["article"] = articles

	c.Layout = "admin/layout.tpl"
	c.TplName = "admin/list.tpl"
}

// 添加文章
func (c *BlogCoontroller) Add() {
	username := c.GetSession("username")
	c.Data["user"] = username

	c.Data["title"] = "发表文章"
	// 获取分类
	cate := models.Category{}
	cates := []models.Category{}
	query := cate.Query()
	query.All(&cates)
	beego.Info(cates)

	// 获取标签
	tag := models.Tag{}
	tags := []models.Tag{}
	querys := tag.Query()
	querys.All(&tags)
	beego.Info(tags)

	c.Data["category"] = cates
	c.Data["tag"] = tags

	if c.Ctx.Request.Method == "POST" {
		title := strings.TrimSpace(c.GetString("title"))
		categoryName := c.GetString("category")
		tags_id := c.GetStrings("tag_id")
		content := c.GetString("content")

		beego.Info("标题:", title)
		beego.Info("文章分类：", categoryName)
		beego.Info("标签：", tags_id)
		beego.Info("文章内容:", content)

		// 数据已经拿到开始处理

		category := models.Category{}
		category.CategoryName = categoryName
		err := category.Read("category_name")
		if err != nil {
			beego.Info("非法的分类")
		}

		article := models.Article{
			Title:    title,
			Category: &category,
			Content:  content,
		}

		// 1 获取多对多对象
		o := orm.NewOrm()
		m2m := o.QueryM2M(&article, "Tags")

		tag := models.Tag{}
		for _, val := range tags_id {
			data, err := strconv.Atoi(val)
			if err != nil {
				beego.Info("字符串转整形失败,", err)
			}
			tag.Id = data
			// 插入文章表
			err = article.Insert()
			if err != nil {
				beego.Info("插入文章表失败,", err)
			}
			// 多对多插入
			_, err = m2m.Add(&tag)
			if err != nil {
				beego.Info("写入失败,", err)
			}
			// 添加成功跳到列表页面
			c.Redirect("/admin/blog/", http.StatusFound)
		}

	}

	c.Layout = "admin/layout.tpl"
	c.TplName = "admin/article.tpl"
}
