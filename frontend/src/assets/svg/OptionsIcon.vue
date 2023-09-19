
<template>
<svg class="wrapper" viewBox="0 0 32 32" fill="none" :width=props.width :height=props.height xmlns="http://www.w3.org/2000/svg">
   <path 
        d="M10 12a2 2 0 1 0 4 0 2 2 0 0 0-4 0zm0-6a2 2 0 1 0 4 0 2 2 0 0 0-4 0zm0 12a2 2 0 1 0 4 0 2 2 0 0 0-4 0z"
        :stroke=props.stroke 
        :fill=props.stroke
        stroke-width="1.5" 
        stroke-linecap="round" 
        stroke-linejoin="round">
    </path> 
</svg>
</template>

<script setup lang="ts">
const props = defineProps({
    stroke: String,
    fill: String,
    width: String,
    height: String
})
</script>

<style>

</style>

