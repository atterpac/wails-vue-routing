

<template>
<svg class="wrapper" viewBox="0 0 32 32" fill="none" :width=props.width :height=props.height xmlns="http://www.w3.org/2000/svg">
   <path 
        d="M5.975 17.504l14.287.001-6.367 6.366L16.021 26l10.004-10.003L16.029 6l-2.128 2.129 6.367 6.366H5.977z"
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
.wrapper {
    margin: 5px;
}
</style>

