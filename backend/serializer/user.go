package serializer

import "backend/model"

// User 用户序列化器，这里也可以直接使用model里面的定义加上json tag
type User struct {
	ID          uint   `json:"id"`
	UserName    string `json:"user_name"`
	Nickname    string `json:"nickname"`
	Status      string `json:"status"`
	Avatar      string `json:"avatar"`
	CreatedAt   int64  `json:"created_at"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

// UserResponse 单个用户序列化
type UserResponse struct {
	Response
	Data User `json:"data"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:          user.ID,
		UserName:    user.UserName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Status:      user.Status,
		Avatar:      user.Avatar,
		CreatedAt:   user.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User, resp Response) UserResponse {
	return UserResponse{
		Response: resp,
		Data:     BuildUser(user),
	}
}
