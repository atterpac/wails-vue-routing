<template>
<div class="cue-table">
    <table>
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
        <tr class='data' v-for="cue in CueData">
            <td>{{cue.number}}</td>
            <td>{{cue.part}}</td>
            <td>{{cue.time}}</td>
            <td>{{cue.delay}}</td>
            <td>{{cue.hang}}</td>
            <td>{{cue.follow}}</td>
            <td>{{cue.block}}</td>
            <td>{{cue.description}}</td>
            <td>{{cue.placement}}</td>
            <td>{{cue.scene_id}}</td>
            <td>{{cue.notes}}</td>
        </tr>
    </table>
</div>
</template>

<script setup lang="ts">
import { CueData } from './cuedata'
</script>

<style scoped>

.cue-table {
    width: 100%;
    padding: 0px 20px 20px 20px;
    overflow: auto;
}


table {
    width: 100%;
    text-align: center;
    border-collapse: collapse;
}


table th {
    padding-top: 20px;
    background-color:#0d0d0d;
    position: sticky;
    top: 0;
}

tr {
    font-size: 16px;
    height: 36px;
}

.data:nth-child(odd) td:first-child {
    border-top-left-radius: 5px;
    border-bottom-left-radius: 5px;
}

.data:nth-child(even) {
    background-color: #151515;   
}

.data:nth-child(even) td:first-child {
    border-bottom-left-radius: 5px;
    border-top-left-radius: 5px;
}

.data:nth-child(even) td:last-child {
    border-bottom-right-radius: 5px;
    border-top-right-radius: 5px;
}
</style>

