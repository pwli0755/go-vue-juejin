package model

import (
	"github.com/satori/go.uuid"
)

// 文章模型
type Article struct {
	Base
	Category Category `gorm:"foreignkey:CategoryName;association_foreignkey:Name"` // 分类
	CategoryName string
	CommentsCount int `gorm:"not null;default 0"`  // 这里的默认值只能标识出该字段在数据库有默认值
	ArticleContent ArticleContent `gorm:"foreignkey:ArticleID"`
	Hot bool
	HotIndex float64
	LikeCount int `gorm:"not null;default 0"`
	Original bool
	RankIndex float64
	Screenshot string // 封面url
	SummaryInfo string `gorm:"size:100;type:varchar(100)"` // 摘要
	Tags []Tag `gorm:"many2many:article_tags;association_foreignkey:Title;association_autoupdate:false"`
	Title string `gorm:"size:100;type:varchar(100)"`
	User User
	UserID uint
}

type ArticleContent struct {
	Base
	ArticleID uuid.UUID
	RawContent string `gorm:"type:MEDIUMTEXT"` // MD原文
	HTMLContent string `gorm:"type:MEDIUMTEXT"` // HTML
}