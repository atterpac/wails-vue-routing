import { createRouter, createWebHashHistory } from 'vue-router';
import HomeView from './views/HomeView.vue'
import DashboardView from './views/DashboardView.vue'
import CueView from './views/CueView.vue'

const routes = [
    {path:"/", name:"Home" , component: HomeView },
    {path:"/dash", name:"Dashboard" , component: DashboardView},
    {path:"/cues", name:"Cues" , component: CueView }
]

export const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

