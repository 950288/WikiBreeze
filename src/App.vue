<script setup lang="ts">
import Menu from "@/components/Menu.vue";
import { RouterLink, RouterView } from "vue-router";
import { routes } from "@/main";

</script>

<template>
  <Menu  :routes="routes" />
  <router-view />
</template>

<style lang="scss">

@import "@/style/index.scss";
// @include set-color-mode(light);
* {
  margin: 0;
  padding: 0;
  // background-color: var(--has-background-success);
}
:root {
  background-color: var(--has-background-white);
}
</style>