<template>
<div class="app">
    <Header style="widows: 1" @navCollapse="toggleCollapse" @drawerCollapse="toggleDrawer"/>
    <div class="container noselect">
        <LeftDrawer :isCollapse="isNavCollapse ? true : false"/>
            <div class="routerview noselect">
              <router-view />
            </div>
        <RightDrawer class='right-drawer' :isCollapse="isDrawerCollapse ? true : false" v-if="!isDrawerCollapse"/> 
    </div>
    <Footer v-if="false"/>
</div>
</template>

<script lang="ts" setup>
import { RouterView } from 'vue-router';
import LeftDrawer from './components/layout/LeftDrawer.vue'
import Header from './components/layout/Header.vue'
import RightDrawer from './components/layout/RightDrawer.vue'
import Footer from './components/layout/Footer.vue'
import { ref } from 'vue';

const isNavCollapse = ref<Boolean>(true)
const isDrawerCollapse = ref<Boolean>(true)

const toggleCollapse = () => {
  isNavCollapse.value = !isNavCollapse.value
}

const toggleDrawer = () => {
  isDrawerCollapse.value = !isDrawerCollapse.value
}
</script>

<style>
.app {
    width: 100vw;
    height: 100vh;
    display: flex;
    flex-direction: column;
    background-color: var(--bg-darker);
    position: absolute;
    top: 0;
    left: 0;
    color: #D3D3D3;
}


.container {
    flex: 1; /* Grow and shrink container */
    display: flex;
    justify-content: space-between; /* Add space between left and right drawers and routerview */
    flex-direction: row;
    overflow: hidden;
}

.routerview {
    background-color: var(--bg-dark-darker);
    padding-top: 0px;
    padding-left: 10px;
    border-radius: 8px;
    display: flex;
    flex: 5;
    box-shadow: 7px 7px 10px rgba(0, 0, 0, 0.2);
    min-width: 300px;
    overflow: hidden;
    margin-right: 10px;
    margin-bottom: 10px;
}

.right-drawer {
    width: 100%;
    max-width: 300px;
    background-color: var(--bg-darker);
    border-radius: 8px;
    height: 100%;
    margin: 0px 10px 0px 10px;
    align-items: stretch;
    display: flex;
    flex: 1;
    justify-content: space-between;
    box-shadow: 7px 7px 10px rgba(0, 0, 0, 0.2); 
}
</style>
