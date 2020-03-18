package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/hzde0128/hblog/models"
	_ "github.com/hzde0128/hblog/routers"
	_ "github.com/hzde0128/hblog/utils"
)

func main() {
	// 上一页/下一页
	beego.AddFuncMap("prepage", prepage)
	beego.AddFuncMap("nextpage", nextpage)

	// 调试输出sql语句
	orm.Debug = true

	beego.Run()
}

// prepage 上一页
func prepage(idx int) (page int) {
	if idx > 1 {
		page = idx - 1
	} else {
		page = idx
	}
	return
}

// nextpage 下一页
func nextpage(idx, count int) (page int) {
	if idx < count {
		page = idx + 1
	} else {
		page = count
	}
	return
}