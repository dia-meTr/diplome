<template>
    <div class="dish-info">
        <img :src="dish.photoUrl" 
        :alt="dish.name" 
        class="dish-image"/>
        <div class="dish-details">
            <h4>{{ dish.name }}</h4>
            <p>Price: {{ dish.price }}</p>
        </div>
        <div class="admin-buttons">
            <my-button @click="decreaseQuantity">Додати тег</my-button>
            <span>{{ amount }}</span>
            <my-button @click="increaseQuantity">Додати продукт</my-button>
        </div>
        <div class="admin-buttons">
            <my-button @click="decreaseQuantity">Видалити страву</my-button>
            <span>{{ amount }}</span>
            <my-button @click="increaseQuantity">Редагувати страву</my-button>
        </div>
    </div>

</template>


<script>
export default {
    name: 'DishItemAdmin',
    props: {
        dish: {
            type: Object,
            required: true
        }
    }
}
</script>

<style scoped>

.dish-info {
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

.admin-buttons {
    display: flex;
    flex-direction: column;
    gap: 5px;
    margin: 5px;
}
</style>