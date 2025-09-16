<script setup>
import { useAuthStore } from "@/stores/auth";
import { useRouter } from "vue-router";

const authStore = useAuthStore();
const router = useRouter();

function logout() {
  authStore.handleLogout();
}

function goToProfile() {
  router.push("/profile");
}

function goToPayments() {
  router.push("/payments");
}
</script>

<template>
  <div class="content-container status-view-content">
    <h1>Selamat Datang Kembali!</h1>
    <p>Halo, {{ authStore.user }}! Anda telah berhasil login.</p>
    <button @click="goToPayments" class="button button-primary">
      Lihat Tabel Pembayaran
    </button>
    <button @click="goToProfile" class="button button-secondary">
      Ubah Password
    </button>
    <button @click="logout" class="button button-danger">Keluar</button>
  </div>
</template>
