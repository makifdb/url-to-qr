<template>
    <div class="fixed right-3 bottom-3">
        <Toggle
            v-model="darkTheme"
            @click="toggleTheme"
            activeBgClass="bg-gray-600"
            BgClass="bg-gray-200"
        >
            <template #icon>
                <svg
                    class="w-6 h-6 text-white"
                    aria-hidden="true"
                    fill="currentColor"
                    viewBox="0 0 20 20"
                    v-if="darkTheme"
                >
                    <path d="M17.293 13.293A8 8 0 016.707 2.707a8.001 8.001 0 1010.586 10.586z" />
                </svg>
                <svg
                    class="w-6 h-6 text-yellow-500"
                    aria-hidden="true"
                    fill="currentColor"
                    viewBox="0 0 20 20"
                    v-else
                >
                    <path
                        fill-rule="evenodd"
                        d="M10 2a1 1 0 011 1v1a1 1 0 11-2 0V3a1 1 0 011-1zm4 8a4 4 0 11-8 0 4 4 0 018 0zm-.464 4.95l.707.707a1 1 0 001.414-1.414l-.707-.707a1 1 0 00-1.414 1.414zm2.12-10.607a1 1 0 010 1.414l-.706.707a1 1 0 11-1.414-1.414l.707-.707a1 1 0 011.414 0zM17 11a1 1 0 100-2h-1a1 1 0 100 2h1zm-7 4a1 1 0 011 1v1a1 1 0 11-2 0v-1a1 1 0 011-1zM5.05 6.464A1 1 0 106.465 5.05l-.708-.707a1 1 0 00-1.414 1.414l.707.707zm1.414 8.486l-.707.707a1 1 0 01-1.414-1.414l.707-.707a1 1 0 011.414 1.414zM4 11a1 1 0 100-2H3a1 1 0 000 2h1z"
                        clip-rule="evenodd"
                    />
                </svg>
            </template>
        </Toggle>
    </div>
</template>
<script setup>

import { onBeforeMount, computed, ref, watch } from "vue";
import Toggle from '@/components/toggle.vue';

const darkTheme = ref(false)
const theme = computed(() => (darkTheme.value ? 'dark' : 'light'))

onBeforeMount(() => {
    darkTheme.value = localStorage.getItem('theme') === 'dark'
})

function toggleTheme() {
    theme.value = darkTheme.value ? 'light' : 'dark'
    localStorage.setItem('theme', theme.value)
}

watch(theme, (theme) => {
    theme === "light"
        ? document.querySelector("html").classList.remove("dark")
        : document.querySelector("html").classList.add("dark");
})
</script>