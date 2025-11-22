import { fileURLToPath, URL } from 'node:url'
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'

export default defineConfig({
  plugins: [
    vue(),
    vueDevTools(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
  },
  server: {
    port: 3000,
    proxy: {
      // Redirige solo otras rutas del backend, no /api/login
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        secure: false,
        // Ignora exactamente '/api/login' para que Vue Router maneje esta ruta
        bypass: (req) => {
          if (req.url === '/api/login') {
            console.log('Enpoint local /api/login')
            return req.url
          }
        }
      },
    },
  },
})
