package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// Category 分类表
type Category struct {
	Id           int
	CategoryName string     `orm:"unique"` // 分类名称
	Order        int        // 排序
	Status       bool       // 启动或禁用
	Article      []*Article `orm:"reverse(many)"`               // 分类下面的文章
	CreateAt     time.Time  `orm:"auto_now_add;type(datetime)"` // 添加时间
}

func (c *Category) TableName() string {
	return "category"
}

func (c *Category) Insert() error {
	if _, err := orm.NewOrm().Insert(c); err != nil {
		return err
	}
	return nil
}

func (c *Category) Read(fields ...string) error {
	if err := orm.NewOrm().Read(c, fields...); err != nil {
		return err
	}
	return nil
}

func (c *Category) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(c, fields...); err != nil {
		return err
	}
	return nil
}

func (c *Category) Delete() error {
	if _, err := orm.NewOrm().Delete(c); err != nil {
		return err
	}
	return nil
}

func (c *Category) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(c)
}
