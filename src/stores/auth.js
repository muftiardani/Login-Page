import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import * as api from '@/api/auth';

export const useAuthStore = defineStore('auth', () => {
    // State
    const user = ref(localStorage.getItem('username') || null);
    const token = ref(localStorage.getItem('token') || null);
    const router = useRouter();

    // Getters
    const isLoggedIn = computed(() => !!token.value);

    // Actions
    async function handleLogin(credentials) {
        try {
            const data = await api.login(credentials);
            token.value = data.token;
            user.value = credentials.username;
            localStorage.setItem('token', data.token);
            localStorage.setItem('username', credentials.username);
            router.push('/'); // Arahkan ke halaman status setelah login
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
        localStorage.removeItem('username');
        router.push('/auth'); // Arahkan ke halaman login setelah logout
    }

    return { user, token, isLoggedIn, handleLogin, handleRegister, handleLogout };
});