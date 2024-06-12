<!-- src/views/Home.vue -->
<template>
    <div>
      <h1>Welcome to Our Site</h1>
      <button @click="redirectToOAuth">Login with Google</button>
    </div>
  </template>
  
  <script>
  import axios from 'axios';

  export default {
    name: 'Home',
    code: '1',
    methods: {
    async redirectToOAuth() {
    try {
        this.isLoading = true
        const response = await axios.get(this.$hostname + '/api/v1/oauth2callback')
        this.code = response.data
      } catch (e) {
        console.error(e)
      } finally {
        console.log(this.code)  
        this.isLoading = false
      }
    }
  },

  }
  </script>
  
  <style scoped>
  button {
    background-color: #4285f4;
    color: white;
    border: none;
    padding: 10px 20px;
    font-size: 16px;
    cursor: pointer;
  }
  </style>
  