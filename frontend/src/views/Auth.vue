<template>
  <div v-if="showAuth">
        <div class="auth-modal-box">
    <auth-form v-if="current=='login'" />
    <reg-form :isPop=true v-else-if="current=='reg'" style="width:320px" />
        </div>
  </div>
</template>

<script>
import AuthForm from "../components/AuthForm.vue";
import RegForm from "../components/RegForm.vue";

export default {
  components: {
    AuthForm,
    RegForm
  },

  data() {
    return {};
  },
  props: {},
  methods: {},
  computed: {
    showAuth() {
      return this.$store.getters.isShow;
    },
    current() {
      return this.$store.getters.getCurrent;
    }
  }
};
</script>

<style lang="less" scoped>
.auth-modal-box {
  display: flex;
  align-items: center;
  justify-content: center;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.3);
  z-index: 500;
}
</style>