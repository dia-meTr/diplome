<template>
  <TheHeader />
  <div class="auth-form-container flex-center-center">
    <div class="inner">
      <h1 class="text-center">
        Реєстрація
      </h1>
      <div class="inner__card">
        <h2 class="text-center">
          Будь ласка, введіть інформацію про учасника
        </h2>
        <div class="form">
          <input
            :value="displayName"
            placeholder="Введіть своє ім'я, будь ласка"
            type="text" 
            autofocus 
            @input="displayName = $event.target.value" />
          <input
            v-model="email"
            placeholder="Будь ласка, введіть свою електронну пошту"
            type="text" />
          <input
            v-model="password"
            placeholder="Будь ласка, введіть свій пароль (8 символів або більше, з урахуванням регістру)"
            type="password"
            @keydown.enter="useSignUp" />
          <input
            type="submit"
            value="Зареєструватись"
            @click="useSignUp" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'
import TheHeader from '~/components/TheHeader'

export default {
  components: {
    TheHeader
  },
  data() {
    return {
      displayName: '',
      email: '',
      password: ''
    }
  },
  computed: {
    ...mapState('user', [
      'user'
    ])
  },
  methods: {
    ...mapActions('user', [
      'signUp'
    ]),
    async useSignUp() {
      if (!this.displayName.trim()) {
        return alert("Введіть своє ім'я, будь ласка")
      } else if (!this.email.trim()) {
        return alert('Будь ласка, введіть свою електронну пошту')
      } else if (!this.password.trim()) {
        return alert('Будь ласка, введіть пароль')
      } else if (!this.email.includes('@') && !this.email.includes('.')) {
        return alert('Неправильний формат електронної пошти')
      } else if (this.password.trim().length <= 7) {
        return alert('Будь ласка, встановіть пароль щонайменше з 8 символів.')
      }
      this.signUp({
        displayName: this.displayName,
        email: this.email,
        password: this.password
      })
      this.displayName = ''
      this.email = ''
      this.password = ''
    }
  }
}

</script>

<style lang="scss" scoped>
.auth-form-container {
  background: url("../assets/login_bg.jpg") no-repeat center center fixed;
  background-size: cover;
  height: 90vh;
  position: absolute;
  top: 10vh;
  width: 100%;
  .inner {
    h1 {
      color: $color-white;
    }
    &__card {
      width: 50rem;
      margin: 0 auto;
      border-radius: 0.6rem;
      background-color: $color-white;
      box-shadow: 0.2rem 0.2rem 2rem rgba(0, 0, 0, .3);
      color: $color-font;
      h2 {
        padding: 3rem;
        border-bottom: 0.1rem solid $color-header-border;
        strong {
          font-weight: $font-medium;
          color: $color-primary;
        }
      }
      .form {
        padding: 3rem 2.2rem;
        [type="submit"] {
          background-color: $color-primary;
          border: none;
          color: $color-white;
          font-size: 2rem;
          cursor: pointer;
        }
        input {
          width: 100%;
          margin-bottom: 1.2rem;
          padding: 1.5rem;
          border: 0.1rem solid $color-header-border;
          border-radius: 0.4rem;
          outline: none;
          font-size: 1.4rem;
          &:not(:last-of-type) {
            &:focus {
              border-color: $color-primary;
              background-color: rgba($color-primary, .03);
            }
          }
        }
      }
    }
  }
}
</style>
