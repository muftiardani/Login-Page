<script setup>
import { ref } from "vue";

const emit = defineEmits(["submit-login"]);

const email = ref("");
const password = ref("");
const isPasswordVisible = ref(false);

function handleSubmit() {
  emit("submit-login", { email: email.value, password: password.value });
}

function togglePasswordVisibility() {
  isPasswordVisible.value = !isPasswordVisible.value;
}
</script>

<template>
  <form @submit.prevent="handleSubmit" class="auth-form">
    <h2>Selamat Datang</h2>
    <div class="input-group">
      <label for="login-email">Alamat Email</label>
      <input
        type="email"
        id="login-email"
        v-model="email"
        required
        autocomplete="email"
      />
    </div>
    <div class="input-group">
      <label for="login-password">Kata Sandi</label>
      <div class="password-wrapper">
        <input
          :type="isPasswordVisible ? 'text' : 'password'"
          id="login-password"
          v-model="password"
          required
          autocomplete="current-password"
        />
        <span @click="togglePasswordVisibility" class="password-toggle-icon">
          {{ isPasswordVisible ? '🙈' : '👁️' }}
        </span>
      </div>
    </div>
    <button type="submit" class="submit-button">Masuk</button>
    <div class="toggle-view">
      <p>
        Belum punya akun?
        <router-link to="/register">Buat akun baru</router-link>
      </p>
    </div>
  </form>
</template>