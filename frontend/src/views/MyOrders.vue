<template>
  <div>
    <h1>Мої замовлення</h1>

    <div v-for="order in orders" :key="order.id">
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
              name: 'Product 1',
              price: 100,
              photo: 'https://bestofood-storage.s3.us-east-2.amazonaws.com/food-basic.jpg',
              amount: 1
            },
            {
              id: 2,
              name: 'Product 2',
              price: 200,
              photo: 'https://bestofood-storage.s3.us-east-2.amazonaws.com/food-basic.jpg',
              amount: 2
            }
          ],
          status: 'new'
        },
        {
          id: 2,
          dishes: [
            {
              id: 3,
              name: 'Product 3',
              price: 300,
              photo: 'https://bestofood-storage.s3.us-east-2.amazonaws.com/food-basic.jpg',
              amount: 3
            },
            {
              id: 4,
              name: 'Product 4',
              price: 400,
              photo: 'https://bestofood-storage.s3.us-east-2.amazonaws.com/food-basic.jpg',
              amount: 4
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