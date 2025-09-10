<template>
  <div class="dashboard-container">
    <div class="stat-card">
      <h4>Total Stock</h4>
      <p>{{ totalStock }}</p>
      <span>units of all products</span>
    </div>
    <div class="stat-card">
      <h4>Product SKUs</h4>
      <p>{{ productCount }}</p>
      <span>unique products</span>
    </div>
    <div class="stat-card low-stock">
      <h4>Low Stock Items</h4>
      <p>{{ lowStockCount }}</p>
      <span>items needing attention</span>
    </div>
    <div class="stat-card out-of-stock">
      <h4>Out of Stock</h4>
      <p>{{ outOfStockCount }}</p>
      <span>items to restock</span>
    </div>
  </div>
</template>

<script>
export default {
  name: "InventoryDashboard",
  props: {
    products: {
      type: Array,
      required: true,
    },
  },
  computed: {
    totalStock() {
      return this.products.reduce((sum, product) => sum + product.quantity, 0);
    },
    productCount() {
      return this.products.length;
    },
    lowStockCount() {
      return this.products.filter((p) => p.status === "Low Stock").length;
    },
    outOfStockCount() {
      return this.products.filter((p) => p.status === "Out of Stock").length;
    },
  },
};
</script>

<style scoped>
.dashboard-container {
  display: flex;
  justify-content: space-around;
  gap: 20px;
  padding: 20px;
  margin: 0 20px;
  background-color: #f8f9fa;
  border-radius: 8px;
}
.stat-card {
  background-color: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  text-align: center;
  flex-grow: 1;
}
.stat-card h4 {
  margin: 0 0 10px 0;
  color: #6c757d;
}
.stat-card p {
  margin: 0 0 5px 0;
  font-size: 2.5em;
  font-weight: bold;
  color: #343a40;
}
.stat-card span {
  color: #6c757d;
  font-size: 0.9em;
}
.stat-card.low-stock p {
  color: #ffc107;
}
.stat-card.out-of-stock p {
  color: #dc3545;
}
</style>
