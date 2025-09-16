<script setup>
defineProps({
  payments: {
    type: Array,
    required: true,
  },
});

// Fungsi untuk format mata uang dan tanggal
const formatCurrency = (value) => {
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
  }).format(value);
};

const formatDate = (dateString) => {
  const options = {
    year: "numeric",
    month: "long",
    day: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  };
  return new Date(dateString).toLocaleDateString("id-ID", options);
};
</script>

<template>
  <div class="table-container">
    <table v-if="payments.length > 0">
      <thead>
        <tr>
          <th>Nama Pelanggan</th>
          <th>Jumlah</th>
          <th>Status</th>
          <th>Tanggal Pembayaran</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="payment in payments" :key="payment.id">
          <td>{{ payment.customer_name }}</td>
          <td>{{ formatCurrency(payment.amount) }}</td>
          <td>
            <span
              :class="[
                'status-badge',
                `status-${payment.status.toLowerCase()}`,
              ]"
            >
              {{ payment.status }}
            </span>
          </td>
          <td>{{ formatDate(payment.payment_date) }}</td>
        </tr>
      </tbody>
    </table>
    <div v-else class="no-data">
      <p>Tidak ada data pembayaran yang ditemukan.</p>
    </div>
  </div>
</template>

<style scoped>
.table-container {
  width: 100%;
  overflow-x: auto;
}
table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 1.5rem;
}
th,
td {
  padding: 12px 15px;
  border: 1px solid #ddd;
  text-align: left;
}
th {
  background-color: var(--primary-color);
  color: white;
  font-weight: 600;
}
tbody tr:nth-of-type(even) {
  background-color: #f3f3f3;
}
.no-data {
  text-align: center;
  margin-top: 2rem;
  color: #666;
}
.status-badge {
  padding: 5px 10px;
  border-radius: 12px;
  color: white;
  font-size: 0.85rem;
  font-weight: 600;
}
.status-lunas {
  background-color: #28a745;
}
.status-tertunda {
  background-color: #ffc107;
  color: #333;
}
.status-dibatalkan {
  background-color: #dc3545;
}
</style>
