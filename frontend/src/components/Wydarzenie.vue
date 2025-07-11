<script setup>
    import {ref, onMounted, reactive, watch} from 'vue';
    import dayjs from 'dayjs'
    import 'dayjs/locale/pl'
    import { useDisplay } from 'vuetify/lib/composables/display';
    import { useRouter, useRoute } from 'vue-router';
    import { initQuill } from '@/quill/quill';

    const display = useDisplay();
    dayjs.locale('pl');
    const router = useRouter();
    const route = useRoute();
    let alias = route.params.alias;
    watch(() => route.fullPath, () => {
        window.location.reload()
    })

    let today = new Date();
    let month = ref(today.toLocaleString('pl-PL', {month: 'short'}));
    month.value = month.value.slice(0,1).toUpperCase() + month.value.slice(1); // month --> Month
    let year = ref(today.toLocaleString('pl-PL', {year: 'numeric'})); 
    let day = ref(today.toLocaleString('pl-PL', {day: 'numeric'}));
    
    const dataWydarzenia = reactive({
            Nazwa: '',
            Alias: '',
            Opis: '',
            DataStart: today,
            DataStop: today,
            Organizator: 'Akademia Łomżyńska',
            Lokalizacja: 'Akademicka',


    })
    const dateRange = ref([]);
    const formatDatePicker = (date) => {
        return `${dayjs(dateRange.value[0]).format('DD.MM.YYYY HH:mm')} - ${dayjs(dateRange.value[1]).format('DD.MM.YYYY HH:mm')}`
    }
    function resetFormData() {
        dataWydarzenia.Alias = '';
        dataWydarzenia.Opis = '';
        dataWydarzenia.DataStart = today;
        dataWydarzenia.DataStop = today;
        dataWydarzenia.Organizator = 'Akademia Łomżyńska';
        dataWydarzenia.Lokalizacja = 'Akademicka';
        dataWydarzenia.Aktywne = false;
        if (quill) quill.innerHTML = '';
    }
    const timeStart = ref(new Date)
    const timeStop = ref(new Date)
    
    watch(() => route.fullPath, () => {
        resetFormData();
    })

    async function fillForm(data) {
        dataWydarzenia.Alias = data.Alias || '';
        dataWydarzenia.Nazwa = data.Nazwa || '';
        dataWydarzenia.Opis = data.Opis || '';
        dataWydarzenia.Organizator = data.Organizator || '';
        dataWydarzenia.Lokalizacja = data.Lokalizacja || '';
        dataWydarzenia.DataStart = data.DataStart || '';
        dataWydarzenia.DataStop = data.DataStop || '';
        dataWydarzenia.Aktywne = data.Aktywne || false;
        dateRange.value[0] = dataWydarzenia.DataStart;
        dateRange.value[1] = dataWydarzenia.DataStop;
    }

    watch(dateRange, ()=>{
        dataWydarzenia.DataStart=dateRange.value[0].toISOString();
        dataWydarzenia.DataStop=dateRange.value[1].toISOString();
    })

    async function GetDataWydarzenia() {
        if (alias) {
            const response = await fetch(`/api/${alias}`);
            console.log(response);
            const data = await response.json();
            await fillForm(data);
            console.log(data);
        }
    }
    
    let quill = null;

    async function saveEvent() {
        if (route.path === '/kalendarium/dodaj') {
          await fetch(`/api/zapisz`, {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(dataWydarzenia)
                }) 
                router.replace('/kalendarium')
        } else {
            await fetch(`/api/${dataWydarzenia.Alias}`, {
                method: "PUT",
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(dataWydarzenia)
                }) 
                router.replace('/kalendarium')
            }
        }

    async function visibilytyEvent() {
        const response = await fetch(`/api/visibility-switcher/${dataWydarzenia.Alias}`, {
            method: 'PUT',
        });

        if (response.ok) {
            dataWydarzenia.Aktywne = !dataWydarzenia.Aktywne
        }
    }

    async function deleteEvent() {
        await fetch(`/api/${alias}`, {
            method: 'DELETE',
            headers: {
                'Content-type': 'application/json',
            },
        });
        await router.replace('/kalendarium');
    }

    onMounted(async () => {
        if (route.path == '/kalendarium/dodaj') { 
            resetFormData();
        }
        if (alias){
            await GetDataWydarzenia()
            today = dayjs(dataWydarzenia.DataStart)
            month.value = today.format('MMMM')
            month.value = month.value.slice(0,1).charAt(0).toUpperCase() + month.value.slice(1,3) // month --> Month
            year.value = today.format('YYYY')
            day.value = today.format('D')
            
            console.log(month.value)
            console.log(today)
        }
        quill = initQuill('#quill', dataWydarzenia.Opis);
        quill.on('text-change', () => {
            dataWydarzenia.Opis = quill.root.innerHTML;
            console.log(dataWydarzenia.Opis)
        });
        
    })
</script>



<template>
    <link rel="stylesheet" href="/src/assets/css/wydarzenieContent.css"/>
    <link rel="stylesheet" href="/src/assets/css/footer.css"/>
    <!--  -->
   

    <main class="content">
        <div class="buttons ms-20">
            <v-btn icon="$edit" v-if="display.smAndDown.value"
            class="vbuttons"
            rounded="xs"
            color="#000080"
            
            @click="saveEvent"
            variant="flat">
            </v-btn>
            <v-btn v-else class="vbuttons" 
            
            density="comfortable"
            rounded="xs"
            color="#000080"
            
            @click="saveEvent"
            variant="flat"
            >
            <v-icon icon="$edit" start/>
            Zapisz
            </v-btn>
            <v-btn :icon= "dataWydarzenia.Aktywne ? '$radioOn':'$radioOff'" v-if="display.smAndDown.value"
            class="vbuttons"

            rounded="xs"
            color="orange-darken-2"
            variant="flat"
            @click="visibilytyEvent"></v-btn>
            <v-btn v-else class="vbuttons" 
            prepend-icon=""
            density="comfortable"
            rounded="xs"
            color="orange-darken-2"
            variant="flat"
            @click="visibilytyEvent"
            >
            <v-icon :icon= "dataWydarzenia.Aktywne ? '$radioOn':'$radioOff'" start=""/>
            Widoczny</v-btn>

            <v-btn icon="$cancel" v-if="display.smAndDown.value"
            class="vbuttons"

            rounded="xs"
            color="red"
            @click="deleteEvent"
            :variant= "flat"></v-btn>
            <v-btn v-else class="vbuttons"
            density="comfortable"
            rounded="xs"
            color="red"
            @click="deleteEvent"
            :variant= "flat"
            >
            <v-icon icon="$cancel" start=""/>
            Usuń</v-btn>
        </div>
        <div id="alias">
            <span class="module-event-date">{{day}} {{month}} {{year}}</span>
            <input v-model="dataWydarzenia.Nazwa" id="inputAlias" placeholder="Tytuł" />
        </div>
        <div class=info>
            <p>Godzina: 
                <VueDatePicker v-model="dateRange"
                :format="formatDatePicker" 
                locale="pl" 
                range 
                class="date-picker">

            </VueDatePicker></p>
            
            <p>Lokalizacja: <input v-model="dataWydarzenia.Lokalizacja" class="infoSpan" contenteditable="true" spellcheck="false" @input="autoWidth"></input></p>
            <p>Organizator: <input v-model="dataWydarzenia.Organizator" class="infoSpan" contenteditable="true" spellcheck="false" @input="autoWidth"></input></p>
            <p id="opis-label">Opis wydarzenia: </p>
            <div id="quill"></div>

        </div>
    </main>
    
</template>

<style>

    .vbuttons {
        margin-right: 10px;
        margin-top: 30px;

    }
    .v-card-text, .v-card-subtitle{
        color: black;
    }
   
    .date-picker {
        width: 35%;
        min-width: 320px;
    }

    @media (max-width:960px) {
        .vbuttons {
            margin-bottom: 30px;
            margin-top: 0
        }
        .infoSpan {
            min-width: 200px;

        }
        .buttons {
            text-align: center;
        }
    }
</style>


