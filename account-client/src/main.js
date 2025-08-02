import { createApp } from 'vue';
import App from './App.vue';
<<<<<<< HEAD
import router from './routes';
import './index.css';

createApp(App).use(router).mount('#app');
=======
import { createAuthStore } from '../src/store/auth';
import router from './routes';
import './validators';
import './index.css';

const authStore = createAuthStore({
  onAuthRoute: '/',
  requireAuthRoute: '/authenticate',
});

createApp(App).use(authStore).use(router).mount('#app');
>>>>>>> 3999c5b64d9e82200d4e850d267e0d3ed56f0643
