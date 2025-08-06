import "vuetify/styles"
import { createVuetify } from "vuetify"
import { VBtn, VSelect, VCard, VPagination, VFooter} from "vuetify/components"
import * as directives from "vuetify/directives"
import "@mdi/font/css/materialdesignicons.css"

const vuetify = createVuetify({
    components:{
        VBtn,
        VSelect,
        VCard,
        VPagination,
        VFooter,
    }, 
    directives,
    icons: {
        defaultSet: 'mdi'
    }
})

export default vuetify