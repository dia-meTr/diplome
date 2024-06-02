<template>
  <div>
    <h1> Пости </h1>
    <div class="app__btns">
      <my-button 
      @click="fetchPosts"
      >
      Завантажити пости
    </my-button>

    <my-select
      v-model="selectedSort"
      :options="sortOptions"
    />
    </div>
    
    <my-button @click="dialogVisible = true">Додати пост</my-button>
    <my-dialog v-model:show="dialogVisible">
      <PostForm @create="addPost"></PostForm>
    </my-dialog>
    
    <PostList 
    :posts="posts"
    @remove="removePost" 
    v-if="isLoading === false"
    ></PostList>

    <div v-else>Завантаження...</div>
    
  </div>
</template>

<script>  
import PostForm from '@/components/PostForm.vue'
import PostList from '@/components/PostList.vue'
import axios from 'axios'
//import MyButton from './components/UI/MyButton.vue';


export default {
  components: {
    PostForm,
    PostList
  },
  data() {
    return {
      posts: [],
      dialogVisible: false,
      isLoading: false,
      selectedSort: '',
      sortOptions: [
        { value: 'title', name: 'По назві' },
        { value: 'body', name: 'По опису'}] 
    }
  },
  methods: {
    addPost (post) {
      this.posts.push(post)
      this.dialogVisible = false
    },
    removePost (post) {
      this.posts = this.posts.filter(p => p.id !== post.id)
    },
    async fetchPosts() {
      try {
        this.isLoading = true
        const response = await axios.get('https://jsonplaceholder.typicode.com/posts?_limit=10')
        this.posts = response.data
      } catch (e) {
        console.error(e)
      } finally {
        this.isLoading = false
      }
    }
  },
  mounted() {
    this.fetchPosts()
  },
}
</script>

<style>

.app__btns {
  margin: 15px 0;
  display: flex;
  justify-content: space-between;
}

</style>
