import axios from 'axios';
import ProductList from './components/ProductList.vue';
import ProductForm from './components/ProductForm.vue';
import EditProductModal from './components/EditProductModal.vue';
import InventoryDashboard from './views/Dashboard.vue';

const API_URL = 'http://localhost:8081';

export default {
  name: 'App',
  components: {
    ProductList,
    ProductForm,
    EditProductModal,
    InventoryDashboard
  },
  data() {
    return {
      view: 'products',
      products: [],
      currentFilter: '',
      isEditModalVisible: false,
      productToEdit: null,
      isAuthenticated: localStorage.getItem('isAuthenticated') === 'true', // Read from localStorage
    }
  },
  created() {
    if (this.isAuthenticated) {
      this.fetchProducts();
    }
  },
  watch: {
    view(newView) {
      if (newView === 'products') {
        this.fetchProducts();
      }
    }
  },
  methods: {
    // Auth Methods
    handleLogin() {
      this.isAuthenticated = true;
      localStorage.setItem('isAuthenticated', 'true');
      this.fetchProducts();
    },
    handleLogout() {
      this.isAuthenticated = false;
      localStorage.removeItem('isAuthenticated');
      this.products = [];
      this.$router.push('/login'); // Redirect to login page
    },

    // Modal Methods
    openEditModal(product) {
      this.productToEdit = product;
      this.isEditModalVisible = true;
    },
    closeEditModal() {
      this.isEditModalVisible = false;
      this.productToEdit = null;
    },

    // API Methods for Products
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
    }
  }
}
