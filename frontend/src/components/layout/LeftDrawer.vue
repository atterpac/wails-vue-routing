<template>
<div class="drawer-container noselect">
 <div class="drawer-left" :style="{ width: props.isCollapse ?  '24px' : '160px' }"> 
     <div class="nav" :style="{ width: props.isCollapse ? '24px' : '160px'}">
         <div class="nav-item" v-for="item in navItems" :style="{ width: props.isCollapse ? '36px' : '160px' }" @click="onClick(item.path)" :class="{'is-active': $route.path === item.path }">
        <Component :is=item.icon fill="none" :style="{ stroke: iconColor}" width="24px" height="24px" />
                <p v-if="!props.isCollapse">{{item.label}}</p>
            </div>
        </div>
     </div>
</div>
</template>

<style>
.drawer-container {
    display: flex;
    margin-right: 10px;
}
.drawer-left {
    padding: 10px; /* Add padding for spacing */
}

.nav {
    display: flex;
    width: 100%;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 5px;
    margin-left: 5px;
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
    margin: 5px;
    margin-top: 10px;
    padding: 5px;
}


.nav-item:hover {
    background-color: var(--gray-8);
}

.is-active {
    background-color: var(--gray-9);
}
</style>

<script setup lang="ts">
import { ref } from 'vue';
import type { Ref } from 'vue';
import { router } from '../../router';
import type { Component } from 'vue';
import HomeIcon from '../../assets/svg/HomeIcon.vue';
import CueListIcon from '../../assets/svg/CueListIcon.vue'
import LightBulbIcon from '../../assets/svg/LightBulbIcon.vue'
import TargetIcon from '../../assets/svg/TargetIcon.vue'
import { RouteLocation, RouteLocationPathRaw, RouteLocationRaw } from 'vue-router';

interface NavItem {
    label: String,
    icon: Component,
    path: RouteLocationRaw,
}

const props = defineProps({
    isCollapse: Boolean
})


const onClick = (route: RouteLocationRaw) => {
    router.push(route)
}

const navItems: NavItem[] = [
{label: "Home", icon: HomeIcon, path:"/"},
{label: "Cues", icon: CueListIcon, path:"/cues"},
{label: "Fixtures", icon: LightBulbIcon, path:"/fixtures"},
{label: "Followspots", icon: TargetIcon, path:"/followspots"},

]

const iconColor: Ref<String> = ref("var(--primary-accent)")

const isCollapse: Ref<boolean> = ref(true);
// Resize 
const currentWidth = ref(80);
</script>
