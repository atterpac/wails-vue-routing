<script lang="ts" setup>
import { RouterView } from 'vue-router';
import LeftDrawer from './components/layout/LeftDrawer.vue'
import Header from './components/layout/Header.vue'
import RightDrawer from './components/layout/RightDrawer.vue'
import Footer from './components/layout/Footer.vue'
import { ref } from 'vue';

const isNavCollapse = ref<Boolean>(false)
const isDrawerCollapse = ref<Boolean>(true)

const toggleCollapse = () => {
  isNavCollapse.value = !isNavCollapse.value
}

const toggleDrawer = () => {
  isDrawerCollapse.value = !isDrawerCollapse.value
}
</script>

<template>
<div class="app">
    <Header style="widows: 1" @navCollapse="toggleCollapse" @drawerCollapse="toggleDrawer"/>
    <div class="container noselect">
        <LeftDrawer :isCollapse="isNavCollapse ? true : false"/>
        <div class="view-container noselect">
            <div class="routerview noselect">
              <router-view />
            </div>
        </div>
        <RightDrawer :isCollapse="isDrawerCollapse ? true : false"/> 
    </div>
    <Footer v-if="false"/>
</div>
</template>

<style>
  .app {
    width: 100vw;
    height: 100vh;
    display: flex;
    flex-direction: column;
    background-color: var(--gray-9);
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

  .view-container {
    flex: 1;
    margin: 3px;
   }

  .routerview {
    background-color: var(--gray-8); /* Adjust background color as needed */
    border-radius: 8px;
    border: 0.25px solid #242629;
    height: 97%;
    margin: 10px 10px 10px 0px;
    align-items: stretch;
    display: flex;
    justify-content: space-between; /* Add space between left and right drawers */
    box-shadow: 7px 7px 10px rgba(0, 0, 0, 0.2);
  }
</style>
