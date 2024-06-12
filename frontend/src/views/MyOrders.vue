<template>
  <div class="order-page">
    <h1>Мої замовлення</h1>

    <div v-for="order in orders" :key="order.id" class="order-content">
      <OrderCard :order="order" />
    </div>

  </div>
</template>

<script>
import OrderCard from '@/components/OrderCard.vue';

export default {
  data() {
    return {
      orders: [
        {
          id: 1,
          dishes: [
            {
              id: 1,
              name: 'Борщ',
              price: 200,
              photo: 'https://bestofood-storage.s3.us-east-2.amazonaws.com/Borshch.jpg',
              amount: 1
            },

          ],
          status: 'new'
        },
        {
          id: 2,
          dishes: [
            {
              id: 3,
              name: 'Овочевий суп',
              price: 300,
              photo: 'https://bestofood-storage.s3.us-east-2.amazonaws.com/Ovochevyy_sup.jpg',
              amount: 2
            },
            {
              id: 4,
              name: 'Суп-пюре з гарбуза',
              price: 400,
              photo: 'https://bestofood-storage.s3.us-east-2.amazonaws.com/Sup-pyure_z_harbuza.jpg',
              amount: 1
            }
          ],
          status: 'delivered'
        }
      ] // Replace this with your actual list of orders
    };
  },

  components: {
    OrderCard
  }
};
</script>

<style>
.order-content{
  display: flex;
  flex-direction: column;
  gap: 20px;
  width: 50%;
}
.order-page {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
}
</style>