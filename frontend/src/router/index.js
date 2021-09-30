import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('../views/Home.vue')
  },
  {
    path: '/detail',
    name: 'Detail',
    component: () => import('../views/Detail.vue')
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('../views/About.vue')
  },
  {
    path: '/mypokemon',
    name: 'MyPokemon',
    component: () => import('../views/MyPokemon.vue')
  },
  {
    path: '/mypokemondetail',
    name: 'MyPokemonDetail',
    component: () => import('../views/MyPokemonDetail.vue')
  },

  {
    path: '/renamePokemon',
    name: 'RenamePokemon',
    component: () => import('../views/RenamePokemon.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
