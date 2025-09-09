<script setup>
import { useAuthStore } from '@/stores/auth';

const authStore = useAuthStore();

function logout() {
  authStore.handleLogout();
}
</script>

<template>
  <div class="status-view">
    <h1>Selamat Datang Kembali!</h1>
    <p>Halo, {{ authStore.user }}! Anda telah berhasil login.</p>
    <button @click="logout" class="logout-button">Keluar</button>
  </div>
</template>