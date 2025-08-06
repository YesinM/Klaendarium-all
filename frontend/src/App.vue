<script setup>
function loginCAS() {
  const returnUrl = encodeURIComponent(window.location.origin + "/frontend")
  window.location.href = `/api/login`
}


</script>

<template class="page-wrapper">
    <v-app>
    <link rel="stylesheet" href="/src/assets/css/footer.css"/>
    <link rel="stylesheet" href="/src/assets/css/wydarzenieContent.css"/>
    <header>
        <div id="banner">
            <img id="baner" style="height: 100%; width: 100%;" src="../src/assets/wydarzenia_baner.jpg">
        </div>
        <nav>
            <v-toolbar class="w-100 justify-center padding" color="white">
                <div class="w-100 justify-center">
                    <v-btn href="https://al.edu.pl" target="_self" class="me-2"
                    variant="flat" 
                    color="#BE1E2D"
                    >
                    Start</v-btn>
                    <v-btn to="/kalendarium/dodaj"
                    variant="flat"
                    color="grey-lighten-5">
                        Dodaj
                        <v-icon icon="$plus"></v-icon>
                    </v-btn>
                    <v-btn @click="loginCAS">
                        Zaloguj się
                    </v-btn>
                </div>
            </v-toolbar>
        <!-- <div class="middle4">
            <ul class="breadcrumb">
                <li class="active">Jesteś tutaj:&nbsp;&nbsp;</li>
                <li> <a class="pathway" href="https://al.edu.pl">Start</a></li>
                <li class="active">Kalendarium</li>
                <li>
                    <RouterLink to="/kalendarium/dodaj" v-if="zalogowany">
                        <span>Dodaj wydarzenie</span>
                    </RouterLink>
                </li>
            </ul>
        </div> -->
        </nav>
    </header>
    <div class="content">
        <router-view></router-view>
    </div>
    <v-footer class="bg-grey-lighten-4 py-0 flex-column w-100 " height="auto">

    <div class="text-body-2 left-0 d-flex ga-2 my-3">
        <div>
        &copy; 2025 –
        <a href="https://al.edu.pl/" class="text-primary font-weight-medium">Akademia Łomżyńska</a> –
        Wszelkie prawa zastrzeżone.
        Wykonanie:
        <a href="https://al.edu.pl/dsk/kontakt" target="_blank" class="text-primary font-weight-medium">
            Dział Systemów Komputerowych
        </a>
        </div>

        <v-divider vertical class="ma-0"></v-divider>

        <div class="d-flex flex-wrap justify-center ga-3">
        <a href="http://al.edu.pl/deklaracja-dostepnosci" class="text-decoration-none text-secondary">Deklaracja dostępności</a>
        <a href="http://usosweb.al.edu.pl" class="text-decoration-none text-secondary">USOSweb</a>
        <a href="http://poczta.al.edu.pl" class="text-decoration-none text-secondary">Poczta</a>
        <a href="http://al.edu.pl/mapa-strony" class="text-decoration-none text-secondary">Mapa strony</a>
        <a href="http://al.edu.pl/pomoc" class="text-decoration-none text-secondary">Pomoc</a>
        <a href="http://al.edu.pl/kontankt" class="text-decoration-none text-secondary">Kontakt</a>
        </div>
    </div>
    </v-footer>
    </v-app>
</template>

<style>
    
a {
        text-decoration: none !important;   
    }
nav {
    padding: 10px 0px 20px 195px;

}
li .module-event-date {
    max-height: 60px;
}
.page-wrapper {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
}

.content {
    flex-grow: 1;
}

html, body {
    height: 100%;
    margin: 0;
    padding: 0;
}

.v-application {
    display: flex !important;
    flex-direction: column !important;
    min-height: 100vh !important;
}

.v-application > main {
    flex-grow: 1 !important;
}

.v-footer {
    flex-shrink: 0 !important;
}
#logo{
    width: 25%;
    padding: 0 10px 0 10px;
}
@media (max-width: 680px){
        nav {
            padding: 10px 0px 20px 20px;
            text-align: center;
        }
    }
</style>
