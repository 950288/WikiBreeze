<template>
    <div :class="`notification notification${count}`" :style="`transform: translateX(-50%) translateY(${trsY}%);`">
        <div :class="`body ${notifyType}`">
            <div class="infoContent">
                {{ msg }}
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import { ref, onMounted } from 'vue'
console.log('notification')
const trsY = ref(100);
const notifyType = ref("info")
setTimeout(() => { trsY.value = 0; }, 1)
const props = defineProps({
    duration: {
        type: Number,
        default: 3000
    },
    msg: {
        type: String,
        default: ''
    },
    title: {
        type: String,
        default: ""
    },
    type: {
        type: String,
        default: 'info'
    },
    promise: {
        type: Promise<{ success: boolean, notify: string | undefined }>,
        default: null
    },
    count: {
        type: Number,
        default: 0
    }
})
const title = ref(props.title)
const msg = ref(props.msg)
console.log(props)
console.log(props.duration)
async function java(a: number) {
    console.log(a + 1)
    return a + 1
}
java(1).then((a) => {
    console.log(a)
})
const duration = props.duration
onMounted(() => {
    console.log(duration)
    if (duration != 0) {
        destory(duration)
    } else {
        props.promise?.then((value) => {
            console.log(value)
            if (value.notify) {
                msg.value = value.notify
            } else {
                console.log('no notify')
            }
            if(value.success){
                notifyType.value = "success"
            } else {
                notifyType.value = "danger"
            }
            destory()
        })
    }
})

function destory(duration: number = 2500) {
    console.log('translate' + '.notification' + props.count)
    setTimeout(() => {
        trsY.value = 100;
        setTimeout(() => {
            console.log('destory' + '.notification' + props.count)
            const notification = document.querySelector('.notification' + props.count)
            if (notification) {
                notification.remove()
            }
        }, 1000)
    }, duration)
}
</script>

<style lang="scss" scoped>
.notification {
    position: fixed;
    bottom: 0;
    width: 60%;
    left: 50%;
    // transform: translateX(-50%) translateY(100%);
    transition: all .2s ease;

    .body {
        height: 60px;
        width: 100%;
        // background: #d5e157;
        margin: 0 auto 20px auto;
        border-radius: 10px;
        transition: background-color .2s linear;

        .infoContent {
            width: 100%;
            // background: #000000;
            border-radius: 10px 10px 0 0;
            // color: #fff;
            font-size: 2em;
            line-height: 60px;
            text-align: center;
            // transform: translateY(50%);
        }
    }
}

.info {
    background-color: #ffe08a;
    * {
        color: rgba(0, 0, 0, 1);
    }
}

.success {
    background-color: #48c78e;
    * {
        color: #fff;
    }

}

.danger {
    background-color: #f14668;
    * {
        color: #fff;
    }

}
</style>

//Q. to copilot :"thank you so much"
//A. "you're welcome"
//Q. "how are you?"
//A. "I'm good"
//Q. "what's your name?"
//A. "I'm copilot"
