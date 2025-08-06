import 'vuetify/styles'
import { createApp } from 'vue'
import { createVuetify } from "vuetify"
import { createPinia } from 'pinia'
import '@mdi/font/css/materialdesignicons.css'
import VueDatePicker from "@vuepic/vue-datepicker";
import '@vuepic/vue-datepicker/dist/main.css'
import App from './App.vue'
import router from './router'

const pinia = createPinia()


const vuetify = createVuetify()

const app = createApp(App);

app.component('VueDatePicker', VueDatePicker)

app.use(router).use(pinia).use(vuetify).mount('#app')
