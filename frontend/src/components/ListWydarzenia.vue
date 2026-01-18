<script setup>
import { reactive, onMounted, ref, watch, defineComponent } from 'vue'
import dayjs from 'dayjs'
import 'dayjs/locale/pl'
import { computed } from 'vue';
import { useDisplay } from 'vuetify/lib/composables/display';
import { useAuthStore } from '@/stores/userStore';
import { useCalendarFiltersStore } from '@/stores/calendarStore';

const auth = useAuthStore()
const filtersStore = useCalendarFiltersStore() 

const display = useDisplay();
const currentPage = ref(1)
const dateList = ref([])
const numberOfPages = ref(0)

const searchQuery = computed(() => filtersStore.searchQuery)
const chosedYears = computed(() => filtersStore.chosedYears)

watch(searchQuery, (newVal) => {
  filtersStore.setSearchQuery(newVal)
  currentPage.value = 1
  fetchItems()
})

watch(chosedYears, (newVal) => {
  filtersStore.setChosedYears(newVal)
  currentPage.value = 1
  fetchItems()
})

dayjs.locale('pl')

async function fetchItems() {
  numberOfPages.value = 0;

  const params = new URLSearchParams();

  if (filtersStore.searchQuery && filtersStore.searchQuery.length > 0) {
    params.append("title", filtersStore.searchQuery);
  }

  if (filtersStore.chosedYears && Array.isArray(filtersStore.chosedYears)) {
    filtersStore.chosedYears.forEach(year => params.append("year", year));
  }

  try {
    const response = await fetch(`/api/all?${params.toString()}`);
    if (!response.ok) throw new Error("Request error");

    const data = await response.json();
    dateList.value = data;

    numberOfPages.value = Math.ceil(data.length / 12);
  } catch (err) {
    console.error(err);
  }
}

function getOptions(first = 2015, last = getCurrentYear()) {
  const options = [];
  for (let i = first; i <= last + 1; i++) {
    options.push(i);
  }
  return options
}

function localDate(rawDate) {
  const day = dayjs(rawDate).format('D');
  const month = dayjs(rawDate).format('MMM');
  const year = dayjs(rawDate).format('YYYY')

  return day + ' ' + month.charAt(0).toUpperCase() + month.slice(1) + ' ' + year // maj ==> Maj
}

function getCurrentYear() {
  return new Date().getFullYear()
}

const filterActive = computed(() => {
  return filtersStore.searchQuery || (filtersStore.chosedYears && filtersStore.chosedYears.length > 0)
})

const filteredList = computed(() => {
  return dateList.value.filter(ev => {
    if (auth.isAdmin) return true          
    if (filterActive.value) return true

    return ev.Aktywne   
  })
})

const paginatedList = computed(() => {
  const start = (currentPage.value - 1) * 12
  return filteredList.value.slice(start, start + 12)
})

function onClear() {
  filtersStore.setSearchQuery("")
  filtersStore.setChosedYears([]) 
  currentPage.value = 1
  fetchItems()
}

function fetchOnEnter(event) {
  if (event.key === 'Enter') {
    fetchItems()
  }
}

onMounted(async () => {
  await fetchItems();
})
</script>

<template>
  <main class="content">
    <div id="findHeader">
      <v-select 
        class="h-33"
        style="width: 150px; border-radius: 25px; flex: 2"
        :items="getOptions()"
        v-model="filtersStore.chosedYears" 
        label="Rok"
        chips
        multiple
        clearable
        density="compact"
        variant="outlined"
        @update:model-value="fetchItems" 
      ></v-select>
      
      <v-text-field
        v-model="filtersStore.searchQuery"
        label="Szukaj wydarzeń"
        variant="outlined"
        density="compact"
        append-inner-icon="mdi-magnify"
        style="flex: 5;"
        @click:append-inner="fetchItems"
        @keyup.enter="fetchItems"
        clearable
        @click:clear="onClear"
        @keyup="fetchOnEnter"
      >
      </v-text-field>
    </div>
    
    <div class="eventList">
      <div v-for="date in paginatedList" :key='date.id'>
        <div> 
          <router-link :to="`/kalendarium/${date.Alias}`">
            <v-card
              :variant="date.Aktywne ? 'flat' : 'plain'"
              class="eventCard"
            >
              <v-card-title></v-card-title>
              <v-card-subtitle
                style="padding-top: 5px;"
              >   
                {{ localDate(date.DataStart) }}
              </v-card-subtitle>
              <v-card-text class="two-lines-ellipsis"
                style="font-size: 1rem; font-weight: bold;"
              >
                {{ date.Nazwa }}
              </v-card-text>
              <v-card-actions>
                <router-link :to="`/kalendarium/${date.Alias}`">
                  <v-btn 
                    class="red-link" 
                    variant="flat"
                    style="text-transform: none;"
                  >
                    <v-icon icon="$next" start/>
                    Więcej
                  </v-btn>
                </router-link>
              </v-card-actions>
            </v-card>
          </router-link>  
        </div>
      </div>
    </div>
  </main>
  
  <v-pagination 
    class="pagination"
    v-if="numberOfPages > 1"
    v-model="currentPage"
    :length="numberOfPages"
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

.v-card{
  flex: 1 1;
  height: 100%
}

.eventList {
  display: flex;
  flex-flow: row wrap;
  align-items: flex-start;
  align-content: flex-start;
  min-height: 640px;
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