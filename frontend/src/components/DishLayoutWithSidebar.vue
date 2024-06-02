<template>
  <div class="dish-layout">
    <div class="sidebar">
      <SidebarFilter />
    </div>
    <div class="main-content">
      <DisplayProducts 
        :dishes="dishes"
      />
    </div>
  </div>
</template>

<script>
//import { mapState, mapActions } from 'vuex'
import SidebarFilter from './SidebarFilter.vue'
import DisplayProducts from './DisplayProducts.vue'
import axios from 'axios'

export default {
    name: 'DishLayoutWithSidebar',
    components: { 
		SidebarFilter,
    DisplayProducts
	},
  data() {
    return {
      dishes: [],
      isLoading: false,
      selectedSort: '',
    }
  },
  methods: {
    async fetchPosts() {
    try {
        this.isLoading = true
        const response = await axios.get('http://0.0.0.0:8080/api/v1/dish')
        this.dishes = response.data
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

<style scoped>

.dish-layout {
    display: flex;
    height: 100%; /* Takes the remaining height of the container */
    margin: 0;
}

.sidebar {
    width: 250px;
    background-color: #fff;
    border-right: 1px solid #ddd;
    
    box-sizing: border-box;
    overflow-y: auto;
    height: 100%; /* Takes the full height of the parent (dish-layout) */
    position: relative;
}

.main-content {
    /* margin-left: 250px; Indent to not overlap sidebar */
    padding: 20px;
    flex: 1;
    overflow-y: auto;
}
</style>
