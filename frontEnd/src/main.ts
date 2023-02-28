import { cloneVNode, createApp, h, render } from "vue";
import App from "./App.vue";
import { createRouter, createWebHashHistory } from "vue-router";
import notification from "@/components/notification.vue";
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

let notifyCount = 0;
app.config.globalProperties.$notify = (time: Number,message: string,type: string, recall: Promise<{success:boolean,notify:string | undefined}>) => {console.log("nnnnnnnnnnnnnnnnnnnnotification");render(h(h(notification,{key:notifyCount++}),{time : time,msg: message,type: type, Promise: recall}), document.body)};

app.mount("#app");
