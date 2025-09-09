<script setup>
import { ref, watch } from "vue";
import { useRoute } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import LoginForm from "@/components/LoginForm.vue";
import RegisterForm from "@/components/RegisterForm.vue";

const route = useRoute();
const authStore = useAuthStore();

const currentView = ref(route.path);
const notification = ref({
  message: "",
  type: "",
  show: false,
});

watch(
  () => route.path,
  (newPath) => {
    currentView.value = newPath;
    notification.value.show = false;
  }
);

async function handleLogin(credentials) {
  const result = await authStore.handleLogin(credentials);
  if (!result.success) {
    showNotification(result.message, "error");
  }
}

async function handleRegister(credentials) {
  const result = await authStore.handleRegister(credentials);
  showNotification(result.message, result.success ? "success" : "error");
  if (result.success) {
  }
}

function showNotification(message, type) {
  notification.value.message = message;
  notification.value.type = type;
  notification.value.show = true;
  setTimeout(() => {
    notification.value.show = false;
  }, 5000);
}
</script>

<template>
  <div class="auth-container">
    <Transition name="slide-fade" mode="out-in">
      <div v-if="notification.show" :key="notification.message" :class="['notification', notification.type]">
        <span class="notification-icon">
          <template v-if="notification.type === 'success'">✅</template>
          <template v-else>❌</template>
        </span>
        {{ notification.message }}
      </div>
    </Transition>

    <Transition name="fade" mode="out-in">
      <LoginForm v-if="currentView === '/login'" @submit-login="handleLogin" key="login" />
      <RegisterForm v-else-if="currentView === '/register'" @submit-register="handleRegister" key="register" />
    </Transition>
  </div>
</template>

<style scoped>
/* Transisi untuk notifikasi */
.slide-fade-enter-active,
.slide-fade-leave-active {
  transition: all 0.5s ease;
}
.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateY(-20px);
  opacity: 0;
}

/* Transisi untuk formulir */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>