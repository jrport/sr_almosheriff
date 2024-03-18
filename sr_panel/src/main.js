import { createApp } from 'vue'
import './style.css'
import router from './router'
import Chart from "vue-frappe-chart"
import App from './App.vue'

const app = createApp(App)
app.use(router)
app.use(Chart)
app.mount("#app")
