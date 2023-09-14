<template>
    <div class="content is-large">
        <br>

        <div class="box" v-for="(page, index) in data">
            <h2>{{ index }}</h2>
            <div class="buttons">
                <button v-for="content in page" class="button is-primary is-outlined"
                    @click="editcontent(index.toString(), content)">
                    {{ content }}
                </button>
            </div>
        </div>
        <article v-if="error" class="message is-warning is-large">
            <div class="message-body">
                <h3>fetch page list error!</h3>
                <p>{{ error }}</p>
            </div>
        </article>
        <progress v-if="!data && !error" class="progress is-large is-info" max="100">Loading</progress>
    </div>
</template>
<script setup lang="ts">
import { useFetch } from "@vueuse/core";
import { router } from "@/main";
import { requestUrl } from "@/App.vue";
let { data, error } = useFetch(requestUrl.value + "/list").get().json();
function editcontent(index: string, content: string) {
    router.push({
        name: "Editor",
        state: {
            index: index,
            content: content,
        }
    })
}
</script>
<style lang="scss" scoped>
.box {
    margin: 0;
    padding: 10px;
    display: flex;
    background-color: var(--has-background-white);
    box-shadow: none;
    border: 1px solid var(--has-background-light-grey);
    position: relative;
    margin-bottom: 15px;


    h2 {
        display: block;
        height: 100%;
        line-height: 115%;
        margin: 0;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;

    }

    .buttons {
        display: flex;
        justify-content: right;
        right: 10px;
        top: 10px;
        margin-left: auto;

        .button {
            color: var(--has-text-primary);
            background: var(--has-background-primary);
            margin-left: 0.5rem;
            margin-right: 0;
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