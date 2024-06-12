<template>
  <div class="admin-tag-page">
    <h1>Теги</h1>
    <my-button @click="dialogVisible = true">Додати тег</my-button>
    <my-dialog v-model:show="dialogVisible">
      <TagForm @create="addTag"></TagForm>
    </my-dialog>
    <div class="tag-list">
      <TagItemAdmin 
      v-for="tag in tagItems" 
      :key="tag.id" 
      :tag="tag" 
      @delete="deleteTag"/>
    </div>
  </div>
</template>

<script>
import TagItemAdmin from '@/components/items/TagItemAdmin.vue';
import TagForm from '@/components/forms/TagForm.vue';
import axios from 'axios'

export default {
  components: {
    TagItemAdmin,
    TagForm
  },
  data() {
    return {
      tagItems: [],
      dialogVisible: false,
    };
  },
  methods: {
    async addTag (tag) {
        try {
          const axios = require('axios');

          console.log(3)        
        const response = await axios.post(this.$hostname + '/api/v1/tag', tag);

        this.fetchTags();
        } catch (error) {
        console.error(error);
        } finally {
        this.dialogVisible = false;
        }
      
      this.dialogVisible = false
    },
    async deleteTag (id) {
      try {
        const axios = require('axios');
        const options = {
          method: 'DELETE',
          url: this.$hostname + '/api/v1/tag/' + id,
        };
        axios
        .request(options)
        .then(function (response) {
          console.log(response.data);
        })
        .catch(function (error) {
          console.error(error);
        });
        this.fetchTags();
      } catch (error) {
        console.error(error);
      }
    },
    async fetchTags() {
    try {
        this.isLoading = true
        const response = await axios.get(this.$hostname + '/api/v1/tag')
        this.tagItems = response.data
      } catch (e) {
        console.error(e)
      } finally {
        this.isLoading = false
      }
    }
  },
  mounted() {
    this.fetchTags()
  },
};
</script>

<style scoped>
h1 {
  margin-bottom: 20px;
}
.tag-list{
  display: flex;
  flex-direction: column;
  gap: 20px;
  width: 50%;
}
.admin-tag-page {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
}

</style>
