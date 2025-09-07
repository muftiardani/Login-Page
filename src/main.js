import { createApp } from 'vue'
import { createPinia } from 'pinia'
import router from './router' // Impor router
import App from './App.vue'
import './style.css'

const app = createApp(App)
const pinia = createPinia()

app.use(pinia) // Gunakan Pinia untuk state management
app.use(router) // Gunakan Vue Router untuk navigasi
app.mount('#app')