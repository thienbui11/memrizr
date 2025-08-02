<<<<<<< HEAD
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  server: {
    host: true,
    port: 3000,
    allowedHosts: ['malcorp.test'],
  },
})
=======
import vue from '@vitejs/plugin-vue';

/**
 * https://vitejs.dev/config/
 * @type {import('vite').UserConfig}
 */
export default {
  base: '/account/',
  plugins: [vue()],
};
>>>>>>> 3999c5b64d9e82200d4e850d267e0d3ed56f0643
