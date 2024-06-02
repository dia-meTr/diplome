import { createStore } from 'vuex'
import { postModule } from './postModule'

export default createStore({
  state: {
    likes: 1,
    isAuth: false
  },
  getters: {
  },
  mutations: {
    incrementLikes(state) {
      state.likes++
    },
    decrementLikes(state) {
      state.likes--
    }
  },
  actions: {
  },
  modules: {
    post: postModule
  }
})
