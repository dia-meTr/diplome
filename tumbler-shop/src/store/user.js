import { publicRequest } from '~/api/publicRequest'
import router from '~/routes'

export default {
  namespaced: true,
  state: () => ({
    user: {},
    isLogIn: false, // За допомогою цього значення перевірте, чи ввійшов поточний користувач.
    isAdmin: false  // За допомогою цього значення перевірте, чи є поточний користувач адміністратором.
  }),
  mutations: {
    // Оновлення статусу
    updateState(state, payload) {
      Object.keys(payload).forEach(key => {
        state[key] = payload[key]
      })
    }
  },
  actions: {
    // Зареєструватись
    async signUp({ commit }, payload = {}) { // Параметри отримуються у формі об’єкта. Якщо даних немає, виникає помилка, тому значення за замовчуванням вказується як `{}`.
      try {
        const res = await publicRequest({
          url: 'auth/signup',
          method: 'POST',
          body: {
            ...payload
          }
        })
        const { user, accessToken } = res
        if (!accessToken) throw new Error(res)

        window.localStorage.setItem('token', accessToken)
        commit('updateState', {
          user
        })
        alert('Реєстрацію завершено')
        router.push({
          name: 'mainpage'
        })
      } catch (err) {
        alert(err)
      }
    },
    // Увійти
    async signIn({ commit }, payload = {}) {
      try {
        const res = await publicRequest({
          url: 'auth/login',
          method: 'POST',
          body: {
            ...payload
          }
        })
        const { user, accessToken } = res
        if (!accessToken) throw new Error(res) 

        window.localStorage.setItem('token', accessToken) // Збережіть accessToKen у localStorage
        commit('updateState', { // Збережіть інформацію користувача для зберігання, збережіть за допомогою isLogIn: true
          user,
          isLogIn: true
        })
        router.push({
          name: 'mainpage'
        })
      } catch (err) {
        alert(err)
      }
    },
    // Вийити
    async signOut({ commit }) {
      await publicRequest({
        url: 'auth/logout',
        method: 'POST'
      })
      commit('updateState', {
        user: {},
        isLogIn: false,
        isAdmin: false
      })
      alert('로그아웃 되었습니다')
    },
    // Редагувати інформацію про користувача
    async editUserInfo({ commit }, payload = {}) {
      try {
        const res = await publicRequest({
          url: 'auth/user',
          method: 'PUT',
          body: {
            ...payload
          }
        })
        console.log(res)
        if (!res.displayName) throw new Error(res)

        commit('updateState', {
          ...res
        })
        alert('Інформація про користувача змінена')
        location.reload()
      } catch (err) {
        alert(err)
      }
    },
    // Надіслати запит на оновлення статусу
    requestUpdateState({ commit }, payload = {}) {
      commit('updateState', payload)
    }
  }
}
