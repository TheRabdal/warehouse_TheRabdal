<template>
  <div>
    <div class="header">
      <h1>Warehouse Management</h1>
      <button @click="handleLogout" class="logout-button">Logout</button>
    </div>

    <div class="navigation-tabs">
      <button @click="view = 'products'" :class="{ active: view === 'products' }">Product Management</button>
    </div>

    <!-- Product Management View -->
    <div v-if="view === 'products'">
      <!-- InventoryDashboard was the Dashboard.vue itself, removed recursive import -->
      <ProductForm @add-product="addProduct" />
      <div class="filter-container">
        <span>Filter by status:</span>
        <button @click="filterByStatus('')">All</button>
        <button @click="filterByStatus('In Stock')">In Stock</button>
        <button @click="filterByStatus('Low Stock')">Low Stock</button>
        <button @click="filterByStatus('Out of Stock')">Out of Stock</button>
                    
      </div>
      <ProductList
        :products="products"
        @delete-product="deleteProduct"
        @edit-product="openEditModal"
      />
      
    </div>
  </div>
  <EditProductModal
    v-if="isEditModalVisible"
    :product="productToEdit"
    @save-update="updateProduct"
    @close-modal="closeEditModal"
  />
</template>

<script>
import axios from 'axios';
import ProductList from '../components/ProductList.vue';
import ProductForm from '../components/ProductForm.vue';
import EditProductModal from '../components/EditProductModal.vue';

const API_URL = 'http://localhost:8081';

export default {
  name: 'WarehouseDashboard',
  components: {
    ProductList,
    ProductForm,
    EditProductModal,
  },
  data() {
    return {
      view: 'products',
      products: [],
      currentFilter: '',
      isEditModalVisible: false,
      productToEdit: null,
      isAuthenticated: localStorage.getItem('isAuthenticated') === 'true',
    }
  },
  created() {
    this.fetchProducts();
  },
  watch: {
    view(newView) {
      if (newView === 'products') {
        this.fetchProducts();
      }
    }
  },
  methods: {
    handleLogin() {
      this.isAuthenticated = true;
      localStorage.setItem('isAuthenticated', 'true');
      this.fetchProducts();
    },
    handleLogout() {
      this.isAuthenticated = false;
      localStorage.removeItem('isAuthenticated');
      this.products = [];
      this.$router.push('/login');
    },
    openEditModal(product) {
      this.productToEdit = product;
      this.isEditModalVisible = true;
    },
    closeEditModal() {
      this.isEditModalVisible = false;
      this.productToEdit = null;
    },
    fetchProducts() {
      let url = `${API_URL}/api/products`;
      if (this.currentFilter) {
        url += `?status=${this.currentFilter}`;
      }
      axios.get(url)
        .then(response => {
          this.products = response.data || [];
        })
        .catch(error => {
          console.error('Error fetching products:', error);
        });
    },
    filterByStatus(status) {
      this.currentFilter = status;
      this.fetchProducts();
    },
    addProduct(product) {
      axios.post(`${API_URL}/api/products`, product)
        .then(response => {
          this.products.push(response.data);
        })
        .catch(error => {
          console.error('Error adding product:', error);
        });
    },
    deleteProduct(sku) {
      axios.delete(`${API_URL}/api/products/${sku}`)
        .then(() => {
          this.products = this.products.filter(p => p.sku !== sku);
        })
        .catch(error => {
          console.error('Error deleting product:', error);
        });
    },
    updateProduct(updatedProduct) {
      axios.put(`${API_URL}/api/products/${updatedProduct.sku}`, updatedProduct)
        .then(response => {
          const index = this.products.findIndex(p => p.sku === updatedProduct.sku);
          if (index !== -1) {
            this.products.splice(index, 1, response.data);
          }
          this.closeEditModal();
        })
        .catch(error => {
          console.error('Error updating product:', error);
        });
    },
    
  }
}
</script>

<style scoped>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-smoothing: grayscale;
  color: #2c3e50;
}
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px;
  background-color: #f8f9fa;
}
.logout-button {
  padding: 8px 15px;
  background-color: #dc3545;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}
.logout-button:hover {
  background-color: #c82333;
}
.filter-container {
  margin: 0 20px 20px 20px;
  display: flex;
  gap: 10px;
  align-items: center;
}
.filter-container button {
  padding: 5px 10px;
  border: 1px solid #ccc;
  background-color: #f0f0f0;
  cursor: pointer;
  border-radius: 5px;
}
.filter-container button:hover {
  background-color: #e0e0e0;
}
.navigation-tabs {
  padding: 10px 20px;
  background-color: #e9ecef;
  border-bottom: 1px solid #dee2e6;
}
.navigation-tabs button {
  padding: 10px 15px;
  border: none;
  background: none;
  cursor: pointer;
  font-size: 1em;
}
.navigation-tabs button.active {
  border-bottom: 2px solid #42b983;
  font-weight: bold;
}
.export-button {
  padding: 5px 10px;
  border: 1px solid #007bff;
  background-color: #007bff;
  color: white;
  cursor: pointer;
  border-radius: 5px;
}
.export-button:hover {
  background-color: #0056b3;
}
</style>