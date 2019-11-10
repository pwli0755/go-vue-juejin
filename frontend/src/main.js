import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugins/element'
import store from './store'//引入store

import './mock/'


Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')

