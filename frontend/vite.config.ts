import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import dotenv from 'dotenv'
import { resolve } from 'path'
import path from 'path'

dotenv.config()

export default defineConfig({
  server: {
    port: Number(process.env.CLIENT_PORT) || 3002,
  },
  define: {
    __SERVER_PORT__: process.env.SERVER_PORT,
  },
  build: {
    outDir: path.join(__dirname, 'dist/client'),
    cssCodeSplit: true,
    manifest: true,
  },
  plugins: [react()],
  publicDir: 'public',
  resolve: {
    alias: {
      '@': resolve(__dirname, 'src'),
      '@components': resolve(__dirname, 'src/components'),
      /*'@core': resolve(__dirname, 'src/core'),
      '@game': resolve(__dirname, 'src/game'),*/
      '@assets': resolve(__dirname, 'src/assets'),
      '@src': resolve(__dirname, 'src'),
      //'@pages': resolve(__dirname, 'src/pages'),
    },
  },
  css: {
    preprocessorOptions: {
      scss: {
        silenceDeprecations: ['legacy-js-api'],
      },
    },
  },
})