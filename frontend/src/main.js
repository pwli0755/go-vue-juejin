import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugins/element'
import store from './store'//引入store

import './mock/'
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
Vue.use(mavonEditor)

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')

