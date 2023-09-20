<template>
<section class="table-wrapper">

    <div class="fixture-table">
        <Table :maxHeight=90 :columns="FixtureColumns" :tableData="FixtureData" title="Fixtures"/>
    </div>
</section>
</template>

<script setup lang="ts">
import { Ref, ref } from 'vue';
import Table from '../components/tables/CustomTable.vue';
import { Column } from '../components/tables/column';
import { FixtureData, FixtureColumns } from '../components/tables/fixtures/fixtureData';

</script>

<style scoped>

section {
    width: 100%;
    overflow: auto;
}

.header {
    position: sticky;
    background-color: var(--gray-8);
    height: 40px;
    border-radius: 10px;
}

h1 {
    margin-left: 30px;
    float: left;
}

button {
    float: right;
    margin-right: 20px;
    width: 30px;
    height: 30px;
    font-size: 18px;
    font-weight: 700;
    background-color: #409292;
    border: none;
    border-radius: 5px;
}

button:hover{
    border:  1px solid #a0a0a0;
    width: 35px;
    height: 35px;
    transition: 50ms;
}

.fixture-table {
    width: 97%;
}

</style>
