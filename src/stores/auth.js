import { defineStore } from "pinia";
import { ref } from "vue";
import * as api from "@/api/auth.js";
import { useToast } from "vue-toastification";

export const useAuthStore = defineStore("auth", () => {
  const user = ref(localStorage.getItem("email") || null);
  const isAuthenticated = ref(!!user.value);
  const isLoading = ref(false);
  const toast = useToast();

  async function handleLogin(credentials) {
    isLoading.value = true;
    try {
      const data = await api.login(credentials);
      user.value = credentials.email;
      isAuthenticated.value = true;
      localStorage.setItem("email", credentials.email);
      const router = (await import("@/router")).default;
      await router.push({ name: "Dashboard" });
      toast.success(data.message);
      return { success: true };
    } catch (error) {
      isAuthenticated.value = false;
      let userMessage = error.message || "Terjadi kesalahan pada server.";
      toast.error(userMessage);
      return { success: false, message: userMessage };
    } finally {
      isLoading.value = false;
    }
  }

  async function handleRegister(credentials) {
    isLoading.value = true;
    try {
      const data = await api.register(credentials);
      toast.success(data.message);
      return { success: true, message: data.message };
    } catch (error) {
      const userMessage = error.message || "Gagal melakukan pendaftaran.";
      toast.error(userMessage);
      return { success: false, message: userMessage };
    } finally {
      isLoading.value = false;
    }
  }

  async function handleLogout() {
    await api.logout();
    user.value = null;
    isAuthenticated.value = false;
    localStorage.removeItem("email");

    window.location.href = "/auth/login";
  }

  return {
    user,
    isAuthenticated,
    isLoading,
    handleLogin,
    handleRegister,
    handleLogout,
  };
});
