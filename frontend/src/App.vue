<template>
  <div id="app">
  <!-- 路由出口 -->
  <!-- 路由匹配到的组件将渲染在这里 -->
  <router-view></router-view>
  </div>
</template>

<script>

// import Home from "@/views/Home.vue";
// import About from "@/views/About.vue";


export default {
  name: "app",
  components: {
    // Home,
    // About,
  },
  data() {
    return {
   
    }
  },
  methods:{}

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
