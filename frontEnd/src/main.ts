import { createApp, h, render, createVNode, ref } from "vue";
import App from "./App.vue";
import { createRouter, createWebHashHistory, createWebHistory } from "vue-router";
import notification from "@/components/Notification.vue";
import input from "@/components/InputInfo.vue";
import routes_list from "@/routes.json";
import { useColorMode } from "@vueuse/core";


routes_list.routes.forEach((route: { component: any; name: string, path: string }, index) => {
  route.component = () => import(`@/views/${route.name.toLowerCase()}.vue`);
  route.path = route.path;
});
export const routes: any = routes_list.routes;
export const router = createRouter({
  history: createWebHistory(),
  // history: createWebHashHistory(),
  // 404 => home page
  routes: routes,
});
// import { useColorMode } from "@vueuse/core";
export const themeMode = useColorMode();

const app = createApp(App);

app.use(router);

let notifycount = 0;
// Define the global Vue property $notify
app.config.globalProperties.$notify = (duration: Number, title:string ,msg: string, type: string, recall: Promise<{ success: boolean, notify: string | undefined }>) => {
  const notificationInstance = h(<any>notification, {
    duration:duration,
    msg,
    title,
    type,
    promise: recall,
    count: notifycount++,
  });
  console.log(notificationInstance.props)
  // Render the notification component
  const vnode = createVNode(notificationInstance);
  // Mount the VNode to an element outside of the app root
  const container = document.createElement('div');
  document.body.appendChild(container);
  render(vnode, container);
};

let inputcount = 0;
app.config.globalProperties.$input = (
  title: string,
  upload: boolean
) => {
  let recall = ref();
  const inputInstance = h(<any>input ,{
    title,
    upload,
    count: inputcount++,
    recall
  });
  console.log(recall)
  console.log(inputInstance.props)
  // Render the notification component
  const vnode = createVNode(inputInstance);
  // Mount the VNode to an element outside of the app root
  const container = document.createElement('div');
  document.body.appendChild(container);
  render(vnode, container);
  return recall;
}


app.mount("#app");
