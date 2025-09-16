<script setup>
import { ref } from "vue";
import { useAuthStore } from "@/stores/auth.js";
import * as api from "@/api/auth.js";
import { useToast } from "vue-toastification";
import { usePasswordValidator } from "@/composables/usePasswordValidator.js";

const authStore = useAuthStore();
const toast = useToast();
const oldPassword = ref("");
const confirmNewPassword = ref("");

const {
  password: newPassword,
  passwordStrength,
  validatePassword,
} = usePasswordValidator();

async function handleChangePassword() {
  if (!oldPassword.value || !newPassword.value || !confirmNewPassword.value) {
    toast.error("Semua kolom harus diisi");
    return;
  }
  if (newPassword.value !== confirmNewPassword.value) {
    toast.error("Konfirmasi kata sandi baru tidak cocok.");
    return;
  }
  if (oldPassword.value === newPassword.value) {
    toast.error("Kata sandi baru tidak boleh sama dengan kata sandi lama.");
    return;
  }
  const passwordError = validatePassword(newPassword.value);
  if (passwordError) {
    toast.error(passwordError);
    return;
  }
  try {
    const payload = {
      email: authStore.user,
      oldPassword: oldPassword.value,
      newPassword: newPassword.value,
    };
    const result = await api.changePassword(payload);
    toast.success(result.message);
    oldPassword.value = "";
    newPassword.value = "";
    confirmNewPassword.value = "";
  } catch (error) {
    toast.error(error.message);
  }
}
</script>

<template>
  <div class="card profile-card">
    <form @submit.prevent="handleChangePassword" class="auth-form">
      <div class="input-group">
        <label for="old-password">Kata Sandi Lama</label>
        <input
          type="password"
          id="old-password"
          v-model="oldPassword"
          required
          autocomplete="current-password"
        />
      </div>
      <div class="input-group">
        <label for="new-password">Kata Sandi Baru</label>
        <input
          type="password"
          id="new-password"
          v-model="newPassword"
          required
          autocomplete="new-password"
        />
        <div v-if="newPassword.length > 0" class="password-strength-container">
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
      </div>
      <div class="input-group">
        <label for="confirm-new-password">Konfirmasi Kata Sandi Baru</label>
        <input
          type="password"
          id="confirm-new-password"
          v-model="confirmNewPassword"
          required
          autocomplete="new-password"
        />
      </div>
      <button type="submit" class="button button-primary">
        Simpan Perubahan
      </button>
    </form>
  </div>
</template>

<style scoped>
.card {
  background-color: var(--card-background);
  border-radius: var(--border-radius);
  padding: 2rem;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.05);
  border: 1px solid var(--border-color);
}
.profile-card {
  max-width: 700px;
}
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
</style>
