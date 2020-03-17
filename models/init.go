package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hzde0128/hblog/utils"
	"strconv"
)

func init() {
	dsn := utils.Cfg.MysqlConf.Username + ":" + utils.Cfg.Password + "@" + "tcp(" + utils.Cfg.MysqlConf.Address + ":" + strconv.Itoa(utils.Cfg.MysqlConf.Port) + ")/" + utils.Cfg.MysqlConf.Dbname + "?charset=utf8"
	beego.Debug(dsn)
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(User), new(Tag), new(Category), new(Article))
	orm.RunSyncdb("default", false, true)

	// 初始化用户
	user := User{}
	user.UserName = "admin"
	data := []byte("admin888")
	has := md5.Sum(data)
	user.Password = fmt.Sprintf("%x", has)
	_, _, err := user.ReadOrCreate("username", "password")
	if err != nil {
		beego.Info("读取或创建失败,", err)
		return
	}

}
