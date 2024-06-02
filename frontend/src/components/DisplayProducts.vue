<template>
  <div class="product-header flex-center-vertically">
    <h1>Страви</h1>
 <!--   <div>
      <input
        type="text"
        placeholder="Пошук"
        class="search"
        @input="search = $event.target.value" />
    </div>-->
  </div>
  <div class="product-content">
    <SingleProduct 
      v-for="dish in dishes"
      :key="dish.id"
      :dish="dish" />
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'
import SingleProduct from '@/components/SingleProduct.vue'

export default {
	components: { 
		SingleProduct
	},
	data() {
		return {
			search: '',
		}
	},
	props: {
    	dishes: {
      		type: Array,
      		required: true
    	}
	},
}
</script>

<style scoped>
.product-header {
	margin-bottom: 4em;
	&.flex-center-vertically {
			justify-content: space-between;
	}
}
.product-content {
	display: flex;
	flex-wrap: wrap;
	justify-content: space-evenly;
	gap: 7em 4em;
	margin-bottom: 4em;
}

@media only screen and (min-width: 1540px) {
	.product-content {
		gap: 10em 4em;
	}
}
</style>
