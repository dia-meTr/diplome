<template>
  
  <div class="container">
    <DishLayoutWithSidebar />
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'
import TheHeader from '@/components/TheHeader'
import DishLayoutWithSidebar from '@/components/DishLayoutWithSidebar.vue'

export default {
  components: {
    TheHeader,
    DishLayoutWithSidebar
  },
  computed: {
    ...mapState('user', [
      'user',
      'isLogIn',
      'isAdmin'
    ])
  },
  methods: {
    ...mapActions('user', [
      'requestUpdateState'
    ])
  }
}
</script>

<style>
body {
    font-family: Arial, sans-serif;
    background-color: #f4f4f9;
    margin: 0;
    padding: 0;
}

.container{
    width: 100%;
    height: 100%;
}
</style>
