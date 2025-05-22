<script setup>
import {ref} from 'vue'

const zalogowany = ref(true);
const isLoaded = ref(false);

window.addEventListener('load', () => {
    isLoaded.value = true;
})
</script>

<template v-if="isLoaded">
    <link rel="stylesheet" href="/src/assets/css/bootstrap.min.css"/>
    <link rel="stylesheet" href="/src/assets/css/template.css"/>
    <link rel="stylesheet" href="/src/assets/css/kalendarium.css"/>
    <link rel="stylesheet" href="/src/assets/css/footer.css"/>
    <!--  -->
    <header>
        <div id="banner">
            <img id="baner" style="height: 100%; width: 100%;" src="../src/assets/wydarzenia_baner.jpg">
        </div>
        <nav>
        <div class="middle4">
            <ul class="breadcrumb">
                <li class="active">Jesteś tutaj:&nbsp;&nbsp;</li>
                <li> <a class="pathway" href="https://al.edu.pl">Start</a></li>
                <li class="active">Kalendarium</li>
                <li>
                    <RouterLink to="/kalendarium/dodaj" v-if="zalogowany">
                        <span>
                            Dodaj wydarzenie    
                        </span>
                    </RouterLink>
                </li>
            </ul>
        </div>
        </nav>
    </header>
    <router-view></router-view>
    <footer>
        <div class="social-media">
            <h3> <span>Social media</span> </h3>
            <a href="/"><img src="/src/assets/f.png"></a>
            <a href="/"><img src="/src/assets/i.png"></a>
            <a href="/"><img src="/src/assets/linkedin.png"></a>
            <a href="/"><img src="/src/assets/y.png" alt=""></a>
        </div>
        <div class="contacts">
            <h2>Kontakt</h2>
            <span>Akademia Łomżyńska</span>
            <span>Akademicka</span>
            <span>NIP: 718-19-47-148</span>
            <span><a href="tel:+48862155950">+48 86 215 59 50</a></span>
            <span><a href="mailto:biuro@al.edu.pl">biuro@al.edu.pl</a></span>
        </div>
        <div class="footer-bottom">
            <div>Copyright @2025 - <a href="https://al.edu.pl/">Akademia Łomżyńska</a> - Wszelkie prawa zastrzeżone. Wykonanie: <a href="https://al.edu.pl/dsk/kontakt">Dział Systemów Komputerowych</a></div>
            <div><a href="/">Deklaracja dostępności </a> - <a href="/">USOSweb</a> - <a href="/">Poczta</a> - <a href="/">Mapa strony</a> - <a href="/">Pomoc</a> - <a href="/">Kontakt</a></div>
        </div>
    </footer>
</template>

<style>
    li .module-event-date{
        max-height: 60px;
    }
</style>


