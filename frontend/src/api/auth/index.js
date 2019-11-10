import HTTP from '@/api/'

// 注册
const register = form => HTTP.post('/api/v1/user/register', form);

// 登录
const login = form => HTTP.post(`/api/v1/user/login`, form);

// 获取用户信息
const getUserInfo = userID => HTTP.get('/api/v1/user/me', { params: { userID,  } });

// 激活
const activateUser = token => HTTP.get('/api/v1/activate', {params: {token}})

export {
  register,
  login,
  getUserInfo,
  activateUser,
};
