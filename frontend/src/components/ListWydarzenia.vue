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
    const day = dayjs(rawDate).format('D');
    const month = dayjs(rawDate).format('MMM');
    const year = dayjs(rawDate).format('YYYY')

    return day + ' '+ month.charAt(0).toUpperCase() + month.slice(1) + ' ' + year // maj ==> Maj
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
const dateList = ref([])

const chosedYears = ref([])
async function fetchItems(){
    console.log('Żyję!')
    const params = new URLSearchParams();
    chosedYears.value.forEach(year => {
        params.append("year", year)
    });

    const response = await fetch(`/api/all?${params.toString()}`);
    const data = await response.json();
    dateList.value = data;
    console.log(data);

}

watch(chosedYears, () => {
    fetchItems();
})

onMounted(()=>{
    fetchItems();
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
            <div v-for="date in dateList" :key='date.id'>
                
                <v-card 
                    class="eventCard"
                    :title= 'dayjs(date.DataStart).format("HH:mm")'
                    :subtitle="localDate(date.DataStart)"
                    :text="date.Nazwa"
                    variant="flat">
                    <v-card-actions>
                        <router-link :to="`/kalendarium/${date.Alias}`">
                        <v-btn>
                            <v-icon icon="$next" start/>
                            Zobacz więcej
                        </v-btn>
                        </router-link>
                    </v-card-actions>
                </v-card>
                <!-- <div class="events-title" v-if="filteringByYear(date)"> filtrowanie za rokiem 
                    <span class="module-event-date">
                        {{localDate(date.DataStart)}}
                        &nbsp;
                    </span> 
                    
                </div>          -->
            
            </div>
        </div>
    </main>
    <v-pagination></v-pagination>
</template>

<style>

    .eventList {
        display: flex;
        flex-flow: row wrap;
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
    
</style>


    