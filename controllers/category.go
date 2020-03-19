package controllers

import (
	"github.com/astaxie/beego"
	"github.com/hzde0128/hblog/models"
	"github.com/hzde0128/hblog/utils"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Category() {
	username := c.GetSession("username")
	c.Data["user"] = username

	//获取网站标题
	system := models.System{}
	systeminfo := models.System{}
	system.Query().One(&systeminfo)

	c.Data["basic"] = systeminfo

	c.Data["title"] = "分类管理"
	cate := models.Category{}
	cates := []models.Category{}
	query := cate.Query()
	//query.All(&cates)
	beego.Info(cates)

	// 获取页面信息，向上取整
	page, err := strconv.Atoi(c.GetString("page"))
	if err != nil {
		page = 1
	}

	count, err := query.Count()
	if err != nil {
		beego.Info("查询失败,", err)
		return
	}
	beego.Info("Count:", count)

	pageSize := 5
	// 起始页
	start := pageSize * (page - 1)
	// 总页数
	pageCount := int(math.Ceil(float64(count) / float64(pageSize)))

	query.Limit(pageSize, start).All(&cates)
	beego.Info("分页信息:", utils.Cfg.Paginate)
	c.Data["pageCount"] = pageCount
	c.Data["page"] = page
	c.Data["category"] = cates

	if c.Ctx.Request.Method == "POST" {
		cate_name := strings.TrimSpace(c.GetString("cate_name"))
		beego.Info("CategoryName: ", cate_name)
		if cate_name == "" {
			c.Data["errMsg"] = "请输入分类名称"
		} else {

			// 检查是否存在，不存在则添加
			cate.CategoryName = cate_name
			err := cate.Read("category_name")
			if err == nil {
				beego.Info("该分类已存在")
				c.Data["errMsg"] = "该分类已存在，请勿重复添加"
				//c.Redirect("/admin/category/", http.StatusFound)
			} else {

				// 添加到数据库
				err = cate.Insert()
				if err != nil {
					beego.Info("插入数据库失败，", err)
					c.Redirect("/admin/category/", http.StatusFound)
				}

				// 添加成功后更新order 默认以id为顺序

				beego.Info(cate_name, "添加成功")
				// 跳转刷新分类页面
				c.Redirect("/admin/category/", http.StatusFound)
			}
		}
	}

	c.Layout = "admin/layout.tpl"
	c.TplName = "admin/category.tpl"
}
