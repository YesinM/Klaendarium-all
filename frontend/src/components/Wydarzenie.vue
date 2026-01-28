<script setup>
    import {ref, onMounted, reactive, watch, nextTick} from 'vue';
    import dayjs from 'dayjs'
    import utc from 'dayjs/plugin/utc'
    import 'dayjs/locale/pl'
    import { useDisplay } from 'vuetify/lib/composables/display';
    import { useRouter, useRoute } from 'vue-router';
    import { initQuill } from '@/quill/quill';
    import 'quill/dist/quill.snow.css'
    import 'quill-table-better/dist/quill-table-better.css'
    import '../assets/css/quill_style.css'
    import { useAuthStore } from '@/stores/userStore';
    
    const auth = useAuthStore();
    const display = useDisplay();
    dayjs.extend(utc);
    dayjs.locale('pl');
    const router = useRouter();
    const route = useRoute();
    let alias = route.params.alias;
    
    const isLoading = ref(false)
    const eventNotFound = ref(false)
    
    const showAlert = ref(false)
    const alertTitle = ref('')
    const alertMessage = ref('')

    let today = new Date();
    let month = ref(today.toLocaleString('pl-PL', {month: 'short'}));
    month.value = month.value.slice(0,1).toUpperCase() + month.value.slice(1);
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
        dataWydarzenia.Nazwa = '';
        dataWydarzenia.Alias = '';
        dataWydarzenia.Opis = '';
        dataWydarzenia.DataStart = today;
        dataWydarzenia.DataStop = today;
        dataWydarzenia.Organizator = 'Akademia Łomżyńska';
        dataWydarzenia.Lokalizacja = 'Akademicka';
        dataWydarzenia.Aktywne = false;
        dateRange.value = [];
        if (quill) {
            quill.root.innerHTML = ''
        }
    }
    
    function showCustomAlert(title, message) {
        alertTitle.value = title
        alertMessage.value = message
        showAlert.value = true
    }
    
    watch(() => route.fullPath, () => {
        resetFormData();
    })

    watch(() => route.params.alias, async (newAlias) => {
    if (newAlias) {
        alias = newAlias
        isLoading.value = true
        eventNotFound.value = false
        
        resetFormData()
        
        await GetDataWydarzenia()
        
        await nextTick()
        if (quill) {
            quill.disable()
            quill = null
        }

        quill = initQuill('#quill', dataWydarzenia.Opis, auth.isAdmin)
        quill.on('text-change', () => {
            dataWydarzenia.Opis = quill.root.innerHTML
        })
    }
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
        if (dateRange.value && dateRange.value[0] && dateRange.value[1]) {
            dataWydarzenia.DataStart = dateRange.value[0].toISOString();
            dataWydarzenia.DataStop = dateRange.value[1].toISOString();
        }
    })

    async function GetDataWydarzenia() {
        isLoading.value = true
        eventNotFound.value = false
        
        try {
            const response = await fetch(`/api/${alias}`);
            
            if (response.status === 404) {
                eventNotFound.value = true
                return
            }
            
            if (!response.ok) throw new Error("Request error");
            
            const data = await response.json();
            await fillForm(data);
            
            if (dataWydarzenia.DataStart) {
                today = dayjs(dataWydarzenia.DataStart)
                month.value = today.format('MMMM')
                month.value = month.value.slice(0,1).charAt(0).toUpperCase() + month.value.slice(1,3)
                year.value = today.format('YYYY')
                day.value = today.format('D')
            }
            
        } catch (err) {
            console.error(err);
            eventNotFound.value = true
        } finally {
            isLoading.value = false
        }
    }
    
    let quill = null;

    async function saveEvent() {
        if (!dataWydarzenia.Nazwa || dataWydarzenia.Nazwa.trim() === '') {
            showCustomAlert('Błąd walidacji', 'Pole "Tytuł" jest wymagane!')
            return;
        }
        
        if (quill) {
            quill.blur();
            const activeCells = document.querySelectorAll('.ql-table-cell-active, .ql-cell-focused');
            activeCells.forEach(cell => {
                cell.classList.remove('ql-table-cell-active', 'ql-cell-focused');
                cell.blur();
            });
        }

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
        if (!confirm('Czy na pewno chcesz usunąć to wydarzenie?')) return;
        
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

    onMounted(async () => {
        await nextTick()
        
        if (alias){
            await GetDataWydarzenia()

            quill = initQuill('#quill', dataWydarzenia.Opis, auth.isAdmin);
            quill.on('text-change', () => {
            dataWydarzenia.Opis = quill.root.innerHTML;});

            
            const activeCells = document.querySelectorAll('.ql-cell-focused');
            activeCells.forEach(cell => {
                cell.classList.remove('ql-cell-focused');
                cell.blur();
            });
        } else {
            quill = initQuill('#quill', dataWydarzenia.Opis, auth.isAdmin);
            quill.on('text-change', () => {
            dataWydarzenia.Opis = quill.root.innerHTML;
            });
        }
    })
</script>

<template>
    <div v-if="eventNotFound && alias">
        <v-container class="fill-height">
            <v-responsive class="align-center text-center fill-height">
                <v-icon 
                    size="120" 
                    color="grey-lighten-1"
                    class="mb-6"
                    icon="mdi-calendar-remove"
                ></v-icon>
                
                <v-empty-state
                    headline="404"
                    title="Wydarzenie nie zostało znalezione"
                    :text="`Wydarzenie '${alias}' nie istnieje.`"
                    :image="false"
                >
                    <template #actions>
                        <div class="d-flex flex-wrap justify-center gap-3">
                            <v-btn icon="mdi-arrow-left"
                                color="#BE1E2D"
                                to="/kalendarium"
                                density="comfortable"
                                rounded="xs"
                            >
                            </v-btn>
                        </div>
                    </template>
                </v-empty-state>
            </v-responsive>
        </v-container>
    </div>
    
    <div v-else-if="isLoading">
        <v-container class="fill-height">
            <v-row justify="center">
                <v-col cols="12" class="text-center">
                    <v-progress-circular
                        indeterminate
                        color="primary"
                        size="64"
                    ></v-progress-circular>
                    <p class="mt-4 text-body-1">Ładowanie wydarzenia...</p>
                </v-col>
            </v-row>
        </v-container>
    </div>
    
    <div v-else>
        <link rel="stylesheet" href="/src/assets/css/wydarzenieContent.css"/>
        <link rel="stylesheet" href="/src/assets/css/footer.css"/>
        
        <main class="content">
            <div class="buttons">
                <div v-if="auth.isAdmin" class="editButtons">
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
                    {{ dataWydarzenia.Aktywne ? 'Widoczny' : 'Ukryty' }}</v-btn>

                    <v-btn icon="$cancel" v-if="display.smAndDown.value"
                    class="vbuttons"
                    rounded="xs"
                    color="red"
                    @click="deleteEvent"
                    variant="flat"></v-btn>
                    <v-btn v-else class="vbuttons"
                    density="comfortable"
                    rounded="xs"
                    color="red"
                    @click="deleteEvent"
                    variant="flat"
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
                    variant="flat"></v-btn>
                    <v-btn v-else class="vbuttons"
                    density="comfortable"
                    rounded="xs"
                    color="green"
                    @click="addToGoogleCalendar"
                    variant="flat"
                    >
                    <v-icon icon="$calendar" start=""/>
                    Dodaj do kalendarza Google</v-btn>

                    <v-btn icon="$close" 
                    class="vbuttons"
                    rounded="xs"
                    color="grey-lighten-2"
                    @click="backButton"
                    density="comfortable"
                    variant="flat"></v-btn>
                </div>
            </div>

            <div id="alias">
                <input v-model="dataWydarzenia.Nazwa" id="inputAlias" placeholder="Tytuł" :readOnly="!auth.isAdmin"/>
            </div>
            <div class=info>
                <span style="display:inline-flex; align-items: center;"> 
                    <VueDatePicker v-model="dateRange"
                    :disabled="!auth.isAdmin"
                    :format="formatDatePicker" 
                    locale="pl" 
                    range 
                    class="date-picker">
                    
                </VueDatePicker></span>
                
                <p>Lokalizacja: <input :readOnly="!auth.isAdmin" v-model="dataWydarzenia.Lokalizacja" class="infoSpan" contenteditable="true" spellcheck="false"></input></p>
                <p>Organizator: <input :readOnly="!auth.isAdmin" v-model="dataWydarzenia.Organizator" class="infoSpan" contenteditable="true" spellcheck="false"></input></p>
                <p id="opis-label">Opis wydarzenia: </p>
                <div id="quill"></div>
            </div>
        </main>
        
        <v-dialog v-model="showAlert" max-width="500">
            <v-card>
                <v-card-title class="error white--text">
                    <v-icon left>mdi-alert-circle</v-icon>
                    {{ alertTitle }}
                </v-card-title>
                <v-card-text class="pt-4 text-body-1">
                    {{ alertMessage }}
                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn
                        variant="flat" 
                        color="#BE1E2D" @click="showAlert = false">
                        OK
                    </v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </div>
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
    
    .fill-height {
        min-height: 70vh;
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