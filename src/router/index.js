import { createRouter, createWebHistory } from 'vue-router'
import WelcomePage from '@/views/WelcomePage.vue'
import MainLayout from '@/views/MainLayout.vue'

const routes = [
    { path: '/', component: WelcomePage },      // 纯欢迎页
    { path: '/main', component: MainLayout }    // 纯功能页（无横幅）
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router
