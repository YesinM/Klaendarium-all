<script setup>

import { reactive, onMounted, ref, watch, defineComponent} from 'vue'
import dayjs from 'dayjs'
import 'dayjs/locale/pl'
import { computed } from 'vue';
import { useDisplay } from 'vuetify/lib/composables/display';


const display = useDisplay();

dayjs.locale('pl')

function getOptions(first = 2015, last = getCurrentYear()) {
    const options = [];
    for (first; first<=last+1; first++)
        options.push(first);
    return options
}

function localDate(rawDate){
    const day = dayjs(rawDate).format('D');
    const month = dayjs(rawDate).format('MMM');
    const year = dayjs(rawDate).format('YYYY')

    return day + ' '+ month.charAt(0).toUpperCase() + month.slice(1) + ' ' + year // maj ==> Maj
}

function getCurrentYear() {
  return new Date().getFullYear()
}



const currentPage = ref(1)

const dateList = ref([])
const numberOfPages = ref(0)

const chosedYears = ref([])
async function fetchItems(){
    numberOfPages.value = 0
    console.log('Żyję!')
    const params = new URLSearchParams();
    chosedYears.value.forEach(year => {
        params.append("year", year)
    });

    const response = await fetch(`/api/all?${params.toString()}`);
    const data = await response.json();
    dateList.value = await data;
    numberOfPages.value = await Math.ceil((dateList.value.length) / 12)
}

function isActive(active){
    return 
}

const paginatedList = computed (()=>{
    const s = (currentPage.value - 1) * 12;
    const e = s + 12;
    return dateList.value.slice(s, e);
})

watch(chosedYears, () => {
    currentPage.value = 1
    fetchItems();
    
})
console.log(display.smAndDown.value)
onMounted(()=>{
    fetchItems();
    console.log(display.smAndDown.value)
})


   
</script>

<template>
    <main class="content">
        <v-select 
            class = "h-33"
            style = "width: 150px; border-radius: 25px;"
            :items = getOptions()
            v-model = chosedYears
            label="Rok"
            chips
            multiple
            clearable
            density="compact"
            variant="outlined"
        ></v-select>
        <div class="eventList">
            <div v-for="date in paginatedList" :key='date.id'>
                <router-link :to="`/kalendarium/${date.Alias}`">
                <v-card :variant = "date.Aktywne ? 'flat' : 'plain'"
                    class="eventCard">
                    <v-card-title></v-card-title>
                    <v-card-subtitle
                    style="padding-top: 5px;"
                    >   
                        {{ localDate(date.DataStart) }}
                    </v-card-subtitle>
                    <v-card-text class="two-lines-ellipsis"
                    style="font-size: 1rem; font-weight: bold;">
                        {{ date.Nazwa }}
                    </v-card-text>
                    <v-card-actions>

                        <router-link :to="`/kalendarium/${date.Alias}`">
                        <v-btn class = "red-link" 
                            variant="flat"
                            style="text-transform: none;">
                            <v-icon icon="$next" start/>
                            Więcej
                        </v-btn>
                        </router-link>
                    </v-card-actions>
                </v-card>
                </router-link>  
                <!-- <div class="events-title" v-if="filteringByYear(date)"> filtrowanie za rokiem 
                    <span class="module-event-date">
                        {{localDate(date.DataStart)}}
                        &nbsp;
                    </span> 
                    
                </div>          -->
            
            </div>
        </div>
    </main>
        <v-pagination class="pagination"
            v-if="numberOfPages > 1"
            v-model = currentPage
            :length = numberOfPages
            :total-visible="display.smAndDown.value ? 3 : 5"  
            rounded="circle"
            size="small"
        >
        </v-pagination>
</template>

<style>
    .two-lines-ellipsis {
        display: -webkit-box;
        -webkit-line-clamp: 2;
        -webkit-box-orient: vertical;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: normal;
        line-height: 2rem;
        max-height: 3.9rem;  

    }

 /* .v-pagination .v-btn--disabled {
        margin: 0 2px !important; 
        padding: 4px 8px !important; 
        min-width: 32px !important; 
    } */
    .v-card{
        flex: 1 1;
        height: 100%
    }
    .eventList {
        display: flex;
        flex-flow: row wrap;
        align-items: stretch;
    }
    .eventList>div{
        margin: 5px;
    }
    li {
        text-decoration: none;
    }
    .testKalendarz {
        display: inline;
        position: sticky;
        height: 100px;
        width: 100px;
        border: 10px;
        background-color: aqua;
    }
    .module-event-date {
        margin-bottom: 10px;
    }
    .pagination {
        max-width: 1000px
    }
    @media (max-width:680px) {
    .eventList{
        display: block;    
    }
    }
    
</style>


    