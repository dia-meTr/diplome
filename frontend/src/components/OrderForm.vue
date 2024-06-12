<template>
  <form @submit.prevent>
    <my-input 
      v-model="post.title"
      placeholder="Адреса"
    />
    
    <my-input 
      v-model="post.body"
      placeholder="Дата даставки"
    />

    <my-input 
      v-model="post.body"
      placeholder="Номер телефону"
    />

    <my-button @click="addPost">Замовити</my-button>
  </form>
</template>

<script>

export default {
  name: 'PostForm',

  data() {
    return {
      post : {
        title: '',
        body: ''
      }
    }
  },
  methods: {
    addPost() {
      
      this.post.id = Date.now()
      this.$emit('create', this.post)

      this.post = {
        title: '',
        body: ''
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
</style>
