<template>
    <svg class="wrapper" viewBox="0 0 24 24" :fill=props.stroke  :width=props.width :height=props.height xmlns="http://www.w3.org/2000/svg">
        <path 
            d="M12 1a2 2 0 1 1-2 2 2 2 0 0 1 2-2zm8.79 4.546L14.776 6H9.223l-6.012-.454a.72.72 0 0 0-.168 1.428l6.106.97a.473.473 0 0 1 .395.409L10 12 6.865 22.067a.68.68 0 0 0 .313.808l.071.04a.707.707 0 0 0 .994-.338L12 13.914l3.757 8.663a.707.707 0 0 0 .994.338l.07-.04a.68.68 0 0 0 .314-.808L14 12l.456-3.647a.473.473 0 0 1 .395-.409l6.106-.97a.72.72 0 0 0-.168-1.428z"
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
