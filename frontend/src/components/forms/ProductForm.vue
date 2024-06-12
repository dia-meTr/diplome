<template>
  <form @submit.prevent>
    <my-input 
      v-model="tag.title"
      placeholder="Назва"
    />

    <my-button @click="addTag">Додати</my-button>
  </form>
</template>

<script>

export default {
  name: 'TagForm',

  data() {
    return {
      tag : {
        title: ''
      }
    }
  },
  methods: {
    addTag() {
      
      this.$emit('create', this.tag)

      this.tag = {
        title: ''
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
