<template>
  <form @submit.prevent>
    <my-input 
      v-model="dish.title"
      placeholder="Назва страви"
    />
    
    <my-input 
      v-model="dish.description"
      placeholder="Опис страви"
    />

    <my-input 
      v-model="dish.price"
      placeholder="Ціна страви"
    />

    <my-file-input 
      v-model="dish.photo"
      class="upload-input"
    >
      Завантажити фото
    </my-file-input>

    <my-button @click="addDish">Додати страви</my-button>
  </form>
</template>

<script>

export default {
  name: 'PostForm',

  data() {
    return {
      dish : {
        title: '',
        description: '',
        price: '',
        photo: null
      }
    }
  },
  methods: {
    addDish() {
      this.$emit('create', this.dish)

      this.dish = {
        title: '',
        description: '',
        price: '',
        photo: null
      } 
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
form {
  display: flex;
  flex-direction: column;
}
.upload-input {
  width: 100%;
  border: 1px solid teal;
  padding: 10px 15px;
  margin-top: 15px;
}

</style>
