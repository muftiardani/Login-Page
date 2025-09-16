<script setup>
import { ref, computed } from "vue";
import { useAuthStore } from "@/stores/auth.js";
import * as api from "@/api/auth.js";
import { useToast } from "vue-toastification";

const authStore = useAuthStore();
const toast = useToast();
const oldPassword = ref("");
const newPassword = ref("");
const confirmNewPassword = ref("");

const passwordStrength = computed(() => {
  const p = newPassword.value;
  if (p.length === 0) return { score: 0, text: "", color: "#ccc", width: "0%" };
  let score = 0;
  let checks = 0;
  if (p.length >= 8) {
    score += 20;
    checks++;
  }
  if (/[a-z]/.test(p)) {
    score += 20;
    checks++;
  }
  if (/[A-Z]/.test(p)) {
    score += 20;
    checks++;
  }
  if (/[0-9]/.test(p)) {
    score += 20;
    checks++;
  }
  if (/[^a-zA-Z0-9]/.test(p)) {
    score += 20;
    checks++;
  }
  score = Math.min(score, 100);
  let text = "Sangat Lemah";
  let color = "#dc3545";
  if (checks === 5) {
    text = "Sangat Kuat";
    color = "#28a745";
  } else if (checks === 4) {
    text = "Kuat";
    color = "#8fce00";
  } else if (checks === 3) {
    text = "Kurang Kuat";
    color = "#ffc107";
  } else if (checks === 2) {
    text = "Lemah";
    color = "#fd7e14";
  }
  return { score, text, color, width: `${score}%` };
});

function validatePassword(password) {
  if (password.length < 8)
    return "Kata sandi baru harus memiliki minimal 8 karakter.";
  if (!/[a-z]/.test(password))
    return "Kata sandi baru harus mengandung setidaknya satu huruf kecil.";
  if (!/[A-Z]/.test(password))
    return "Kata sandi baru harus mengandung setidaknya satu huruf besar.";
  if (!/[0-9]/.test(password))
    return "Kata sandi baru harus mengandung setidaknya satu angka.";
  if (!/[^a-zA-Z0-9]/.test(password))
    return "Kata sandi baru harus mengandung setidaknya satu karakter spesial.";
  return null;
}

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
