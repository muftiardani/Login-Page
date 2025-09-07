<template>
  <div class="card-container">
    <div v-if="message" :class="['message', messageClass]">
      {{ message }}
    </div>

    <LoginForm v-if="props.initialView === 'login'" @submit-login="onLogin" />
    <RegisterForm v-else @submit-register="onRegister" />
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useAuthStore } from '@/stores/auth';
import LoginForm from '@/components/LoginForm.vue';
import RegisterForm from '@/components/RegisterForm.vue';

const props = defineProps({
  initialView: {
    type: String,
    required: true,
  }
});

const authStore = useAuthStore();
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
    message.value = `✅ ${result.message} Silakan login.`;
    messageClass.value = 'success';
  } else {
    message.value = `❌ ${result.message}`;
    messageClass.value = 'error';
  }
}
</script>