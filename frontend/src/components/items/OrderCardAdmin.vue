<template>
    <div class="order-card">
        <div class="order-card-header">
            <h3>Деталі замовлення</h3>
            <p>Статус: {{ order.status }}</p>
        </div>
        <div class="order-card-body">
            <div class="order-card-body-items">
                <div v-for="dish in order.dishes" :key="dish.id">
                    <div class="dish-info">
                        <img :src="dish.photo" :alt="dish.name" />
                        <div class="dish-details">
                            <h4>{{ dish.name }}</h4>
                            <p>Price: {{ dish.price }}</p>
                            <p>Amount: {{ dish.amount }}</p>
                        </div>
                    </div>
                </div>
                <my-button @click="changeStatus">Змінити статус</my-button>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'OrderCard',
    props: {
        order: {
            type: Object,
            required: true
        }
    }
}
</script>

<style scoped>
.order-card {
    border: 1px solid #ccc;
    padding: 10px;
    margin-bottom: 10px;
}

.order-card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 10px;
}

.order-card-body {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 10px;
    
}

.dish-info {
    display: flex;
    align-items: center;
    margin-bottom: 10px;
}

.dish-info img {
    width: 50px;
    height: 50px;
    margin-right: 10px;
}

.dish-details {
    display: flex;
    flex-direction: column;
}
</style>