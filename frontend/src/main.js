import 'vuetify/styles'
import { createApp } from 'vue'
import { createVuetify } from "vuetify"
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import '@mdi/font/css/materialdesignicons.css'
// import VueDatePicker from "@vuepic/vue-datepicker";
// import '@vuepic/vue-datepicker/dist/main.css'
import App from './App.vue'
import router from './router'



const vuetify = createVuetify({
    components, 
    directives,
    icons: {
        defaultSet: 'mdi'
    }
})

let app = createApp(App);

// App.component('VueDatePicker', VueDatePicker)

app.use(router).use(vuetify).mount('#app')
