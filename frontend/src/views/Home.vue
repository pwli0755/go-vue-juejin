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
    <auth />
  </div>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'
import MyHeader from "@/components/MyHeader.vue";
import TagBar from "@/components/TagBar.vue";
import FeedHeader from "@/components/FeedHeader.vue";
import FeedData from "@/components/FeedData.vue";
import Auth from "@/views/Auth.vue";
import RegForm from "@/components/RegForm.vue";

export default {
  name: "home",
  components: {
    MyHeader,
    TagBar,
    FeedHeader,
    FeedData,
    RegForm,
    Auth
  },
  data() {
    return {};
  },
  methods: {}
};
</script>

<style lang="less" scoped>
.container {
  margin-top: 105px;
  display: flex;
  justify-content: center;
  .feed {
    display: flex;
    flex-direction: column;
  }
}
</style>
