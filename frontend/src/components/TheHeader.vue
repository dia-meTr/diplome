<template>
  <header
    id="header"
    class="flex-center-vertically">
    <div class="container">
      <div class="flex-center-vertically product-detail">
        <!--<img
          class="logo"
          src="../assets/logo.png"
          alt="logo"
          @click="$router.push({ name: 'mainpage'})" /> -->
        <div class="link-container">
          <template v-if="!isLogIn">
            <!--<RouterLink to="/signup">
              Реєстрація
            </RouterLink> -->
            <span class="menu-divider"> | </span>
            <RouterLink to="/signin">
              Увійти
            </RouterLink>
          </template>
          <template v-else>
            <template v-if="isAdmin">
              <RouterLink to="/admin/orders">
                AdminPage
              </RouterLink>
              <span class="menu-divider"> | </span>
            
              <RouterLink to="/mycart">
                Моя корзина
              </RouterLink>
              <span class="menu-divider"> | </span>
              
              <RouterLink to="/myorders">
                Мої замовлення
              </RouterLink>
              <span class="menu-divider"> | </span>
            </template>
            <RouterLink
              :to="{name: 'mainpage'}"
              @click="signOut">
              Вийти
            </RouterLink>
          </template>
        </div>
      </div>
    </div>
  </header>
</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
  data() {
    return {
      isLogIn: true,
      isAdmin: true
    }
  },
  //computed: {
  //  ...mapState('user', [
  //    'isLogIn',
  //    'isAdmin'
  //  ])
  //},
  methods: {
  //  ...mapActions('user', {
  //    useSignOut: 'signOut'
  //  }),
    signOut() {
      this.useSignOut()
    }
  }
}
</script>

<style lang="scss" scoped>
header {
  width: 100vw;
  height: 10vh;
  background-color: whitesmoke;
  border: 0.1rem solid whitesmoke;
  padding: 0.75em 0;
  /*margin-bottom: 1em;*/
  &.flex-center-vertically {
    justify-content: space-between;
  }
  &.sticky {
    position: fixed;
    top: 0
  }
  .flex-center-vertically {
    justify-content: space-between;
    display: flex;
    align-items: center;
  }
  a, .menu-divider {
    margin-right: 2rem;
    &:last-child {
      margin-right: 0;
    }
  }
  .container {
    width: 1200px;
    margin: 0 auto;
  }
  @media only screen and (min-width: 1540px) {
    .container {
      width: 1366px;
    }
  }
}
</style>
