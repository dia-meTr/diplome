<template>
  <div class="single-product__wrap">
    <div 
      class="layer"
      @click="$emit('layer-popup')">
      <i class="fa-solid fa-xmark"></i>
    </div>
    <div class="single-product__wrap--img flex-center-vertically">
      <img
        :src="dish.thumbnail"
        :alt="dish.title" />
    </div>
    <div class="single-product__wrapdata">
      <div class="single-product__wrapdata--title">
        <h1>{{ dish.title }}</h1>
      </div>
      <div class="single-product__wrapdata--price">
        <h3>{{ dish.price.toLocaleString('ko-KR') }}Один</h3>
      </div>
      <div class="single-product__wrapdata--description">
        {{ dish.description }}
      </div>
      <div class="single-product__wrapdata--button">
        <button
          class="status-badge confirm-badge"
          @click="test">
          Модифікація продукту
        </button>
        <div
          v-if="modal"
          class="black">
          <div class="white">
            <EditProduct
              :dish="dish"
              @layer-popup="modal = false" />
          </div>
        </div>
        <div>
          <button 
            class="status-badge cancel-badge"
            @click="deleteProduct(dish.id)">
            Видалити товар
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { publicRequest } from '../api/publicRequest'
import EditProduct from '../components/EditProduct.vue'
export default {
  components: {
    EditProduct,
  },
  props: {
    dish: {
      type: Object,
      default: () => ({})
    }
  },
  emits: ['layer-popup'],
  data() {
    return {
      modal: false
    }
  },
  mounted() {
    this.singleProduct()
  },
  methods: {
    // Виправлення продукту
    async singleProduct() {
      const res = await publicRequest({
        url: `products/${this.product.id}`,
        method: 'GET'
      })
      console.log(res)
      this.product
    },
    // Видалити товар
    async deleteProduct() {
      const res = await publicRequest({
        url: `products/${this.product.id}`,
      method: 'DELETE'
    })
      console.log(res)
      alert('Видалити товар')
      this.$router.go(0)
    },
    test() {
      this.modal = true
    },
    a() {
      this.$router.go(0)
    }
  },
}
</script>

<style lang="scss" scoped>
  .single-product {
    &__wrap {
      display: flex;
      align-content: center;
      height: 100%;
      width: 100%;
      &--img {
        img {
          height: 30vh;
        }
      }
      .layer {
        position: absolute;
        right: 3rem;
        top: 2vh;
        &:hover {
          cursor: pointer;
        }
        .fa-solid {
          font-size: 3rem;
        }
      }
    }
    &__wrapdata {
      width: 100%;
      margin: auto;
      &--title {
        h1 {
          margin: 0;
          padding-bottom: 15px;
          border-bottom: 1px solid;
        }
      }
      &--price {
        padding: 15px 0;
        color: $color-primary;
      }
      &--description {
        padding: 5px 0;
        line-height: 1.7;
      }
      &--button {
        display: flex;
        margin-top: 3vh;
        .confirm-badge{
          background-color: $color-primary;
          color: $color-primary-light;
        }
        .cancel-badge {
          margin-left: 10px;
          background-color: $color-danger;
          color:  $color-danger-light;
        }
      }
        .white {
          width: 50vw;
          height: 64vh;
          background-color: white;
          position: absolute;
          top: 50%;
          left: 50%;
          -webkit-transform: translate(-50%, -50%);
          transform: translate(-50%, -50%);
          border-radius: 5px;
          padding: 2em;
  }
    }
  }
</style>
