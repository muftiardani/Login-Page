import { createRouter, createWebHistory } from "vue-router";
import { useAuthStore } from "@/stores/auth.js";

import MainLayout from "@/layouts/MainLayout.vue";
import AuthView from "@/views/AuthView.vue";

const routes = [
  {
    path: "/auth/login",
    name: "Login",
    component: AuthView,
    meta: { requiresGuest: true },
  },
  {
    path: "/auth/register",
    name: "Register",
    component: AuthView,
    meta: { requiresGuest: true },
  },
  {
    path: "/",
    component: MainLayout,
    meta: { requiresAuth: true },
    children: [
      {
        path: "",
        name: "Dashboard",
        component: () => import("@/views/DashboardView.vue"),
        meta: {
          title: "Dashboard",
          subtitle: "Selamat datang kembali!",
        },
      },
      {
        path: "payments",
        name: "Payments",
        component: () => import("@/views/PaymentView.vue"),
        meta: {
          title: "Tabel Pembayaran",
          subtitle: "Berikut adalah daftar semua transaksi pembayaran.",
        },
      },
      {
        path: "profile",
        name: "Profile",
        component: () => import("@/views/ProfileView.vue"),
        meta: {
          title: "Profil Pengguna",
          subtitle: "Ubah kata sandi dan kelola akun Anda.",
        },
      },
    ],
  },
  {
    path: "/:pathMatch(.*)*",
    redirect: { name: "Dashboard" },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore();
  const isLoggedIn = authStore.isAuthenticated;

  if (to.meta.requiresAuth && !isLoggedIn) {
    next({ name: "Login" });
  } else if (to.meta.requiresGuest && isLoggedIn) {
    next({ name: "Dashboard" });
  } else {
    next();
  }
});

export default router;
