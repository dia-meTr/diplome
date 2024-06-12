<template>
    <div class="cart-page">
      <h1> Моя Корзина </h1>

      <div v-if="!cart.length"> Корзина пуста </div>

      <div class="shopping-cart" v-else>
        <ShoppingCartItem
          v-for="item in cart"
          :key="item.id"
          :dish="item.dish"
          :amount="item.amount"
        />
        <div>
          <h2> Total: 200 </h2>
          <my-button @click="dialogVisible = true"> Оформити замовлення </my-button>
          <my-dialog v-model:show="dialogVisible">
            <OrderForm @create="addPost"></OrderForm>
          </my-dialog>
      </div>
    </div>
    </div>
</template>

<script>
import { mapState } from 'vuex'
import ShoppingCartItem from '@/components/ShoppingCartItem.vue'
import OrderForm from '@/components/OrderForm.vue';

export default {
  components: {
    ShoppingCartItem,
    OrderForm
  },
  data() {
    return {
      dialogVisible: false,
      cart: [
        {
          id: 1,
          dish: {
            name: 'Борщ',
            price: 200,
            photoUrl: 'https://bestofood-storage.s3.us-east-2.amazonaws.com/Borshch.jpg'
          },
          amount: 1
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
.shopping-cart{
  display: flex;
  flex-direction: column;
  gap: 20px;
  width: 50%;
}
.cart-page {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
}
</style>
