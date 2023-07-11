<script setup lang="ts">
import Menu from "@/components/Menu.vue";
import { RouterView } from "vue-router";
import { routes } from "@/main";

</script>
<script  lang="ts">
import { ref } from "vue";
export const requestUrl = ref(process.env.NODE_ENV === 'development' ? "http://localhost:5001" : window.location.origin);
</script>

<template>
  <Menu :routes="routes" />
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