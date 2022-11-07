<template>
    <div class="relative inline-block text-left" ref="langdrop">
        <button
            type="button"
            class="inline-flex justify-center w-full rounded-md border border-gray-300 dark:border-gray-600 px-2 py-1 bg-gray-50 dark:bg-gray-800 text-sm focus:outline-none"
            id="menu-button"
            aria-expanded="true"
            aria-haspopup="true"
            @click="show = !show"
        >
            {{ locales.filter(locale => currentLocale === locale.id)[0].name }}
            <svg
                class="-mr-1 ml-2 h-5 w-5 transition"
                xmlns="http://www.w3.org/2000/svg"
                :class="[show ? 'transform rotate-180' : null]"
                viewBox="0 0 20 20"
                fill="currentColor"
                aria-hidden="true"
            >
                <path
                    fill-rule="evenodd"
                    d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
                    clip-rule="evenodd"
                />
            </svg>
        </button>

        <transition
            enter-active-class="transition ease-out duration-100"
            enter-from-class="transform opacity-0 scale-95"
            enter-to-class="transform opacity-100 scale-100"
            leave-active-class="transition ease-in duration-75"
            leave-from-class="transform opacity-100 scale-100"
            leave-to-class="transform opacity-0 scale-95"
        >
            <div
                class="origin-top-left absolute left-0 mt-2 w-56 rounded-md shadow-lg bg-gray-200 dark:bg-gray-900  focus:outline-none"
                role="menu"
                aria-orientation="vertical"
                aria-labelledby="menu-button"
                tabindex="-1"
                v-if="show"
            >
                <div class="py-1" role="none">
                    <button
                        href="#"
                        class="hover:bg-gray-300 dark:hover:bg-gray-700 w-full block px-4 py-2 text-sm transition"
                        role="menuitem"
                        tabindex="-1"
                        id="menu-item-0"
                        v-for="(locale, index) in locales"
                        :key="index"
                        @click="setLang(locale.id)"
                    >{{ locale.name }}</button>
                </div>
            </div>
        </transition>
    </div>
</template>

<script setup>
import { localeOptions } from '@/constant/config';
import { ref, onBeforeUnmount, onBeforeMount } from 'vue';
import { useI18n } from 'vue-i18n';
const { locale } = useI18n({ useScope: 'global' })

const locales = localeOptions
const currentLocale = ref(localStorage.getItem('locale') || 'en')
const show = ref(false)
const langdrop = ref(null)

onBeforeMount(() => {
    locale.values = currentLocale.value
})

function setLang(id) {
    localStorage.setItem('locale', id)
    currentLocale.value = id
    locale.value = id
    console.log(currentLocale.value)
    show.value = false
}

function clickOutside(e) {
    if (langdrop.value && !langdrop.value.contains(e.target)) {
        show.value = false;
    }
}

document.addEventListener('click', clickOutside);
onBeforeUnmount(() => {
    document.removeEventListener('click', clickOutside);
});
</script>