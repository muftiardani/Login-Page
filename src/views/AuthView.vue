<template>
  <div class="card-container">
    <div v-if="message" :class="['message', messageClass]">
      {{ message }}
    </div>

    <Transition name="fade-slide" mode="out-in">
      <LoginForm v-if="props.initialView === 'login'" @submit-login="onLogin" />
      <RegisterForm v-else @submit-register="onRegister" />
    </Transition>
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
    message.value = `Ups! ${result.message}`;
    messageClass.value = 'error';
  }
}

async function onRegister(credentials) {
  const result = await authStore.handleRegister(credentials);
  if (result.success) {
    message.value = `Hore! ${result.message} Silakan masuk untuk melanjutkan.`;
    messageClass.value = 'success';
  } else {
    message.value = `Maaf, ${result.message}`;
    messageClass.value = 'error';
  }
}
</script>