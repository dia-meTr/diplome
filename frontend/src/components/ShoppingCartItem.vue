<template>
    <div class="shopping-cart-item">
        <img :style="{ backgroundImage: `url(${ dish.photoUrl })`}"
         :alt="dish.name" 
         class="product-image" 
         />
        <div class="product-details">
            <h3 class="product-name">{{ dish.name }}</h3>
            <p class="product-price">{{ dish.price }}</p>
            <p class="product-quantity">Quantity: {{ amount }}</p>
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

.product-image {
    width: 100px;
    height: 100px;
    object-fit: cover;
    margin-right: 20px;
}

.product-details {
    flex-grow: 1;
}

.product-name {
    font-size: 18px;
    margin-bottom: 5px;
}

.product-price {
    font-size: 16px;
    margin-bottom: 5px;
}

.product-quantity {
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