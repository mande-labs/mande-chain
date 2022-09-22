import starportLibrary from '@starport/vue'
import { createApp } from 'vue'
import Toast from "vue-toastification";
import "vue-toastification/dist/index.css";

import App from './App.vue'
import router from './router'
import store from './store'

const options = {
    // You can set your default options here
};

const app = createApp(App)
app.use(store).use(router).use(starportLibrary).use(Toast, options).mount('#app')
