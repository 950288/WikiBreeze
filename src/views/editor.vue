<template>
    <div class="content is-large">
        <br>
        <tiptap @save="save" />
        <br>
    </div>
</template>

<script setup lang="ts">
import { useFetch } from "@vueuse/core";
import Tiptap from '@/components/Tiptap.vue'
import { ref, watch ,onMounted } from "vue";
const props = defineProps(['index', 'node'])
onMounted(() => {
    console.log(props.index);
    console.log(props.node);
    let { data, error } = useFetch("http://127.0.0.1:8082/getnode", {
        method: 'POST',
        body: JSON.stringify({
            // page: index,
            // node: node
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
});
function save(datas :JSON){
    console.log(datas);
    let { data, error } = useFetch("http://127.0.0.1:8082/getnode", {
        method: 'POST',
        body: JSON.stringify({
            // page: index,
            // node: node
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