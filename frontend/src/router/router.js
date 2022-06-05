import {createRouter, createWebHistory} from 'vue-router'
import Home from '../components/pages/home.vue'
import Login from '../components/pages/login.vue'
import Loto6Results from '../components/pages/loto6results.vue'
import RegisterLoto6Results from '../components/pages/registerloto6results.vue'

const routes = [
    {
        path: '/',
        name: 'home',
        component: Home,
    },
    {
        path: '/login',
        name: 'login',
        component: Login
    },
    {
        path: '/loto6results',
        name: 'loto6results',
        component: Loto6Results
    },
    {
        path: '/registerloto6results',
        name: 'registerloto6results',
        component: RegisterLoto6Results
    }
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router