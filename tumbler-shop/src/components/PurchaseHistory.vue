<template>
  <div class="table-page">
    <h1>Історія покупок</h1>
    <div>
      <input
        class="search"
        type="text"
        placeholder="Пошук товару"
        @input="search = $event.target.value" />
    </div>
    <div class="table-wrapper">
      <table>
        <TableSkeletonUI />
        <tr>
          <th>Номер</th>
          <th>Назва продукту</th>
          <th>Ціна продукту</th>
          <th>Дата замовлення</th>
          <th>Підтвердження покупки</th>
          <th>Скасувати покупку</th>
          <th>Статус покупки</th>
        </tr>
        <tr
          v-for="product in searchedProducts"
          :key="product.id">
          <td>{{ searchedProducts.indexOf(product) + 1 }}</td>
          <td class="flex-center-vertically">
            <img :src="product.thumbnail" /><span>{{ product.title }}</span>
          </td>
          <td>{{ product.price }} 원</td>
          <td>{{ product.time }}</td>
          <td>
            <span
              v-show="product.status === '구매신청'"
              class="action-btn confirm-btn"
              @click="confirmPurchase(product)">구매확정</span>
          </td>
          <td>
            <span
              v-show="product.status === '구매신청'"
              class="action-btn cancel-btn"
              @click="cancelPurchase(product)">구매취소</span>
          </td>
          <td>
            <span
              v-if="product.status === '구매신청'"
              class="status-badge request-badge">{{ product.status }}</span>
            <span
              v-else-if="product.status === '구매확정'"
              class="status-badge confirm-badge">{{ product.status }}</span>
            <span
              v-else
              class="status-badge cancel-badge">{{ product.status }}</span>
          </td>
        </tr>
      </table>
    </div>
  </div>
</template>

<script>
import { publicRequest } from '~/api/publicRequest'
import TableSkeletonUI from '~/components/TableSkeletonUI.vue'

export default {
	components: {
		TableSkeletonUI
	},
	data() {
		return {
			search: '',
			purchasedProducts: [],
		}
	},
	computed: {
		searchedProducts() {
			return this.purchasedProducts.filter(product => {
				return product.title.match(this.search)
			})
		},
		myPageMenuList() {
			return this.$store.state.menu.myPageMenuList
		}
	},
	created() {
		this.$store.dispatch('menu/isShowLoading', true)
		this.getPurchaseHistory()
	},
	methods: {
		async getPurchaseHistory() {
			try {
				const res = await publicRequest({
					url: 'products/transactions/details',
					method: 'GET',
				})
				if (res[0].product) {
					this.createPurchasedProducts(res)
				} else throw new Error(res)
			} catch (err) {
				console.log(err)
			}	finally {
				this.$store.dispatch('menu/isShowLoading', false)
			}
		},
		createPurchasedProducts(res) {
			const purchaseHistoryLists = res.sort((a, b) => new Date(a.timePaid) - new Date(b.timePaid))
			purchaseHistoryLists.forEach(purchaseHistoryList => {
				let { product: {title, thumbnail, price}, timePaid, done, isCanceled, detailId } = purchaseHistoryList
				price = price.toLocaleString('ko-KR')
				timePaid = this.$dayjs(timePaid).format('YYYY년 MM월 DD일 HH시 mm분')
				let status
				if(done === false) {
					status = '구매신청'
					if(isCanceled === true) {
						status = '구매취소'
					}
				} else if(done === true) {
					status = '구매확정'
				}
				const info =  {
					title: title,
					thumbnail: thumbnail,
					price: price,
					time: timePaid,
					status: status,
					id: detailId
				}
				this.purchasedProducts.push(info)
			})
		},
		async confirmPurchase(purchasedProduct) {
			const body = {
				detailId: purchasedProduct.id
			}
			const res = await publicRequest({
				url: 'products/ok',
				method: 'POST',
				body: body
			})
			if(res === true) {
				alert('구매가 확정되었습니다.')
				this.$router.go()
			} else {
				alert(res.error)
			}
		},
		async cancelPurchase(purchasedProduct) {
			const body = {
				detailId: purchasedProduct.id
			}
			const res = await publicRequest({
				url: 'products/cancel',
				method: 'POST',
				body: body
			})
			if(res === true) {
				alert('구매가 취소되었습니다.')
				this.$router.go()
			} else {
				alert(res.error)
			}
		}
	}
}
</script>

<style lang="scss" scoped>
.flex-center-vertically {
	img {
		width: 30px;
    margin-right: 0.5em;
	}
}
</style>
