<template>
<div class="drawer-container noselect">
 <div class="drawer-left" :style="{ width: isCollapse ? '160px' : '36px' }"> 
    <div class="nav">
        <div class="nav-item" v-for="item in navItems" :style="{ width: isCollapse ? '160px' : '36px' }" @click="onClick(item.path)">
        <Component :is=item.icon fill="none" stroke="#409292" width="24px" height="24px" />
                <p v-if="isCollapse">{{item.label}}</p>
            </div>
        </div>
     </div>
 <div class="resize" @mousedown="startResize"></div>
</div>
</template>

<style>
.drawer-container {
    display: flex;
}
  .drawer-left {
    padding: 10px; /* Add padding for spacing */
  }

  .resize {
    width: 5px;
    border-radius: 3px;
    background-color: #3d3d3d;
    cursor: ew-resize;
    display: none;
  }

  .drawer-container:hover .resize{
    display: block;
    resize: both;
}

.nav {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;
}

.nav-item {
    display: flex;
    flex-direction: row;
    font-size: 16px;
    color: #fbfbfb;
    align-items: center;
    justify-content: flex-start;
    height: 35px;
    border-radius: 10px;
    margin: 2px;
    margin-top: 10px;
}


.nav-item:hover {
    background-color: #4d4d4d;
}
</style>

<script setup lang="ts">
import { ref } from 'vue';
import type { Ref } from 'vue';
import { router } from '../../router';
import type { Component } from 'vue';
import { onMounted, onUnmounted } from 'vue';
import HomeIcon from '../../assets/svg/HomeIcon.vue';
import CueListIcon from '../../assets/svg/CueListIcon.vue'
import LightBulbIcon from '../../assets/svg/LightBulbIcon.vue'
import TargetIcon from '../../assets/svg/TargetIcon.vue'

interface NavItem {
    label: String,
    icon: Component,
    path: String,
}


const onClick = (route: RouteLocationRaw) => {
    router.push(route)
}

const navItems: NavItem[] = [
{label: "Home", icon: HomeIcon, path:"/"},
{label: "Cues", icon: CueListIcon, path:"/cues"},
{label: "Fixtures", icon: LightBulbIcon, path:"/fixtures"},
{label: "Followspots", icon: TargetIcon, path:"/followspots"},

]

const isCollapse: Ref<boolean> = ref(true);
// Resize 
const currentWidth = ref(80);
let isResizing = ref(false);
let startX = 0;
let startWidth = 0;

const startResize = (event) => {
    isResizing.value = true;
    startX = event.clientX;
    startWidth = currentWidth.value;

    document.addEventListener('mousemove', resize);
    document.addEventListener('mouseup', stopResize);
};

const resize = (event) => {
    if (!isResizing.value) return;

    const newWidth = startWidth + event.clientX - startX;

    // Limit the minimum and maximum width (adjust these values as needed)
    isCollapse.value = newWidth > startX ;
};

const stopResize = () => {
    isResizing.value = false;
    document.removeEventListener('mousemove', resize);
    document.removeEventListener('mouseup', stopResize);
};


onMounted(() => {
        document.addEventListener('mouseup', stopResize);
        });

onUnmounted(() => {
        document.removeEventListener('mouseup', stopResize);
        });
</script>
