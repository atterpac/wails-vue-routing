<template>
  <div class="cue-table">
    <div class="table-container" style="max-height: props.maxHeight + vh">
      <div class="header">
            <h1>Cues</h1>
            <button>+</button>
          </div>
      <table>
      <thead>
          <tr>
            <th>Cue</th>
            <th>Part</th>
            <th>Time</th>
            <th>Delay</th>
            <th>Hang</th>
            <th>Follow</th>
            <th>Block</th>
            <th>Description</th>
            <th>Placement</th>
            <th>Scene ID</th>
            <th>Notes</th>
          </tr>
        </thead>
        <tbody>
          <tr class="data" v-for="cue in CueData" :key="cue.number">
            <td>{{ cue.number }}</td>
            <td>{{ cue.part }}</td>
            <td>{{ cue.time }}</td>
            <td>{{ cue.delay }}</td>
            <td>{{ cue.hang }}</td>
            <td>{{ cue.follow }}</td>
            <td>{{ cue.block }}</td>
            <td>{{ cue.description }}</td>
            <td>{{ cue.placement }}</td>
            <td>{{ cue.scene_id }}</td>
            <td>{{ cue.notes }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { CueData } from './cuedata'

const props = defineProps({
    maxHeight: Number,
    }) 
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
    background-color: var(--gray-8);
    height: 40px;
    width: 100%;
    position: sticky;
    top: 0;
    z-index: 2;
    border-radius: 10px;
}

table {
  width: 100%;
  text-align: center;
  border-collapse: collapse;
}

thead {
  padding-top: 20px;
  border-radius: 3px;
  background-color: var(--gray-9);
  position: sticky;
  top: 60px;
  z-index: 2; /* Ensure the header is above the table body */
}

tr {
  font-size: 16px;
  height: 36px;
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
  background-color: var(--gray-9);
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
    background-color: var(--purple-8);
    border: none;
    border-radius: 5px;
}

button:hover{
    border:  1px solid #a0a0a0;
    width: 35px;
    height: 35px;
    transition: 50ms;
}
</style>

