<script setup>
    import {ref, onMounted, reactive, watch} from 'vue';
    import { createRouter, useRoute } from 'vue-router';
    import { initQuill } from '@/quill/quill';

    const route = useRoute();
    let wydarzenieID = route.query.id;
    
    const dataWydarzenia = reactive({
            //id:
            //nazwa: title,
            Alias: '',
            Opis: '',
            DataStart: '',
            DataStop: '',
            Organizator: 'Akademia Łomżyńska',
            Lokalizacja: 'Akademicka',


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
        if (wydarzenieID) {
            const response = await fetch(`/api/wydarzenie?id=${wydarzenieID}`);
            console.log(response);
            const data = await response.json();
            await fillForm(data);
            console.log(data);
        }
    }
    
    onMounted(async () => {
        if (wydarzenieID){
            await GetDataWydarzenia()
        }  
    })
</script>



<template>
    <link rel="stylesheet" href="/src/assets/css/bootstrap.min.css"/>
    <link rel="stylesheet" href="/src/assets/css/template.css"/>
    <link rel="stylesheet" href="/src/assets/css/kalendarium.css"/>
    <link rel="stylesheet" href="/src/assets/css/przyciski.css"/>
    <link rel="stylesheet" href="/src/assets/css/wydarzenieContent.css"/>
    <link rel="stylesheet" href="/src/assets/css/footer.css"/>
   
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
            <div v-html="dataWydarzenia.Opis"></div>

        </info>
    </content>
    
</template>



