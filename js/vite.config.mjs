import { defineConfig } from 'vite'
import tailwindcss from '@tailwindcss/vite'

export default defineConfig({
  plugins: [
    tailwindcss(),
  ],
  server: {
    port: 8080,
  },
  build: {
    lib: {
      entry: 'src/main.ts',
      name: 'hexago',
      formats: ['iife'],
    },
    rollupOptions: {
      output: {
        entryFileNames: 'app.js',
      },
    },
    outDir: '../static/dist',
  },
})

