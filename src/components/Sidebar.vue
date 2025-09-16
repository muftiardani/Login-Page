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
.sidebar {
  width: 260px;
  background: var(--primary-color);
  color: white;
  display: flex;
  flex-direction: column;
  padding: 1.5rem;
  flex-shrink: 0;
}
.logo {
  font-size: 1.5rem;
  font-weight: 700;
  letter-spacing: 2px;
  margin-bottom: 2rem;
  text-align: center;
}
.navigation {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}
.nav-item {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 0.9rem 1rem;
  border-radius: 8px;
  color: #a7c0cd;
  text-decoration: none;
  transition: all 0.2s ease;
  width: 100%;
  background: none;
  border: none;
  cursor: pointer;
  font-family: var(--font-family);
  font-size: 1rem;
}
.nav-item:hover,
.router-link-exact-active {
  background-color: var(--secondary-color);
  color: white;
}
.logout-section {
  margin-top: auto;
}
.logout-button:hover {
  background-color: #c82333;
  color: white;
}
</style>
