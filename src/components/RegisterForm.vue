<script setup>
import { ref } from "vue";
import { useAuthStore } from "@/stores/auth.js";
import { usePasswordValidator } from "@/composables/usePasswordValidator.js";

const emit = defineEmits(["submit-register"]);
const authStore = useAuthStore();

const email = ref("");
const confirmPassword = ref("");
const isPasswordVisible = ref(false);

const { 
  password, 
  passwordStrength, 
  validatePassword 
} = usePasswordValidator();

const passwordValidationError = ref("");

function handleSubmit() {
  passwordValidationError.value = validatePassword(password.value);
  if (passwordValidationError.value) {
    return;
  }
  emit("submit-register", {
    email: email.value,
    password: password.value,
    confirmPassword: confirmPassword.value,
  });
}

function togglePasswordVisibility() {
  isPasswordVisible.value = !isPasswordVisible.value;
}
</script>

<template>
  <form @submit.prevent="handleSubmit" class="auth-form">
    <h2>Buat Akun Baru</h2>

    <div class="input-group">
      <label for="register-email">Email</label>
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
      <div class="password-wrapper">
        <input
          :type="isPasswordVisible ? 'text' : 'password'"
          id="register-password"
          v-model="password"
          required
          autocomplete="new-password"
        />
        <span @click="togglePasswordVisibility" class="password-toggle-icon">
          {{ isPasswordVisible ? "" : "" }}
        </span>
      </div>

      <div v-if="password.length > 0" class="password-strength-container">
        <div class="strength-bar-background">
          <div
            class="strength-bar-fill"
            :style="{
              width: passwordStrength.width,
              backgroundColor: passwordStrength.color,
            }"
          ></div>
        </div>
        <div
          class="password-strength-text"
          :style="{ color: passwordStrength.color }"
        >
          Kata Sandi {{ passwordStrength.text }}
        </div>
      </div>

      <div v-if="passwordValidationError" class="validation-error">
        {{ passwordValidationError }}
      </div>
    </div>

    <div class="input-group">
      <label for="register-confirm-password">Konfirmasi Kata Sandi</label>
      <input
        type="password"
        id="register-confirm-password"
        v-model="confirmPassword"
        required
        autocomplete="new-password"
      />
    </div>

    <button
      type="submit"
      class="button button-primary"
      :disabled="authStore.isLoading"
    >
      {{ authStore.isLoading ? "Memproses..." : "Daftar Sekarang" }}
    </button>

    <div class="toggle-view">
      <p>
        Sudah punya akun?
        <router-link to="/auth/login">Masuk di sini</router-link>
      </p>
    </div>
  </form>
</template>

<style scoped>
.password-strength-container {
  margin-top: 10px;
  text-align: left;
}
.strength-bar-background {
  width: 100%;
  height: 8px;
  background-color: #e0e0e0;
  border-radius: 4px;
  overflow: hidden;
  margin-bottom: 5px;
}
.strength-bar-fill {
  height: 100%;
  transition: width 0.3s ease-in-out, background-color 0.3s ease-in-out;
  border-radius: 4px;
}
.password-strength-text {
  font-size: 0.85rem;
  font-weight: 600;
}
.validation-error {
  color: var(--error-color);
  font-size: 0.85rem;
  margin-top: 8px;
  text-align: left;
}
</style>