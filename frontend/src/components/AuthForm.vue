<template>
  <form class="auth-form">
    <div class="panfish">
      <img src="https://b-gold-cdn.xitu.io/v3/static/img/normal.0447fe9.png" class="normal" style />
      <img
        src="https://b-gold-cdn.xitu.io/v3/static/img/greeting.1415c1c.png"
        class="greeting"
        style="display: none;"
      />
      <img
        src="https://b-gold-cdn.xitu.io/v3/static/img/blindfold.58ce423.png"
        class="blindfold"
        style="display: none;"
      />
    </div>
    <i title="关闭" class="close-btn ion-close-round" @click="hideAuth"></i>
    <div class="panel">
      <h1 class="title">登录</h1>
      <!---->
      <div class="input-group">
        <div class="input-box">
          <input
            v-model="user_name"
            name="loginPhoneOrEmail"
            maxlength="64"
            placeholder="请输入用户名或邮箱"
            class="input"
          />
        </div>
        <div class="input-box">
          <input
            v-model="password"
            name="loginPassword"
            type="password"
            maxlength="64"
            placeholder="请输入密码"
            class="input"
          />
        </div>
      </div>
      <button class="btn" :disabled="loginbtn!='登录'" @click.prevent="login">{{loginbtn}}</button>
      <div class="prompt-box">
        没有账号？
        <span class="clickable" @click="switchAuthType('reg')">注册</span>
        <a href="/reset-password" class="right clickable">忘记密码</a>
      </div>
    </div>
    <div class="oauth-box">
      <div class="hint">第三方账号登录：</div>
      <div class="oauth">
        <div class="oauth-bg">
          <img
            title="微博"
            alt="微博"
            src="https://b-gold-cdn.xitu.io/v3/static/img/weibo.fa758eb.svg"
            class="oauth-btn"
          />
        </div>
        <div class="oauth-bg">
          <img
            title="微信"
            alt="微信"
            src="https://b-gold-cdn.xitu.io/v3/static/img/wechat.e0ff124.svg"
            class="oauth-btn"
          />
        </div>
        <div class="oauth-bg">
          <img
            title="GitHub"
            alt="GitHub"
            src="https://b-gold-cdn.xitu.io/v3/static/img/github.547dd8a.svg"
            class="oauth-btn"
          />
        </div>
      </div>
    </div>
    <div class="agreement-box">
      注册登录即表示同意
      <a href="/terms" target="_blank">用户协议</a>、
      <a href="/privacy" target="_blank">隐私政策</a>
    </div>
  </form>
</template>

<script>
import { mapActions, mapMutations } from "vuex";
import * as API from "@/api/auth/";
export default {
  data() {
    return {
      user_name: "",
      password: "",
      loginbtn: "登录"
    };
  },
  methods: {
    ...mapActions([
      "hideAuth", // 将 `this.hideAuth()` 映射为 `this.$store.dispatch('hideAuth')`
      "switchAuthType"
    ]),
    ...mapMutations(["setLogin"]),
    login() {
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
      }
      // 灰化登录按钮
      this.loginbtn = "登录中...";
      API.login({ user_name: this.user_name, password: this.password })
        .then(res => {
          this.$notify({
            message: res.data.msg || res.data.error,
            customClass: "error",
            showClose: false,
            duration: 2000
          });
          if (res.data.msg && res.data.msg == "登录成功") {
            this.$store.state.user = res.data.data;
            this.setLogin();
            this.hideAuth();
          } 
        })
        .catch(() => {
          this.$notify({
            message: "网络开小差了，请稍后再试",
            customClass: "error",
            showClose: false,
            duration: 2000
          });
        }).finally(()=>{
          this.loginbtn = "登录";
        });
    }
  }
};
</script>

<style lang="less" scoped>
.auth-form {
  position: relative;
  padding: 24px;
  width: 320px;
  max-width: 100%;
  font-size: 14px;
  background-color: #fff;
  border-radius: 2px;
  box-sizing: border-box;
}
.panfish .normal {
  position: absolute;
  top: 0;
  left: 50%;
  width: 10rem;
  transform: translate(-50%, -83%);
  z-index: 1;
}
.panfish .greeting {
  position: absolute;
  top: 0;
  left: 50%;
  width: 10rem;
  transform: translate(-50%, -75.8%);
  z-index: 1;
}
.panfish .blindfold {
  position: absolute;
  top: 0;
  left: 50%;
  width: 8.6rem;
  transform: translate(-50%, -75%);
  z-index: 1;
}
.close-btn {
  float: right;
  cursor: pointer;
  opacity: 0.4;
}
.ion-close-round:before {
  content: "×";
}
i:hover {
  font-weight: 600;
}
.title {
  font-size: 18px;
  margin: 0 0 20px;
}
.input-group {
  margin-bottom: 5px;
  overflow: hidden;
}
.input-box {
  position: relative;
  margin-bottom: 8px;
}
.input {
  padding: 10px;
  width: 100%;
  border: 1px solid #e9e9e9;
  border-radius: 2px;
  outline: none;
  box-sizing: border-box;
}
a,
button,
input {
  margin: initial;
}
button,
input {
  overflow: visible;
}
.btn {
  width: 100%;
  height: 40px;
  color: #fff;
  background-color: #007fff;
  border-radius: 2px;
  outline: none;
  box-sizing: border-box;
  cursor: pointer;
  border: none;
}

.prompt-box {
  margin: 12px 0 0;
  color: #767676;
}
.prompt-box .clickable {
  color: #007fff;
  cursor: pointer;
}
.prompt-box .right {
  float: right;
}
a {
  text-decoration: none;
  cursor: pointer;
  color: #909090;
  margin: initial;
}
.oauth-box {
  margin-top: 14px;
  line-height: 23px;
  color: #767676;
}
.oauth-box .oauth {
  display: flex;
  align-items: center;
  justify-content: space-around;
  margin-top: 15px;
}
.oauth-box .oauth-bg {
  width: 45px;
  height: 45px;
  border-radius: 50%;
  background-color: #f4f8fb;
  display: flex;
  align-items: center;
  justify-content: center;
}
.oauth-box .oauth-btn {
  height: 23px;
  vertical-align: bottom;
  cursor: pointer;
}
img {
  width: auto;
  height: auto;
}
img {
  border-style: none;
}
.agreement-box {
  margin-top: 37px;
  color: #767676;
}
.agreement-box a {
  color: #007fff;
}
</style>
<style lang="less">
.el-notification.error {
  display: flex;
  align-items: baseline;
  background-color: #e6f3ff;
  width: 20%;
  border-radius: 0;
  padding: 10px;
  p {
    color: #007fff;
  }
}
</style>