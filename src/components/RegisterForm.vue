<script setup>
import { ref } from "vue";

const emit = defineEmits(["submit-register", "switch-to-login"]);

const username = ref("");
const password = ref("");

function handleSubmit() {
  emit("submit-register", {
    username: username.value,
    password: password.value,
  });
}
</script>

<template>
  <form @submit.prevent="handleSubmit" class="auth-form">
    <h2>Registrasi</h2>
    <div class="input-group">
      <label for="register-username">Username</label>
      <input
        type="text"
        id="register-username"
        v-model="username"
        required
        autocomplete="username"
      />
    </div>
    <div class="input-group">
      <label for="register-password">Password</label>
      <input
        type="password"
        id="register-password"
        v-model="password"
        required
        autocomplete="new-password"
      />
    </div>
    <button type="submit" class="submit-button">Registrasi</button>
    <div class="toggle-view">
      <p>
        Sudah punya akun?
        <a @click.prevent="emit('switch-to-login')" href="#">Login</a>
      </p>
    </div>
  </form>
</template>

<style scoped>
/* Scoped styles for RegisterForm */
</style>
