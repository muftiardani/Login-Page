<script setup>
import { useAuthStore } from '@/stores/auth';
import { useRouter } from 'vue-router';

const authStore = useAuthStore();
const router = useRouter();

function logout() {
  authStore.handleLogout();
}

function goToProfile() {
  router.push('/profile');
}
</script>

<template>
  <div class="status-view">
    <h1>Selamat Datang Kembali!</h1>
    <p>Halo, {{ authStore.user }}! Anda telah berhasil login.</p>
    <button @click="goToProfile" class="submit-button" style="background: var(--accent-color); margin-bottom: 1rem;">Ubah Password</button>
    <button @click="logout" class="logout-button">Keluar</button>
  </div>
</template>