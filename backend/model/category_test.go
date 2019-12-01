package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"testing"
	"time"
)

var db *gorm.DB

// init db
func setup() {
	var err error
	db, err = gorm.Open("mysql", os.Getenv("go_vue_test_dsn"))
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	//空闲
	db.DB().SetMaxIdleConns(10)
	//打开
	db.DB().SetMaxOpenConns(10)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * time.Duration(100))
}

var categoryNames = []string{"后端", "前端", "Android", "IOS", "人工智能", "开发工具", "代码人生", "阅读"}

// 父测试流程
func TestCategoryWorkFlow(t *testing.T) {
	setup()
	defer db.Close()
	t.Run("Create", testCreateCategory)
	t.Run("Get", testGetCategory)
	t.Run("Del", testDeleteCategory)
	t.Run("Reget", testRegetCategory)
}

// 子测试
func testCreateCategory(t *testing.T) {
	for _, name := range categoryNames {
		if err:=db.Model(&Category{}).Unscoped().Where("name=?", name+"_test").Error; err != gorm.ErrRecordNotFound {
			db.Unscoped().Where("name=?", name+"_test").Delete(&Category{})
		}
		if err := db.Create(&Category{Name: name+"_test"}).Error; err != nil {
			t.Errorf("Error of CreateCategory: %v", err)
		}
	}
}

func testGetCategory(t *testing.T) {
	for _, name := range categoryNames {
		var category Category
		if err := db.Where(&Category{Name:name+"_test"}).First(&category).Error; err != nil {
			t.Errorf("Error of GetCategory: %v", err)
		}
	}
}

func testDeleteCategory(t *testing.T) {
	for _, name := range categoryNames {
		if err := db.Where("name = ?", name+"_test").Delete(&Category{}).Error; err != nil {
			t.Errorf("Error of DeleteCategory: %v", err)
		}
	}
}

func testRegetCategory(t *testing.T) {
	for _, name := range categoryNames {
		var category Category
		if err := db.Where(&Category{Name:name+"_test"}).First(&category).Error; err == nil {
			t.Errorf("Error of RegetCategory: %v", err)
		}
	}
}
