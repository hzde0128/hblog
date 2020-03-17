package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/hzde0128/hblog/models"
	_ "github.com/hzde0128/hblog/routers"
	"github.com/hzde0128/hblog/utils"
)

func main() {
	// 上一页/下一页
	beego.AddFuncMap("prepage", utils.Prepage)
	beego.AddFuncMap("nextpage", utils.Nextpage)

	// 调试输出sql语句
	orm.Debug = true

	beego.Run()
}
