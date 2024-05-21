<template>
  <TheHeader />
  <div class="container">
    <div class="product-detail flex-center-vertically">
      <div
        class="thumbnail">
        <img
          :src="product.thumbnail"
          :alt="product.title" />
      </div>
      <div class="content">
        <h2>{{ product.title }}</h2>
        <hr />
        <h3 class="price">
          {{ parseInt(product.price).toLocaleString('ko-KR') }}Гривень
        </h3>
        <p>{{ product.description }}</p>
        <button
          @click="purchase">
          Покупка
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import TheHeader from '~/components/TheHeader'
import { publicRequest } from '~/api/publicRequest'
import { mapState } from 'vuex'

export default {
	components: {
		TheHeader
	},
	data() {
		return {
			product: {}
		}
	},
	computed: {
    ...mapState('user', [
      'isLogIn',
    ]),
  },
	created() {
		this.getProductDetail()
	},
	methods: {
		async getProductDetail() {
			const res = await publicRequest({
				url: `products/${this.$route.params.id}`,
				method: 'GET',
			})
			this.product = res
		},
		purchase() {
			if(!this.isLogIn) {
        alert('Ви вийшли з системи. Будь ласка, увійдіть знову')
				this.$router.push('/signin')
				return
			}
			this.$router.push({ name: 'requestpurchase', params: { id: this.product.id }, query: {title: this.product.title, price: this.product.price, thumbnail: this.product.thumbnail} })
		}
	}
}
</script>

<style lang="scss" scoped>
.product-detail {
	padding: 5em 10em 0;
	&.flex-center-vertically {
		gap: 5em;
	}
}
</style>
