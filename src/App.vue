<template>
  <router-view v-if="!isAuthenticated" @login-success="handleLogin" />
  <div v-else>
    <div class="header">
      <h1>Warehouse Management</h1>
      <button @click="handleLogout" class="logout-button">Logout</button>
    </div>

    <div class="navigation-tabs">
      <button @click="view = 'products'" :class="{ active: view === 'products' }">Product Management</button>
    </div>

    <!-- Product Management View -->
    <div v-if="view === 'products'">
      <InventoryDashboard :products="products" />
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

<script src="./App.js"></script>
<style src="./App.css"></style>