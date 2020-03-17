package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/hzde0128/hblog/models"
	"net/http"
	"strings"
	"time"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Login() {

	if c.Ctx.Request.Method == "POST" {
		username := strings.TrimSpace(c.GetString("username"))
		password := strings.TrimSpace(c.GetString("password"))

		beego.Info(username, password)

		user := models.User{}
		user.UserName = username
		data := []byte(password)
		has := md5.Sum(data)
		user.Password = fmt.Sprintf("%x", has)
		err := user.Read("username", "password")
		if err != nil {
			beego.Info("用户名或密码错误")
			c.Redirect("/admin/login", http.StatusFound)
		}

		// 验证成功，设置Cookie和Session
		c.Ctx.SetCookie("name", username, time.Second*3600)
		c.SetSession("name", username)

		c.Redirect("/admin/", http.StatusFound)
	}
	c.TplName = "admin/login.tpl"
}

func (c *UserController) Logout() {
	c.DelSession("name")

	c.Redirect("/", http.StatusFound)
}

func (c *UserController) User() {

	c.Redirect("/admin/user", http.StatusFound)
}
