<template>
  <div id="app">
    <MyHeader />
    <TagBar />

    <div class="container">
      <div class="feed">
        <feed-header />
        <feed-data />
      </div>
      <div class="aside">
        <reg-form />
      </div>
    </div>
    <auth/>

    <!-- <img src="./assets/logo.png" />
    <div>
      <p>
        If Element is successfully added to this project, you'll see an
        <code v-text="'<el-button>'"></code>
        below
      </p>
      <el-button>el-button</el-button>
    </div>-->
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
  </div>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'
import MyHeader from "./components/MyHeader.vue";
import TagBar from "./components/TagBar.vue";
import FeedHeader from "./components/FeedHeader.vue";
import FeedData from "./components/FeedData.vue";
import Auth from "./views/Auth.vue";
import RegForm from "./components/RegForm.vue";
// import AuthForm from "./components/AuthForm.vue";




export default {
  name: "app",
  components: {
    // HelloWorld,
    MyHeader,
    TagBar,
    FeedHeader,
    FeedData,
    RegForm,
    Auth
  },
  data() {
    return {
      showLogin:false,

    };
  },
  methods:{

  }
};
</script>

<style lang="less">
html {
  font-size: 14px;
}
body {
  margin: 0;
  padding: 0;
  background-color: #f5f5f5;
}
a {
  text-decoration: none;
}
#app {
  font-family: -apple-system, system-ui, BlinkMacSystemFont, Helvetica Neue,
    PingFang SC, Hiragino Sans GB, Microsoft YaHei, Arial, sans-serif;
}
.container {
  display: flex;
  justify-content: center;
  .feed {
    display: flex;
    flex-direction: column;
  }
}
</style>
