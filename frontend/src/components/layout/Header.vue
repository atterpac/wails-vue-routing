
<script setup lang="ts">
import { ref } from 'vue';
import SearchIcon from '../../assets/svg/SearchIcon.vue';
import SwatchIcon from '../../assets/svg/SwatchIcon.vue';
import CloseIcon from '../../assets/svg/CloseIcon.vue';
import MaximizeIcon from '../../assets/svg/MaximizeIcon.vue';
import MinimizeIcon from '../../assets/svg/MinimizeIcon.vue';
import SpotlightIcon from '../../assets/svg/SpotlightIcon.vue';
import CogIcon from '../../assets/svg/CogIcon.vue';
const search = ref("")
</script>

<template>
  <div class="header notselect">
    <div class="left">
        <div class="action"><CogIcon width="24px" height="24px" stroke="#adadad" fill="none" /> </div>
    </div>
    <div class="search notselect">
       <input class="search-bar notselect" placeholder="Search..." v-model="search"/>      
       <SearchIcon width="24px" height="24px" stroke="#6d6d6d" fill="none"/>
    </div>
    <div class="right">
        <div class="action"> <SwatchIcon width="24px" height="24px" stroke="#adadad" fill="none" /> </div>
       <div class="action"> <SpotlightIcon width="24px" height="24px" stroke="#adadad" fill="none" /> </div>
        <div class="window-actions">
            <div class="action"><MinimizeIcon width="18px" height="18px" stroke="#adadad" fill="#adadad" /></div>
            <div class="action"><MaximizeIcon width="18px" height="18px" stroke="#adadad" fill="adadad" /></div>
            <div class="action"><CloseIcon width="18px" height="18px" stroke="#adadad" fill="adadad" /></div>
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
    padding-left: 10px;
    padding-right: 10px;
    border: 2px none #838383;
    background-color: #0d0d0d;
    justify-self: center;
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
    border-radius: 5px;
}

.search:focus-within {
    border: 2px solid #838383;
    border-radius: 5px;
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

.left {
    flex: 1;
    display: flex;
    height: 44px;
    align-items: center;
    margin-right: auto;
    justify-content: space-between;
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

.right .action:last-child:hover {
    background-color: #b22222;
}

.action:hover{
    background-color: #2d2d2d;
}

</style>
