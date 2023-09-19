

<template>
    <svg class="wrapper" viewBox="0 0 24 24" :fill=props.stroke  :width=props.width :height=props.height xmlns="http://www.w3.org/2000/svg">
        <path 
            d="M3 3h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1zm6 2v14h11V5H9z"
            :stroke=props.stroke 
            stroke-width=1
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
    