// import Vue.js module
import Vue from 'vue'
// import Vue-resource module
import VueResource from 'vue-resource'
// import components
import Page from './components/page/page.vue' // page template
// activating Vue-resource
Vue.use(VueResource)

new Vue({
  el: "#base",
  render: h => h(Page)
})
