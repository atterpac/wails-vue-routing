<template>
<svg class="wrapper" viewBox="0 0 24 24" fill="none"  :width=props.width :height=props.height xmlns="http://www.w3.org/2000/svg">
    <path 
        d="M22 5C22 3.34315 20.6569 2 19 2H5C3.34315 2 2 3.34315 2 5V19C2 20.6569 3.34315 22 5 22H19C20.6569 22 22 20.6569 22 19V5ZM20 5C20 4.44772 19.5523 4 19 4H5C4.44772 4 4 4.44772 4 5V19C4 19.5523 4.44772 20 5 20H19C19.5523 20 20 19.5523 20 19V5Z"
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
