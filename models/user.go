package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// User 用户表结构体
type User struct {
	Id         int       // 自增Id
	UserName   string    `orm:"unique;size(32)"`             // 用户名
	Password   string    `orm:"size(32)"`                    // 密码
	Email      string    `orm:"size(50)"`                    // Email
	LastLogin  time.Time `orm:"auto_now_add;type(datetime)"` // 上次登录时间
	LoginCount int       // 登录次数
	LastIp     string    `orm:"size(32)"` // 上次登录IP
	Active     int8      // 激活状态
}

// 指定表名 user
func (m *User) TableName() string {
	return "user"
}

func (m *User) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *User) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	if created, id, err := orm.NewOrm().ReadOrCreate(m, field, fields...); err != nil {
		return false, 0, err
	} else {
		return created, id, nil
	}
}
