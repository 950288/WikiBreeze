<template>
    <div class="content is-large">
        <br>

        <div class="box" v-for="(page, index) in data">
            <h2>{{ index }}</h2>
            <div class="buttons">
                <button v-for="node in page" class="button is-primary is-outlined"
                    @click="editnode(index, node)">{{ node }}</button>
                <!-- <button class="button is-primary is-outlined">About_2</button> -->
            </div>
        </div>
        <!-- <div class="box">
            <h2>About</h2>
            <div class="buttons">
                <button class="button is-primary is-outlined">About_1</button>
                <button class="button is-primary is-outlined">About_2</button>
            </div>
        </div> -->
        <article v-if="error" class="message is-warning is-large">
            <div class="message-body">
                <h3>fetch page list error!</h3>
                <p>{{ error }}</p>
                <p>please refresh this page</p>
            </div>
        </article>
        <progress v-if="!data" class="progress is-large is-info" max="100">60%</progress>
        <!-- <div>
            {{ data }}<br>
        </div> -->
    </div>
</template>
<script setup lang="ts">
import { useFetch } from "@vueuse/core";
import { ref, watch } from "vue";
const getdata = ref(false);
const geterror = ref(false);
let { data, error } = useFetch("http://127.0.0.1:8082/list").get().json();
function editnode(index, node) {
    let { data, error } = useFetch("http://127.0.0.1:8082/getnode", {
        method: 'POST',
        body: JSON.stringify({
            page: index,
            node: node
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
<style lang="scss" scoped>
.box {
    height: 40px;
    margin: 0;
    padding: 10px;
    background-color: var(--has-background-white);
    box-shadow: none;
    border: 1px solid var(--has-background-light-grey);
    position: relative;
    margin-bottom: 15px;

    h2 {
        display: block;
        height: 100%;
        line-height: 100%;
        margin: 0;
    }

    .buttons {
        display: block;
        position: absolute;
        right: 10px;
        top: 10px;

        .button {
            color: var(--has-text-primary);
            background: var(--has-background-primary);

            transition: ease 0.2s;
        }

        .button:hover {
            transform: scale(1.2);
        }
    }
}

.message {
    background: var(--is-warning-gb);

    .message-body * {
        // color: #000;
    }
}
</style>