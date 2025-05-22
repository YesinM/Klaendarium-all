<script setup>
    import {ref, onMounted, reactive, watch} from 'vue';
    import { createRouter, useRoute } from 'vue-router';
    import { initQuill } from '@/quill/quill';

    const route = useRoute();
    let wydarzenieID = route.query.id;

    const today = new Date();
    let month = today.toLocaleString('pl-PL', {month: 'short'});
    month = month.slice(0,1).toUpperCase() + month.slice(1); // month --> Month
    let year = today.toLocaleString('pl-PL', {year: 'numeric'}); 
    let day = today.toLocaleString('pl-PL', {day: 'numeric'});
    
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
    }
    
    watch(() => route.fullPath, () => {
        resetFormData();
    })

    onMounted(async function () {
        if (route.path == '/kalendarium/dodaj') resetFormData();
        if (wydarzenieID) {
            const response = await fetch(`/api/wydarzenie?id=${wydarzenieID}`);
            console.log(response);
            const data = await response.json();
            console.log(data);

            dataWydarzenia.Alias = data.Nazwa || '';
            dataWydarzenia.Opis = data.Opis || '';
            dataWydarzenia.Organizator = data.Organizator || '';
            dataWydarzenia.Lokalizacja = data.Lokalizacja || '';
            dataWydarzenia.DataStart = data.DataStart || '';
            dataWydarzenia.DataStop = data.DataStop || '';
        }
            const quill = initQuill('#quill', dataWydarzenia.Opis);
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
   

    <content>
        <buttons>
            <button id="b_zapisz">Zapisz</button>
            <button id="b_widoczny">Widoczny</button>
            <button id="b_usun">Usuń</button>
        </buttons>
        <alias>
            <span class="module-event-date">{{day}} {{month}} {{year}}</span>
            <input v-model="dataWydarzenia.Alias" id="inputAlias" placeholder="Tytuł" />
        </alias>
        <info>
            <p>Godzina: </p>
            <p>Lokalizacja: <input v-model="dataWydarzenia.Lokalizacja" class="infoSpan" contenteditable="true" spellcheck="false" @input="autoWidth"></input></p>
            <p>Organizator: <input v-model="dataWydarzenia.Organizator" class="infoSpan" contenteditable="true" spellcheck="false" @input="autoWidth"></input></p>
            <p>Opis wydarzenia: </p>
            <div id="quill"></div>

        </info>
    </content>
    
</template>



