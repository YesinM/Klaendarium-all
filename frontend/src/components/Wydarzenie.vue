<script setup>
    import {ref, onMounted, reactive, watch} from 'vue';
    import dayjs from 'dayjs'
    import 'dayjs/locale/pl'
    import { createRouter, useRoute } from 'vue-router';
    import { initQuill } from '@/quill/quill';

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
            //id:
            //nazwa: title,
            Alias: '',
            Opis: '',
            DataStart: today,
            DataStop: today,
            Organizator: 'Akademia Łomżyńska',
            Lokalizacja: 'Akademicka',


    })
    
    function resetFormData() {
        dataWydarzenia.Alias = '';
        dataWydarzenia.Opis = '';
        dataWydarzenia.DataStart = today;
        dataWydarzenia.DataStop = today;
        dataWydarzenia.Organizator = 'Akademia Łomżyńska';
        dataWydarzenia.Lokalizacja = 'Akademicka';
        if (quill) quill.innerHTML = '';
    }
    
    watch(() => route.fullPath, () => {
        resetFormData();
    })

    async function fillForm(data) {
        dataWydarzenia.Alias = data.Nazwa || '';
        dataWydarzenia.Opis = data.Opis || '';
        dataWydarzenia.Organizator = data.Organizator || '';
        dataWydarzenia.Lokalizacja = data.Lokalizacja || '';
        dataWydarzenia.DataStart = data.DataStart || '';
        dataWydarzenia.DataStop = data.DataStop || '';
    }

    async function GetDataWydarzenia() {
        if (alias) {
            const response = await fetch(`/api/${alias}`);
            console.log(response);
            const data = await response.json();
            await fillForm(data);
            console.log(data);
        }
    }
    
    function makeVisibleDatepicker(){

    }
    let quill = null;

    function saveEvent() {

    }

    function visibilytyEvent() {
        
    }

    function deleteEvent() {

    }

    onMounted(async () => {
        if (route.path == '/kalendarium/dodaj') { 
            resetFormData();
        }
        if (alias){
            await GetDataWydarzenia()
            today = new Date(dataWydarzenia.DataStart)
            month.value = today.toLocaleString('pl-PL', {month: 'short'});
            month.value = month.value.slice(0,1).toUpperCase() + month.value.slice(1); // month --> Month
            year.value = today.toLocaleString('pl-PL', {year: 'numeric'}); 
            day.value = today.toLocaleString('pl-PL', {day: 'numeric'});
            console.log(month.value)
            console.log(today)
        }
        quill = initQuill('#quill', dataWydarzenia.Opis);
        quill.on('text-change', () => {
            dataWydarzenia.Opis = quill.root.innerHTML;
        });
        
    })
</script>



<template>
    <link rel="stylesheet" href="/src/assets/css/bootstrap.min.css"/>
    <link rel="stylesheet" href="/src/assets/css/template.css"/>
    <link rel="stylesheet" href="/src/assets/css/kalendarium.css"/>
    <link rel="stylesheet" href="/src/assets/css/przyciski.css"/>
    <link rel="stylesheet" href="/src/assets/css/wydarzenieContent.css"/>
    <link rel="stylesheet" href="/src/assets/css/footer.css"/>
    <!--  -->
   

    <main class="content">
        <buttons class="ms-20">
            <v-btn class="vbuttons" 
            
            density="comfortable"
            rounded="xs"
            color="#000080"
            
            @click="saveEvent"
            variant="flat"
            >
            <v-icon icon="$edit" start/>
            Zapisz
            </v-btn>
            <v-btn class="vbuttons" 
            prepend-icon=""
            density="comfortable"
            rounded="xs"
            color="orange-darken-2"
            variant="flat"
            @click="visibilytyEvent"
            >
            <v-icon icon="$radioOff" start=""/>
            Widoczny</v-btn>
            <v-btn class="vbuttons"
            density="comfortable"
            rounded="xs"
            color="red"
            @click="deleteEvent"
            variant="flat"
            >
            <v-icon icon="$cancel" start=""/>
            Usuń</v-btn>
        </buttons>
        <div id="alias">
            <span class="module-event-date">{{day}} {{month}} {{year}}</span>
            <input v-model="dataWydarzenia.Alias" id="inputAlias" placeholder="Tytuł" />
        </div>
        <info>
            <p>Godzina: <v-btn variant="flat" @click="makeVisibleDatepicker"></v-btn></p>
            <p>Lokalizacja: <input v-model="dataWydarzenia.Lokalizacja" class="infoSpan" contenteditable="true" spellcheck="false" @input="autoWidth"></input></p>
            <p>Organizator: <input v-model="dataWydarzenia.Organizator" class="infoSpan" contenteditable="true" spellcheck="false" @input="autoWidth"></input></p>
            <p>Opis wydarzenia: </p>
            <div id="quill"></div>

        </info>
    </main>
    
</template>

<style>

    .vbuttons {
        margin-right: 10px;
        margin-top: 30px;

    }
</style>


