<template>
    <div class="shopping-cart-item">
        <img :src=dish.photoUrl
         :alt="dish.name" 
         class="dish-image" 
         />
        <div class="dish-details">
            <h3 class="dish-name">{{ dish.name }}</h3>
            <p class="dish-price">{{ dish.price }}</p>
            <p class="dish-quantity">Quantity: {{ amount }}</p>
            <div class="quantity-buttons">
                <button class="quantity-button" @click="decreaseQuantity">-</button>
                <button class="quantity-button" @click="increaseQuantity">+</button>
            </div>
        </div>
        <button class="remove-button" @click="removeItem">Remove</button>
    </div>
</template>

<script>
export default {
    props: {
        dish: {
            type: Object,
            required: true
        },
        amount: {
            type: Number,
            required: true
        }
    },
    methods: {
        removeItem() {
            // Implement the logic to remove the item from the shopping cart
        }
    }
};
</script>

<style scoped>
.shopping-cart-item {
    display: flex;
    align-items: center;
    margin-bottom: 20px;
}

.dish-image {
    width: 100px;
    height: 100px;
    object-fit: cover;
    margin-right: 20px;
}

.dish-details {
    flex-grow: 1;
}

.dish-name {
    font-size: 18px;
    margin-bottom: 5px;
}

.dish-price {
    font-size: 16px;
    margin-bottom: 5px;
}

.dish-quantity {
    font-size: 14px;
    color: gray;
}

.remove-button {
    background-color: red;
    color: white;
    border: none;
    padding: 5px 10px;
    cursor: pointer;
}
</style>