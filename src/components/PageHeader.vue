<script setup>
defineProps({
  title: {
    type: String,
    required: true,
  },
  subtitle: {
    type: String,
    default: "",
  },
});
</script>

<template>
  <header class="page-header">
    <div>
      <h1>{{ title }}</h1>
      <p v-if="subtitle" class="subtitle">{{ subtitle }}</p>
    </div>
    <div class="actions">
      <slot></slot>
    </div>
  </header>
</template>

<style scoped>
.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
  padding-bottom: 1rem;
  border-bottom: 1px solid var(--border-color);
}
h1 {
  font-size: 2.2rem;
  color: var(--primary-color);
  font-weight: 700;
  margin: 0;
}
.subtitle {
  color: var(--secondary-button-color);
  margin-top: 0.25rem;
}
.actions {
  display: flex;
  gap: 1rem;
}
</style>
