<script setup>

import {computed, } from "vue";

const { w, h} = defineProps({
  w: {
    type: Number,
    default: 20,
  },
  h: {
    type: Number,
    default: 20,
  },
})

const style = computed(() => {
  return {
    display: 'inline-block',
    verticalAlign: 'middle',
    width: w + 'px',
    height: h + 'px',
  }
})
</script>

<template>
  <img src="~@/assets/aichat.png" class="cursor-pointer" :style="style"/>
</template>

<style scoped lang="scss">

</style>