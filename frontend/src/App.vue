<template>
    <LocaleSelect class="p-2" />
    <div class="flex h-screen p-4 bg-gray-10 -mt-12">
        <div class="mx-auto my-auto">
            <div class="flex flex-col flex-1 mb-12">
                <h1
                    class="text-6xl font-bold text-center text-transparent bg-clip-text bg-gradient-to-r from-purple-400 to-pink-600">
                    {{ $t('title') }}
                </h1>
                <span class="text-2xl text-center text-gray-500 dark:text-gray-200">
                    {{ $t('description') }}
                </span>
            </div>
            <div class="flex flex-row flex-1 space-x-1">
                <input type="url" v-model="form.url"
                    class="w-full px-4 py-2 text-gray-700 bg-white border border-gray-300 rounded-md focus:outline-none focus:border-purple-500"
                    :placeholder="$t('input')" />
                <ColorInput v-model="form.bg_color" disable-text-inputs format="hex8 string" position="bottom left" />
                <ColorInput v-model="form.fg_color" disable-text-inputs format="hex8 string" position="bottom left" />
            </div>
            <div class="flex flex-row justify-center mt-4 space-x-4">
                <button @click="ShortURL()"
                    class="px-6 py-2 text-white bg-purple-500 border border-purple-500 rounded-md hover:bg-purple-600 focus:outline-none focus:bg-purple-600">
                    {{ $t('create') }}
                </button>
            </div>
            <div class="flex flex-col flex-1 mt-12 space-y-4" v-if="loaded">
                <span class="text-2xl text-center text-gray-500 dark:text-gray-200">
                    {{ $t('after') }}
                </span>
                <img alt="QR Code" class="mx-auto dark:bg-gray-300" :src="QRCode" />
            </div>
            <NightToggle />
        </div>
    </div>
</template>

<script setup>
import { ref, reactive, inject } from "vue";
import NightToggle from './components/night-toggle.vue';
import LocaleSelect from '@/components/locale-select.vue';
import ColorInput from 'vue-color-input'
import axios from "axios";

const QRCode = ref('')
const loaded = ref(false)
const toast = inject('toast');
const form = reactive({
    url: '',
    bg_color: '#ffffff00',
    fg_color: '#993dbdff'
})

axios.defaults.baseURL = 'http://localhost:3000/api/v1'

// Shorten URL
function ShortURL() {
    if (form.url.length > 0) {
        axios.post('/convert', {
            url: form.url,
            bg_color: form.bg_color,
            fg_color: form.fg_color
        }).then((res) => {
            QRCode.value = res.data.image
            loaded.value = true
        }).catch((error) => {
            toast.error(error.response.data.error)
        })
    }
}

</script>