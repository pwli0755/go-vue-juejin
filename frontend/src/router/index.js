import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
    path: '/email/confirm',
    name: 'activate',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import('../views/Loading.vue')
  },
  {
    path: '/editor/drafts/new',
    name: 'newPost',
    component:() => import('../views/Editor.vue')
  },
  {
    path: '/',
    name: 'home',
    component: Home
  },
]

const router = new VueRouter({
  routes
})

export default router
