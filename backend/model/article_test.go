package model

import (
	"testing"
)
var article Article
var tags []Tag
// 父测试流程
func TestArticleWorkFlow(t *testing.T) {
	setup()
	defer db.Close()
	db.Where("title in (?)", []string{"golang", "vue", "github"}).Find(&tags)
	article = Article{
		CategoryName:"后端",
		CommentsCount:8,
		ArticleContent:ArticleContent{RawContent:"# RawContent", HTMLContent:"<h1>HTMLContent<h1>"},
		Hot:true,
		HotIndex:1.2,
		LikeCount:66,
		Original:true,
		RankIndex:22,
		Screenshot:"https://user-gold-cdn.xitu.io/2019/11/29/16eb551ff1574d77?w=1226&h=732&f=png&s=191910",
		SummaryInfo:"this is a test",
		Tags:tags,
		Title:"推荐几个不错的console调试技巧",
		UserID:60,
	}

	t.Run("Create", testCreateArticle)
	t.Run("Get", testGetArticle)
	t.Run("Del", testDeleteArticle)
	t.Run("Reget", testRegetArticle)
}




// 子测试
func testCreateArticle(t *testing.T) {
	if err := db.Create(&article).Error; err != nil {
		t.Errorf("Error of CreateArticle: %v", err)
	}
}

func testGetArticle(t *testing.T) {
	var article Article
	if err:=db.Set("gorm:auto_preload", true).Where("title='推荐几个不错的console调试技巧'").First(&article).Error; err != nil {
		t.Errorf("Error of GetArticle: %v", err)
	}
	println("")
}

func testDeleteArticle(t *testing.T) {
	if err:=db.Set("gorm:auto_preload", true).Where("title='推荐几个不错的console调试技巧'").Limit(1).Delete(&Article{}).Error; err != nil {
		t.Errorf("Error of GetArticle: %v", err)
	}
	//println("")
}

func testRegetArticle(t *testing.T) {

}
