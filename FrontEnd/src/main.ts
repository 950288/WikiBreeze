import { createApp, h, render, createVNode, ref } from "vue";
import App from "./App.vue";
import { createPinia, defineStore } from 'pinia'
import { createRouter, createWebHashHistory } from "vue-router";
import notification from "@/components/Notification.vue";
import input from "@/components/Input.vue";
import Home from "@/views/home.vue";
import Editor from "@/views/editor.vue";
const pinia = createPinia()

import { useColorMode } from "@vueuse/core";
export const routes = [
  { name: "Home", path: "/", component: Home, props: true },
  { name: "Editor", path: "/editor", component: Editor }
];
export const router = createRouter({
  history: createWebHashHistory(),
  routes: routes,
});

export const themeMode = useColorMode();

const app = createApp(App);

export const useNotifyStore = defineStore('notify', {
  state: () => ({ notifyId: 0 }),
  actions: {
    notify(duration: Number, title: string, msg: string, type: string, recall: Promise<{ success: boolean, notify: string | undefined }> | null) {
      const notificationInstance = h(<any>notification, {
        duration: duration,
        msg,
        title,
        type,
        promise: recall,
        count: this.notifyId++,
      });
      // Render the notification component
      const vnode = createVNode(notificationInstance);
      // Mount the VNode to an element outside of the app root
      const container = document.createElement('div');
      document.body.appendChild(container);
      render(vnode, container);
    }
  }
})


export const useInputStore = defineStore('input', {
  state: () => ({ inputId: 0 }),
  actions: {
    input(title: string, upload: boolean) {
      let recall = ref();
      let inputInstance = h(<any>input, {
        title,
        uploadEnable: upload,
        count: this.inputId++,
        recall
      });
      // Render the notification component
      const vnode = createVNode(inputInstance);
      // Mount the VNode to an element outside of the app root
      const container = document.createElement('div');
      document.body.appendChild(container);
      render(vnode, container);
      return recall;
    }
  }
})

app.use(router);
app.use(pinia)
app.mount("#app");
