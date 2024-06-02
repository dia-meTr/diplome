<template>
  <div class="post">
    <div>
      <div> <strong>Назва</strong> {{ post.title }} </div>
      <div> <strong>Опис</strong> {{ post.body }} </div>
    </div>
    <div class="post_btn" > 
      <my-button 
        @click="$emit('remove', post)"
        >
        Видалити
      </my-button>
    
      <my-button 
        @click="$router.push(`posts/${post.id}`)"
        >
        Відкрити 
      </my-button>
    </div>
  </div>
</template>

<script>

export default {
  name: 'post-item',

  props: {
    post :{
      type: Object,
      required: true
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.post {
  padding: 15px;
  border: 2px solid teal;
  margin-top: 15px;
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.post__btns {
  display: flex;
}
</style>
