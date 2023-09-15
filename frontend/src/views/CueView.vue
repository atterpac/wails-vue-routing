<template>
<section>
    <div class="table">
        <CueTable :maxHeight=900 />
    </div>
</section>
</template>

<script setup lang="ts">
import CueTable from '../components/tables/cues/CueTable.vue';
</script>

<style>

section {
    width: 100%;
    padding: 10px;
}

.table {
    width: 97%;
}

</style>
