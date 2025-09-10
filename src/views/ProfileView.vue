<script setup>
import { ref } from 'vue';
import { useAuthStore } from '@/stores/auth';
import * as api from '@/api/auth';
import router from '@/router';

const authStore = useAuthStore();
const oldPassword = ref('');
const newPassword = ref('');
const notification = ref({ message: '', type: '', show: false });

async function handleChangePassword() {
  if (!oldPassword.value || !newPassword.value) {
    showNotification('Semua kolom harus diisi', 'error');
    return;
  }

  try {
    const payload = {
      email: authStore.user,
      oldPassword: oldPassword.value,
      newPassword: newPassword.value,
    };
    const result = await api.changePassword(authStore.token, payload);
    showNotification(result.message, 'success');
    oldPassword.value = '';
    newPassword.value = '';
  } catch (error) {
    showNotification(error.message, 'error');
  }
}

function showNotification(message, type) {
  notification.value = { message, type, show: true };
  setTimeout(() => {
    notification.value.show = false;
  }, 5000);
}

function goBack() {
  router.push('/');
}
</script>

<template>
  <div class="auth-container">
    <Transition name="slide-fade" mode="out-in">
        <div v-if="notification.show" :key="notification.message" :class="['notification', notification.type]">
            {{ notification.message }}
        </div>
    </Transition>

    <h2>Ubah Kata Sandi</h2>
    <form @submit.prevent="handleChangePassword" class="auth-form">
      <div class="input-group">
        <label for="old-password">Kata Sandi Lama</label>
        <input type="password" id="old-password" v-model="oldPassword" required autocomplete="current-password" />
      </div>
      <div class="input-group">
        <label for="new-password">Kata Sandi Baru</label>
        <input type="password" id="new-password" v-model="newPassword" required autocomplete="new-password" />
      </div>
      <button type="submit" class="submit-button">Simpan Perubahan</button>
      <button type="button" @click="goBack" class="submit-button" style="margin-top: 1rem; background: #6c757d;">Kembali</button>
    </form>
  </div>
</template>

<style scoped>
.slide-fade-enter-active,
.slide-fade-leave-active {
  transition: all 0.5s ease;
}
.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateY(-20px);
  opacity: 0;
}
</style>