import { createRouter, createWebHistory } from 'vue-router';
import AuthView from '@/views/AuthView.vue';
import StatusView from '@/views/StatusView.vue';

const routes = [
    {
        path: '/login',
        name: 'Login',
        component: AuthView,
        props: { initialView: 'login' },
        meta: { requiresGuest: true }
    },
    {
        path: '/register',
        name: 'Register',
        component: AuthView,
        props: { initialView: 'register' },
        meta: { requiresGuest: true }
    },
    {
        path: '/',
        name: 'Status',
        component: StatusView,
        meta: { requiresAuth: true }
    },
    {
        path: '/:pathMatch(.*)*',
        redirect: '/'
    }
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

router.beforeEach((to, from, next) => {
    const isLoggedIn = !!localStorage.getItem('token');

    if (to.meta.requiresAuth && !isLoggedIn) {
        next('/login');
    } else if (to.meta.requiresGuest && isLoggedIn) {
        next('/');
    } else {
        next();
    }
});

export default router;