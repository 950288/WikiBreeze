<template>
    <div class="Input">
        <div class="background" @click="destory()"></div>
        <div class="body">
            <div class="heading">
                <p>{{ props.title }}</p>
                <div class="close"><button @click="destory()">
                        <img src="/src/assets/close.svg" alt="">
                    </button></div>
            </div>
            <div class="getLink">
                <div class="msg">
                    <p>{{ 'Please input the image url:' }}</p>
                    <input type="text" v-model="url" />
                </div>
            </div>
            <div class="confirm">
                <button class="button is-primary" @click="confirm()">{{ "confirm" }}</button>
            </div>
            <div v-if="upload">
                <div class="upload">
                    <p>{{ 'Or upload the image' }}</p>
                    <input type="file" />
                </div>
            </div>
            <div class="confirm" v-if="upload">
                <button class="button is-primary">{{ "upload" }}</button>
                <!-- <button class="button is-danger" @click="destory()">{{ "cancel" }}</button> -->
            </div>
        </div>
    </div>
</template>
<script lang="ts" setup>

import { ref, onMounted } from 'vue'

const url = ref("")
const props = defineProps({
    title: {
        type: String,
        // default: "Title"
    },
    upload:{
        type: Boolean,
        // default: falses
    },
    count: {
        type: Number,
        default: 0
    },
    recall: {
        type: ref
    }
})
function confirm(){
    destory()
    if(props.recall)
    props.recall.value = url.value;
}
function destory() {
    const notification = document.querySelector('.Input')?.parentElement
    if (notification) {
        notification.remove()
    }
}
</script>
<style lang="scss" scoped>
.Input {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 100;


    .background {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background-color: rgba(0, 0, 0, 0.2);
        backdrop-filter: blur(5px);
    }

    .body {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        width: 400px;
        background-color: var(--has-background-lightest-grey);
        border-radius: 3.5px;
        box-shadow: 0 0 10px rgba(0, 0, 0, .5);

        .heading {
            width: 100%;
            height: 50px;
            border-top-left-radius: 3.5px;
            border-top-right-radius: 3.5px;
            font-size: 2rem;
            font-family: 'Microsoft YaHei';
            background-color: #333;


            p {
                width: 100%;
                height: 100%;
                line-height: 50px;
                text-align: center;
                color: #fff;
            }

            .close {
                position: absolute;
                right: 10px;
                top: 10px;
                width: 30px;
                height: 30px;
                border-radius: 25%;
                cursor: pointer;
                transition: ease 0.1s;

                &:hover {
                    background-color: rgba(0, 0, 0, 0.8);
                }

                button {
                    position: absolute;
                    top:0;
                    width: 100%;
                    height: 100%;
                    border: none;
                    outline: none;
                    background-color: transparent;
                    cursor: pointer;
                    img {
                        width: 100%;
                        height: 100%;
                    }
                }
            }
        }

        .getLink {
            width: 100%;
            border-top-left-radius: 5px;
            border-top-right-radius: 5px;

            .msg {
                width: 100%;
                line-height: 50px;
                text-align: left;
                font-size: 1.25rem;
                margin-left: 10px;

                .p {
                    width: 100%;
                    height: 50px;
                    line-height: 50px;
                    text-align: left;
                    font-size: 1.25rem;
                }

                input {
                    border-radius: 5px;
                    outline: none;
                    border-bottom: 1px solid #333;
                    font-size: 1.25rem;
                    width: calc(100% - 30px);
                    color: #000;
                    height: 30px;
                    padding-left: 10px;
                }
            }
        }

        .upload {
            width: 100%;
            margin-left: 10px;

            p {
                width: 100%;
                height: 50px;
                line-height: 50px;
                text-align: left;
                font-size: 1.25rem;
            }

            input {
                margin-bottom: 15px;
            }
        }

        .confirm {
            text-align: center;
            margin-bottom: 18px;

            button {
                margin-left: 25px;
                margin-right: 25px;
                width: 100px;
                transition: ease 0.2s;
            }
            button:hover{
                transform: scale(1.2);
            }
        }
    }
}</style>