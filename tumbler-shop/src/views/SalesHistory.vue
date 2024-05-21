<template>
  <h1>історія продажів</h1>
  <div class="sales-history">
    <button @click="salesProduct">
      Повна історія продажів
    </button>
  </div>
  <div class="table-page">
    <div class="table-wrapper">
      <table>
        <TableSkeletonUI />
        <tr>
          <th>Номер</th>
          <th>Покупець</th>
          <th>Назва продукту</th>
          <th>Ціна продукту</th>
          <th>Дата замовлення</th>
          <th>Статус покупки</th>
        </tr>
        <salesList
          v-for="(prodect,index) in salesDetails"
          :key="prodect"
          :index="index"
          :sales="prodect" />
      </table>
    </div>
  </div>
</template>

<script>
import salesList from '~/components/SalesList.vue'
import TableSkeletonUI from '~/components/TableSkeletonUI.vue'

export default {
  components: {
    salesList,
    TableSkeletonUI
  },
  computed: {
    salesDetails() {
      return this.$store.state.admin.salesDetails
    },
  },
  created() {
    this.salesProduct()
  },
  methods: {
    async salesProduct() {
      this.$store.dispatch('menu/isShowLoading', true)
      try {
        await this.$store.dispatch('admin/salesProduct')
      } catch (err) {
        alert(err)
      } finally {
        this.$store.dispatch('menu/isShowLoading', false)
      }
    }
  }
}
</script>

<style lang="scss" scoped>
  button {
    margin-top: 0;
  }
</style>
