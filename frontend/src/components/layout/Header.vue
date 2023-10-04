
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
import { GetUsers } from '../../../wailsjs/go/controllers/App'
const search = ref("")
const spotlightClicked = ref<Boolean>(false)

const emits = defineEmits(["drawerCollapse"])

const spotlight = (e: Event) => {
    emits('drawerCollapse')
    const spotlightElement = document.querySelector('.spotlight')
    if (spotlightElement?.classList.contains('active')) {
        spotlightElement.classList.remove('active') 
    } else {
        spotlightElement?.classList.add('active')
    }
}

const getUsers = () => {
    GetUsers().then((users) => {
        console.log(users)
    })
    console.log("user gotten")
}
const iconColor = ref<String>("var(--gray-4)")
</script>

<template>
  <div class="header noselect" style="widows: 1;">
    <div class="left">
        <div class="action" @click='getUsers'><CogIcon width="24px" height="24px" :style="{ stroke: iconColor, fill: iconColor}" fill="none" /> </div>
        <div class="action" @click="$emit('navCollapse')"><SideNavIcon width="24px" height="24px" :style="{ stroke: iconColor, fill: iconColor}" fill="none" /> </div>
    </div>
    <div class="search noselect" style="widows: 2">
       <input class="search-bar notselect" placeholder="Search..." v-model="search"/>      
       <SearchIcon class="search-icon" width="24px" height="24px" :style="{ stroke: iconColor, fill: iconColor}" fill="none"/>
    </div>   
    <div class="right">
        <div class="right-actions">
            <div class="action"> <SwatchIcon width="24px" height="24px" :style="{ stroke: iconColor}" fill="none" /> </div>
            <div class="action spotlight " @click="spotlight"> <SpotlightIcon width="24px" height="24px" :style="{ stroke: iconColor, fill: iconColor}" /> </div>
        </div>       
        <div class="window-actions">
            <div class="action"><MinimizeIcon width="16px" height="16px" :style="{ stroke: iconColor, fill: iconColor}" /></div>
            <div class="action"><MaximizeIcon width="16px" height="16px" :style="{ stroke: iconColor}" /></div>
            <div class="action cancel"><CloseIcon width="16px" height="16px" :style="{ stroke: iconColor, fill: iconColor}" /></div>
        </div>
    </div>
  </div>
</template>

<style>
  .header {
    height: 20px; /* Adjust height as needed */
    display: flex;
    align-items: center; /* Center content vertically */
    justify-content: space-between; /* Add space between header elements */
    padding: 10px; /* Add padding for spacing */
    margin-bottom: 5px;
  }

  .search-bar {
    padding-left: 20px;
    padding-right: 30px;
    margin-left: 1px;
    border: 2px none #838383;
    border-top-left-radius: 8px;
    border-bottom-left-radius: 8px;
    justify-self: center;
    background-color: var(--gray-b);
    height: 28px;
    
    outline: none;
    width: 325px;
    color: #D3D3D3;
    flex: 1;
}

.search {
    display: flex;
    align-items: center;
    border: 0.5px solid var(--gray-9);
    background-color: var(--gray-c);
    border-radius: 10px;
    height: 30px;
    margin-top: 5px;
}

.search:focus-within {
    widows: 0;
    border: 2px solid var(--primary-accent);
    border-radius: 10px;
    background-color:var(--primary-accent);
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
    align-items: center;
    margin-right: auto;
    justify-content: left;
}

.window-actions {
    height: 30px;
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
    border-radius: 0;
    background-color: #b22222;
}

.action:hover{
    background-color: var(--gray-8)
}

.header-active {
    border: 2px solid var(--primary-accent-dark);
    background-color: var(--gray-8);
}

</style>
