package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/hzde0128/hblog/controllers"
	"net/http"
)

func init() {
	beego.InsertFilter("/admin/*", beego.BeforeRouter, filterFunc)
	//beego.InsertFilter("/admin/(user|info|blog/|category/|tag/)", beego.BeforeRouter, filterFunc)
	// 定义路由
	beego.Router("/", &controllers.MainController{})
	// 登录页面
	beego.Router("/login", &controllers.UserController{}, "*:Login")
	// 登出页面
	beego.Router("/admin/logout", &controllers.UserController{}, "*:Logout")
	// 修改密码
	beego.Router("/admin/user", &controllers.UserController{}, "*:User")
	// 后台首页
	beego.Router("/admin/", &controllers.AdminController{})
	// 后台信息首页
	beego.Router("/admin/info", &controllers.AdminController{}, "*:Info")
	// 文章管理
	beego.Router("/admin/blog/", &controllers.BlogCoontroller{})
	// 添加文章
	beego.Router("/admin/blog/add", &controllers.BlogCoontroller{}, "*:Add")
	// 分类管理
	beego.Router("/admin/category/", &controllers.CategoryController{}, "*:Category")
	// 标签管理
	beego.Router("/admin/tag/", &controllers.TagController{}, "*:Tag")
}

var filterFunc = func(ctx *context.Context) {
	username := ctx.Input.Session("username")
	if username == nil {
		ctx.Redirect(http.StatusFound, "/login")
	}
}
