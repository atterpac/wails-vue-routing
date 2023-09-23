<template>
    <div class='frame-wrapper'>
        <div class='frames'>
            <div class="frame" :id='frame.id' v-for="frame in frames" :style="{backgroundColor: frame.color}" @click='selectFrame(frame)'>
                {{frame.id}}
            </div>
        </div>
        <div>
            <p> {{activeFrame}}</p>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

type SpotFrames = {
    id: number,
    label: string,
    color: string,
}

const activeFrame = ref("Open")

const frames: SpotFrames[] = [
{id: 1, label: "Red", color: "red"},
{id: 2, label: "Blue", color: "blue"},
{id: 3, label: "Green", color: "green"},
{id: 4, label: "Yellow", color: "yellow"},
{id: 5, label: "Purple", color: "purple"},
]

const selectFrame = (frame: SpotFrames) => {
    const frames = document.querySelectorAll('.frame')
    for (let i = 0; i < frames.length; i++) {
        if (frames[i].id == frame.id) {
            // Set class as active
            frames[i].classList.add('active')
            activeFrame.value = frame.label
        } else {
            frames[i].classList.remove('active')
        }
    }
    console.log(frame)
}
</script>

<style scoped>
.frame-wrapper {
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
}

.frames {
    width: 150px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
}

.frame {
    width: 20px;
    height: 60px;
    flex: 1;
    margin: 5px;
    border: 1px solid var(--gray-8);
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 5px;
    opacity: 0.1;
}

.frame.active {
    opacity: 1;
    flex: 6;
    border: 2px solid var(--primary-accent);
    margin: 3px;
}

.frame:hover {
    flex: 6;
    border-color: var(--primary-accent);
    opacity: 1;
}

p {
    margin: 0;
    padding: 0;
    font-size: 16px;
    color: var(--gray-4);
}
</style>
