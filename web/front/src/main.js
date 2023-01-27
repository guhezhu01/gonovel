import { createApp } from 'vue'
import App from './App.vue'
import './registerServiceWorker'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from "./router"
import store from "./store"
import axios from "axios"
import './assets/css/style.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import "./assets/font_3864366_7ysxdfbnyx6/iconfont.css"


axios.defaults.baseURL = "http://localhost:3000/api/v1"

const app = createApp(App)

axios.interceptors.request.use(config => {
  config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
  return config
})
app.config.globalProperties.$http = axios
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}
app.use(router).use(ElementPlus, {
  locale: zhCn,
}).use(store).mount('#app')

