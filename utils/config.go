package utils

import (
	"github.com/astaxie/beego"
	"gopkg.in/ini.v1"
)

var Cfg = new(AppConf)

type AppConf struct {
	MysqlConf `ini:"mysql"`
	Paginate  int `ini:"paginate"`
}

type MysqlConf struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"user"`
	Password string `ini:"password"`
	Dbname   string `ini:"dbname"`
}

// 加载配置文件
func init() {
	err := ini.MapTo(Cfg, "conf/app.conf")
	if err != nil {
		beego.Error("读取配置文件失败,", err)
		return
	}
}
