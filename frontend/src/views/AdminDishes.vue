<template>
  <div class="admin-dish-page">
    <h1>Страви</h1>
    <my-button @click="dialogVisible = true">Додати страву</my-button>
    <my-dialog v-model:show="dialogVisible">
      <DishForm @create="addDish"></DishForm>
    </my-dialog>
    <div class="dish-list">
      <DishItemAdmin v-for="dish in dishItems" :key="dish.id" :dish="dish" />
    </div>
  </div>
</template>

<script>
import DishItemAdmin from '@/components/items/DishItemAdmin.vue';
import DishForm from '@/components/forms/DishForm.vue';
import axios from 'axios'

export default {
  components: {
    DishItemAdmin,
    DishForm
  },
  data() {
    return {
      dishItems: [],
      dialogVisible: false,
    };
  },
  methods: {
    async addDish(dish) {
      try {
        const axios = require('axios');
        const options = {
          method: 'POST',
          url: this.$hostname + '/api/v1/dish',
          data: dish,
        };

        
//      const response = await axios.post('http://0.0.0.0:8080/api/v1/dish', dish);
      axios
      .request(options)
      .then(function (response) {
        console.log(response.data);
      })
      .catch(function (error) {
        console.error(error);
      });
      console.log(response.data);

      this.fetchDishes();
      } catch (error) {
      console.error(error);
      } finally {
      this.dialogVisible = false;
      }
    },
    async fetchDishes() {
    try {
        this.isLoading = true
        const response = await axios.get('http://0.0.0.0:8080/api/v1/dish?tags=1')
        this.dishItems = response.data
      } catch (e) {
        console.error(e)
      } finally {
        this.isLoading = false
      }
    }
  },
  mounted() {
    this.fetchDishes()
  },
};
</script>

<style scoped>
h1 {
  margin-bottom: 20px;
}
.dish-list{
  display: flex;
  flex-direction: column;
  gap: 20px;
  width: 50%;
}
.admin-dish-page {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
}

</style>
