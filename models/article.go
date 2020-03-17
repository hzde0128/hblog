package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Article struct {
	Id       int       // 文章ID
	Title    string    // 文章标题
	Tags     []*Tag    `orm:"reverse(many)"` // 文章标签 多对多 一篇文章可能有多个标签，一个标签对应多篇文章
	Category *Category `orm:"rel(fk)"`       // 文章分类，一对多
	Content  string    `orm:"size(10000)"`   // 文章内容
	Status   bool      // 状态 草稿/发布
	CreateAt time.Time `orm:"auto_now_add;type(datetime)"` // 创建时间
}

func (a *Article) TableName() string {
	return "article"
}

func (a *Article) Insert() error {
	if _, err := orm.NewOrm().Insert(a); err != nil {
		return err
	}
	return nil
}

func (a *Article) Read(fields ...string) error {
	if err := orm.NewOrm().Read(a, fields...); err != nil {
		return err
	}
	return nil
}

func (a *Article) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}

func (a *Article) Delete() error {
	if _, err := orm.NewOrm().Delete(a); err != nil {
		return err
	}
	return nil
}

func (a *Article) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(a)
}
