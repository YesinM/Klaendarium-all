import 'vuetify/styles'
import { createApp } from 'vue'
import { createVuetify } from "vuetify"
import { createPinia } from 'pinia'
import '@mdi/font/css/materialdesignicons.css'
import VueDatePicker from "@vuepic/vue-datepicker";
import '@vuepic/vue-datepicker/dist/main.css'
import App from './App.vue'
import router from './router'
import { setupCalendar, Calendar, DatePicker } from 'v-calendar';
import 'v-calendar/style.css';

const pinia = createPinia()
const vuetify = createVuetify()
const app = createApp(App);

app.use(router)
app.use(pinia)
app.use(vuetify)
app.use(setupCalendar, {})

app.component('VueDatePicker', VueDatePicker)
app.component('VCalendar', Calendar)
app.component('DatePicker', DatePicker)


app.mount('#app')
