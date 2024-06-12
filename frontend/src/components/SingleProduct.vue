<template>
  <div class="product-wrapper flex-center-vertically">
    <RouterLink
      :to="{
        name: 'dishdetail',
        params: { id: dish.id },
      }"
      class="dishcard-parent">
      <div
        class="dishcard-child"
        :style="{ backgroundImage: `url(${ dish.photoUrl })` }"></div>
    </RouterLink>
    <p class="product-title text-center">
      {{ dish.name }}
    </p>
    <my-button
      class="btn btn-primary"
      @click="addToCart(dish)">
      Додати до кошика
    </my-button>
  </div>
</template>

<script>
export default {
  props: {
    dish: {
      type: Object,
      default: () => ({})
    }
  }
}
</script>

<style scoped>
.product-wrapper.flex-center-vertically {
  display: flex;
  flex-direction: column;
}
.dishcard-parent {
  width: 13em; 
  height: 13em;
  margin-bottom: 3em;
  &:hover .dishcard-child,
  &:focus .dishcard-child {
    transform: scale(1.1);
  }
  .dishcard-child {
    width: 100%;
    height: 100%;
    background-color: white;
    background-position: center;
    background-size: cover;
    transition: all .5s;
  }
}

@media only screen and (min-width: 1540px) {
  .dishcard-parent {
    width: 230px; 
    height: 230px;
  }
}
</style>
