<template>
  <div>
    <!--<h1> Пости </h1>
    <h2>{{ $store.state.likes }}</h2>
    <div></div>
      <my-button @click="$store.commit('incrementLikes')">+</my-button>
      <my-button @click="$store.commit('decrementLikes')">-</my-button>
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
  -->
  </div>
</template>

<script>  
import PostForm from '@/components/PostForm.vue'
import PostList from '@/components/PostList.vue'
import axios from 'axios'
//import MyButton from './components/UI/MyButton.vue';
import { mapState, mapGetters, mapActions, mapMutations } from 'vuex'


export default {
  components: {
    PostForm,
    PostList
  },
  data() {
    return {

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
