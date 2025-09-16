<script setup>
import { onMounted, ref } from "vue";
import * as api from "@/api/auth.js";
import SummaryCard from "@/components/dashboard/SummaryCard.vue";
import RecentPayments from "@/components/dashboard/RecentPayments.vue";
import PaymentsChart from "@/components/dashboard/PaymentsChart.vue";

const summary = ref(null);
const recentPayments = ref([]);
const chartData = ref([]);
const isLoading = ref(true);

const formatCurrency = (value) => {
  if (typeof value !== "number") return "Rp 0";
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    minimumFractionDigits: 0,
  }).format(value);
};

onMounted(async () => {
  try {
    const [summaryData, paymentsData, chartApiData] = await Promise.all([
      api.getDashboardSummary(),
      api.getPayments(),
      api.getChartData(),
    ]);
    summary.value = summaryData;
    recentPayments.value = paymentsData.slice(0, 5);
    chartData.value = chartApiData;
  } catch (error) {
    console.error("Gagal memuat data dashboard:", error);
  } finally {
    isLoading.value = false;
  }
});
</script>

<template>
  <div v-if="isLoading" class="loading-state">
    <p>Memuat data dashboard...</p>
  </div>

  <div v-else class="dashboard-content">
    <div class="summary-grid">
      <SummaryCard
        title="Total Pendapatan (Lunas)"
        :value="formatCurrency(summary?.total_revenue)"
        icon="💰"
      />
      <SummaryCard
        title="Pembayaran Selesai"
        :value="summary?.completed_payments"
        icon="✅"
      />
      <SummaryCard
        title="Pembayaran Tertunda"
        :value="summary?.pending_payments"
        icon="⏳"
      />
    </div>

    <div class="main-grid">
      <PaymentsChart :chartData="chartData" />
      <RecentPayments :payments="recentPayments" />
    </div>
  </div>
</template>

<style scoped>
.summary-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1.5rem;
  margin-bottom: 2rem;
}
.main-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1.5rem;
}
@media (min-width: 992px) {
  .main-grid {
    grid-template-columns: 2fr 1fr;
  }
}
.loading-state {
  text-align: center;
  padding: 3rem;
  color: var(--secondary-button-color);
}
</style>
