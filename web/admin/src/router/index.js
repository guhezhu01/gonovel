import { createRouter, createWebHistory } from 'vue-router'
import Login from "../views/Login"

const routes = [
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/',
    name: 'admin',
    component: () => import('../views/Admin'),
    children: [{
      path: "/user",
      name: 'user',
      component: () => import('../components/user/User')
    },
    {
      path: "/article",
      name: 'article',
      component: () => import('../components/article/Article')
    },
    {
      path: "/comment",
      name: 'comment',
      component: () => import('../components/comment/Comment')
    },
    {
      path: "/self",
      name: 'self',
      component: () => import('../components/user/SelfEdit')
    },
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  const userToken = window.sessionStorage.getItem('token')
  if (to.path === '/login') return next()
  if (!userToken) {
    next('/login')
  } else {
    next()
  }
})

export default router
