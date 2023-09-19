
<script setup lang="ts">
import { ref } from 'vue';
import SearchIcon from '../../assets/svg/SearchIcon.vue';
import SwatchIcon from '../../assets/svg/SwatchIcon.vue';
import CloseIcon from '../../assets/svg/CloseIcon.vue';
import MaximizeIcon from '../../assets/svg/MaximizeIcon.vue';
import MinimizeIcon from '../../assets/svg/MinimizeIcon.vue';
import SpotlightIcon from '../../assets/svg/SpotlightIcon.vue';
import CogIcon from '../../assets/svg/CogIcon.vue';
import SideNavIcon from '../../assets/svg/SideNavIcon.vue'
const search = ref("")

const iconColor = ref<String>("var(--gray-3)")
</script>

<template>
  <div class="header notselect" style="widows: 1;">
    <div class="left">
        <div class="action"><CogIcon width="24px" height="24px" :style="{ stroke: iconColor, fill: iconColor}" fill="none" /> </div>
        <div class="action" @click="$emit('navCollapse')"><SideNavIcon width="24px" height="24px" :style="{ stroke: iconColor, fill: iconColor}" fill="none" /> </div>
    </div>
    <div class="search notselect" style="--wails-draggable:no-drag">
       <input class="search-bar notselect" placeholder="Search..." v-model="search"/>      
       <SearchIcon class="search-icon" width="24px" height="24px" :style="{ stroke: iconColor, fill: iconColor}" fill="none"/>
    </div>   
    <div class="right">
        <div class="right-actions">
            <div class="action"> <SwatchIcon width="24px" height="24px" :style="{ stroke: iconColor}" fill="none" /> </div>
            <div class="action " @click="$emit('drawerCollapse')"> <SpotlightIcon width="24px" height="24px" :style="{ stroke: iconColor, fill: iconColor}" /> </div>
        </div>       
        <div class="window-actions">
            <div class="action"><MinimizeIcon width="18px" height="18px" :style="{ stroke: iconColor, fill: iconColor}" /></div>
            <div class="action"><MaximizeIcon width="18px" height="18px" :style="{ stroke: iconColor}" /></div>
            <div class="action cancel"><CloseIcon width="18px" height="18px" :style="{ stroke: iconColor, fill: iconColor}" /></div>
        </div>
    </div>
  </div>
</template>

<style>
  .header {
    height: 24px; /* Adjust height as needed */
    display: flex;
    align-items: center; /* Center content vertically */
    justify-content: space-between; /* Add space between header elements */
    padding: 10px; /* Add padding for spacing */
  }

  .search-bar {
    padding-left: 20px;
    padding-right: 10px;
    border: 2px none #838383;
    border-top-left-radius: 10px;
    border-bottom-left-radius: 10px;
    justify-self: center;
    background-color: var(--gray-a);
    height: 30px;
    
    outline: none;
    width: 325px;
    color: #D3D3D3;
    flex: 1;
}

.search {
    display: flex;
    align-items: center;
    border: 0.5px solid #232323;
    background-color: #0d0d0d;
    border-radius: 10px;
    margin-top: 10px;
}

.search:focus-within {
    widows: 0;
    border: 1.5px solid var(--purple-8);
    border-radius: 10px;
    background-color:var(--purple-8);
}

.right {
    flex: 1;
    width: 300px;
    display: flex;
    height: 44px;
    align-items: center;
    margin-left: auto;
    justify-content: space-between;
}

.right-actions {
    display: flex;
    align-items: center;
    margin-left: 20%;
}

.left {
    flex: 1;
    display: flex;
    height: 44px;
    align-items: left;
    margin-right: auto;
    justify-content: left;
}

.window-actions {
    height: 44px;
    width: 150px;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.action {
    width: 44px;
    height: 44px;
    margin-top: 5px;
    display: flex;
    justify-content: center;
    align-items: center;
    border-radius: 5px;
}

.action:last-child {
    margin-right: -10px;
}

.right .cancel:last-child:hover {
    background-color: #b22222;
}

.action:hover{
    background-color: #2d2d2d;
}

</style>
