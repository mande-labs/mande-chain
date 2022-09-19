import { createRouter, createWebHistory } from 'vue-router'

import Vote from '../views/Vote.vue'
import Portfolio from '../views/Portfolio.vue'

const routerHistory = createWebHistory()
const routes = [
  { path: '/', component: Portfolio },
  { path: '/portfolio', component: Portfolio },
  { path: '/vote', component: Vote }
]

const router = createRouter({
  history: routerHistory,
  routes
})

export default router
