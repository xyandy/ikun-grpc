import { viteCommonjs } from '@originjs/vite-plugin-commonjs';
import vue from '@vitejs/plugin-vue';
import vueJsx from '@vitejs/plugin-vue-jsx';
import { fileURLToPath, URL } from 'node:url';
import { defineConfig } from 'vite';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), vueJsx(), viteCommonjs()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
      '@wailsjs': fileURLToPath(new URL('./wailsjs', import.meta.url)),
    },
  },
});
