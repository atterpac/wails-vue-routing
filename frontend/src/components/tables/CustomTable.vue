<template>
  <div class="table">
    <div style="max-height: props.maxHeight + vh">
      <div class="header">
            <h1>{{props.title}}</h1>
            <button>+</button>
      </div>
      <div class='table-container'>
          <table>
              <thead>
                  <tr class="table-head">
                    <template v-for="column in props.columns">
                      <th 
                        v-if="column.visible"
                        style="{width: column.width + 'px'}"  
                        :data-col=column.data
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
                    <td class="data-cell" 
                        v-for="value in dataKeys" 
                        :data-col="value" 
                        :data-row="idx" 
                        :data-selected="false" 
                        @click="clickedInput"
                    >
                      {{( data[value] )}}
                    </td>
                    <td class="options" 
                        data-col="actions" 
                        :data-row="idx"
                    >
                      <OptionsIcon 
                        @click="clickedOptions" 
                        :data-row="idx" 
                        class="option-icon" 
                        width="24px" 
                        height="24px"/>
                    </td>
                  </tr>
                </tbody>
          </table>
        </div>
    </div>
    <div class="option-box" v-if="isOptionShown" :style="{top: optionPosition.y, left: optionPosition.x}">
      I am floating box
    </div>
    <div class="input-box" v-if="isInputShown" :style="{top: inputPosition.y, left: inputPosition.x, width: inputPosition.width, height: inputPosition.height}">
        <input v-model="inputData" class="floating-input" id="cell-input"/>
    </div>
  </div> 
</template>

<script setup lang="ts">
import { Column } from './column';
import { computed, ref } from 'vue';
import OptionsIcon from '../../assets/svg/OptionsIcon.vue'
import CogIcon from '../../assets/svg/CogIcon.vue';


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

type Position = {
    x: string,
    y: string,
    top?: string,
    left?: string,
    right?: string,
    bottom?: string,
    width?: string,
    height?: string,
}

type Adjustment = {
    x: number,
    y: number,
}

const isOptionShown = ref(false)

const optionPosition = ref<Position>({
    x: '0px',
    y: '0px'
})

let isClicked = false
const clickedOptions = (event: Event) => {
  // Grab necessary data from event
  const target = event.target as HTMLElement
  const el = target.parentElement as HTMLElement
  const wrapper = document.querySelector('.table-wrapper') as HTMLElement
  if (!wrapper) {
    console.log("Wrapper could not be found");
    return
  }

  // Query data-row actions for nth child of row
  const position = parseFloat(el.getAttribute('data-row') as string) // Get data-row attribute from parent element
  const row = document.querySelector(`tr:nth-child(${position+1})`) as HTMLElement
  if (!row) {
    console.log("Error finding row position");
    return
  }
  row.querySelector('.option-icon')?.classList.add('active')

  // Shift Option panels position to match the clicked element with an X offset
  const bounds = el.getBoundingClientRect();
  optionPosition.value.x =  `${bounds.x - 150}px`
  optionPosition.value.y =  `${bounds.y - 10}px`
  isOptionShown.value = true;

 // Click Listener Setup
  const watchClickListener = () => {
     clickListener(el);
  }

  const clickListener = (el: HTMLElement) => {
    if (!isClicked) { // Check if this is the first click or not so it doesnt automatically trigger the remove
        isClicked = true
        isOptionShown.value = true;
    } else {
        if (el.classList.contains('.option-box')) {
          return
        }
        document.removeEventListener('click', watchClickListener)
        isOptionShown.value = false;
        isClicked = false;
        row.querySelector('.option-icon')?.classList.remove('active')
    }
  }

// Attach Listener
  document.addEventListener('click', watchClickListener)

// Scroll Listener Setup
  const watchScrollFunction = () => {
    watchScroll(el,row, wrapper)
  }

  const watchScroll = (el: HTMLElement, row: HTMLElement, wrapper: HTMLElement) => {
    const boundry = wrapper.getBoundingClientRect();
    const pastBoundry = isPastBoundry(el, boundry, watchScrollFunction, wrapper)
    if (pastBoundry) {
      isOptionShown.value = false;
      row?.querySelector('.option-icon')?.classList.remove('active')
      return
    } 
    const currentPos = optionPosition.value as Position;
    shiftPosition(el, currentPos, 150, 1);
  }
  
// Attach Scroll Listener
  wrapper.addEventListener('scroll', watchScrollFunction, false);
}

const shiftPosition = (match: HTMLElement, element: Position, xShift?: number, yShift?: number, matchPosition?: boolean  ) => {
    console.log("Shifting Position");
    const coords = match.getBoundingClientRect();
    element.x = `${(coords.x - (xShift ? xShift : 0))}px`;
    element.y = `${(coords.y - (yShift ? yShift : 0))}px`;
    if (matchPosition) {
        element.width = `${coords.width}px`
        element.height = `${coords.height}px`
    }
    
}

const RemoveListener = (el: HTMLElement, event: string, listenFunc: EventListener) => {
    el.removeEventListener(event, listenFunc)
}

const isPastBoundry = (el: HTMLElement, boundry: DOMRect, listenFunc: EventListener, listenEl: HTMLElement) => {
    const cords = el.getBoundingClientRect();
    if ((boundry.top > cords.top || boundry.bottom < cords.bottom)) {
      RemoveListener(listenEl,'scroll', listenFunc)
      isInputShown.value = false;
      return true
    }
    return false
} 

const watchScroll = (el: HTMLElement, boundryEl: HTMLElement, listenFunc: EventListener, position: Position, adjustment?: Adjustment) => {
    const coords = el.getBoundingClientRect();
    const limit = boundryEl.getBoundingClientRect();

    isPastBoundry(el, limit, listenFunc, boundryEl)
    position.x = `${coords.x - (adjustment ? adjustment.x : 0)}px`
    position.y = `${(coords.y - (adjustment ? adjustment.y : 0))}px`
}

const isInputShown = ref(false)

const inputData = ref('')

const inputPosition = ref<Position>({
    x: '0px',
    y: '0px',
    width: '0px',
    height: '0px'
})


let isInputClicked = false
const clickedInput = (event: Event) => {
  const el = event.target as HTMLElement
  inputData.value = el.innerText
    
  const bounds = el.getBoundingClientRect();
  inputPosition.value.x = `${bounds.x}px`;
  inputPosition.value.y = `${bounds.y}px`;
  inputPosition.value.width = `${bounds.width}px`;
  inputPosition.value.height = `${bounds.height}px`;
  isInputShown.value = true;
 //Scroll
  const watchScrollFunction = () => {
    watchScroll(el, wrapper, watchScrollFunction, inputPosition.value)
  }
 
  // Grab wrapper and ensure it exists (although it always should)
  const wrapper = document.querySelector('.table-wrapper') as HTMLElement
  if (!wrapper) {
    console.log("No boundry found")
    return
  }

  // Attach scroll listener
  wrapper.addEventListener('scroll', watchScrollFunction, false )


  //Click
  const watchClickFunction = () => {
    watchClick(el)
  }


  const watchClick = (el: HTMLElement) => {
    isInputShown.value = true;
    if (el.classList.contains('input-float')) {
      return
    }
    if (isInputClicked) {
      document.removeEventListener('click', watchClickFunction, false)
      isInputShown.value = false;
      isInputClicked = false;
      return
    } else {
      document.addEventListener('click', watchClickFunction, false);
      isInputClicked = true;
    }
  }

  
  document.querySelector('.floating-input')?.addEventListener('blur', () => {
    if (isInputClicked) {
      document.removeEventListener('click', watchClickFunction, false);
      isInputClicked = false;
    }
  });

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
  padding: 5px 5px 5px 10px;
}

.header {
    position: sticky;
    width: 100%;
    background-color: var(--bg-darker);
    height: 40px;
    width: 100%;
    position: sticky;
    top: 0;
    z-index: 4;
    border-radius: 5px;
    margin-left: -5px;
}

table {
  width: 100%;
  text-align: center;
  border-collapse: collapse;
  overflow: auto;
}

.table-container {
 width: 100%;
 height: 100%;
}


thead {
  margin-top: 10px;
  border-radius: 3px;
  background-color: var(--bg-dark);
  position: sticky;
  top: 60px;
  z-index: 4; /* Ensure the header is above the table body */
}

.table-head {
  height: 45px;
}


tbody {
    overflow: auto;
    margin-top: 10px;
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
  background-color: var(--bg-dark);
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
  border: .75px solid var(--gray-9);
  padding: 10px;
  position: absolute;
  z-index: 2;
  width: 120px;
  height: 200px;
  border-radius: 10px;
}

.input-box {
  background-color: var(--bg-darker);
  border: 0.25px solid var(--primary-accent-dark);
  position: absolute;
  z-index: 1;
  border-radius: 1px;
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
.input-box:focus-within {
  border: 1.5px solid var(--primary-accent-dark);
  box-shadow: 0px 0px 8px 0px var(--primary-accent-dark);
  border-radius: 5px;
  margin-top: -1.5px;
}
</style>

