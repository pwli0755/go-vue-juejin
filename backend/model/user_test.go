package model

import (
	"testing"
)

var user *User
// 父测试流程
func TestUserWorkFlow(t *testing.T) {
	setup()
	defer db.Close()
	t.Run("Create", testCreateUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

// 子测试
func testCreateUser(t *testing.T) {

}

func testGetUser(t *testing.T) {

}

func testDeleteUser(t *testing.T) {

}

func testRegetUser(t *testing.T) {

}
