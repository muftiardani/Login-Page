<script setup>
const props = defineProps({
  username: {
    type: String,
    required: true,
  },
});
const emit = defineEmits(["logout"]);
</script>

<template>
  <div class="status-view">
    <h2>Halo, {{ username }}!</h2>
    <p class="welcome-message">Anda berhasil login.</p>
    <button @click="emit('logout')" class="logout-button">Logout</button>
  </div>
</template>

<style scoped>
/* Scoped styles for StatusView */
</style>
