<script setup>
import { Bar } from "vue-chartjs";
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  BarElement,
  CategoryScale,
  LinearScale,
} from "chart.js";
import { computed } from "vue";

ChartJS.register(
  Title,
  Tooltip,
  Legend,
  BarElement,
  CategoryScale,
  LinearScale
);

const props = defineProps({
  chartData: {
    type: Array,
    required: true,
  },
});

const data = computed(() => ({
  labels: props.chartData.map((d) =>
    new Date(d.label).toLocaleDateString("id-ID", {
      day: "2-digit",
      month: "short",
    })
  ),
  datasets: [
    {
      label: "Pendapatan Harian (Lunas)",
      backgroundColor: "#003366",
      borderRadius: 4,
      data: props.chartData.map((d) => d.value),
    },
  ],
}));

const options = {
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    legend: {
      display: false,
    },
    tooltip: {
      callbacks: {
        label: function (context) {
          let label = context.dataset.label || "";
          if (label) {
            label += ": ";
          }
          if (context.parsed.y !== null) {
            label += new Intl.NumberFormat("id-ID", {
              style: "currency",
              currency: "IDR",
            }).format(context.parsed.y);
          }
          return label;
        },
      },
    },
  },
  scales: {
    y: {
      beginAtZero: true,
      ticks: {
        callback: function (value, index, ticks) {
          return "Rp " + value / 1000 + "k";
        },
      },
    },
  },
};
</script>

<template>
  <div class="card">
    <h3>Pendapatan 7 Hari Terakhir</h3>
    <div class="chart-container">
      <Bar :data="data" :options="options" />
    </div>
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
h3 {
  margin-bottom: 1.5rem;
  color: var(--primary-color);
}
.chart-container {
  height: 300px;
  position: relative;
}
</style>
