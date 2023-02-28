<template>
    <div class="content is-large">
        <br>
        <tiptap v-if="contenetJson && renderConfigJson" @save="save" :contenetjson="contenetJson"
            :renderConfigJson="renderConfigJson" />
        <progress v-if="!(contenetJson && renderConfigJson)" class="progress is-large is-info" max="100">60%</progress>
        <br>
    </div>
</template>

<script setup lang="ts">
import { useFetch } from "@vueuse/core";
import Tiptap from '@/components/Tiptap.vue'
import { ref, watch, onMounted, getCurrentInstance } from "vue";
import { requestUrl } from "@/App.vue";
const contenetJson = ref<JSON>();
const renderConfigJson = ref<JSON>();
const app = <any>getCurrentInstance();


onMounted(() => {
    let { data: PageDate, error: getnodeError } = useFetch(requestUrl.value + "/getdata", {
        method: 'POST',
        body: JSON.stringify({
            page: history.state.index,
            content: history.state.content
        }),
        headers: {
            'Content-Type': 'application/json'
        }
    }).get().json();
    watch(PageDate, (val) => {
        contenetJson.value = val;
        console.log(val);
    });
    watch(getnodeError, (val) => {
        console.log(val);
    });
    let { data: renderData, error: getRenderDataError } = useFetch(requestUrl.value + "/getRenderconfig", {
        method: 'POST',
        body: JSON.stringify({
            page: history.state.index,
            content: history.state.content
        }),
        headers: {
            'Content-Type': 'application/json'
        }
    }).get().json();
    watch(renderData, (val) => {
        renderConfigJson.value = val;
        console.log(val);
    });
    watch(getRenderDataError, (val) => {
        console.log(val);
    });
});
function save(contentjson: JSON, contenthtml: string) {
    // console.log(contentjson);
    let { data: saveReturnMsg, error: saveError } = useFetch(requestUrl.value + "/savedata", {
        method: 'POST',
        body: JSON.stringify({
            page: history.state.index,
            content: history.state.content,
            Contentjson: contentjson,
            Contenthtml: contenthtml
        }),
        headers: {
            'Content-Type': 'application/json'
        }
    }).get().json();
    async function recall() {
        watch(saveReturnMsg, (val) => {
            if (val.success == true) {
                console.log("saved success!");
                Promise.resolve({ success: true, notify: "save success !" })
            } else {
                console.log("saved failed!");
                Promise.resolve({ success: false, notify: saveError.value })
            }
        });
    }
    

    // proxy.$notify({
    //     title: 'save',
    //     content: 'saving...',
    //     duration: 0,
    //     type: 'info',
    //     Promise: recall()
    // });
    console.log("notify");
    // console.log(proxy);
    app?.proxy.$notify();

    console.log(saveReturnMsg.value);
}
</script>