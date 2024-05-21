<template>
  <div class="inner">
    <h1>Редагувати мою інформацію</h1>
    <div class="inner__card">
      <div class="form">
        <input
          :value="displayName"
          placeholder="Будь ласка, введіть нове ім'я"
          type="text" 
          autofocus 
          @input="displayName = $event.target.value" />
        <input
          v-model="oldPassword"
          placeholder="Будь ласка, введіть існуючий пароль"
          type="password" />
        <input
          v-model="newPassword"
          placeholder="Будь ласка, введіть новий пароль"
          type="password"
          @keydown.enter="useEditUserInfo" />
        <input
          type="submit"
          value="Редагувати мою інформацію"
          @click="useEditUserInfo" />
      </div>
    </div>
  </div>  
</template>

<script>
import { mapActions } from 'vuex'

export default {
  data() {
    return {
      displayName: '',
      oldPassword: '',
      newPassword: ''
    }
  },
  methods: {
    ...mapActions('user', [ 
      'editUserInfo' 
    ]),
    async useEditUserInfo() {
      if (!this.displayName.trim()) {
        return alert('Будь ласка, введіть назву для зміни')
      } else if (!this.oldPassword.trim()) {
        return alert('Будь ласка, введіть існуючий пароль')
      } else if (!this.newPassword.trim()) {
        return alert('Будь ласка, введіть новий пароль')
      } else if (this.newPassword.trim().length <= 7) {
        return alert('Будь ласка, встановіть пароль щонайменше з 8 символів.')
      }
      this.editUserInfo({
        displayName: this.displayName,
        oldPassword: this.oldPassword,
        newPassword: this.newPassword
      })
      this.displayName = ''
      this.oldPassword = ''
      this.newPassword = ''
    }
  }
}
</script>

<style lang="scss" scoped>
.inner {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
    h1 {
      color: $color-font;
      font-size: 3rem;
      font-weight: 700;
      text-align: center;
      margin-bottom: 2.4rem;
    }
    &__card {
      width: 50rem;
      margin: 0 auto;
      border-radius: 0.6rem;
      background-color: $color-white;
      border: 0.1rem solid $color-header-border;
      color: $color-font;
      h2 {
        padding: 3rem;
        font-size: 1.8rem;
        text-align: center;
        border-bottom: 0.1rem solid $color-header-border;
        strong {
          font-weight: 700;
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
        p {
          font-size: 1.2rem;
          text-align: center;
          a {
            color: $color-primary;
            font-weight: bold;
            text-decoration: underline;
          }
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
</style>
