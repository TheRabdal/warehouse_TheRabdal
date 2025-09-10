<template>
  <div class="modal-overlay">
    <div class="modal-content">
      <h2>Edit Product</h2>
      <form @submit.prevent="save">
        <div class="form-group">
          <label>Product Name</label>
          <input type="text" v-model="editableProduct.name" required />
        </div>
        <div class="form-group">
          <label>SKU</label>
          <input type="text" v-model="editableProduct.sku" disabled />
        </div>
        <div class="form-group">
          <label>Quantity</label>
          <input
            type="number"
            v-model.number="editableProduct.quantity"
            required
          />
        </div>
        <div class="form-group">
          <label>Location</label>
          <input type="text" v-model="editableProduct.location" required />
        </div>
        <div class="form-group">
          <label>Status</label>
          <select v-model="editableProduct.status">
            <option>In Stock</option>
            <option>Low Stock</option>
            <option>Out of Stock</option>
          </select>
        </div>
        <div class="modal-actions">
          <button type="submit" class="save-btn">Save Changes</button>
          <button type="button" @click="close" class="cancel-btn">
            Cancel
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  name: "EditProductModal",
  props: {
    product: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {
      editableProduct: null,
    };
  },
  created() {
    // Clone the prop to a local data property to avoid mutating the prop directly.
    this.editableProduct = { ...this.product };
  },
  methods: {
    save() {
      this.$emit("save-update", this.editableProduct);
    },
    close() {
      this.$emit("close-modal");
    },
  },
};
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}
.modal-content {
  background: white;
  padding: 20px;
  border-radius: 8px;
  width: 400px;
}
.form-group {
  margin-bottom: 15px;
}
label {
  display: block;
  margin-bottom: 5px;
}
input,
select {
  width: 100%;
  padding: 8px;
  box-sizing: border-box;
}
.modal-actions {
  text-align: right;
  margin-top: 20px;
}
.save-btn {
  background-color: #42b983;
  color: white;
  padding: 10px 15px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}
.cancel-btn {
  background-color: #6c757d;
  color: white;
  padding: 10px 15px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  margin-left: 10px;
}
</style>
