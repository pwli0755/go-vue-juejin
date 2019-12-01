package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
)

var tagNames = []string{"golang", "vue", "Java", "C++", "C", "Python", "JavaScript", "NodeJS",
	"shell", "大数据", "云计算", "Docker", "linux", "mysql", "redis", "k8s", "数据分析", "AI", "git", "github"}

// 父测试流程
func TestTagWorkFlow(t *testing.T) {
	setup()
	defer db.Close()
	t.Run("Create", testCreateTag)
	t.Run("Get", testGetTag)
	t.Run("Del", testDeleteTag)
	t.Run("Reget", testRegetTag)
}

// 子测试
func testCreateTag(t *testing.T) {
	for _, name := range tagNames {
		if err := db.Model(&Tag{}).Unscoped().Where("title=?", name+"_test").Error; err != gorm.ErrRecordNotFound {
			db.Unscoped().Where("title=?", name+"_test").Delete(&Tag{})
		}
		if err := db.Create(&Tag{Title: name + "_test"}).Error; err != nil {
			t.Errorf("Error of CreateTag: %v", err)
		}
	}
}

func testGetTag(t *testing.T) {
	for _, name := range tagNames {
		var tag Tag
		if err := db.Where(&Tag{Title: name + "_test"}).First(&tag).Error; err != nil {
			t.Errorf("Error of GetTag: %v", err)
		}
	}
}

func testDeleteTag(t *testing.T) {
	for _, name := range tagNames {
		if err := db.Where("title = ?", name+"_test").Delete(&Tag{}).Error; err != nil {
			t.Errorf("Error of DeleteTag: %v", err)
		}
	}
}

func testRegetTag(t *testing.T) {
	for _, name := range tagNames {
		var tag Tag
		if err := db.Where(&Tag{Title: name+"_test"}).First(&tag).Error; err == nil {
			t.Errorf("Error of RegetTag: %v", err)
		}
	}
}
