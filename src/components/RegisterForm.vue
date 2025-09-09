<script setup>
import { ref } from "vue";

const emit = defineEmits(["submit-register"]);

const email = ref("");
const password = ref("");

function handleSubmit() {
  emit("submit-register", {
    email: email.value,
    password: password.value,
  });
}
</script>

<template>
  <form @submit.prevent="handleSubmit" class="auth-form">
    <h2>Buat Akun Baru</h2>
    <div class="input-group">
      <label for="register-email">Alamat Email</label>
      <input
        type="email"
        id="register-email"
        v-model="email"
        required
        autocomplete="email"
      />
    </div>
    <div class="input-group">
      <label for="register-password">Kata Sandi</label>
      <input
        type="password"
        id="register-password"
        v-model="password"
        required
        autocomplete="new-password"
      />
    </div>
    <button type="submit" class="submit-button">Daftar Sekarang</button>
    <div class="toggle-view">
      <p>
        Sudah punya akun?
        <router-link to="/login">Masuk di sini</router-link>
      </p>
    </div>
  </form>
</template>