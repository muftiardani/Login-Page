<template>
  <div class="card-container">
    <div v-if="message" :class="['message', messageClass]">
      {{ message }}
    </div>

    <Transition name="fade-slide" mode="out-in">
      <LoginForm
        v-if="currentView === 'login'"
        @submit-login="onLogin"
        @switch-to-register="switchToRegister"
      />
      <RegisterForm
        v-else
        @submit-register="onRegister"
        @switch-to-login="switchToLogin"
      />
    </Transition>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useAuthStore } from '@/stores/auth';
import LoginForm from '@/components/LoginForm.vue';
import RegisterForm from '@/components/RegisterForm.vue';

const authStore = useAuthStore();
const currentView = ref('login');
const message = ref('');
const messageClass = ref('');

async function onLogin(credentials) {
  const result = await authStore.handleLogin(credentials);
  if (!result.success) {
    message.value = `❌ ${result.message}`;
    messageClass.value = 'error';
  }
}

async function onRegister(credentials) {
  const result = await authStore.handleRegister(credentials);
  if (result.success) {
    message.value = `✅ ${result.message}`;
    messageClass.value = 'success';
    // Otomatis pindah ke login setelah registrasi berhasil
    switchToLogin(); 
  } else {
    message.value = `❌ ${result.message}`;
    messageClass.value = 'error';
  }
}

function switchToLogin() {
  currentView.value = 'login';
  message.value = '';
}

function switchToRegister() {
  currentView.value = 'register';
  message.value = '';
}
</script>