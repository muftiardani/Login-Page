<script setup>
import { useAuthStore } from "@/stores/auth.js";
import HomeIcon from "./icons/HomeIcon.vue";
import PaymentsIcon from "./icons/PaymentsIcon.vue";
import ProfileIcon from "./icons/ProfileIcon.vue";
import LogoutIcon from "./icons/LogoutIcon.vue";

const authStore = useAuthStore();
</script>

<template>
  <aside class="sidebar">
    <div class="logo">AplikasiKu</div>
    <nav class="navigation">
      <router-link to="/" class="nav-item">
        <HomeIcon />
        <span>Dashboard</span>
      </router-link>
      <router-link to="/payments" class="nav-item">
        <PaymentsIcon />
        <span>Pembayaran</span>
      </router-link>
      <router-link to="/profile" class="nav-item">
        <ProfileIcon />
        <span>Profil</span>
      </router-link>
    </nav>
    <div class="logout-section">
      <button @click="authStore.handleLogout()" class="nav-item logout-button">
        <LogoutIcon />
        <span>Keluar</span>
      </button>
    </div>
  </aside>
</template>

<style scoped>
/* Style untuk sidebar */
</style>
