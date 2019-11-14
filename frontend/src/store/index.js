import Vue from 'vue';
import Vuex from 'vuex';
Vue.use(Vuex);
const state = {   //要设置的全局访问的state对象
    showAuth: false,
    current: '',
    loginState: false,
    user: {},
};
const getters = {   //实时监听state值的变化(最新状态)
    isShow(state) {  //承载变化的showAuth的值
        return state.showAuth;
    },
    getCurrent() {  //登录or注册
        return state.current;
    },
    getLoginState() {
        return state.loginState;
    }
};
const mutations = {
    show(state) {   //自定义改变state初始值的方法，这里面的参数除了state之外还可以再传额外的参数(变量或对象);
        state.showAuth = true;
    },
    hide(state) {  //同上
        state.showAuth = false;
    },
    changeCurrent(state, nv) { // 改变current
        state.current = nv;
    },
    setLogin(state) {
        state.loginState = true;
    },

};
const actions = {
    hideAuth(context) {  //自定义触发mutations里函数的方法，context与store 实例具有相同方法和属性
        context.commit('hide');
    },
    showAuth(context) {  //同上注释
        context.commit('show');
    },
    switchAuthType(context, nv) {
        context.commit('changeCurrent', nv)
    },
    setLoginState(context) {
        context.commit('setLogin')
    }
};
const store = new Vuex.Store({
    state,
    getters,
    mutations,
    actions
});
export default store;