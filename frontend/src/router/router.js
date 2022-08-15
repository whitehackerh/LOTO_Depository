import {createRouter, createWebHistory} from 'vue-router'
import Store from '@/store/index.js'
import Home from '../components/pages/home.vue'
import Login from '../components/pages/login.vue'
import Loto6Results from '../components/pages/loto6results.vue'
import RegisterLoto6Results from '../components/pages/registerloto6results.vue'
import PredictLoto6 from '../components/pages/predictloto6.vue'
import EditLoto6Result from '../components/pages/editLoto6Result.vue'
import Loto6Statistics from '../components/pages/loto6statistics.vue'
import UsersPredictionsLoto6 from '../components/pages/usersPredictionsLoto6.vue'
import Loto7Results from '../components/pages/loto7results.vue'
import RegisterLoto7Results from '../components/pages/registerloto7results.vue'
import PredictLoto7 from '../components/pages/predictloto7.vue'
import Loto7Statistics from '../components/pages/loto7statistics.vue'
import EditLoto7Result from '../components/pages/editLoto7Result.vue'
import UsersPredictionsLoto7 from '../components/pages/usersPredictionsLoto7.vue'
import UserRegistration from '../components/pages/userRegistration.vue'

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
        path: '/predictloto6',
        name: 'predictloto6',
        component: PredictLoto6,
        meta: { requiresAuth: true }
    },
    {
        path: '/loto6statistics',
        name: 'loto6statistics',
        component: Loto6Statistics
    },
    {
        path: '/editLoto6Result/:id',
        name: 'editLoto6Result',
        component: EditLoto6Result,
    },
    {
        path: '/usersPredictionsLoto6/:id',
        name: 'usersPredictionsLoto6',
        component: UsersPredictionsLoto6,
        meta: { requiresAuth: true },
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
    },
    {
        path: '/predictloto7',
        name: 'predictloto7',
        component: PredictLoto7,
        meta: { requiresAuth: true }
    },
    {
        path: '/loto7statistics',
        name: 'loto7statistics',
        component: Loto7Statistics
    },
    {
        path: '/editLoto7Result/:id',
        name: 'editLoto7Result',
        component: EditLoto7Result,
    },
    {
        path: '/usersPredictionsLoto7/:id',
        name: 'usersPredictionsLoto7',
        component: UsersPredictionsLoto7,
        meta: { requiresAuth: true },
    },
    {
        path: '/userRegistration',
        name: 'userRegistration',
        component: UserRegistration,
    },
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

router.beforeEach((to, from, next) => {
    if (to.matched.some(record => record.meta.requiresAuth) && !Store.state.token) {
      next({ path: '/login', query: { redirect: to.fullPath } });
    } else {
      next();
    }
  });

export default router