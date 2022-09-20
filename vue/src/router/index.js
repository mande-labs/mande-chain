import { createRouter, createWebHistory } from 'vue-router'

import Portfolio from '../views/Portfolio.vue'

const routerHistory = createWebHistory()
const routes = [
  { path: '/', component: Portfolio },
  { path: '/portfolio', component: Portfolio },
]

const router = createRouter({
  history: routerHistory,
  routes
})

export default router
