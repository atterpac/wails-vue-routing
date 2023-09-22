<template>
    <div class="fs-cue">
        <div class='top'>
            <div class='cue-icon'>
                <img class='shrek' src="https://t.ly/r56Sf"/>
            </div>
            <div class='info'>
                <div class='data'>
                    <TargetIcon class='data-icon' width='24px' height='24px' stroke='var(--primary-accent)'/>
                    <input type='text' placeholder="Target" />
                </div>
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
            <div class='time'>
                <TimerIcon class='data-icon' width='48px' height='48px' stroke='var(--gray-7)' fill='var(--gray-b)'/>
            </div>
            <div class='focus'>
                <PersonIcon class='data-icon' width='48px' height='48px' stroke='var(--gray-7)' fill='var(--gray-b)'/>
            </div>
            <div class='focus'>
                I am Gel Frame
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import TargetIcon from '../../assets/svg/TargetIcon.vue' 
import RightArrowIcon from '../../assets/svg/RightArrow.vue'
import InfoIcon from '../../assets/svg/InfoIcon.vue'
import TimerIcon from '../../assets/svg/TimerIcon.vue'
import PersonIcon from '../../assets/svg/PersonIcon.vue'
</script>

<style scoped>
.fs-cue {
    min-width: 400px;
    width: 33%;
    height: 25%;
    background-color: var(--gray-a);
    border-radius: 10px;
    margin: 10px;
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
}

.shrek {
    width: 90px;
    height: 90px;
    margin: 10px;
    border: 1px solid var(--gray-8);
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

.data {
    display: flex;
    align-items: center;
    width: 80%;
    margin: 3px;
}

.data > input {
    height: 36px;
    background-color: var(--gray-b);
    border: 1px solid var(--gray-8);
    border-radius: 8px;
    text-align: center;
    flex: 8;
}

.data-icon {
    flex: 1;
}

.bottom {
    display: flex;
    flex-direction: row;
    justify-content: space-around;
}
</style>
