<script setup>
    import {ref, onMounted, reactive, watch, nextTick} from 'vue';
    import dayjs from 'dayjs'
    import utc from 'dayjs/plugin/utc'
    import 'dayjs/locale/pl'
    import { useDisplay } from 'vuetify/lib/composables/display';
    import { useRouter, useRoute } from 'vue-router';
    import { initQuill } from '@/quill/quill';

    const display = useDisplay();
    dayjs.extend(utc)
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
            const response = await fetch(`/api/${alias}`);
            console.log(response);
            const data = await response.json();
            await fillForm(data);
            console.log(data);
    }
    
    let quill = null;

    async function saveEvent() {
        const url = route.path === '/kalendarium/dodaj' ?
        '/api/zapisz' :
        `/api/${dataWydarzenia.Alias}`;

        const method = route.path === '/kalendarium/dodaj' ?
        'POST' :
        'PUT';
        await fetch(url, {
            method: method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(dataWydarzenia)
            }) 
            router.replace('/kalendarium')
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

    function toGoogleDateFormat(dateString) {
        return dayjs(dateString).utc().format('YYYYMMDDTHHmmss[Z]');
    }


    function addToGoogleCalendar(){
        let start = toGoogleDateFormat(dataWydarzenia.DataStart)
        let stop = toGoogleDateFormat(dataWydarzenia.DataStop)

        const url =`https://calendar.google.com/calendar/render?action=TEMPLATE` +
        `&text=${encodeURIComponent(dataWydarzenia.Nazwa)}` +
        `&dates=${start}/${stop}` +
        `&location=${encodeURIComponent(dataWydarzenia.Lokalizacja || "")}`
        window.open(url, "_blank")
    }

    function backButton() {
        router.push('/kalendarium')
    }
    const adminStatus = ref(false);
    async function isAdmin() {
        const res = await fetch('/api/isAdmin');
        const data = await res.json();
        await console.log("AdminStatus " + data.admin)
        return data.admin
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

        adminStatus.value = await isAdmin();
        console.log("adminStatus ref: " + adminStatus.value)
        
        if (adminStatus.value) {
            await nextTick()
            quill = initQuill('#quill', dataWydarzenia.Opis);
            quill.on('text-change', () => {
                dataWydarzenia.Opis = quill.root.innerHTML;
                console.log(dataWydarzenia.Opis)
           });
        
        }
    })
</script>



<template>
    <link rel="stylesheet" href="/src/assets/css/wydarzenieContent.css"/>
    <link rel="stylesheet" href="/src/assets/css/footer.css"/>
    <!--  -->
   

    <main class="content">
        <div  class="buttons">
            <div v-if="adminStatus" class="editButtons">
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
            <div class="additionalButtons">
                <v-btn icon="$calendar" v-if="display.smAndDown.value"
                class="vbuttons"
                rounded="xs"
                color="green"
                @click="addToGoogleCalendar"
                :variant= "flat"></v-btn>
                <v-btn v-else class="vbuttons"
                density="comfortable"
                rounded="xs"
                color="green"
                @click="addToGoogleCalendar"
                :variant= "flat"
                >
                <v-icon icon="$calendar" start=""/>
                Dodaj do Google Kalendarz</v-btn>

                <v-btn icon="$close" 
                class="vbuttons"
                rounded="xs"
                color="grey-lighten-2"
                @click="backButton"
                density="comfortable"
                :variant= "flat"></v-btn>
            </div>
        </div>

        <div id="alias">
            <input v-model="dataWydarzenia.Nazwa" id="inputAlias" placeholder="Tytuł" :readOnly="!adminStatus"/>
        </div>
        <div class=info>
            <span style="display:inline-flex; align-items: center;"> 
                <VueDatePicker v-model="dateRange"
                :disabled="!adminStatus"
                :format="formatDatePicker" 
                locale="pl" 
                range 
                class="date-picker">
                
            </VueDatePicker></span>
            
            <p>Lokalizacja: <input :readOnly="!adminStatus" v-model="dataWydarzenia.Lokalizacja" class="infoSpan" contenteditable="true" spellcheck="false" @input="autoWidth"></input></p>
            <p>Organizator: <input :readOnly="!adminStatus" v-model="dataWydarzenia.Organizator" class="infoSpan" contenteditable="true" spellcheck="false" @input="autoWidth"></input></p>
            <p id="opis-label">Opis wydarzenia: </p>
            <div id="quill" v-if="adminStatus"></div>
            <div id="user" v-if="!adminStatus" v-html="dataWydarzenia.Opis"></div>

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


