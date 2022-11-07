<template>
    <button
        type="button"
        class=" inline-flex items-center h-8 w-16 border-2 border-transparent rounded-full cursor-pointer transition-colors ease-in-out duration-200 focus:outline-none"
        :class="[modelValue ? activeBgClass : BgClass]"
        role="switch"
        @click="toggle"
        :aria-checked="modelValue"
    >
        <span class="sr-only">Use setting</span>
        <span
            class="pointer-events-none h-6 w-6 rounded-full transform transition ease-in-out duration-200"
            :class="[modelValue ? 'translate-x-9' : 'translate-x-0', modelValue ? ActiveButtonClass : ButtonClass]"
        >
            <span
                class="absolute inset-0 h-full w-full flex items-center justify-center transition-opacity"
                :class="[modelValue ? ' ease-out duration-100' : ' ease-in duration-200']"
                aria-hidden="true"
            >
                <slot name="icon" />
            </span>
        </span>
    </button>
</template>

<script setup>
const props = defineProps({
    modelValue: Boolean,
    activeBgClass: String,
    ActiveButtonClass: String,
    BgClass: String,
    ButtonClass: String,
});

const emits = defineEmits(['update:modelValue']);

function toggle() {
    emits("update:modelValue", !props.modelValue)
}


</script>