package model

// 分类
type Category struct {
	Base
	Name string `sql:"not null; unique_index"`
}

