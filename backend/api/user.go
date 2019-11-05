package api

import (
	"backend/cache"
	"backend/serializer"
	"backend/service"
	"backend/util"
	"reflect"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		if user, err := service.Register(); err != nil {
			c.JSON(200, err)
		} else {
			res := serializer.BuildUserResponse(user, serializer.Response{Msg: "注册成功"})
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		if user, err := service.Login(); err != nil {
			c.JSON(200, err)
		} else {
			// 设置Session
			s := sessions.Default(c)
			// 限制同时只能有一个设备登录
			sessionPrefix := "session_"
			// 通过反射获取session ID
			ID := reflect.ValueOf(s).Elem().FieldByName("session").Elem().FieldByName("ID").String()
			sessionID := sessionPrefix + ID
			// 将上次登录生成的session清空
			sessionIdx := cache.SessionIdxPrefix + strconv.FormatUint(uint64(user.ID), 10)
			// 获取上次登陆的session ID
			if previousSsnID, err := cache.RedisClient.Get(sessionIdx).Result(); err == nil {
				_, err := cache.RedisClient.Del(previousSsnID).Result()
				if err != nil {
					util.Log.Error("Delete previous session failed!")
				}
			}
			s.Clear()
			s.Set("user_id", user.ID)
			s.Save()

			// 将sessionID存储至redis
			cache.RedisClient.Set(sessionIdx, sessionID, 0)

			res := serializer.BuildUserResponse(user, serializer.Response{Msg: "登录成功"})
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUserResponse(*user, serializer.Response{Msg: "用户详情"})
	c.JSON(200, res)
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Status: 0,
		Msg:    "登出成功",
	})
}
