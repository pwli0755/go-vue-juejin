<template>
  <div
    v-loading="loading"
    element-loading-text="拼命加载中"
    element-loading-spinner="el-icon-loading"
    element-loading-background="rgba(0, 0, 0, 0.8)"
  >
    <img src="../assets/loading.gif" alt="cat" />
  </div>
</template>

<script>
import * as API from "@/api/auth";
export default {
  data() {
    return {
      loading: true
    };
  },
  mounted() {
    setTimeout(() => {
      API.activateUser(this.$route.query.token)
        .then(res => {
          this.loading = false;
          if (res.data) {
            this.$notify({
              title: "success",
              message: "激活成功",
              type: "success"
            });
            this.$router.push({ name: "home" });
          } else {
            this.$notify.error({
              title: "error",
              message: "token错误",
              duration: 0
            });
          }
        })
        .catch(err => {
          this.loading = false;
          this.$notify.warning({
            title: "error",
            //   message: "网络开小差了",
            message: err,
            duration: 0
          });
        });
    }, 2000);
  }
};
</script>

<style lang="less" scoped>
div {
  display: flex;
  justify-content: center;
  font-size: 36px;
  img {
    width: 100%;
    height: 100%;
  }
}
</style>