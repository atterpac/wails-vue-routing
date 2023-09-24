<template>
    <section class="followspot-view">
       <div class="followspot" v-for="followspot in followspots">
           <h3>{{ followspot.name }}</h3>
            <FollowSpot :id=followspot.id />
       </div>
    </section>
</template>

<script setup lang='ts'>
import FollowSpot from '../components/followspot/FollowSpotCue.vue'

type followspot = {
    id: string,
    name: string,
    beam_size?: string,
    beam_color?: string,
}

const followspots: followspot[] = [
{
   id: '1', 
   name: 'Followspot 1',
},
{
   id: '2', 
   name: 'Followspot 2',
},
{
   id: '2', 
   name: 'Followspot 3',
}
]

</script>

<style>

.followspot-view {
    width: 100%;
    height: 100%;
    display: flex;
    overflow: hidden;
    flex-direction: row;
    justify-content: space-around;
    align-content: center;
}

.followspot {
    width: 33%;
}
</style>
