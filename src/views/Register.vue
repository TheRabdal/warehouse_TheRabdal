<template>
  <div class="register-container">
    <h2>Register</h2>
    <form @submit.prevent="register">
      <div class="form-group">
        <label for="username">Username:</label>
        <input type="text" id="username" v-model="username" required />
      </div>
      <div class="form-group">
        <label for="password">Password:</label>
        <input type="password" id="password" v-model="password" required />
      </div>
      <div class="form-group">
        <label for="confirmPassword">Confirm Password:</label>
        <input type="password" id="confirmPassword" v-model="confirmPassword" required />
        <p v-if="passwordMismatch" class="error-message">Passwords do not match!</p>
      </div>
      <button type="submit">Register</button>
    </form>
    <p>Already have an account? <router-link to="/login">Login here</router-link></p>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'RegisterView',
  data() {
    return {
      username: '',
      password: '',
      confirmPassword: '',
      passwordMismatch: false,
    };
  },
  methods: {
    async register() {
      this.passwordMismatch = false; // Reset on each attempt

      if (this.password !== this.confirmPassword) {
        this.passwordMismatch = true;
        return;
      }

      try {
        await axios.post('http://localhost:8081/api/register', {
          username: this.username,
          password: this.password,
        });
        alert('Registration successful! Please login.');
        this.$router.push('/login');
      } catch (error) {
        alert('Registration failed: ' + (error.response?.data || error.message));
      }
    },
  },
};
</script>

<style scoped>
.register-container {
  max-width: 420px; /* Slightly wider container */
  margin: 60px auto; /* More vertical margin */
  padding: 30px; /* More padding */
  border: none; /* Remove border, rely on shadow */
  border-radius: 12px; /* More rounded corners */
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15); /* Stronger, softer shadow */
  background: linear-gradient(145deg, #ffffff, #f0f0f0); /* Subtle gradient background */
  font-family: 'Roboto', sans-serif; /* Professional font */
  color: #333; /* Darker text for better contrast */
}

h2 {
  text-align: center;
  color: #2c3e50; /* Darker, professional blue-grey */
  margin-bottom: 30px; /* More space below heading */
  font-size: 32px; /* Larger heading */
  font-weight: 700; /* Bolder heading */
  letter-spacing: -0.5px; /* Tighter letter spacing */
}

.form-group {
  margin-bottom: 22px; /* More space between form groups */
}

label {
  display: block;
  margin-bottom: 8px; /* Consistent space below label */
  color: #555; /* Slightly darker label color */
  font-weight: 500; /* Medium font weight */
  font-size: 15px; /* Slightly smaller label font */
}

input[type='text'],
input[type='password'] {
  width: 100%; /* Full width */
  padding: 14px 15px; /* More padding */
  border: 1px solid #dcdcdc; /* Lighter, cleaner border */
  border-radius: 8px; /* More rounded input fields */
  box-sizing: border-box;
  font-size: 17px; /* Larger font in input */
  transition: border-color 0.3s ease, box-shadow 0.3s ease; /* Smooth transitions for focus */
}

input[type='text']:focus,
input[type='password']:focus {
  border-color: #007bff; /* Highlight on focus */
  box-shadow: 0 0 0 3px rgba(0, 123, 255, 0.25); /* Subtle glow on focus */
  outline: none; /* Remove default outline */
}

button {
  width: 100%;
  padding: 15px; /* More padding */
  background-color: #007bff; /* Primary blue */
  color: white;
  border: none;
  border-radius: 8px; /* More rounded button */
  cursor: pointer;
  font-size: 19px; /* Larger font in button */
  font-weight: 600; /* Bolder button text */
  letter-spacing: 0.5px; /* Slightly spaced out text */
  transition: background-color 0.3s ease, transform 0.2s ease, box-shadow 0.3s ease; /* Smooth transitions */
  box-shadow: 0 4px 10px rgba(0, 123, 255, 0.3); /* Button shadow */
}

button:hover {
  background-color: #0056b3; /* Darker blue on hover */
  transform: translateY(-3px); /* More pronounced lift effect */
  box-shadow: 0 6px 15px rgba(0, 123, 255, 0.4); /* Stronger shadow on hover */
}

button:active {
  transform: translateY(0); /* Reset on click */
  box-shadow: 0 2px 5px rgba(0, 123, 255, 0.2); /* Smaller shadow on click */
}

p {
  text-align: center;
  margin-top: 30px; /* More space above link */
  color: #666; /* Softer grey */
  font-size: 16px;
}

router-link {
  color: #007bff; /* Primary blue for link */
  text-decoration: none;
  font-weight: 600; /* Bolder link */
  transition: color 0.3s ease; /* Smooth color transition */
}

router-link:hover {
  color: #0056b3; /* Darker blue on hover */
  text-decoration: underline;
}

.error-message {
  color: #e74c3c; /* More vibrant red for error messages */
  font-size: 14px;
  margin-top: 8px; /* More space above error message */
  font-weight: 500;
}

/* Global body styling for a cohesive look */
body {
  background-color: #f4f7f6; /* Light grey background for the entire page */
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh; /* Full viewport height */
  margin: 0;
}
</style>
