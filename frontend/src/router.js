import { createRouter, createWebHistory} from 'vue-router';
import ListWydarzenia from './components/ListWydarzenia.vue';
import Wydarzenie from './components/Wydarzenie.vue';

const routes = [
    {path: '/kalendarium/wydarzenie', component: Wydarzenie},
    {path: '/kalendarium', component: ListWydarzenia},
    {path: '/kalendarium/dodaj', component: Wydarzenie}
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router