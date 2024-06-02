<template>

    <button class="btn">
      <slot></slot>
    </button>

</template>

<script>
export default {
  name: 'my-button'
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.btn {
  padding: 10px 15px;
  background: none;
  color: teal;
  border: 1px solid teal;
}
</style>
