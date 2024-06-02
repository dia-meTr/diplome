<template>
    <div>
      <h1> Моя Корзина </h1>

      <div v-if="!cart.length"> Корзина пуста </div>

      <div v-else>
        <ShoppingCartItem
          v-for="item in cart"
          :key="item.id"
          :dish="item.dish"
          :amount="item.amount"
        />


      </div>
    </div>
</template>

<script>
import { mapState } from 'vuex'
import ShoppingCartItem from '@/components/ShoppingCartItem.vue'

export default {
  components: {
    ShoppingCartItem
  },
  data() {
    return {
      cart: [
        {
          id: 1,
          dish: {
            name: 'Product 1',
            price: 100,
            photoUrl: 'https://bestofood-storage.s3.us-east-2.amazonaws.com/food-basic.jpg'
          },
          amount: 1
        },
        {
          id: 2,
          dish: {
            name: 'Product 2',
            price: 200,
            photoUrl: 'https://bestofood-storage.s3.us-east-2.amazonaws.com/food-basic.jpg'
          },
          amount: 2
        },
      ]
    }
  },


}

</script>

<style scoped>
h1 {
  margin-bottom: 20px;
}
</style>
