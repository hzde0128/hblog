package controllers

import (
	"github.com/astaxie/beego"
	"github.com/hzde0128/hblog/models"
	"github.com/hzde0128/hblog/utils"
	"net/http"
	"runtime"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	username := c.GetSession("username")

	// 获取本机IP
	ip, err := utils.GetOutboundIP()
	if err != nil {
		beego.Info("获取IP失败",err)
	}
	// 获取版本信息
	version := utils.Cfg.Version

	c.Data["title"] = "首页"
	c.Data["user"] = username
	c.Data["version"] = version
	c.Data["goVer"] = runtime.Version()
	c.Data["ip"] = ip
	c.Data["os"] = runtime.GOOS
	c.Data["cpuNum"] = runtime.NumCPU()
	c.Data["arch"] = runtime.GOARCH

	c.Layout = "admin/layout.tpl"
	c.TplName = "admin/index.tpl"
}

func (c *AdminController) Setting() {
	c.Data["title"] = "系统设置"
	username := c.GetSession("username")
	c.Data["user"] = username

	// 获取系统配置信息
	system := models.System{}
	systeminfo := models.System{}
	err := system.Query().One(&systeminfo)
	if err != nil {
		beego.Info("暂无信息")
	}
	beego.Info("系统信息：", systeminfo)

	if c.Ctx.Request.Method == "POST"{
		title := c.GetString("title")
		baseurl := c.GetString("baseUrl")
		copyright := c.GetString("copyright")

		if title == "" || baseurl == "" || copyright =="" {
			c.Data["errMsg"] = "字段不能为空"
			c.Layout = "admin/layout.tpl"
			c.TplName = "admin/setting.tpl"
			return
		}

		// 检查之前有没有数据，没有就新增有则修改
		err := system.Query().One(&systeminfo)
		if err != nil {
			beego.Info("找不到信息",err)
			system.Title = title
			system.BaseUrl = baseurl
			system.CopyRight = copyright
			err := system.Insert()
			if err != nil {
				beego.Info("新增失败")
			}
		} else {
			// 更新系统设置
			system.Title = title
			system.BaseUrl = baseurl
			system.CopyRight = copyright
			err := system.Update()
			if err != nil {
				beego.Info("更新失败")
			}
		}
		// 跳转系统设置页面
		c.Redirect("/admin/setting", http.StatusFound)
	}

	c.Data["system"] = systeminfo
	c.Layout = "admin/layout.tpl"
	c.TplName = "admin/setting.tpl"
}
