package service

import (
	"backend/errs"
	"backend/model"
	"backend/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Email    string `form:"email" json:"email" binding:"required,min=5,max=30"`
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=40"`
}

// Valid 验证表单
func (service *UserRegisterService) Valid() *serializer.Response {

	count := 0
	model.DB.Model(&model.User{}).Where("email = ?", service.Email).Count(&count)
	if count > 0 {
		return errs.BuildErrorResponse(errs.ERR_EMAIL)
	}

	count = 0
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return errs.BuildErrorResponse(errs.ERR_USER_NAME)
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() (model.User, *serializer.Response) {
	user := model.User{
		Email:    service.Email,
		UserName: service.UserName,
		Status:   model.Inactive,
	}

	// 表单验证
	if err := service.Valid(); err != nil {
		return user, err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return user, errs.BuildErrorResponse(errs.ERR_GEN_PASSWD)
	}

	// 创建用户
	if err := model.CreateUser(&user); err != nil {
		return user, errs.BuildErrorResponse(errs.ERR_REGISTER)
	}
	// 发送激活邮件
	// TODO

	return user, nil
}
