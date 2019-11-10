import HTTP from '@/api/'

// 获取文章列表
const getArticles = () => HTTP.get('/api/v1/articles');


export {
    getArticles,
};
