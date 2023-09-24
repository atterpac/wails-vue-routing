<template>
    <div class='frame-wrapper'>
        <div class='frames'>
            <div class="frame" 
                :id='frame.id' 
                v-for="frame in frames" 
                :style="{backgroundColor: frame.color}" 
                @click='selectFrame(frame)' 
                @mouseover='showSpan' 
                @mouseleave='hideSpan'>

                <div class='frame-label'>
                    <a>{{frame.id}}</a> 
                    <a v-if='frame.hovered'>{{frame.label}}</a>
                </div>
            </div>
        </div>
        <div>
            <p> {{activeFrame}}</p>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, Ref } from 'vue'

type SpotFrames = {
    id: string,
    label: string,
    color: string,
    hovered?: boolean
}

const activeFrame = ref("Open")



const frames: Ref<SpotFrames[]> = ref([
{id: "0", label: "Open", color: "", hovered: true},
{id: "1", label: "Red", color: "red", hovered: false},
{id: "2", label: "Blue", color: "blue", hovered: false},
{id: "3", label: "Green", color: "green", hovered: false},
{id: "4", label: "Yellow", color: "yellow", hovered: false},
{id: "5", label: "Purple", color: "purple", hovered: false},

])

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

const hovered = ref(false);

const showSpan = (e: Event) => {
// Get ID of target and set hovered to true
    const id = (e.target as HTMLElement).id
    for (let i = 0; i < frames.value.length; i++) {
        if (frames.value[i].id == id) {
            frames.value[i].hovered = true
        }
        }
};

const hideSpan = (e: Event) => {
    const id = (e.target as HTMLElement).id
    for (let i = 0; i < frames.value.length; i++) {
        if (frames.value[i].id == id) {
            frames.value[i].hovered = false
        }
    }
};
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
    width: 200px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
}

.frame {
    width: 20px;
    height: 60px;
    flex: 1;
    margin: 2px;
    border: 1px solid var(--gray-8);
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 5px;
    opacity: 0.1;
}

.frame-label {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    transition: opacity 0.3s ease-in-out;
    opacity: 1;
    margin: 5px;
}
.frame.active {
    opacity: 1;
    flex: 6;
    border: 2px solid var(--gray-4);
    margin: 3px;
}

.frame:hover {
    flex: 8;
    width: 50px;
    opacity: 1;
      /* Animation */
    transition: flex 0.3s ease-in-out, transform 0.3s ease-in-out;
    transform: scale(1.5);
}

p {
    margin: 0;
    padding: 0;
    font-size: 16px;
    color: var(--gray-4);
}
</style>
