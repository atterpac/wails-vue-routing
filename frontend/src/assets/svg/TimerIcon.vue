
<template>
    <svg class="wrapper" viewBox="0 0 24 24"  :width=props.width :height=props.height xmlns="http://www.w3.org/2000/svg">
        <path opacity="1" 
            d="M21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12Z" 
        >
        </path> 
        <path d="M21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12Z" 
            :fill=props.fill
            :stroke=props.stroke 
            stroke-width="2">
        </path> 
        <path d="M12 7L12 12" 
            :stroke=props.stroke 
            stroke-width="2" 
            stroke-linecap="round" 
            stroke-linejoin="round">
        </path> 
        <path d="M21 4L20 3" 
            :stroke=props.stroke 
            stroke-width="2" 
            stroke-linecap="round" 
            stroke-linejoin="round">
        </path> 
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
    
