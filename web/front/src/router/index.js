import { createRouter, createWebHistory } from 'vue-router'
import Home from "../components/home/Home"

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home,
    //设置标题
    meta: { title: 'guhezhu小说网' }
  },
  {
    path: '/type',
    name: 'type',
    meta: { title: '小说分类' },
    component: () => import('../components/novel_type/NovelMain'),
  },
  {
    path: '/admin',
    name: 'admin',
    meta: { title: '个人中心' },
    component: () => import('../components/admin_information/AdminView'),
  },
  {
    path: '/novel_information',
    name: 'novel_information',
    meta: { title: '小说详情' },
    component: () => import('../components/novel_information/NovelInformation'),
  },
  {
    path: '/login',
    name: 'login',
    component: () => import("../components/login/Login"),
  },

 
]

const router = createRouter({
  history: createWebHistory(),
  routes
})


router.beforeEach(async (to, from, next) => {
  
  //设置标题
  if (to.meta.title) {
    document.title = to.meta.title
  }

  next()
})

export default router
