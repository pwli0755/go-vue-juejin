// 引入mock.js
import Mock from "mockjs"

// 获取 mock.Random 对象
const Random = Mock.Random
// mock新闻数据，包括新闻标题title、内容content、创建时间createdTime
const produceArticle = function () {
    let articleList = []
    for (let i = 0; i < 20; i++) {
        let articleObject = {
            title: Random.ctitle(8, 20), // Random.ctitle( min, max ) 随机产生一个中文标题，长度默认在3-7之间
            content: Random.cparagraph(50, 2000), // Random.cparagraph(min, max) 随机生成一个中文段落，段落里的句子个数默认3-7个
            createdTime: Random.date(), // Random.date()指示生成的日期字符串的格式,默认为yyyy-MM-dd；
            author: Random.cname(),  // 作者
            tag: Random.pick(['前端', '后端', '安卓', 'IOS', '人工智能']),  // 标签分类
            createdAt: Random.date('yy-MM-dd'), // 发表日期
            thumb: Random.natural(100, 9999),  // 点赞
            comment: Random.natural(50, 6666),  // 评论
            id: Random.id(), // id

        }
        articleList.push(articleObject)
    }
    return articleList.slice(0, articleList.length);
}

const produceActivate=function(){
    let valid = Random.boolean;
    let msg = valid?"激活成功":"token无效";
    return {valid, msg}
}
// 请求该url，就可以返回articleList
Mock.mock('/api/v1/articles','get', produceArticle)

Mock.mock(/\/api\/v1\/activate.*/,'get', produceActivate)




/*
Mock.mock( rurl, rtype, template|function( options ) )
rurl
可选。
表示需要拦截的 URL，可以是 URL 字符串或 URL 正则。例如 '/domian/list.json'。
rtype
可选。
表示需要拦截的 Ajax 请求类型。例如 GET、POST、PUT、DELETE 等。
template
可选。
表示数据模板，可以是对象或字符串。
数据模板中的每个属性由 3 部分构成：属性名、生成规则、属性值：
// 属性名 name
// 生成规则 rule
// 属性值 value
'name|rule': value
例如：'name|1-10':1 会产生一个1-10之间的整数，详细规则参见官方文档
function(options)
可选。
表示用于生成响应数据的函数。
options
指向本次请求的 Ajax 选项集，含有 url、type 和 body 三个属性
*/