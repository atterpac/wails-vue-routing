<template>
  <div class="cue-table">
    <div class="table-container" style="max-height: props.maxHeight + vh">
      <div class="header">
            <h1>{{props.title}}</h1>
            <button>+</button>
          </div>
      <table>
      <thead>
          <tr>
            <template v-for="column, idx in props.columns">
              <th 
                v-if="column.visible"
                style="{width: column.width + 'px'}"  
              >
                  {{column.label}}
              </th>
          </template>
            <th>
              <CogIcon width="24px" height="24px" :style="{ stroke: 'var(--gray-4)', fill: 'var(--gray-4)'}" fill="none" />
            </th>
          </tr>
        </thead>
        <tbody>
          <tr class="data" v-for="data, idx in tableData" :key="idx" >
            <td v-for="value in dataKeys">
              {{( data[value] )}}
            </td>
            <td class="options">
              <OptionsIcon class="option-icon" width="24px" height="24px"/>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div> 
</template>

<script setup lang="ts">
import { Column } from './column';
import { computed } from 'vue';
import OptionsIcon from '../../assets/svg/OptionsIcon.vue'
import CogIcon from '../../assets/svg/CogIcon.vue';

const props = defineProps({
    title: String,
    maxHeight: Number,
    columns: {
      type: Array<Column>,
        required: true,
    },
    tableData: {
      type: Array<any>,
      required: true,
    },
}) 

const dataKeys = computed(() => {
  let keys = props.tableData[0] ? props.columns.filter((column) => column.visible === true).map((column) => column.data) : [];
  return keys as string[];
});
</script>

<style scoped>
.cue-table {
  width: 100%;
  height: 90%;
  padding: 0px 20px 20px 20px;
}

.header {
    position: sticky;
    width: 100%;
    background-color: var(--bg-dark);
    height: 40px;
    width: 100%;
    position: sticky;
    top: 0;
    z-index: 2;
    border-radius: 5px;
    margin-left: -5px;
}

table {
  width: 100%;
  text-align: center;
  border-collapse: collapse;
}

thead {
  padding-top: 20px;
  border-radius: 3px;
  background-color: var(--bg-darker);
  position: sticky;
  top: 60px;
  z-index: 2; /* Ensure the header is above the table body */
}

tr {
  font-size: 16px;
  height: 36px;
  font-weight: 500;
  align-items: center;
}


th:first-child {
  border-top-left-radius: 5px;
  border-bottom-left-radius: 5px;
}

th:last-child {
  border-top-right-radius: 5px;
  border-bottom-right-radius: 5px ;
}
.data:nth-child(odd) td:first-child {
  border-top-left-radius: 5px;
  border-bottom-left-radius: 5px;
}

.data:nth-child(even) {
  background-color: var(--bg-darker);
}

.data:nth-child(even) td:first-child {
  border-bottom-left-radius: 5px;
  border-top-left-radius: 5px;
}

.data:nth-child(even) td:last-child {
  border-bottom-right-radius: 5px;
  border-top-right-radius: 5px;
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
    background-color: var(--primary-accent);
    border: none;
    border-radius: 5px;
}

button:hover{
    border:  1px solid #a0a0a0;
    width: 35px;
    height: 35px;
    transition: 50ms;
}

.options {
  align-items: center;
  justify-content: flex-end;
}

.option-icon {
  margin-bottom: -2px;
  stroke: var(--gray-6);
  fill: var(--gray-6);
}

.option-icon:hover {
  stroke: var(--primary-accent);
  fill: var(--primary-accent);
}
</style>

