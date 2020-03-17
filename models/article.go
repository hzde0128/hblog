package models

import "time"

type Article struct {
	Id       int       // 文章ID
	Title    string    // 文章标题
	Tags     []*Tag    `orm:"rel(m2m)"` // 文章标签 多对多 一篇文章可能有多个标签，一个标签对应多篇文章
	Category *Category  `orm:"rel(fk)"`  // 文章分类，一对多
	Content  string    // 文章内容
	Status   bool      // 状态 草稿/发布
	CreateAt time.Time // 创建时间
}

func (a *Article) TableName() string {
	return "article"
}
