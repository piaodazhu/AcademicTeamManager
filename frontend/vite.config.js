import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    host: '127.0.0.1',
    port: 8062
  },
  css: {
    preprocessorOptions: {
          less: {
              modifyVars: {
                  'primary-color': '#476FFF',
                  'link-color': '#476FFF',
                  'border-radius-base': '5px'
              },
              javascriptEnabled: true
          }
      }
  },
})
