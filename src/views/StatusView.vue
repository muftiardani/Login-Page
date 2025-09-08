<template>
  <div class="card-container status-view">
    <h2>Halo, {{ authStore.user }}!</h2>
    <p class="welcome-message">Selamat datang kembali! Anda berhasil masuk.</p>
    <button @click="logout" class="logout-button">Keluar</button>
  </div>
</template>

<script setup>
import { useAuthStore } from '@/stores/auth';

const authStore = useAuthStore();

function logout() {
  authStore.handleLogout();
}
</script>