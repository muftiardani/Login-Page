<script setup>
import { ref } from "vue";

const emit = defineEmits(["submit-login"]);

const username = ref("");
const password = ref("");

function handleSubmit() {
  emit("submit-login", { username: username.value, password: password.value });
}
</script>

<template>
  <form @submit.prevent="handleSubmit" class="auth-form">
    <h2>Login</h2>
    <div class="input-group">
      <label for="login-username">Username</label>
      <input
        type="text"
        id="login-username"
        v-model="username"
        required
        autocomplete="username"
      />
    </div>
    <div class="input-group">
      <label for="login-password">Password</label>
      <input
        type="password"
        id="login-password"
        v-model="password"
        required
        autocomplete="current-password"
      />
    </div>
    <button type="submit" class="submit-button">Login</button>
    <div class="toggle-view">
      <p>
        Belum punya akun?
        <router-link to="/register">Daftar</router-link>
      </p>
    </div>
  </form>
</template>