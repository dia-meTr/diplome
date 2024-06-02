

export const postModule = {
    state: () => ({
        posts: [],
        dialogVisible: false,
        isLoading: false,
        selectedSort: '',
        sortOptions: [
          { value: 'title', name: 'По назві' },
          { value: 'body', name: 'По опису'}] 
    }),
    getters: {
        getPosts(state) {
            return state.posts
        },
        getPost(state) {
            return state.post
        },
        sortPosts(state) {
            state.posts = [...state.posts].sort((a, b) => {
                return a[state.selectedSort].localeCompare(b[state.selectedSort])
            })
        },
    },
    mutations: {
        setPosts(state, posts) {
            state.posts = posts
        },
        setLoading(state, bool) {
            state.isLoading = bool
        },
        setDialogVisible(state, bool) {
            state.dialogVisible = bool
        },
        setSelectedSort(state, value) {
            state.selectedSort = value
        },
    },
    actions: {
        async fetchPosts() {
            try {
                commit('setLoading', true)
                
                const response = await axios.get('https://jsonplaceholder.typicode.com/posts?_limit=10')
                
                commit('setPosts', response.data)
            } catch (e) {
                console.error(e)
            } finally {
                commit('setLoading', false)
            }
          },
        async fetchPost({commit}, id) {
            const response = await fetch(`https://jsonplaceholder.typicode.com/posts/${id}`)
            const post = await response.json()
            commit('setPost', post)
        }
    }
}