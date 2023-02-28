<template>
    <div class="notification">
        <div class="body">
            <div class="title">
                <slot name="title">java</slot>
            </div>
            <div class="content">
                <slot name="content">asdfghj</slot>
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import { booleanLiteral } from '@babel/types';
import { ref, onMounted} from 'vue'
console.log('notification')
const props = defineProps({
    time: {
        type: Number,
        default: 3000
    },
    msg:{
        type: String,
        default: 'java'
    },
    type: {
        type: String,
        default: 'success'
    },
    Promise: {
        type: Promise<{success:boolean,notify:string | undefined}>,
        default: null
    }
})
async function java(a: number){
    console.log(a + 1)
    return a + 1
} 
java(1).then((a) => {
    console.log(a)
})
const time = props.time

onMounted(() => {
    console.log(time)
    if (time) {
        setTimeout(() => {
            destory()
        }, time)
    }

    props.Promise?.then((returns) => {
        if(returns.notify){
            console.log(returns.notify)
        }else{
            console.log('no notify')
        }
        setTimeout(() => {
            destory()
        }, 1500)
    })
})

function destory() {
    document.body.removeChild(document.querySelector('.notification') as Node)
}
</script>

<style lang="scss" scoped>
.notification {
    position: fixed;
    bottom: 0;
    width: 100%;

    .body {
        height: 100px;
        width: 60%;
        background: #000000;
        margin: 0 auto 20px auto;
        border-radius: 10px;

        .title {
            height: 30px;
            width: 100%;
            background: #000000;
            border-radius: 10px 10px 0 0;
            color: #fff;
            font-size: 20px;
            line-height: 30px;
            text-align: center;
        }
    }
}</style>

//Q. to copilot :"thank you so much"
//A. "you're welcome"
//Q. "how are you?"
//A. "I'm good"
//Q. "what's your name?"
//A. "I'm copilot"
