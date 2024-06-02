<template>
  <TheHeader />
  <div class="app">
    <router-view></router-view>
  </div>
</template>

<script>
//import router from '@/router';
import TheHeader from '@/components/TheHeader.vue';
export default {
  components: {
    TheHeader
  },
  //router
} 


</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
.app {
  padding: 20px;
}
</style>