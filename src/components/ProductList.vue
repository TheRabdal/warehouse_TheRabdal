<template>
  <div class="product-list">
    <div class="product-list-header">
      <h2>Product Inventory</h2>
      <button @click="exportProductsCSV" class="export-button">Export to CSV</button>
    </div>
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
import axios from 'axios'; // Import axios

export default {
  name: "ProductList",
  props: {
    products: {
      type: Array,
      required: true,
    },
  },
  methods: {
    async exportProductsCSV() {
      try {
        const response = await axios.get('http://localhost:8081/api/products/export/csv', {
          responseType: 'blob', // Important for downloading files
        });

        // Create a Blob from the response data
        const blob = new Blob([response.data], { type: 'text/csv' });

        // Create a link element
        const link = document.createElement('a');
        link.href = window.URL.createObjectURL(blob);
        link.download = 'products.csv'; // Set the filename

        // Append to the body and click it
        document.body.appendChild(link);
        link.click();

        // Clean up
        document.body.removeChild(link);
        window.URL.revokeObjectURL(link.href);

        alert('Products exported successfully!');
      } catch (error) {
        console.error('Error exporting products:', error);
        alert('Failed to export products. Please try again.');
      }
    },
  },
};
</script>

<style scoped>
.product-list {
  margin: 20px;
}

.product-list-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px; /* Add some space below the header */
}

.export-button {
  background-color: #4CAF50; /* Green */
  color: white;
  padding: 10px 15px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 16px;
  transition: background-color 0.3s ease;
}

.export-button:hover {
  background-color: #45a049;
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

<style scoped>
.product-list {
  margin: 20px;
}

.product-list-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px; /* Add some space below the header */
}

.export-dropdown-container {
  position: relative;
  display: inline-block;
}

.export-button {
  background-color: #4CAF50; /* Green */
  color: white;
  padding: 10px 15px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 16px;
  transition: background-color 0.3s ease;
}

.export-button:hover {
  background-color: #45a049;
}

.export-dropdown-menu {
  position: absolute;
  top: 100%; /* Position below the button */
  right: 0; /* Align to the right of the button */
  background-color: white;
  box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
  z-index: 1;
  border-radius: 5px;
  overflow: hidden; /* Ensures rounded corners for children */
  min-width: 160px; /* Adjust as needed */
}

.export-dropdown-menu button {
  color: black;
  padding: 12px 16px;
  text-decoration: none;
  display: block;
  border: none;
  background: none;
  width: 100%;
  text-align: left;
  cursor: pointer;
  font-size: 14px;
}

.export-dropdown-menu button:hover {
  background-color: #f1f1f1;
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
