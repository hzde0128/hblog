package models

// Tag 标签
type Tag struct {
	Id       int
	TagName string
	Article  []*Article `orm:"reverse(many)"`
}

func (t *Tag) TableName() string {
	return "tag"
}
