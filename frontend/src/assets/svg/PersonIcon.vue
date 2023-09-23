<template>
<svg id="person" xmlns="http://www.w3.org/2000/svg" :width="props.width" :height="props.height"
    xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 24 24">
    <path id="person"
        fill='var(--gray-6)'
        d="M12,1c1.104569,0,2,.895431,2,2s-.895431,2-2,2-2-.895431-2-2s.895431-2,2-2Zm8.79,4.546L14.776,6h-5.553L3.211,5.546c-.381613-.023857-.715318.254732-.759993.634471s.215263.72817.591993.793529l6.106.97c.208062.033147.369117.19991.395.409L10,12L6.865,22.067c-.102993.309545.028356.648616.313.808l.071.04c.17476.097529.382668.116422.572146.051992s.342772-.206146.421854-.389992L12,13.914l3.757,8.663c.079082.183846.232376.325562.421854.389992s.397386.045537.572146-.051992l.07-.04c.285025-.159065.416822-.498212.314-.808L14,12l.456-3.647c.025883-.20909.186938-.375853.395-.409l6.106-.97c.37673-.065359.636668-.41379.591993-.793529s-.37838-.658328-.759993-.634471Z"
        transform="matrix(.68225 0 0 0.682247 3.812999 3.813772)" />
    <path id="person-head" d="M0,0h24v24h-24Z" fill="none" />
    <ellipse 
        v-if="props.beam_size == `tight`"
        id="tight" 
        rx="8.846191" 
        ry="8.846191"
        transform="matrix(.899449 0 0 0.880634 12 12.286276)" 
        :fill="props.beam_color"
        fill-opacity="0.1"
        stroke-width="0" />
    <ellipse 
        v-if="props.beam_size == `wide`"
        id="wide" 
        rx="8.846191" 
        ry="8.846191"
        transform="matrix(1.356513 0 0 1.289325 12.000027 12.002316)" 
        :fill="props.beam_color"
        fill-opacity="0.1" stroke-width="0" />
    <ellipse 
        v-if="props.beam_size == `full-body`"
        id="full-body" 
        rx="8.846191" 
        ry="8.846191"
        transform="translate(11.936184 11.94076)" 
        :fill="props.beam_color"
        fill-opacity="0.1" 
        stroke-width="0" 
    />
    <ellipse 
        v-if="props.beam_size == `3-4`"
        id="3-4" 
        rx="8.846191" 
        ry="8.846191"
        transform="matrix(.802725 0 0 0.762962 12.000001 10.270529)" 
        :fill="props.beam_color"
        fill-opacity="0.1" 
        stroke-width="0" 
    />
    <ellipse 
        v-if="props.beam_size == `half-body`"
        id="half-body" 
        rx="8.846191" 
        ry="8.846191"
        transform="matrix(.652469 0 0 0.620605 11.999999 9.011211)" 
        :fill="props.beam_color"
        fill-opacity="0.1" 
        stroke-width="0" 
    />
    <ellipse
        v-if="props.beam_size == `shoulder`"
        id="head-sholders" 
        rx="8.846191" 
        ry="8.846191"
        transform="matrix(.452944 0 0 0.430502 11.999997 7.249285)" 
        :fill="props.beam_color"
        fill-opacity="0.1" 
        stroke-width="0" />
    <ellipse 
        v-if="props.beam_size == `headshot`"
        id="headshot" 
        rx="8.846191" 
        ry="8.846191"
        transform="matrix(.268324 0 0 0.271073 11.999997 5.919184)" 
        :fill="props.beam_color"
        fill-opacity="0.1" />
</svg>
</template>
    
<script setup>
    const props = defineProps({
        fill: String,
        stroke: String,
        height: String,
        width: String,
        beam_size: String,
        beam_color: String,
    })
</script>
    
<style scoped>
    .wrapper {
        margin: 5px;
    }
</style>
