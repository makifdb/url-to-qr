import { createApp } from 'vue'
import { i18n } from "@/plugins/i18n"
import App from './App.vue'
import Toast from '@/plugins/toast';

import '@/assets/css/main.css'

const app = createApp(App);

app.use(i18n);
app.use(Toast);

app.mount('body');