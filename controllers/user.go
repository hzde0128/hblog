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
	username := c.Ctx.GetCookie("username")
	if username != "" {
		c.Data["username"] = username
	}

	if c.Ctx.Request.Method == "POST" {
		username := strings.TrimSpace(c.GetString("username"))
		password := strings.TrimSpace(c.GetString("password"))

		remember := c.GetString("remember")
		beego.Info("记住密码: ", remember)
		beego.Info(username, password)

		// 验证数据有效性
		if username != "" && password != "" {
			user := models.User{}
			user.UserName = username
			data := []byte(password)
			has := md5.Sum(data)
			user.Password = fmt.Sprintf("%x", has)
			err := user.Read("username", "password")
			if err != nil {
				c.Data["errMsg"] = "用户名或密码错误,请重新输入"
			} else {

				// 更新登录次数和IP以及时间
				user.LastIp = strings.Split(c.Ctx.Request.RemoteAddr, ":")[0]
				user.LoginCount += 1
				// 获取时区
				loc, _ := time.LoadLocation("Local")
				timeStr := time.Now().Format("2006-01-02 15:04:05")
				t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
				user.LastLogin = t
				user.Update()

				// 验证成功，设置Cookie和Session
				if remember == "on" {
					c.Ctx.SetCookie("username", username, time.Second*3600)
				} else {
					c.Ctx.SetCookie("username", username)
				}
				c.SetSession("username", username)
				c.Redirect("/admin/", http.StatusFound)
			}
		}
	}

	c.TplName = "login.tpl"
}

func (c *UserController) Logout() {
	c.DelSession("username")
	username := c.Ctx.GetCookie("username")
	c.Ctx.SetCookie("username", username, -1)
	c.Redirect("/login", http.StatusFound)
}

// 修改密码页面
func (c *UserController) User() {
	username := c.GetSession("username")
	c.Data["user"] = username

	c.Data["title"] = "修改密码"
	if c.Ctx.Request.Method == "POST" {
		oldpass := strings.TrimSpace(c.GetString("oldpass"))
		old := []byte(oldpass)
		hash := md5.Sum(old)

		user := models.User{}
		user.UserName = "admin"
		user.Password = fmt.Sprintf("%x", hash)
		err := user.Read("username", "password")
		if err != nil {
			beego.Info("旧密码不正确", err)
			c.Data["errMsg"] = "旧密码不正确"
			c.Layout = "admin/layout.tpl"
			c.TplName = "admin/user.tpl"
			return
		}

		newpass := strings.TrimSpace(c.GetString("newpass"))
		repass := strings.TrimSpace(c.GetString("repass"))
		if newpass != repass {
			beego.Info("两次密码不一致")
			c.Data["errMsg"] = "两次密码不一致"
			c.Layout = "admin/layout.tpl"
			c.TplName = "admin/user.tpl"
			return
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

	c.Layout = "admin/layout.tpl"
	c.TplName = "admin/user.tpl"
}
