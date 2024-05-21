export default {
  namespaced: true,
  state: () => ({
    myPageMenuList: [
      { name: 'Історія покупок', isShow: true },
      { name: 'Мій рахунок', isShow: false},
      { name: 'Редагувати мою інформацію', isShow: false}
    ],
    adminPageMenuList: [
      { name: 'Запит продукту', isShow: true },
      { name: 'Історія продажів', isShow: false }
    ],
    loading: false
  }),
  mutations: {
    updateMenu(state, menuName) {
      const whatMenu = 
      state.myPageMenuList.some(menu => menu.name === menuName)
      ? state.myPageMenuList 
      : state.adminPageMenuList
      
      whatMenu.forEach(menu => {
        menu.isShow = false
        if (menu.name === menuName) {
          menu.isShow = true
        }
      })
    },
    updateLoading(state, boolean) {
      state.loading = boolean
    }
  },
  actions: {
		isShowMenu({ commit }, menuName) {
      window.sessionStorage.setItem('menu', menuName)
			commit('updateMenu', menuName)
    },
    isShowLoading({ commit }, boolean) {
      commit('updateLoading', boolean)
    }
  }
}












