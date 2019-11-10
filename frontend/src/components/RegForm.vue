<template>
  <div class="auth-form">
    <section class="shadow section auth-section">
      <template v-if="isPop">
        <div class="isPop">
          <h1 class="title">注册</h1>
          <i title="关闭" class="close-btn ion-close-round" @click="hideAuth">×</i>
        </div>
      </template>
      <template v-else>
        <div class="title">掘金 - juejin.im</div>
        <div class="slogan">一个帮助开发者成长的社区</div>
      </template>

      <div class="input-group">
        <div class="input-box">
          <input
            name="registerUsername"
            maxlength="20"
            placeholder="用户名"
            class="input"
            v-model="user_name"
          />
        </div>
        <div class="input-box">
          <input
            name="registerPhoneNumber"
            maxlength="64"
            placeholder="邮箱"
            class="input"
            v-model="email"
          />
        </div>
        <!---->
        <!---->
        <div class="input-box">
          <input
            name="registerPassword"
            type="password"
            maxlength="64"
            autocomplete="new-password"
            placeholder="密码（不少于 6 位）"
            class="input"
            v-model="password"
          />
        </div>
      </div>
      <button class="btn submit-btn" @click="register" :disabled="regbtn!='立即注册'">{{regbtn}}</button>
      <div class="switch" v-if="isPop" @click="switchAuthType('login')">已有账号登录</div>
      <div class="oauth-box">
        第三方登录：
        <img
          title="微博"
          alt="微博"
          src="https://b-gold-cdn.xitu.io/v3/static/img/weibo.fa758eb.svg"
          class="oauth-btn"
        />
        <img
          title="GitHub"
          alt="GitHub"
          src="https://b-gold-cdn.xitu.io/v3/static/img/github.547dd8a.svg"
          class="oauth-btn"
        />
        <img
          title="微信"
          alt="微信"
          src="https://b-gold-cdn.xitu.io/v3/static/img/wechat.e0ff124.svg"
          class="oauth-btn"
        />
      </div>
      <div class="agreement-box">
        注册登录即表示
        <br />同意
        <a href="/terms" target="_blank">用户协议</a>、
        <a href="/privacy" target="_blank">隐私政策</a>
      </div>
    </section>
  </div>
</template>

<script>
import { mapActions } from "vuex";
import * as API from "@/api/auth/";
export default {
  data() {
    return {
      user_name: "",
      password: "",
      email: "",
      regbtn: "立即注册"
    };
  },
  props: {
    isPop: {
      type: Boolean,
      default: false
    }
  },
  methods: {
    ...mapActions([
      "hideAuth", // 将 `this.hideAuth()` 映射为 `this.$store.dispatch('hideAuth')`
      "switchAuthType"
    ]),
    register() {
      if (this.user_name == "") {
        this.$notify({
          message: "请输入用户名",
          customClass: "error",
          showClose: false,
          duration: 2000
        });
        return;
      } else if (this.password == "") {
        this.$notify({
          message: "请输入密码",
          customClass: "error",
          showClose: false,
          duration: 2000
        });
        return;
      } else if (this.email == "") {
        this.$notify({
          message: "请输入密码",
          customClass: "error",
          showClose: false,
          duration: 2000
        });
        return;
      }
      var reg = /^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(.[a-zA-Z0-9_-]+)+$/;
      if (!reg.test(this.email)) {
        this.$notify({
          message: "请输入正确的邮箱",
          customClass: "error",
          showClose: false,
          duration: 2000
        });
        return;
      }
      // 灰化注册按钮
      this.regbtn = "请稍后...";
      API.register({
        user_name: this.user_name,
        password: this.password,
        email: this.email
      })
        .then(res => {
          this.$notify({
            message: res.data.msg || res.data.error,
            customClass: "error",
            showClose: false,
            duration: 2000
          });
          if (res.data.msg && res.data.msg == "注册成功") {
            this.$store.state.user = res.data.data;
            // this.changeLoginState();
            this.hideAuth();
          }
          this.regbtn = "立即注册";
        })
        .catch((err) => {
          console.log(err)
          this.$notify({
            message: err,
            customClass: "error",
            showClose: false,
            duration: 2000
          });
          this.regbtn = "立即注册";
        });
    }
  }
};
</script>

<style lang="less" scoped>
.auth-form {
  width: 240px;
  min-width: 240px;
  margin-left: 20px;
  margin-top: 25px;
}
.section {
  background-color: #fff;
  border-radius: 2px;
  margin-bottom: 1.5rem;
  overflow: hidden;
}
.auth-section {
  padding: 1.333rem;
  background-color: #fff;
}
.shadow {
  box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
}
.slogan,
.title {
  margin-bottom: 0.5rem;
  font-size: 14px;
  color: #2e3135;
}
.title {
  font-weight: 600;
}

.input-box {
  position: relative;
  margin-bottom: 0.833rem;
}
.input {
  padding: 0.7rem 0.6rem;
  width: 100%;
  font-size: 14px;
  background-color: #fbfbfb;
  border: 1px solid #f4f4f4;
  border-radius: 2px;
  outline: none;
  box-sizing: border-box;
  margin: initial;
  overflow: visible;
}

.submit-btn {
  padding: 0.7rem 0;
  width: 100%;
  font-size: 14px;
}

.button,
button {
  -webkit-appearance: none;
  appearance: none;
  background-color: #007fff;
  color: #fff;
  border-radius: 2px;
  border: none;
  padding: 0.5rem 1.3rem;
  outline: none;
  transition: background-color 0.3s, color 0.3s;
  cursor: pointer;
}
button {
  list-style: 1;
}
button {
  margin: initial;
}
.oauth-box {
  margin-top: 1.5rem;
  font-size: 14px;
  line-height: 1.9rem;
  color: #767676;
}

.oauth-box .oauth-btn {
  margin-left: 6px;
  height: 1.9rem;
  vertical-align: bottom;
  cursor: pointer;
}
img {
  width: auto;
  height: auto;
}

.agreement-box {
  margin-top: 1.667rem;
  font-size: 14px;
  line-height: 1.5;
  color: #767676;
}
.agreement-box a {
  color: #007fff;
}

a {
  text-decoration: none;
  cursor: pointer;
  color: #909090;
}
.isPop {
  display: flex;
  justify-content: space-between;
  align-items: baseline;
  h1.title {
    font-size: 18px;
    margin: 0 0 20px;
  }

  i {
    line-height: 18px;
    &:hover {
      font-weight: 600;
      cursor: pointer;
    }
  }
}
.switch {
  color: #007fff;
  cursor: pointer;
  font-size: 14px;
  text-align: center;
  margin: 6px 0;
}
</style>