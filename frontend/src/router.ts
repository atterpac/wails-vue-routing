import { createRouter, createWebHashHistory } from 'vue-router';
import HomeView from './views/HomeView.vue'
import DashboardView from './views/DashboardView.vue'
import CueView from './views/CueView.vue'
import LoginView from './views/auth/LoginView.vue'
import RegisterView from './views/auth/RegisterView.vue'
import PassResetView from './views/auth/PassResetView.vue'

const routes = [
    {path:"/", name:"Home" , component: HomeView },
    {path:"/dash", name:"Dashboard" , component: DashboardView},
    {path:"/cues", name:"Cues" , component: CueView },
    {path:"/login", name:"Login" , component: LoginView },
    {path:"/register", name:"Register" , component: RegisterView },
    {path:"/reset", name:"PasswordReset" , component: PassResetView }
]

export const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

