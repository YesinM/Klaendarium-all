<script setup>

import { reactive, onMounted, ref, watch, defineComponent} from 'vue'

import dayjs from 'dayjs'
import 'dayjs/locale/pl'


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


const dateList = ref([])
onMounted(async () => {
    const response = await fetch('/api/all');
    const data = await response.json();
    dateList.value = data;  
}) 


    
</script>

<template>
    
    
    <main class="content">
        <input>
        <v-select 
            class="h-25 w-33"
            :items = getOptions()
            label="Rok"
            chips
            multiple
            clearable
            density="compact"
            variant="outlined"
        ></v-select>
        <ul>    
            <li v-for="date in dateList" :key='date.id' class="events-title">
                <span class="module-event-date">
                    {{localDate(date.DataStart)}}
                    &nbsp;
                </span> 
                <RouterLink :to="{path:'/kalendarium/wydarzenie', query: {id: date.ID}}"> {{date.Nazwa}}</RouterLink>
                    
                
            </li>
        </ul>
    </main>
    <div class="testKalendarz"></div>
    
</template>

<style>
    .testKalendarz {
        display: inline;
        position: sticky;
        height: 100px;
        width: 100px;
        border: 10px;
        background-color: aqua;
    }
</style>


    