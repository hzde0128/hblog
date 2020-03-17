package routers

import (
	"github.com/hzde0128/hblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 定义路由
    beego.Router("/", &controllers.MainController{})
    // 登录页面
	beego.Router("/admin/login", &controllers.UserController{}, "*:Login")
    // 登出页面
	beego.Router("/admin/logout", &controllers.UserController{}, "*:Logout")
    // 修改密码
	beego.Router("/admin/user", &controllers.UserController{}, "*:User")
    // 后台首页
	beego.Router("/admin/", &controllers.AdminController{})
	// 后台信息首页
	beego.Router("/admin/info", &controllers.AdminController{},"*:Info")
    // 文章管理
	beego.Router("/admin/blog/", &controllers.BlogCoontroller{})
    // 分类管理
	beego.Router("/admin/category/", &controllers.CategoryController{})
    // 标签管理
    beego.Router("/admin/tag/", &controllers.TagController{})
}
