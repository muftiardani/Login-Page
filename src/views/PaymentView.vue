<script setup>
import { onMounted, ref } from "vue";
import * as api from "@/api/auth.js";
import PaymentTable from "@/components/PaymentTable.vue";

const payments = ref([]);
const isLoading = ref(true);
const error = ref(null);

onMounted(async () => {
  try {
    payments.value = await api.getPayments();
  } catch (err) {
    error.value = "Gagal memuat data pembayaran. Silakan coba lagi nanti.";
    console.error(err);
  } finally {
    isLoading.value = false;
  }
});
</script>

<template>
  <div class="card">
    <div v-if="isLoading" class="loading-state">
      <p>Memuat data...</p>
    </div>
    <div v-else-if="error" class="error-state">
      <p>{{ error }}</p>
    </div>
    <PaymentTable v-else :payments="payments" />
  </div>
</template>

<style scoped>
.card {
  background-color: var(--card-background);
  border-radius: var(--border-radius);
  padding: 1.5rem;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.05);
  border: 1px solid var(--border-color);
}
.loading-state,
.error-state {
  text-align: center;
  padding: 3rem;
  color: var(--secondary-button-color);
}
</style>
