import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import components from '@/components/UI'

const app = createApp(App)

app.config.globalProperties.$hostname = 'http://0.0.0.0:8080'

components.forEach(component => {
    app.component(component.name, component)
})

app.use(store).use(router).mount('#app')
