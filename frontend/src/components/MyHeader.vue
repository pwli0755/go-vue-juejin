<template>
  <div class="header">
    <div class="container">
      <a href="index.htm" tppabs="https://juejin.im/" class="logo">
        <img src="../assets/logo.svg" alt="掘金" class="logo-img" />
      </a>
      <nav class="main-nav">
        <ul class="nav-list">
          <li class="main-nav-list">
            <ul class="main-nav-list">
              <li class="nav-item link-item" data-v-4bf12e3a>
                <a href="index.htm">首页</a>
              </li>
              <li class="nav-item link-item" data-v-4bf12e3a>
                <a href="index.htm">沸点</a>
              </li>
              <li class="nav-item link-item" data-v-4bf12e3a>
                <a href="index.htm">话题</a>
              </li>
              <li class="nav-item link-item" data-v-4bf12e3a>
                <a href="index.htm">小册</a>
              </li>
              <li class="nav-item link-item" data-v-4bf12e3a>
                <a href="index.htm">活动</a>
              </li>
              <li class="nav-item link-item" data-v-4bf12e3a>
                <a href="index.htm">1024</a>
              </li>
            </ul>
          </li>
          <li class="nav-item search">
            <form role="search" class="search-form">
              <input type="search" maxlength="32" placeholder="搜索掘金" value class="search-input" />
              <img src="../assets/search.svg" alt="搜索" class="search-icon" />
            </form>
          </li>
          <!-- 文章 -->
          <li class="nav-item submit">
            <img src="../assets/article.svg" class="icon" />
            <span>写文章</span>
            <!---->
          </li>
          <li class="nav-item auth">
            <span class="login" @click="showLogin">登录</span>
            <span class="register" @click="showReg">注册</span>
          </li>
        </ul>
      </nav>
    </div>
  </div>
</template>

<script>
import logoUrl from "../assets/logo.svg";
import { mapActions } from 'vuex'

export default {
  data() {
    return {
      logoUrl: logoUrl,

    };
  },
  methods:{
     ...mapActions([
      'showAuth', // 将 `this.increment()` 映射为 `this.$store.dispatch('increment')`

      // `mapActions` 也支持载荷：
      'switchAuthType' // 将 `this.incrementBy(amount)` 映射为 `this.$store.dispatch('incrementBy', amount)`
    ]),
    showLogin(){
      this.switchAuthType('login');
      this.showAuth();
    },
    showReg(){
      this.switchAuthType('reg');
      this.showAuth();
    }
  }
};
</script>


<style lang="less" scoped>
.header {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  border-bottom: 1px solid #f1f1f1;
  color: #909090;
  height: 60px;
  z-index: 250;
    background-color:  rgb(255, 255, 255);
}
.container {
  max-width: 960px;
  margin: auto;
  display: flex;
  -webkit-box-align: center;
  align-items: center;
  height: 100%;
}
.logo {
  min-width: 98px;
  margin-right: 2rem;
}
.main-nav {
  height: 100%;
  flex: 1 0 auto;
  flex-grow: 1;
  flex-shrink: 0;
  flex-basis: auto;
}
.nav-list {
  display: flex;
  -webkit-box-align: center;
  align-items: center;
  justify-content: flex-end;
  position: relative;
  height: 100%;
  margin: 0;
}
ul {
  padding: 0;
  margin: 0;
}
.main-nav-list {
  display: flex;
}
.nav-item.link-item {
  padding: 0 1.5rem;
  height: 60px;
}

.nav-item {
  color: #71777c;
  padding: 0 1.2rem;
  // font-size: 1.33rem;
  margin: 0;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
}
.nav-item {
  position: relative;
  cursor: pointer;
}
li {
  list-style: none;
}
.nav-item a {
  color: #71777c;
}

a {
  text-decoration: none;
  cursor: pointer;
  color: #909090;
}
a:hover {
  color: #007fff;
}
.nav-item.search {
  flex: 1 1 auto;
  justify-content: flex-end;
  cursor: auto;
}

.search-form {
  border: 1px solid hsla(0, 0%, 59.2%, 0.2);
  background-color: rgba(227, 231, 236, 0.2);
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-radius: 2px;
}
.search-form .search-input {
  border: none;
  width: 10rem;
  padding: 0.6rem 1rem;
  box-shadow: none;
  outline: none;
  color: #666;
  background-color: transparent;
}
.search-form .search-icon {
  padding: 0 0.5rem;
  cursor: pointer;
}

.nav-item.submit {
  color: #007fff;
  position: relative;
}
.nav-item {
  color: #71777c;
  padding: 0 1.2rem;
  margin: 0;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
}
.nav-item.submit:after {
    content: "|";
    position: absolute;
    top: 20px;
    left: 100%;
    color: hsla(0,0%,59.2%,.4);
}

.nav-item.auth {
    color: #007fff;
}
.nav-item:last-child {
    padding-right: 0;
}
.nav-item.auth .login:after {
    content: "·";
    margin: 0 .4rem;
}
</style>