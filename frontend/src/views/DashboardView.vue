<template>
    <Dropdown :options="DropdownItems" />
</template>

<script setup lang="ts">
import Dropdown from '../components/common/Dropdown.vue'

type Option = {
    value: string
    label: string
}

const DropdownItems: Option[] = [
    {label: "Option 1", value: "1"},
    {label: "Option 2", value: "2"},
    {label: "Option 3", value: "3"},
    {label: "Option 4", value: "4"},
    {label: "Option 5", value: "5"},
]
</script>

<style scoped>
</style>

