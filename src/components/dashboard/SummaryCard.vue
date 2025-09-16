<script setup>
defineProps({
  title: String,
  value: [String, Number],
  icon: String,
});
</script>

<template>
  <div class="summary-card">
    <div class="icon-wrapper">
      <span class="icon">{{ icon }}</span>
    </div>
    <div class="data">
      <div class="value">{{ value }}</div>
      <div class="title">{{ title }}</div>
    </div>
  </div>
</template>

<style scoped>
.summary-card {
  background-color: var(--card-background);
  border-radius: var(--border-radius);
  padding: 1.5rem;
  display: flex;
  align-items: center;
  gap: 1.5rem;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.05);
  border: 1px solid var(--border-color);
}
.icon-wrapper {
  background-color: #f0f4f8;
  border-radius: 50%;
  width: 60px;
  height: 60px;
  display: flex;
  justify-content: center;
  align-items: center;
}
.icon {
  font-size: 2rem;
}
.data {
  display: flex;
  flex-direction: column;
}
.value {
  font-size: 1.75rem;
  font-weight: 700;
  color: var(--primary-color);
}
.title {
  font-size: 0.9rem;
  color: var(--secondary-button-color);
}
</style>
