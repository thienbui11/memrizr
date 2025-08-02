<template>
<<<<<<< HEAD
  <div class="text-4xl font-bol text-center my-2">App Component</div>
  <div class="flex justify-around my-4">
    <router-link to="/">Home/Details</router-link>
    <router-link to="/authenticate">Auth</router-link>
    <router-link to="/notaroute">Not a Route</router-link>
  </div>
  <router-view></router-view>
</template>

<script>
import { defineComponent } from 'vue';
=======
  <div class="mx-8 my-8">
    <Loader
      v-if="isLoading"
      :height="512"
      class="animate-spin stroke-current text-blue-500 mx-auto"
    />
    <router-view v-else></router-view>
  </div>
</template>

<script>
import { defineComponent, onMounted } from 'vue';
import Loader from './components/ui/Loader.vue';
import { useAuth } from './store/auth';
>>>>>>> 3999c5b64d9e82200d4e850d267e0d3ed56f0643

// Wrapping exported object in define component
// gives us typing help! Woot!
export default defineComponent({
  name: 'App',
<<<<<<< HEAD
=======
  components: {
    Loader,
  },
  setup() {
    const { initializeUser, isLoading } = useAuth();

    onMounted(() => {
      initializeUser();
    });

    return {
      isLoading,
    };
  },
>>>>>>> 3999c5b64d9e82200d4e850d267e0d3ed56f0643
});
</script>