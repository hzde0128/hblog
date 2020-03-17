package models

// Category 分类表
type Category struct {
	Id      int
	CategoryName    string     `orm:"unique"` // 分类名称
	Order   int        // 排序
	Status  bool       // 启动或禁用
	Article []*Article `orm:"reverse(many)"`
}

func (c *Category) TableName() string {
	return "category"
}