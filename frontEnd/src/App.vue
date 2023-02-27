<script setup lang="ts">
import Menu from "@/components/Menu.vue";
import { defineAsyncComponent, h, render } from "vue";
import notification from "@/components/notification.vue";
import { RouterView } from "vue-router";
import { routes } from "@/main";
const AsyncComp = defineAsyncComponent(() =>
  import('@/components/MyComponent.vue')
)
render(h(h(notification),{time:100,}), document.body)
</script>
<script  lang="ts">
import { ref } from "vue";
export const requestUrl = ref(process.env.NODE_ENV === 'development' ? "http://127.0.0.1:8080" : window.location.origin);
</script>

<template>
  <Menu  :routes="routes" />
  <router-view />
  <!-- <notification /> -->
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