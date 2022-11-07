<template>
    <transition enter-active-class="transition ease-out duration-100" enter-from-class="transform opacity-0 scale-95"
        enter-to-class="transform opacity-100 scale-100" leave-active-class="transition ease-in duration-75"
        leave-from-class="transform opacity-100 scale-100" leave-to-class="transform opacity-0 scale-95">
        <div class="max-w-sm rounded-lg p-3 bg-gray-200 dark:bg-dark cursor-pointer" @click="removeEl(randNum)"
            v-show="isActive" ref="root">
            <div class="flex sm:flex-row flex-col items-center space-x-2 text-sm">
                <svg v-if="type === 'success'" class="w-7 h-7 text-green-500 flex-shrink-0" fill="none"
                    stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <svg v-if="type === 'error'" class="w-7 h-7 text-red-500 flex-shrink-0" fill="none"
                    stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <svg v-if="type === 'info'" class="w-7 h-7 text-blue-500 flex-shrink-0" fill="none"
                    stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <svg v-if="type === 'warning'" class="w-7 h-7 text-yellow-500 flex-shrink-0" fill="none"
                    stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                </svg>
                <p class="whitespace-pre-wrap break-words text-xs max-w-sm capitalize">{{ message }}</p>
            </div>
        </div>
    </transition>
</template>


<script setup>
import { ref } from '@vue/reactivity'
import { onMounted } from '@vue/runtime-core';
const props = defineProps({
    message: {
        type: String,
        required: true
    },
    type: {
        type: String,
        default: 'default'
    },
    duration: {
        type: Number,
        default: 5000
    },
})
const parentDiv = document.getElementById('toast') ? document.getElementById('toast') : document.createElement('div')
const isActive = ref(false)
const randNum = ref(Math.floor(Math.random() * 10000))
const root = ref(null)


parentDiv.className = 'fixed top-10 md:right-3 space-y-3 z-50'
parentDiv.id = 'toast'

setTimeout(() => {
    removeEl(randNum.value)
}, props.duration);

function removeEl(id) {
    isActive.value = false
    setTimeout(() => {
        var toast = document.getElementById(id);
        if (toast != null) {
            toast.remove()
        }
    }, 200);
}
onMounted(() => {
    isActive.value = true
    const container = document.body
    container.appendChild(parentDiv)
    root.value.setAttribute("id", randNum.value)
    parentDiv.appendChild(root.value)
})
</script>