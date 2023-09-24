<template>
    <div class="fssetup">
        <div class='select-tray'>
            <div class="setting-select">
                <button class='tab-option active' id='operators' @click='selectTab'>Operators</button>
                <button class='tab-option' id='followspots' @click='selectTab'>Followspots</button>
                <button class='tab-option' id='targets' @click='selectTab'>Targets</button>
            </div>
        </div>
        <div class="settings" >
            Settings
            <div v-if='activeTab == "operators"'>
                <div>
                    <input type="text" placeholder="Operator Name">
                    <button>Add Operator</button>
                </div>
            </div>
            <div v-if='activeTab == "followspots"'>
                <div class='setting-list'>
                    <div class='setting'>
                        <label for="type">Fixture</label>
                        <input type="text" placeholder="Fixture Type" id='type'>
                    </div>
                    <div class='setting'>
                        <label for='framecount'>Number of Frames</label>
                        <input  class='framecount' v-model='numOfFrames' type="number" min='0' max='20' placeholder="Number of Frames" id='framecount'>
                        <span @click='frameCollapse = !frameCollapse'>{{frameCollapse ? '-' : '+'}}</span>
                    </div>
                    <div class='frames' v-if='frameCollapse'>
                        <div class='setting frame'>
                            <label for="open">Open Frame</label>
                            <input class='colorpicker' type="color" id='frame0' value='#ffffff'>
                            <input type="text" placeholder="Open" id='frame0' value='Open'>
                        </div>
                        <div class='setting frame' v-for='frame in numOfFrames'>
                            <label for="frame{{frame}}">Frame {{frame}}</label>
                            <input class='colorpicker' type="color" id='frame{{frame}}'>
                            <input type="text" placeholder="Frame Label" id='frame{{frame}}'>
                        </div>
                    </div>
                </div>
            </div>
            <div v-if='activeTab == "targets"'>
                Targets Tab
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, Ref } from 'vue'
const activeTab = ref("operators");
const pureColor = ref<ColorInputWithoutInstance>("red");
const gradientColor = ref("linear-gradient(0deg, rgba(0, 0, 0, 1) 0%, rgba(0, 0, 0, 1) 100%)");
const numOfFrames = ref(0);
const frameCollapse = ref(false);

const selectTab = (e:Event) => {
    const tab = e.target as HTMLElement
    document.querySelectorAll('.tab-option').forEach(tab => {
        tab.classList.remove('active')
    })
    tab.classList.add('active')
    activeTab.value = tab.id
}

</script>

<style scoped>
.fsssetup {
    width: 100%;
    height: 100%;
    background-color: var(--gray-1);
    display: flex;
    flex-direction: row;
    align-content: flex-start;
}

.select-tray {
    width: 200px;
    height: 100%;
    display: flex;
    flex-direction: column;
    border-right: 1px solid var(--gray-9);
}

.setting-select {
    width: 150px;
    display: flex;
    flex-direction: column;
    margin-right: auto;
    flex: 1;
    padding: 0px 10px;
    border-right: 1px solid var(--gray-7);
}


.setting-select button {
    border: none;
    outline: none;
    width: 100%;
    cursor: pointer;
    background-color: var(--gray-9);
    font-size: 1rem;
    color: var(--gray-1);
    padding: 14px 16px;
    margin: 3px 0px;
}

.setting-select button:hover, .setting-select button.active {
    background-color: var(--gray-b);
    border-radius: 5px;
}

.settings {
    width: 100%;
    max-height: 90vh;
    overflow-y: auto;
}

.frame {
    margin-left: 20px;
}

.framecount {
    margin-left: auto;
    margin-right: 10px;
}


.folowspots {
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    overflow: hidden;
}

.setting {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    border-bottom: 1.5px solid var(--gray-7);
    padding: 15px;
    font-size: 18px;
    overflow: hidden;
}

.setting-list {
    overflow: hidden;
}


.setting > input {
    padding: 5px 10px;
    background-color: var(--gray-b);
    border: 1px solid var(--gray-7);
    color: var(--gray-1);
    border-radius: 5px;
    outline: none;
    font-size: 1rem;
}

.colorpicker {
    margin-left: auto;
    margin-right: 15px;
}

input:focus {
    border: 2px solid var(--primary-accent);
    /* Add Box shadow */
    box-shadow: 0 0 5px var(--primary-accent);
    color: var(--gray-1);
}

input[type="color"] {
	-webkit-appearance: none;
	border: none;
    padding: 0;
	width: 48px;
	height: 32px;
}
input[type="color"]::-webkit-color-swatch-wrapper {
	padding: 0;
}
input[type="color"]::-webkit-color-swatch {
	border: none;
}
</style>
