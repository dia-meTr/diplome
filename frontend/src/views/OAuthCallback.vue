<!-- src/views/OAuthCallback.vue -->
<template>
    <div>
      <p>Redirecting...</p>
    </div>
  </template>
  
  <script>
  import Cookies from 'js-cookie';
  
  export default {
    created() {
      // Отримання параметрів з URL
      const urlParams = new URLSearchParams(window.location.search);
      const accessToken = urlParams.get('access_token');
      const refreshToken = urlParams.get('refresh_token');
  
      // Збереження токенів у cookies
      if (accessToken) {
        Cookies.set('access_token', accessToken, { expires: 1 }); // 1 день
      }
      if (refreshToken) {
        Cookies.set('refresh_token', refreshToken, { expires: 7 }); // 7 днів
      }
  
      // Перенаправлення на головну сторінку
      this.$router.push('/');
    }
  }
  </script>
  
  <style scoped>
  p {
    font-size: 18px;
    color: #333;
  }
  </style>
  