<template>
  <button id="button" class="button is-success" @click="themeMode = themeMode === 'dark' ? 'light' : 'dark'">
    {{ themeMode }}
  </button>
</template>
<script setup lang="ts">
import { themeMode } from "@/main";
</script>
<style lang="scss" scoped>
#button {
  display: flex;
  align-items: center;
  /* Vertically center the button element */
  justify-content: center;
  /* Horizontally center the button element */
  height: 100%;
  position: relative;
  color: var(--has-text-success);
  background: var(--has-background-success);
  transition: ease 0.2s;
}

#button:hover {
  transform: scale(1.2);
}
</style>
