<template>
  <div class="user-management-container">
    <h3>User Management</h3>
    <p>This table lists all registered users in the system.</p>
    <table>
      <thead>
        <tr>
          <th>User ID</th>
          <th>Username</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="user in users" :key="user.id">
          <td>{{ user.id }}</td>
          <td>{{ user.username }}</td>
          <td>
            <button @click="$emit('delete-user', user.id)" class="delete-btn">
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
  name: "UserManagement",
  props: {
    users: {
      type: Array,
      required: true,
    },
  },
};
</script>

<style scoped>
.user-management-container {
  margin: 20px;
  padding: 20px;
  background-color: #fff;
  border: 1px solid #ddd;
  border-radius: 8px;
}
table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 15px;
}
th,
td {
  border: 1px solid #ddd;
  padding: 12px;
  text-align: left;
}
th {
  background-color: #f2f2f2;
}
.delete-btn {
  background-color: #dc3545;
  color: white;
  border: none;
  padding: 8px 12px;
  border-radius: 5px;
  cursor: pointer;
}
.delete-btn:hover {
  background-color: #c82333;
}
</style>
