<template>
<svg class="wrapper" viewBox="0 0 24 24" :fill=props.stroke  :width=props.width :height=props.height xmlns="http://www.w3.org/2000/svg">
    <path 
        d="M2.5,12 C2.5,11.1716 3.17157,10.5 4,10.5 L20,10.5 C20.8284,10.5 21.5,11.1716 21.5,12 C21.5,12.8284 20.8284,13.5 20,13.5 L4,13.5 C3.17157,13.5 2.5,12.8284 2.5,12 Z"
        :stroke=props.stroke 
        stroke-width="2"
        stroke-linejoin="round"
    />
</svg>
</template>

<script setup>
const props = defineProps({
    fill: String,
    stroke: String,
    height: String,
    width: String,
})
</script>

<style scoped>
.wrapper {
    margin: 5px;
}
</style>
