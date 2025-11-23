import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/LoginView.vue'
import DashboardView from '../views/DashboardView.vue'
import DashboardProveedor from '../views/DashboardProveedor.vue' // 👈 NUEVA VISTA

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/api/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/dashboard', // ✅ SOLO ADMIN
      name: 'dashboard',
      component: DashboardView
    },
    {
      path: '/api/proyectos', // ✅ NUEVA RUTA
      name: 'dashboard-proveedor',
      component: DashboardProveedor
    },
    {
      path: '/',
      redirect: '/api/login'
    }
  ]
})

export default router
