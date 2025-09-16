<script setup>
defineProps({
  payments: {
    type: Array,
    required: true,
  },
});

const formatCurrency = (value) => {
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    minimumFractionDigits: 0,
  }).format(value);
};
</script>

<template>
  <div class="card">
    <div class="card-header">
      <h3>Pembayaran Terakhir</h3>
      <router-link to="/payments" class="view-all-link"
        >Lihat Semua &rarr;</router-link
      >
    </div>
    <div class="table-container">
      <table v-if="payments.length > 0">
        <thead>
          <tr>
            <th>Nama Pelanggan</th>
            <th>Jumlah</th>
            <th>Status</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="payment in payments" :key="payment.id">
            <td>{{ payment.customer_name }}</td>
            <td class="amount">{{ formatCurrency(payment.amount) }}</td>
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
          </tr>
        </tbody>
      </table>
      <div v-else class="no-data">
        <p>Belum ada transaksi pembayaran.</p>
      </div>
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
  display: flex;
  flex-direction: column;
}
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
}
h3 {
  margin: 0;
  color: var(--primary-color);
}
.view-all-link {
  color: var(--accent-color);
  text-decoration: none;
  font-weight: 600;
  font-size: 0.9rem;
}
.table-container {
  width: 100%;
  overflow-x: auto;
}
table {
  width: 100%;
  border-collapse: collapse;
}
th,
td {
  padding: 12px 15px;
  text-align: left;
  border-bottom: 1px solid var(--border-color);
}
th {
  font-weight: 600;
  color: var(--secondary-button-color);
  font-size: 0.9rem;
  text-transform: uppercase;
}
tbody tr:last-child td {
  border-bottom: none;
}
.amount {
  font-weight: 600;
  color: var(--text-color);
}
.status-badge {
  padding: 5px 12px;
  border-radius: 12px;
  color: white;
  font-size: 0.85rem;
  font-weight: 600;
  white-space: nowrap;
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

.no-data {
  text-align: center;
  padding: 2rem;
  color: #888;
}
</style>
