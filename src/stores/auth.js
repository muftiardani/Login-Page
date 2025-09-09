import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import router from '@/router';
import * as api from '@/api/auth';

export const useAuthStore = defineStore('auth', () => {
    const user = ref(localStorage.getItem('email') || null);
    const token = ref(localStorage.getItem('token') || null);

    const isLoggedIn = computed(() => !!token.value);
    const isAuthenticated = computed(() => !!token.value);

    async function handleLogin(credentials) {
        try {
            const data = await api.login(credentials);
            token.value = data.token;
            user.value = credentials.email;
            localStorage.setItem('token', data.token);
            localStorage.setItem('email', credentials.email);
            await router.push('/');
            return { success: true, message: data.message };
        } catch (error) {
            return { success: false, message: error.message };
        }
    }

    async function handleRegister(credentials) {
        try {
            const data = await api.register(credentials);
            return { success: true, message: data.message };
        } catch (error) {
            return { success: false, message: error.message };
        }
    }

    function handleLogout() {
        token.value = null;
        user.value = null;
        localStorage.removeItem('token');
        localStorage.removeItem('email');
        router.push('/login');
    }

    return { user, token, isLoggedIn, isAuthenticated, handleLogin, handleRegister, handleLogout };
});