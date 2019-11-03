package model_test

import (
	_ "backend/conf" // 初始化
	"backend/model"
	"strconv"
	"testing"
	"time"
)

var user *model.User
// 父测试流程
func TestUserWorkFlow(t *testing.T) {
	randStr := strconv.FormatInt(time.Now().UnixNano(), 10)
	user = &model.User{Email: "763045276@qq.com_test" + randStr,
		UserName:       "pwli_test" + randStr,
		PasswordDigest: "$2a$12$j2LuAFsRKwES211YqpLpY.Ao4Nm56dUVVrsIT4IKgoC5b6AQDwqoG"}
	t.Run("Create", testCreateUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

// 子测试
func testCreateUser(t *testing.T) {
	err := model.CreateUser(user)
	if err != nil {
		t.Errorf("Error of CreateUser: %v", err)
	}

}

func testGetUser(t *testing.T) {
	u, err := model.GetUser(user.ID)
	if u == (model.User{}) || err != nil {
		t.Errorf("Error of GetUser: %v", err)
	}
}

func testDeleteUser(t *testing.T) {
	if *user != (model.User{}) {
		err := model.DeleteUser(user)
		if err != nil {
			t.Errorf("Error of DeleteUser: %v", err)
		}
	}
}

func testRegetUser(t *testing.T) {
	u, err := model.GetUser(user.ID)
	if u != (model.User{}) && err == nil {
		t.Errorf("Error of RegetUser: %s", "删除失败")
	}
}
