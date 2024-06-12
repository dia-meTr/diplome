<template>
  <form @submit.prevent>
    <my-input 
      v-model="tag.name"
      placeholder="Назва Тегу"
    />

    <my-button @click="addTag">Додати Тег</my-button>
  </form>
</template>

<script>

export default {
  name: 'TagForm',

  data() {
    return {
       tag : {
        name: '',
      }
    }
  },
  methods: {
    addTag() {
      console.log(1)
      console.log(this.tag)

      this.$emit('create', this.tag)

      this.tag = {
        name: '',

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
