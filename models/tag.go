package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// Tag 标签
type Tag struct {
	Id       int
	TagName  string     // 标签名称
	Article  []*Article `orm:"rel(m2m)"`                    // 对应的文章
	CreateAt time.Time  `orm:"auto_now_add;type(datetime)"` // 添加时间
}

func (t *Tag) TableName() string {
	return "tag"
}

func (t *Tag) Insert() error {
	if _, err := orm.NewOrm().Insert(t); err != nil {
		return err
	}
	return nil
}

func (t *Tag) Read(fields ...string) error {
	if err := orm.NewOrm().Read(t, fields...); err != nil {
		return err
	}
	return nil
}

func (t *Tag) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

func (t *Tag) Delete() error {
	if _, err := orm.NewOrm().Delete(t); err != nil {
		return err
	}
	return nil
}

func (t *Tag) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(t)
}
