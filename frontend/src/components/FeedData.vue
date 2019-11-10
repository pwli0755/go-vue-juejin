<template>
  <div class="feed-data">
    <ul class="entry-list">
      <li class="item" v-for="article in articles" :key="article.id">
        <div class="entry-box">
          <div class="entry">
            <a :href="/post/+article.id" target="_blank" class="entry-link">
              <!-- 条目内容 -->
              <div class="content-box">
                <div class="info-box">
                  <!-- meta信息 专栏/作者/日期/tag -->
                  <div class="info-row meta-row">
                    <p class="hot">热</p>
                    <p class="post">专栏</p>
                    <p class="username">
                      <a href="/userhome" target="_blank">{{article.author}}</a>
                    </p>
                    <p class="time">{{article.createdAt}}</p>
                    <p class="tag">
                      <a href="/tag" target="_blank">{{article.tag}}</a>
                    </p>
                  </div>
                  <!-- title 标题-->
                  <div class="title-row">
                    <a class="title">{{article.title}}</a>
                  </div>
                  <!-- action 点赞 评论 -->
                  <div class="action-row">
                    <div class="item">
                      <p class="like">
                        <a href="ilike" target="_blank">
                          <img src=https://b-gold-cdn.xitu.io/v3/static/img/zan.e9d7698.svg alt="" class="icon">
                          <span
                            class="count"
                          >{{article.thumb}}</span>
                        </a>
                      </p>
                    </div>
                    <div class="item">
                      <p class="comment">
                        <a href="comment" target="_blank">
                          <img
                            src="https://b-gold-cdn.xitu.io/v3/static/img/comment.4d5744f.svg"
                            alt
                            class="icon"
                          />
                          <span class="count">{{article.comment}}</span>
                        </a>
                      </p>
                    </div>
                    <div class="item">
                      <p class="share">
                        <a href="share" target="_blank">
                          <img
                            src="https://b-gold-cdn.xitu.io/v3/static/img/share.1d55e69.svg"
                            alt
                            class="icon"
                          />
                        </a>
                      </p>
                    </div>
                  </div>
                </div>
              </div>
            </a>
          </div>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
import * as Article from '@/api/article/'
export default {
  data() {
    return {
      articles: [],
    };
  },
  mounted(){
    this.getArticles();
  },
  methods:{
    getArticles(){
      Article.getArticles().then(res=>{
        console.log(res);
        this.articles = res.data;
      })
    },
  },
};
</script>

<style lang="less" scoped>
.entry-list {
  width: 100%;
  background-color: #fff;
}
ul {
  padding: 0;
  margin: 0;
}
// .item:not(:last-child) {
//   border-bottom: 1px solid rgba(178, 186, 194, 0.15);
// }
li {
  list-style: none;
}
.entry-box {
  width: 700px;
  height: 115px;
}
.entry {
  width: 100%;
  height: 100%;
}
.content-box {
  display: flex;
  align-items: center;
  padding: 14px 16px;
  border-bottom: 1px solid rgba(178, 186, 194, 0.15);
}
.meta-row {
  display: flex;
  color: #b2bac2;
}
.meta-row a {
  color: #b2bac2;
}
.meta-row .hot {
  color: red;
  font-weight: 500;
}
.meta-row .post {
  color: #b71ed7;
}
.meta-row p {
  margin-right: 1px;
  font-size: 12px;
}
.meta-row p:not(:last-of-type)::after {
  content: "·";
  margin: 4px;
}
.title-row {
  padding: 3px 0 6px 0;
}
.title-row a.title {
  font-size: 18px;
  font-weight: 600;
  line-height: 1.2;
  color: #2e3135;
}
.action-row {
  display: flex;
  justify-content: flex-start;
  margin-top: 5px;
}
.action-row a,
.action-row span {
  color: #b2bac2;
  margin-left: 3px;
  font-weight: 700;
}
.action-row .item {
  margin-left: -1px;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 24px;
  border: 1px solid #edeeef;
  border-radius: 1px;
  padding: 0 2px;
}
.action-row p,
.action-row a {
  display: inline-flex;
  justify-content: center;
  padding: 0 2px;
}
</style>