package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// User 用户表结构体
type User struct {
	Id         int       // 自增Id
	UserName   string    `orm:"unique;size(32)"`         // 用户名
	Password   string    `orm:"size(32)"`                // 密码
	Email      string    `orm:"size(50)"`                // Email
	LastLogin  time.Time `orm:"auto_now;type(datetime)"` // 上次登录时间
	LoginCount int       // 登录次数
	LastIp     string    `orm:"size(32)"` // 上次登录IP
	Active     int8      // 激活状态
}

// 指定表名 user
func (m *User) TableName() string {
	return "user"
}

func (u *User) Insert() error {
	if _, err := orm.NewOrm().Insert(u); err != nil {
		return err
	}
	return nil
}

func (u *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *User) Delete() error {
	if _, err := orm.NewOrm().Delete(u); err != nil {
		return err
	}
	return nil
}

func (u *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(u)
}

func (u *User) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	if created, id, err := orm.NewOrm().ReadOrCreate(u, field, fields...); err != nil {
		return false, 0, err
	} else {
		return created, id, nil
	}
}
