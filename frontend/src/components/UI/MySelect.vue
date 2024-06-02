<template>
  <select :value="modelValue" @change="changeOption" class="select">
    <option disabled value="">Оберіть опцію</option>
    <option 
      v-for="option in options" 
      :key="option.value" 
      :value="option.value">
      {{ option.name }}
    </option>
  </select>
</template>

<script>

export default {
  name: 'my-select',

  props: {
    modelValue: {
      type: String,
    },
    options: {
      type: Array,
      default: () => []
    }
  },

  data() {
    return {
      post : {
        title: '',
        body: ''
      }
    }
  },
  methods: {
    changeOption(event) {
      this.$emit('update:modelValue', event.target.value)
    } 
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
