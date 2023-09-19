<template>
<div class="drawer-container" >
 <div class="drawer-right" :style="{ width: props.isCollapse ?  '1px' : '160px' }"> 
    <p v-if="!props.isCollapse">Right Drawer</p>
 </div>
</div>
</template>

<style scoped>

.drawer-container {
    display: flex;
}
.drawer-right {
  max-width: 240px;
  padding: 10px; /* Add padding for spacing */
}

</style>

<script setup lang="ts">
import { ref } from 'vue';
import { onMounted, onUnmounted } from 'vue';

const props = defineProps({
  isCollapse: Boolean
})
// const currentWidth = ref(200);

// let isResizing = ref(false);
// let startX = 0;
// let startWidth = 0;

// const startResize = (event) => {
//   isResizing.value = true;
//   startX = event.clientX;
//   startWidth = currentWidth.value;

//   document.addEventListener('mousemove', resize);
//   document.addEventListener('mouseup', stopResize);
// };

// const resize = (event) => {
//   if (!isResizing.value) return;

//   const newWidth = startWidth - event.clientX + startX;

//   // Limit the minimum and maximum width (adjust these values as needed)
//   currentWidth.value = Math.min(Math.max(newWidth, 5), 600);
// };

// const stopResize = () => {
//   isResizing.value = false;
//   document.removeEventListener('mousemove', resize);
//   document.removeEventListener('mouseup', stopResize);
// };


// onMounted(() => {
//   document.addEventListener('mouseup', stopResize);
// });

// onUnmounted(() => {
//   document.removeEventListener('mouseup', stopResize);
// });
</script>

