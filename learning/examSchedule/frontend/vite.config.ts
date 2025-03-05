import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    port: 3000,
    strictPort: true,
    host: true,
    proxy: {
      '/exam-schedule': {
        target: 'http://localhost:7070',
        changeOrigin: true,
        secure: false
      }
    }
  }
})
