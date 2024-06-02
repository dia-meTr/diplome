<template>
  <div v-if="posts.length > 0">
    
      <h3>Пости</h3>
      
        <post-item
          v-for="post in posts"
          :post="post"
          :key="post.id"
          @remove="$emit('remove', post)"
        />
    
  </div> 
  <h2 v-else>Нажаль страв із заданими тегами немає</h2>
</template>

<script>
import PostItem from '@/components/PostItem.vue'
export default {
  name: 'PostList',
  components: {
    PostItem
  },
  props: {
    posts: {
      type: Array,
      required: true
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>



h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
