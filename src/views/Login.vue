<template>
  <div class="auth-container">
    <h2>Login</h2>
    <form @submit.prevent="login">
      <div class="form-group">
        <label for="username">Username:</label>
        <input type="text" id="username" v-model="username" required>
      </div>
      <div class="form-group">
        <label for="password">Password:</label>
        <input type="password" id="password" v-model="password" required>
      </div>
      <button type="submit">Login</button>
    </form>
    <p v-if="error" class="error-message">{{ error }}</p>
    <p>Don't have an account? <router-link to="/register">Register here</router-link></p>
  </div>
</template>

<script>
import axios from 'axios';

const API_URL = 'http://localhost:8081';

export default {
  name: 'LoginView',
  data() {
    return {
      username: '',
      password: '',
      error: null
    };
  },
  methods: {
    async login() {
      try {
        const response = await axios.post(`${API_URL}/api/login`, {
          username: this.username,
          password: this.password
        });
        if (response.status === 200) {
          localStorage.setItem('isAuthenticated', 'true');
          this.$emit('login-success'); // Emit event to parent (App.vue) or handle directly
          this.$router.push('/dashboard'); // Redirect to dashboard
        } else {
          this.error = 'User tidak ditemukan...';
        }
      } catch (err) {
        if (err.response && (err.response.status === 401 || err.response.status === 400)) {
          this.error = 'User tidak ditemukan...';
        } else {
          this.error = 'Terjadi kesalahan saat login. Silakan coba lagi.'; // Generic error in Indonesian
          console.error('Login error:', err);
        }
      }
    }
  }
};
</script>

<style scoped>
.auth-container {
  max-width: 420px; /* Consistent with register */
  margin: 60px auto; /* Consistent with register */
  padding: 30px; /* Consistent with register */
  border: none;
  border-radius: 12px;
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);
  background: linear-gradient(145deg, #ffffff, #f0f0f0);
  font-family: 'Roboto', sans-serif;
  color: #333;
}

h2 {
  text-align: center;
  color: #2c3e50;
  margin-bottom: 30px;
  font-size: 32px;
  font-weight: 700;
  letter-spacing: -0.5px;
}

.form-group {
  margin-bottom: 22px;
}

label {
  display: block;
  margin-bottom: 8px;
  color: #555;
  font-weight: 500;
  font-size: 15px;
}

input[type='text'],
input[type='password'] {
  width: 100%;
  padding: 14px 15px;
  border: 1px solid #dcdcdc;
  border-radius: 8px;
  box-sizing: border-box;
  font-size: 17px;
  transition: border-color 0.3s ease, box-shadow 0.3s ease;
}

input[type='text']:focus,
input[type='password']:focus {
  border-color: #007bff;
  box-shadow: 0 0 0 3px rgba(0, 123, 255, 0.25);
  outline: none;
}

button[type='submit'] {
  width: 100%;
  padding: 15px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-size: 19px;
  font-weight: 600;
  letter-spacing: 0.5px;
  transition: background-color 0.3s ease, transform 0.2s ease, box-shadow 0.3s ease;
  box-shadow: 0 4px 10px rgba(0, 123, 255, 0.3);
}

button[type='submit']:hover {
  background-color: #0056b3;
  transform: translateY(-3px);
  box-shadow: 0 6px 15px rgba(0, 123, 255, 0.4);
}

button[type='submit']:active {
  transform: translateY(0);
  box-shadow: 0 2px 5px rgba(0, 123, 255, 0.2);
}

p {
  text-align: center;
  margin-top: 30px;
  color: #666;
  font-size: 16px;
}

router-link {
  color: #007bff;
  text-decoration: none;
  font-weight: 600;
  transition: color 0.3s ease;
}

router-link:hover {
  color: #0056b3;
  text-decoration: underline;
}

.error-message {
  color: #e74c3c;
  font-size: 14px;
  margin-top: 8px;
  font-weight: 500;
}
</style>