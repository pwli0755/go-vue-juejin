import axios from 'axios';

// 注册
const register = form => axios.post('/api/v1/user/register', form);

// 登录
const login = form => axios.post(`/api/v1/user/login`, form);

// 获取用户信息
const getUserInfo = userID => axios.get('/api/v1/user/me', { params: { userID,  } });

export {
  register,
  login,
  getUserInfo,
};
