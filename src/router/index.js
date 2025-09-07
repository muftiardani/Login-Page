import { createRouter, createWebHistory } from 'vue-router';
import AuthView from '@/views/AuthView.vue';
import StatusView from '@/views/StatusView.vue';

const routes = [
    {
        path: '/auth',
        name: 'Auth',
        component: AuthView,
        meta: { requiresGuest: true } // Hanya bisa diakses jika belum login
    },
    {
        path: '/',
        name: 'Status',
        component: StatusView,
        meta: { requiresAuth: true } // Hanya bisa diakses jika sudah login
    },
    // Tambahkan fallback route untuk mengarahkan pengguna
    {
        path: '/:pathMatch(.*)*',
        redirect: '/'
    }
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

// Navigation Guard
router.beforeEach((to, from, next) => {
    const isLoggedIn = !!localStorage.getItem('token');

    if (to.meta.requiresAuth && !isLoggedIn) {
        // Jika rute butuh login tapi pengguna belum login, arahkan ke /auth
        next('/auth');
    } else if (to.meta.requiresGuest && isLoggedIn) {
        // Jika rute untuk tamu tapi pengguna sudah login, arahkan ke /
        next('/');
    } else {
        // Lanjutkan navigasi
        next();
    }
});

export default router;