<template>
    <div class="fs-cue">
        <div class='top'>
            <div class='cue-icon'>
                <img class='shrek' src="https://t.ly/r56Sf"/>
                <input type='text' placeholder="Target" />
            </div>
            <div class='info'>
                <div class='data'>
                    <InfoIcon class='data-icon' width='24px' height='24px' stroke='var(--primary-accent)'/>
                    <input type='text' placeholder="Description" />
                </div>
                <div class='data'>
                    <RightArrowIcon class='data-icon' width='24px' height='24px' stroke='var(--primary-accent)'/>
                    <input type='text' placeholder="Placement" />
                </div>
            </div>
        </div>
        <div class='bottom'>
            <div class='info-stack time'>
                <TimerIcon class='time-icon' width='48px' height='48px' stroke='var(--gray-7)' fill='var(--gray-b)'/>
                <input type="number" placeholder="0s" />
            </div>
            <div class='info-stack focus'>
                <PersonIcon class='data-icon' beam_color="#e3b024" :beam_size="beam_size" width='72px' height='72px' stroke='var(--gray-7)' fill='var(--gray-b)' />
                <Dropdown class='dropdown' bg='var(--bg-dark)' :options="DropdownItems" height='28px' @item='updateSelection' />
            </div>
            <div class='frame'>
                <SpotFrame />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import RightArrowIcon from '../../assets/svg/RightArrow.vue'
import InfoIcon from '../../assets/svg/InfoIcon.vue'
import TimerIcon from '../../assets/svg/TimerIcon.vue'
import PersonIcon from '../../assets/svg/PersonIcon.vue'
import Dropdown from '../../components/common/Dropdown.vue'
import SpotFrame from './SpotFrames.vue'
import { ref } from 'vue'

const beam_size = ref("headshot")

type Option = {
    value: string
    label: string
    checked?: boolean
}

const DropdownItems: Option[] = [
    {label: "Wide", value: "wide"},
    {label: "Tight", value: "tight"},
    {label: "3/4 Body", value: "3-4"},
    {label: "Full Body", value: "full-body"},
    {label: "Half Body", value: "half-body"},
    {label: "Head/Shoulder", value: "shoulder"},
    {label: "Headshot", value: "headshot", checked: true},
]

const updateSelection = (item: string) => {
    console.log(item)
    beam_size.value = item
}
</script>

<style scoped>
.fs-cue {
    height: 250px; 
    min-width: 250px;
    width: clamp(200px, 500px, 1000px);
    background-color: var(--bg-dark);
    border: 1px solid var(--gray-8);
    border-radius: 10px;
    margin: 20px;
    padding: 5px;
    display: flex;
    flex-direction: column;
    justify-content: space-around;
}

.top {
    display: flex;
    flex-direction: row;
    align-items: center;
}

.cue-icon {
    margin-left: 20px;
    margin-right: -20px;
    display: flex;
    flex-direction: column;
    align-items: center;
}

.shrek {
    width: 90px;
    height: 90px;
    margin: 10px;
    padding: 10px;
    border-radius: 5px;
}

.info {
    display: flex;
    flex-direction: column;
    flex: 4;
    justify-content: space-around;
    align-items: center;
}

.info-stack {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    margin: 3px;
}

.data {
    display: flex;
    align-items: center;
    width: 80%;
    margin: 3px;
}

.cue-icon > input {
    height: 24px;
    width: 100px;
    font-size: 16px;
    border: none;
    background-color: var(--bg-dark);
    border: 1px solid var(--gray-7);
    border-radius: 8px;
    padding: 5px;
    text-align: center;
    color: var(--gray-4);
    flex: 8;
}

.data > input {
    height: 24px;
    background-color: var(--bg-dark);
    border-radius: 8px;
    border: 1px solid var(--gray-7);
    padding: 5px;
    text-align: center;
    flex: 8;
}

.time {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}

.time-icon {
    margin: 7px;
    align-items: center;
}


.time > input {
    height: 24px;
    width: 48px;
    font-size: 16px;
    border: none;
    background-color: var(--bg-dark);
    border-radius: 8px;
    text-align: center;
    color: var(--gray-4);
    flex: 8;
}

.time > input:focus {
    outline: none;
    border: 2px solid var(--primary-accent);
}

.time > input::placeholder {
    color: var(--gray-3);
}

.data-icon {
    flex: 1;
}


.bottom {
    display: flex;
    flex-direction: row;
    justify-content: space-around;
    padding: 5px 5px;
}

.focus {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100px;
}


p {
    margin-top: -3px;
}

/* Chrome, Safari, Edge, Opera */
input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
  -webkit-appearance: none;
  margin: 0;
}

/* Firefox */
input[type=number] {
  -moz-appearance: textfield;
}
</style>
