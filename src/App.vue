<script setup>
import { ref, onMounted } from "vue";
import LoginForm from "./components/LoginForm.vue";
import RegisterForm from "./components/RegisterForm.vue";
import StatusView from "./components/StatusView.vue";

const message = ref("");
const messageClass = ref("");
const isLoggedIn = ref(false);
const loggedInUser = ref("");
const currentView = ref("login");

const apiBaseUrl = "http://localhost:8080/api";

async function handleLogin(credentials) {
  try {
    const response = await fetch(`${apiBaseUrl}/login`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(credentials),
    });

    const data = await response.json();
    if (!response.ok) {
      throw new Error(data.message || "Login failed");
    }

    localStorage.setItem("token", data.token);
    localStorage.setItem("username", credentials.username);

    isLoggedIn.value = true;
    loggedInUser.value = credentials.username;
    message.value = `✅ ${data.message}`;
    messageClass.value = "success";

    fetchStatus();
  } catch (error) {
    message.value = `❌ ${error.message}`;
    messageClass.value = "error";
  }
}

async function handleRegister(credentials) {
  try {
    const response = await fetch(`${apiBaseUrl}/register`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(credentials),
    });
    const data = await response.json();
    if (!response.ok) {
      throw new Error(data.message || "Registration failed");
    }
    message.value = `✅ ${data.message}`;
    messageClass.value = "success";
    switchToLogin();
  } catch (error) {
    message.value = `❌ ${error.message}`;
    messageClass.value = "error";
  }
}

async function fetchStatus() {
  const token = localStorage.getItem("token");
  if (!token) {
    message.value = "Anda tidak login.";
    messageClass.value = "error";
    return;
  }

  try {
    const response = await fetch(`${apiBaseUrl}/status`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    });

    const data = await response.json();
    if (!response.ok) {
      handleLogout();
      throw new Error(
        data.message || "Sesi tidak valid, silakan login kembali."
      );
    }

    message.value = `API Status: ${data.message}`;
    messageClass.value = "success";
  } catch (error) {
    message.value = `❌ ${error.message}`;
    messageClass.value = "error";
  }
}

function handleLogout() {
  localStorage.removeItem("token");
  localStorage.removeItem("username");

  isLoggedIn.value = false;
  loggedInUser.value = "";
  message.value = "";
  messageClass.value = "";
  switchToLogin();
}

onMounted(() => {
  const token = localStorage.getItem("token");
  const username = localStorage.getItem("username");
  if (token && username) {
    isLoggedIn.value = true;
    loggedInUser.value = username;
    fetchStatus();
  }
});

function switchToRegister() {
  currentView.value = "register";
  message.value = "";
  messageClass.value = "";
}

function switchToLogin() {
  currentView.value = "login";
  message.value = "";
  messageClass.value = "";
}
</script>

<template>
  <div id="app-container">
    <Transition name="fade-slide" mode="out-in">
      <div v-if="isLoggedIn" :key="'status'" class="card-container">
        <div v-if="message" :class="['message', messageClass]">
          {{ message }}
        </div>
        <StatusView :username="loggedInUser" @logout="handleLogout" />
      </div>

      <div v-else :key="'auth-form'" class="card-container">
        <div v-if="message" :class="['message', messageClass]">
          {{ message }}
        </div>
        <LoginForm
          v-if="currentView === 'login'"
          @submit-login="handleLogin"
          @switch-to-register="switchToRegister"
        />
        <RegisterForm
          v-else
          @submit-register="handleRegister"
          @switch-to-login="switchToLogin"
        />
      </div>
    </Transition>
  </div>
</template>

<style>
:root {
  --primary-color: #1f1d59;
  --primary-hover: #1f1d59;
  --text-color: #333;
  --secondary-text-color: #666;
  --background-color: #e9ecef;
  --card-background: #ffffff;
  --border-color: #ddd;
  --focus-shadow: rgba(31, 29, 89, 0.2);
  --success-bg: #d4edda;
  --success-text: #155724;
  --error-bg: #f8d7da;
  --error-text: #721c24;
}

body {
  font-family: "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  margin: 0;
  background-color: var(--background-color);
  color: var(--text-color);
  line-height: 1.6;
}

#app-container {
  width: 100%;
  max-width: 400px;
  padding: 1rem;
  box-sizing: border-box;
}

.card-container {
  background-color: var(--card-background);
  padding: 30px;
  border-radius: 12px;
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.08);
  transition: all 0.3s ease-in-out;
}

h2 {
  text-align: center;
  color: var(--primary-color);
  margin-bottom: 25px;
  font-size: 1.8em;
  font-weight: 600;
}

.input-group {
  margin-bottom: 20px;
}

.input-group label {
  display: block;
  margin-bottom: 8px;
  font-weight: 500;
  color: var(--secondary-text-color);
}

.input-group input[type="text"],
.input-group input[type="password"] {
  width: 100%;
  padding: 12px 15px;
  border: 1px solid var(--border-color);
  border-radius: 8px;
  box-sizing: border-box;
  font-size: 1em;
  transition: border-color 0.2s, box-shadow 0.2s;
}

.input-group input:focus {
  border-color: var(--primary-color);
  box-shadow: 0 0 0 3px var(--focus-shadow);
  outline: none;
}

.submit-button,
.logout-button {
  width: 100%;
  padding: 12px 20px;
  border: none;
  border-radius: 8px;
  background-color: var(--primary-color);
  color: white;
  font-size: 1.1em;
  font-weight: 600;
  cursor: pointer;
  transition: background-color 0.2s ease, transform 0.1s ease;
}

.submit-button:hover,
.logout-button:hover {
  background-color: var(--primary-hover);
  transform: translateY(-2px);
}

.submit-button:active,
.logout-button:active {
  transform: translateY(0);
}

.toggle-view {
  text-align: center;
  margin-top: 25px;
  font-size: 0.9em;
  color: var(--secondary-text-color);
}

.toggle-view a {
  color: var(--primary-color);
  cursor: pointer;
  text-decoration: none;
  font-weight: 600;
  margin-left: 5px;
}

.toggle-view a:hover {
  text-decoration: underline;
}

.message {
  text-align: center;
  padding: 12px;
  margin-bottom: 20px;
  border-radius: 8px;
  font-size: 0.95em;
  font-weight: 500;
}

.success {
  background-color: var(--success-bg);
  color: var(--success-text);
  border: 1px solid var(--success-text);
}
.error {
  background-color: var(--error-bg);
  color: var(--error-text);
  border: 1px solid var(--error-text);
}

/* Status View specific styles */
.status-view {
  text-align: center;
  padding: 30px 20px;
}
.status-view h2 {
  color: var(--primary-color);
  margin-bottom: 15px;
  font-size: 2em;
}
.welcome-message {
  font-size: 1.2em;
  color: var(--secondary-text-color);
  margin-bottom: 30px;
}

/* Vue Transition Styles */
.fade-slide-enter-active,
.fade-slide-leave-active {
  transition: all 0.4s ease;
}

.fade-slide-enter-from {
  opacity: 0;
  transform: translateY(-20px);
}
.fade-slide-leave-to {
  opacity: 0;
  transform: translateY(20px);
}
</style>
