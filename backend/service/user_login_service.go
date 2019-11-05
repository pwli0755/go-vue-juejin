package service

import (
	"backend/errs"
	"backend/model"
	"backend/serializer"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=40"`
}

// Login 用户登录函数
func (service *UserLoginService) Login() (model.User, *serializer.Response) {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return user, errs.BuildErrorResponse(errs.ERR_WRONG_USER_PASSWD)
	}

	if user.CheckPassword(service.Password) == false {
		return user, errs.BuildErrorResponse(errs.ERR_WRONG_USER_PASSWD)
	}
	return user, nil
}
