import { createRouter, createWebHistory } from "vue-router";
import { useAuthStore } from "@/stores/auth.js";
import AuthView from "@/views/AuthView.vue";
import StatusView from "@/views/StatusView.vue";
import ProfileView from "@/views/ProfileView.vue";

const routes = [
  {
    path: "/login",
    name: "Login",
    component: AuthView,
    meta: { requiresGuest: true },
  },
  {
    path: "/register",
    name: "Register",
    component: AuthView,
    meta: { requiresGuest: true },
  },
  {
    path: "/",
    name: "Status",
    component: StatusView,
    meta: { requiresAuth: true },
  },
  {
    path: "/profile",
    name: "Profile",
    component: ProfileView,
    meta: { requiresAuth: true },
  },
  {
    path: "/:pathMatch(.*)*",
    redirect: "/",
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
    next("/login");
  } else if (to.meta.requiresGuest && isLoggedIn) {
    next("/");
  } else {
    next();
  }
});

export default router;
