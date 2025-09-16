<script setup>
import Sidebar from "@/components/Sidebar.vue";
import PageHeader from "@/components/PageHeader.vue";
import { useRoute } from "vue-router";
import { useAuthStore } from "@/stores/auth.js";

const route = useRoute();
const authStore = useAuthStore();

// Membuat subtitle dinamis khusus untuk dashboard
const getSubtitle = () => {
  if (route.name === "Dashboard") {
    return `Selamat datang kembali, ${authStore.user}!`;
  }
  return route.meta.subtitle;
};
</script>

<template>
  <div class="main-layout">
    <Sidebar />
    <main class="content">
      <PageHeader :title="route.meta.title" :subtitle="getSubtitle()" />
      <router-view v-slot="{ Component }">
        <Transition name="fade" mode="out-in">
          <component :is="Component" />
        </Transition>
      </router-view>
    </main>
  </div>
</template>

<style scoped>
.main-layout {
  display: flex;
  min-height: 100vh;
  background-color: #f8f9fa;
}
.content {
  flex-grow: 1;
  padding: 2.5rem;
  overflow-y: auto;
}
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
