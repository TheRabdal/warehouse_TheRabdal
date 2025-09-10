<template>
  <div class="product-list">
    <h2>Product Inventory</h2>
    <table>
      <thead>
        <tr>
          <th>Product Name</th>
          <th>SKU</th>
          <th>Quantity</th>
          <th>Location</th>
          <th>Status</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="product in products" :key="product.sku">
          <td>{{ product.name }}</td>
          <td>{{ product.sku }}</td>
          <td>{{ product.quantity }}</td>
          <td>{{ product.location }}</td>
          <td>{{ product.status }}</td>
          <td>
            <button @click="$emit('edit-product', product)">Update</button>
            <button @click="$emit('delete-product', product.sku)">
              Delete
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: "ProductList",
  props: {
    products: {
      type: Array,
      required: true,
    },
  },
};
</script>

<style scoped>
.product-list {
  margin: 20px;
}
table {
  width: 100%;
  border-collapse: collapse;
}
th,
td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}
th {
  background-color: #f2f2f2;
}
</style>
