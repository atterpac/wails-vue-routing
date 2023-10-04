<template>
    <div class="select" tabindex="1" :style='{width: props.width, height: props.height}'>
        <div v-for='option in options' @click="$emit('item', option.value)">
          <input class="selectopt" name="dropdown" type="radio" :id='option.value' :checked='option.checked'>
          <label :for="option.value" :id="option.value" class="option" @click='setChecked' :style='{background: props.bg}'>{{option.label}}</label>
        </div>
    </div>
</template>

<script setup lang="ts">
import { defineProps, ref } from 'vue'
type Option = {
    value: string
    label: string
    checked?: boolean
}

const props = defineProps({
    options: {
        type: Array<Option>,
        required: false
    },
    width: {
        type: String,
        default: '150px'
    },
    height: {
        type: String,
        default: '30px'
    },
    bg: {
        type: String,
        default: 'var(--bg-dark)'
    }
})

// stuff prop into a reactive value
const options = ref(props.options)


const setChecked = (e: Event) => {
    let checked = e.target as HTMLElement 
    console.log(checked.id)
    if (options.value == undefined) {
        return
    }
    for (let i = 0; i < options.value.length; i++) {
        if (options.value[i].value == checked.id) {
            console.log("checked", options.value[i].value)
            options.value[i].checked = true
        } else {
            console.log("unchecked", options.value[i].value)
            options.value[i].checked = false
        }
    }
}
</script>

<style scoped>

.select {
  display:flex;
  text-align: center;
  flex-direction: column;
  position:relative;
  height: 24px;
  width:75px;
  border-radius: 5px;
  border: 2px solid var(--gray-b);
  padding-bottom: 5px;
}

.select:hover {
    border: 2px solid var(--primary-accent);
}

.option {
  padding:5px 5px 5px 5px;
  height: 30px;
  text-align: center;
  display:flex;
  align-self:center;
  justify-content:center;
  text-align: center;
  background:var(--bg-dark);
  position:absolute;
  top:0;
  width: 100%;
  pointer-events:none;
  order:2;
  z-index:1;
  transition:background .4s ease-in-out;
  box-sizing:border-box;
  overflow:hidden;
  white-space:nowrap;
  
}


.option:first-child {
  border-radius: 5px 5px 0 0;
}

.option:hover {
  background:#666;
}

.select:focus .option {
  text-align: center;
  position:relative;
  pointer-events:all;
}

input { 
  opacity:0;
  position:absolute;
  left:-99999px;
}

input:checked + label {
  order: 1;
  z-index:2;
  border-top:none;
  position:relative;
}
/**/
/* input:checked + label:after { */
/*   content:''; */
/*   width: 0;  */
/* 	height: 0;  */
/* 	border-left: 5px solid transparent; */
/* 	border-right: 5px solid transparent; */
/* 	border-top: 5px solid white; */
/*   position:absolute; */
/*   right:10px; */
/*   top:calc(50% - 2.5px); */
/*   pointer-events:none; */
/*   z-index:3; */
/* } */

input:checked + label:before {
  position:absolute;
  right:0;
  height: 40px;
  width: 40px;
  content: '';
}
</style>

