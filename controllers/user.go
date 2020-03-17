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
			c.Redirect("/login", http.StatusFound)
		}

		// 验证成功，设置Cookie和Session
		c.Ctx.SetCookie("username", username, time.Second*3600)
		c.SetSession("username", username)

		c.Redirect("/admin/", http.StatusFound)
	}
	c.TplName = "login.tpl"
}

func (c *UserController) Logout() {
	c.DelSession("username")

	c.Redirect("/", http.StatusFound)
}

func (c *UserController) User() {
	if c.Ctx.Request.Method == "GET" {
		c.TplName = "admin/user.tpl"
	}

	if c.Ctx.Request.Method == "POST" {
		oldpass := strings.TrimSpace(c.GetString("oldpass"))
		old := []byte(oldpass)
		hash := md5.Sum(old)

		user := models.User{}
		user.UserName = "admin"
		user.Password = fmt.Sprintf("%x", hash)
		err := user.Read("username", "password")
		if err != nil {
			beego.Info("用户密码不正确", err)
			return
		}

		newpass := strings.TrimSpace(c.GetString("newpass"))
		repass := strings.TrimSpace(c.GetString("repass"))
		if newpass != repass {
			beego.Info("两次密码不一致")
		} else {

			// 执行修改操作
			data := []byte(newpass)
			has := md5.Sum(data)
			user.Password = fmt.Sprintf("%x", has)
			err = user.Update("Password")
			if err != nil {
				beego.Info("修改密码失败,", err)
			} else {
				c.DelSession("username")
				// 修改成功，退出登录，使用新密码登录
				c.Redirect("/login", http.StatusFound)
			}
		}
	}

}
