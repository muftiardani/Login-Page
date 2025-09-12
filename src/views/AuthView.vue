<script setup>
import { ref, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useAuthStore } from "@/stores/auth.js";
import LoginForm from "@/components/LoginForm.vue";
import RegisterForm from "@/components/RegisterForm.vue";
import { useToast } from "vue-toastification";

const route = useRoute();
const router = useRouter();
const authStore = useAuthStore();
const toast = useToast();

const currentView = ref(route.path);

watch(
  () => route.path,
  (newPath) => {
    currentView.value = newPath;
  }
);

async function handleLogin(credentials) {
  await authStore.handleLogin(credentials);
}

async function handleRegister(credentials) {
  if (credentials.password !== credentials.confirmPassword) {
    toast.error("Konfirmasi kata sandi tidak cocok.");
    return;
  }
  const result = await authStore.handleRegister(credentials);
  if (result.success) {
    setTimeout(() => {
      router.push("/login");
    }, 2000);
  }
}
</script>

<template>
  <div class="content-container">
    <Transition name="fade" mode="out-in">
      <div :key="currentView">
        <LoginForm
          v-if="currentView === '/login'"
          @submit-login="handleLogin"
        />
        <RegisterForm
          v-else-if="currentView === '/register'"
          @submit-register="handleRegister"
        />
      </div>
    </Transition>
  </div>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>