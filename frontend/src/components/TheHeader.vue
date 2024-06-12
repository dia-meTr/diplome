<template>
    <header class="header">
        <!--<img
          class="logo"
          src="../assets/logo.png"
          alt="logo"
          @click="$router.push({ name: 'mainpage'})" /> -->
          <template v-if="!isLogIn">
            <!--<RouterLink to="/signup">
              Реєстрація
            </RouterLink> -->
            <RouterLink to="/signin"  class="headerLink">
              Увійти
            </RouterLink>
          </template>
          <template v-else>
            <div class="headerLinksContainer">
              <template v-if="isAdmin">
                <RouterLink to="/admin/orders" class="headerLink">
                  Керування замовленнями
                </RouterLink>
                <RouterLink to="/admin/dishes" class="headerLink">
                  Керування стравами
                </RouterLink>
                <RouterLink to="/admin/tags" class="headerLink">
                  Керування тегами
                </RouterLink>
                <RouterLink to="/admin/products" class="headerLink">
                  Керування продуктами
                </RouterLink>
              </template>

                <RouterLink to="/mycart" class="headerLink">
                  Моя корзина
                </RouterLink>
    
                <RouterLink to="/myorders" class="headerLink">
                  Мої замовлення
                </RouterLink>
              
            </div>
            <RouterLink
              class="headerLink"
              :to="{name: 'mainpage'}"
              @click="signOut">
              Вийти
            </RouterLink>
          </template>
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

<style scoped>

.header {
    width: 100%;
    background-color: white;
    padding: 30px 100px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    display: flex;
  }
  
  .headerLinksContainer {
    flex-grow: 1;
    display: flex;
    gap: 40px;
  }
  
  .headerLink {
    color: black;
    text-decoration: none;
    transition: color 0.3s;
  }
  
  .headerLink:hover {
    color: #0B770F;
   }
</style>


