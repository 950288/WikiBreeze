<template>
    <div class="content is-large">
        <br>
        <tiptap v-if="contenetJson && renderConfigJson" @save="save" :contenetjson="contenetJson"
            :renderConfigJson="renderConfigJson" />
        <progress v-if="!(contenetJson && renderConfigJson) && !error" class="progress is-large is-info" max="100">60%</progress>
        <br>
        <article v-if="error" class="message is-warning is-large">
            <div class="message-body">
                <h3>Error!</h3>
                <p>{{ error }}</p>
            </div>
        </article>
    </div>
</template>

<script setup lang="ts">
import { useFetch } from "@vueuse/core";
import Tiptap from '@/components/Tiptap.vue'
import { ref, watch, onMounted, getCurrentInstance } from "vue";
import { requestUrl } from "@/App.vue";
import { router } from "@/main";

const contenetJson = ref<JSON>();
const renderConfigJson = ref<JSON>();
const app = <any>getCurrentInstance();
let error = ref<string>();

onMounted(() => {
    if(!history.state.content){
        router.push({
            name:"Home"
        })
        return
    }
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
        error.value = val;
    });
    let { data: renderData, error: getRenderDataError } = useFetch(requestUrl.value + "/getEditorConfig", {
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
    let RequestBody = JSON.stringify({
        page: history.state.index,
        content: history.state.content,
        Contentjson: contentjson,
        Contenthtml: contenthtml
    });
    // console.log(contentjson);
    let { data: saveReturnMsg, error: saveError } = useFetch(requestUrl.value + "/savedata", {
        method: 'POST',
        body: RequestBody,
        headers: {
            'Content-Type': 'application/json',
            'Content-Length':  RequestBody.length.toString()
        }
    }).get().json();
    async function recall() {
        return new Promise((resolve) => {
            watch(saveReturnMsg, (val) => {
                if (val && val.success == 'true') {
                    console.log("saved success!");
                    resolve({ success: true, notify: "saved successful!" });
                } else {
                    console.log("saved failed!");
                    resolve({ success: false, notify: saveError.value ?  saveError.value : "saved failed!"});
                }
            });
        });
    }
    app?.proxy.$notify(
        0,
        'save',
        'saving...',
        'info',
        recall()
    );
    console.log("notify");

    console.log(saveReturnMsg.value);
}
</script>
<style lang="scss" scoped>
.message {
    background: var(--is-warning-gb);

    .message-body * {
        // color: #000;
    }
}
</style>
