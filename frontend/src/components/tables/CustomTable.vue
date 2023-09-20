<template>
  <div class="table">
    <div class="table-container" style="max-height: props.maxHeight + vh">
      <div class="header">
            <h1>{{props.title}}</h1>
            <button>+</button>
          </div>
      <table>
      <thead>
          <tr class="table-head">
            <template v-for="column, idx in props.columns">
              <th 
                v-if="column.visible"
                style="{width: column.width + 'px'}"  
              >
                  {{column.label}}
              </th>
          </template>
            <th>
              <CogIcon width="24px" height="24px" :style="{ stroke: 'var(--gray-4)', fill: 'var(--gray-4)'}" fill="none" />
            </th>
          </tr>
        </thead>
        <tbody>
          <tr class="data" v-for="data, idx in tableData" :key="idx"  :data-row="idx">
            <td class="data-cell" v-for="value in dataKeys" :data-col="value" :data-row="idx" @click="clickedInput">
              {{( data[value] )}}
            </td>
            <td class="options" @click="clickedOptions" data-col="actions" :data-row="idx">
              <OptionsIcon class="option-icon" width="24px" height="24px"/>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="option-box" v-if="isOptionShown" :style="{top: optionPosition.y, left: optionPosition.x}">
      I am floating box
    </div>
    <div class="input-box" v-if="isInputShown" :style="{top: inputPosition.y, left: inputPosition.x, width: inputPosition.width, height: inputPosition.height}">
      <input v-model="inputData" class="floating-input"/>
    </div>
  </div> 
</template>

<script setup lang="ts">
import { Column } from './column';
import { computed, ref } from 'vue';
import OptionsIcon from '../../assets/svg/OptionsIcon.vue'
import CogIcon from '../../assets/svg/CogIcon.vue';
import { is, isOpaqueType } from '@babel/types';

const props = defineProps({
    title: String,
    maxHeight: Number,
    columns: {
      type: Array<Column>,
        required: true,
    },
    tableData: {
      type: Array<any>,
      required: true,
    },
}) 

const isOptionShown = ref(false)

const optionPosition = ref({
    x: '0px',
    y: '0px'
})

const clickedOptions = (event: Event) => {
  isOptionShown.value = true;
  const target = event.target as HTMLElement
  const el = target.parentElement as HTMLElement
  const position = parseFloat(el.getAttribute('data-row') as string)
  // Query data-col actions for nth child of row
  const query = `tr:nth-child(${position+1})`
  const row = document.querySelector(query)
  row?.querySelector('.option-icon')?.classList.add('active')
  
  console.log(query)
  const coords = el.getBoundingClientRect();
  console.log(coords, isOptionShown.value)

  const shiftX = coords.x - 150;
  const shiftY = coords.y 
  optionPosition.value.x = `${shiftX}px`
  optionPosition.value.y = `${shiftY}px`

  // Check if there is a click outside of the option box
  document.addEventListener('click', (event) => {
    const el = event.target as HTMLElement
    if (el.classList.contains('option-icon')) {
      return
    }
    document.removeEventListener('click', () => {
      console.log('removed')
    })
    isOptionShown.value = false;
    row?.querySelector('.option-icon')?.classList.remove('active')

  }) 

  document.querySelector('.table-wrapper')?.addEventListener('scroll', () => {
    console.log('scrolling')
    const coords = el.getBoundingClientRect();
    const boundry = document.querySelector('.table-wrapper')?.getBoundingClientRect();
    console.log(coords, boundry)
    if (boundry && ((boundry.top+100) > coords.top || (boundry.bottom) < coords.bottom)) {
      isOptionShown.value = false;
      row?.querySelector('.option-icon')?.classList.remove('active')
      return
    } 
    const shiftX = coords.x - 150
    const shiftY = coords.y - 1
    optionPosition.value.x = `${shiftX}px`
    optionPosition.value.y = `${shiftY}px`
  })
}

const isInputShown = ref(false)

const inputData = ref('')

const inputPosition = ref({
    x: '0px',
    y: '0px',
    width: '0px',
    height: '0px'
})


const clickedInput = (event: Event) => {
 
  isInputShown.value = true;
  const el = event.target as HTMLElement

  const position = parseFloat(el.getAttribute('data-row') as string)
  inputData.value = el.innerText
  // Change coords to computed
  const coords = el.getBoundingClientRect();

  // const coords = el.getBoundingClientRect();
  console.log(el, coords, isInputShown.value)
  const shiftX = coords.x 
  const shiftY = coords.y - 1
  const shiftWidth = coords.width 
  const shiftHeight = coords.height
  inputPosition.value.x = `${shiftX}px`
  inputPosition.value.y = `${shiftY}px`
  inputPosition.value.width = `${shiftWidth}px`
  inputPosition.value.height = `${shiftHeight}px`
  // focus on input

  // Add Scroll listener to update inputPosition based on elements new position
  document.querySelector('.table-wrapper')?.addEventListener('scroll', () => {
    console.log('scrolling')
    const coords = el.getBoundingClientRect();
    const boundry = document.querySelector('.table-wrapper')?.getBoundingClientRect();
    console.log(coords, boundry)
    if (boundry?.top && 
        (boundry?.top > coords.top ||
        boundry?.bottom < coords.bottom)) {
      isInputShown.value = false;
      return
    } else {
      isInputShown.value = true;
    } 
    const shiftX = coords.x
    const shiftY = coords.y - 1
    inputPosition.value.x = `${shiftX}px`
    inputPosition.value.y = `${shiftY}px`
  })


  // // Change inner HTML to input data when enter is hit
  // textInput.addEventListener('keyup', (event) => {
  //   console.log("Submitted")
  //   if ((event.key === 'enter' || event.key === 'Escape')&& el) {
  //     console.log("Submitted")
  //     el.innerText = input.value
  //     isInputShown.value = false;
  //     input.removeEventListener('keyup', () => {
  //       console.log('removed')
  //     })
  //   }
  // })


  // Check if there is a click outside of the option box
  document.addEventListener('click', (event) => {
    const el = event.target as HTMLElement
    console.log("clicked", el.classList)
    if (el.classList.contains('.floating-input')) {
      return
    }
    document.removeEventListener('click', () => {
      isInputShown.value = false;
      console.log('removed')
      
    })
  }) 
}

const toggleOption = () => {
    isOptionShown.value = !isOptionShown.value
}

const dataKeys = computed(() => {
  let keys = props.tableData[0] ? props.columns.filter((column) => column.visible === true).map((column) => column.data) : [];
  return keys as string[];
});
</script>

<style scoped>
.table {
  width: 100%;
  height: 90%;
  padding: 0px 20px 20px 20px;
}

.header {
    position: sticky;
    width: 100%;
    background-color: var(--bg-dark);
    height: 40px;
    width: 100%;
    position: sticky;
    top: 0;
    z-index: 2;
    border-radius: 5px;
    margin-left: -5px;
}

table {
  width: 100%;
  text-align: center;
  border-collapse: collapse;
}

thead {
  padding-top: 40px;
  border-radius: 3px;
  background-color: var(--bg-darker);
  position: sticky;
  top: 60px;
  z-index: 2; /* Ensure the header is above the table body */
}

.table-head {
  height: 50px;
}

tr {
  font-size: 16px;
  height: 36px;
  font-weight: 500;
  align-items: center;
}


th:first-child {
  border-top-left-radius: 5px;
  border-bottom-left-radius: 5px;
}

th:last-child {
  border-top-right-radius: 5px;
  border-bottom-right-radius: 5px ;
}
.data:nth-child(odd) td:first-child {
  border-top-left-radius: 5px;
  border-bottom-left-radius: 5px;
}

.data:nth-child(even) {
  background-color: var(--bg-darker);
}

.data:nth-child(even) td:first-child {
  border-bottom-left-radius: 5px;
  border-top-left-radius: 5px;
}

.data:nth-child(even) td:last-child {
  border-bottom-right-radius: 5px;
  border-top-right-radius: 5px;
}


h1 {
    margin-left: 30px;
    float: left;
}

button {
    float: right;
    margin-right: 20px;
    width: 30px;
    height: 30px;
    font-size: 18px;
    font-weight: 700;
    background-color: var(--primary-accent);
    border: none;
    border-radius: 5px;
}

button:hover{
    border:  1px solid #a0a0a0;
    width: 35px;
    height: 35px;
    transition: 50ms;
}

.options {
  align-items: center;
  justify-content: flex-end;
}

.option-icon {
  margin-bottom: -2px;
  stroke: var(--gray-6);
  fill: var(--gray-6);
}

.option-icon.active {
  stroke: var(--primary-accent);
  fill: var(--primary-accent);
}

.option-icon:hover {
  stroke: var(--primary-accent);
  fill: var(--primary-accent);
}

.option-box {
  background-color: var(--bg-darker);
  border: 2px solid var(--primary-accent-dark);
  box-shadow: 0px 0px 8px 0px var(--primary-accent-dark);
  padding: 10px;
  position: absolute;
  z-index: 2;
  width: 120px;
  height: 200px;
  border-radius: 10px;
}

.input-box {
  background-color: var(--bg-darker);
  border: 2px solid var(--primary-accent-dark);
  box-shadow: 0px 0px 8px 0px var(--primary-accent-dark);
  position: absolute;
  z-index: 1;
  border-radius: 8px;
  align-items: center;
  justify-content: center;
}

.floating-input {
  background-color: var(--bg-darker);
  border: none;
  width: 90%;
  height: 90%;
  margin: 2px;
  font-size: 16px;
  border-radius: 3px;
  outline: none;
  text-align: center;
  color: var(--gray-4);
}
</style>

