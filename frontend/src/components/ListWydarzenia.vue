<script setup>

import { reactive, onMounted, ref, watch, defineComponent} from 'vue'
import dayjs from 'dayjs'
import 'dayjs/locale/pl'
import { computed } from 'vue';

dayjs.locale('pl')

function getOptions(first = 2015, last = getCurrentYear()) {
    const options = [];
    for (first; first<=last+1; first++)
        options.push(first);
    return options
}

function localDate(rawDate){
    const day = dayjs(rawDate).format('DD');
    const month = dayjs(rawDate).format('MMM');

    return day + ' '+ month.charAt(0).toUpperCase() + month.slice(1) // maj ==> Maj
}

function getCurrentYear() {
  return new Date().getFullYear()
}

function filteringByYear(date) {
    if (chosedYears.value.length === 0) {
        console.log(chosedYears.length)
        return true;
    }
    return chosedYears.value.includes(+date.year)
}


const chosedYears = ref([])

const dateList = ref([])
onMounted(async () => {
    const response = await fetch('/api/all');
    const data = await response.json();
    dateList.value = data; 

    for (let wydarzenie of dateList.value) {
        wydarzenie.year = dayjs(wydarzenie.DataStart).format('YYYY')
    }
     
}) 


    
</script>

<template>
    
    
    <main class="content">
        <v-select 
            class = "h-33"
            style = "width: 130px; border-radius: 25px;"
            :items = getOptions()
            v-model = chosedYears
            label="Rok"
            chips
            multiple
            clearable
            density="compact"
            variant="outlined"
        ></v-select>
            <div v-for="date in dateList" :key='date.id'  >
                
                <div class="events-title" v-if="filteringByYear(date)"> <!--filtrowanie za rokiem --> 
                    <span class="module-event-date">
                        {{localDate(date.DataStart)}}
                        &nbsp;
                    </span> 
                    <RouterLink :to="{path:'/kalendarium/wydarzenie', query: {id: date.ID}}"> {{date.Nazwa}}</RouterLink>
                </div>         
            
            </div>
    </main>
    <div class="testKalendarz"></div>
    
</template>

<style>
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
</style>


    