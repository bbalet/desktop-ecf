import './assets/main.css'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import { router } from './router'

const baseApiUrl = 'https://cinephoria.cloud'
export { baseApiUrl }

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')
