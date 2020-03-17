package main

import (
	_ "github.com/hzde0128/hblog/routers"
	"github.com/astaxie/beego"
	_ "github.com/hzde0128/hblog/models"
	_ "github.com/hzde0128/hblog/utils"
)

func main() {
	beego.Run()
}

