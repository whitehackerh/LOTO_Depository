import {createRouter, createWebHistory} from 'vue-router'
import Home from '../components/pages/home.vue'
import Login from '../components/pages/login.vue'
import Loto6Results from '../components/pages/loto6results.vue'
import RegisterLoto6Results from '../components/pages/registerloto6results.vue'
import ExpectLoto6 from '../components/pages/expectloto6.vue'
import Loto7Results from '../components/pages/loto7results.vue'
import RegisterLoto7Results from '../components/pages/registerloto7results.vue'

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
    },
    {
        path: '/expectloto6',
        name: 'expectloto6',
        component: ExpectLoto6
    },
    {
        path: '/loto7results',
        name: 'loto7results',
        component: Loto7Results
    },
    {
        path: '/registerloto7results',
        name: 'registerloto7results',
        component: RegisterLoto7Results
    }
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router