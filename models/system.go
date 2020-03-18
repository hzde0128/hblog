package models

import "github.com/astaxie/beego/orm"

type System struct {
	Id        int
	Title     string // 网站标题
	BaseUrl   string // 网站根目录
	CopyRight string // 版权信息
}

// 指定表名 user
func (s *System) TableName() string {
	return "system"
}

func (s *System) Insert() error {
	if _, err := orm.NewOrm().Insert(s); err != nil {
		return err
	}
	return nil
}

func (s *System) Read(fields ...string) error {
	if err := orm.NewOrm().Read(s, fields...); err != nil {
		return err
	}
	return nil
}

func (s *System) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(s, fields...); err != nil {
		return err
	}
	return nil
}

func (s *System) Delete() error {
	if _, err := orm.NewOrm().Delete(s); err != nil {
		return err
	}
	return nil
}

func (s *System) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(s)
}