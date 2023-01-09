import { createApp } from "vue";
import App from "./App.vue";
import { createRouter, createWebHashHistory } from "vue-router";
import routes_list from "@/routes.json";
import { useColorMode } from "@vueuse/core";


routes_list.routes.forEach((route: { component: any; name: string , path : string }, index) => {
    route.component = () => import(`@/views/${route.name.toLowerCase()}.vue`);
    route.path = route.path;
  });
  export const routes: any = routes_list.routes;
export const router = createRouter({
  history: createWebHashHistory(),
  routes: routes,
});
// import { useColorMode } from "@vueuse/core";
export const themeMode = useColorMode();

const app = createApp(App);

app.use(router);

app.mount("#app");
