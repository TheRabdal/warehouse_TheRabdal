<template>
  <div class="product-form">
    <h3>Add New Product</h3>
    <form @submit.prevent="handleSubmit">
      <input
        type="text"
        v-model="product.name"
        placeholder="Product Name"
        required
      />
      <input type="text" v-model="product.sku" placeholder="SKU" required />
      <input
        type="number"
        v-model.number="product.quantity"
        placeholder="Quantity"
        required
      />
      <input
        type="text"
        v-model="product.location"
        placeholder="Location"
        required
      />
      <button type="submit">Add Product</button>
    </form>
  </div>
</template>

<script>
export default {
  name: "ProductForm",
  data() {
    return {
      product: {
        name: "",
        sku: "",
        quantity: 0,
        location: "",
        status: "",
      },
    };
  },
  methods: {
    handleSubmit() {
      // Determine status based on quantity
      if (this.product.quantity > 10) {
        this.product.status = "In Stock";
      } else if (this.product.quantity > 0) {
        this.product.status = "Low Stock";
      } else {
        this.product.status = "Out of Stock";
      }
      this.$emit("add-product", this.product);
      // Reset form
      this.product = {
        name: "",
        sku: "",
        quantity: 0,
        location: "",
        status: "",
      };
    },
  },
};
</script>

<style scoped>
.product-form {
  margin: 20px;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
}
form {
  display: flex;
  flex-direction: column;
  gap: 10px;
}
input {
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 3px;
}
button {
  padding: 10px;
  background-color: #42b983;
  color: white;
  border: none;
  border-radius: 3px;
  cursor: pointer;
}
button:hover {
  background-color: #369f77;
}
</style>
