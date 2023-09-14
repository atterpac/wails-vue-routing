<script lang="ts" setup>
import { RouterView } from 'vue-router';
import LeftDrawer from './components/layout/LeftDrawer.vue'
import Header from './components/layout/Header.vue'
import RightDrawer from './components/layout/RightDrawer.vue'
import Footer from './components/layout/Footer.vue'
</script>

<template>
<div class="app">
    <Header />
    <div class="container noselect">
        <LeftDrawer />
        <div class="view-container">
            <div class="routerview">
              <router-view />
            </div>
        </div>
        <RightDrawer /> 
    </div>
    <Footer />
</div>
</template>

<style>
  .app {
    width: 100vw;
    height: 100vh;
    display: flex;
    flex-direction: column;
    background-color: #151515;
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
   }

  .routerview {
    background-color: #0d0d0d; /* Adjust background color as needed */
    border-radius: 8px;
    border: 0.25px solid #131313;
    height: 97%;
    margin: 8px;
    align-items: stretch;
    display: flex;
    justify-content: space-between; /* Add space between left and right drawers */
    box-shadow: 7px 7px 10px rgba(0, 0, 0, 0.2);
  }
</style>
