<script setup>
import { reactive, onMounted, ref, watch } from 'vue'
import dayjs from 'dayjs'
import 'dayjs/locale/pl'

 dayjs.locale('pl')

function localDate(rawDate){
    const day = dayjs(rawDate).format('DD');
    const month = dayjs(rawDate).format('MMM');
    const year = dayjs(rawDate).format('YYYY');

    return day + ' '+ month.charAt(0).toUpperCase() + month.slice(1) // maj ==> Maj
}

const dateList = ref([])
onMounted(async () => {
    const response = await fetch('/api/all');
    const data = await response.json()
    dateList.value = data
    console.log(dateList.value)
})

    
</script>

<template>
    
    
    <main class="content">
        <select>
            
        </select>
        <button>Dodaj filtr</button>
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


    