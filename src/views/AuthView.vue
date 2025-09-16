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

const currentPath = ref(route.path);

// Mengawasi perubahan path untuk menampilkan komponen yang benar
watch(
  () => route.path,
  (newPath) => {
    currentPath.value = newPath;
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
      router.push({ name: "Login" });
    }, 2000);
  }
}
</script>

<template>
  <div class="auth-container">
    <div class="content-container">
      <Transition name="fade" mode="out-in">
        <LoginForm
          v-if="currentPath === '/auth/login'"
          @submit-login="handleLogin"
          key="login"
        />
        <RegisterForm
          v-else-if="currentPath === '/auth/register'"
          @submit-register="handleRegister"
          key="register"
        />
      </Transition>
    </div>
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
