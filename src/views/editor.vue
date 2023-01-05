<template>
    <div class="content is-large">
        <br>
        <tiptap v-if="contenetJson && renderConfigJson" @save="save" :contenetjson="contenetJson" :renderConfigJson="renderConfigJson"/>
        <progress v-if="!(contenetJson && renderConfigJson)" class="progress is-large is-info" max="100">60%</progress>
        <br>
    </div>
</template>

<script setup lang="ts">
import { useFetch } from "@vueuse/core";
import Tiptap from '@/components/Tiptap.vue'
import { ref, watch ,onMounted } from "vue";
import { requestUrl } from "@/App.vue";
const contenetJson = ref<JSON>();
const renderConfigJson = ref<JSON>();
onMounted(() => {
    let { data , error: getnodeError } = useFetch(requestUrl.value+"/getnode", {
        method: 'POST',
        body: JSON.stringify({
            page: history.state.index,
            content: history.state.content
        }),
        headers: {
            'Content-Type': 'application/json'
        }
    }).get().json();
    watch(data, (val) => {
        contenetJson.value = val;
        console.log(val);
    });
    watch(getnodeError, (val) => {
        console.log(val);
    });
    let { data: renderData , error: getRenderDataError} = useFetch(requestUrl.value+"/getRenderconfig", {
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
function save(contentjson :JSON, contenthtml :string){
    console.log(contentjson);
    let { data, error } = useFetch(requestUrl.value+"/savenode", {
        method: 'POST',
        body: JSON.stringify({
            page: history.state.index,
            content: history.state.content,
            Contentjson : contentjson,
            Contenthtml : contenthtml
        }),
        headers: {
        'Content-Type': 'application/json'
    }
    }).get().json();
    watch(data, (val) => {
        console.log(val);
    });
    watch(error, (val) => {
        console.log(val);
    });
    console.log(data);
}
</script>