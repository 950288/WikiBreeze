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
            <div v-if="uploadEnable">
                <div class="upload">
                    <p>{{ 'Or upload the image(png jpeg jpg svg)' }}</p>
                    <input type="file" @change="handleUploadImage" accept="image/png, image/jpeg, image/jpg, image/svg" />
                </div>
            </div>
            <div class="confirm" v-if="uploadEnable">
                <button class="button is-primary" @click="upload()">{{ "upload" }}</button>
                <!-- <button class="button is-danger" @click="destory()">{{ "cancel" }}</button> -->
            </div>
        </div>
    </div>
</template>
<script lang="ts" setup>
import { ref, onMounted, createVNode, render, h, watch } from 'vue'
import notification from "@/components/Notification.vue";
import { useFetch } from '@vueuse/core';
import { requestUrl } from '@/App.vue';


let notifyCount = 95028;

const notify = (duration: Number, title: string, msg: string, type: string, recall: Promise<{ success: boolean, notify: string | undefined }> | any) => {
    const notificationInstance = h(<any>notification, {
        duration: duration,
        msg,
        title,
        type,
        promise: recall,
        count: `${(notifyCount++).toString()}uploads`,
    });
    // Render the notification component
    const vnode = createVNode(notificationInstance);
    // Mount the VNode to an element outside of the app root
    const container = document.createElement('div');
    document.body.appendChild(container);
    render(vnode, container);
};
const upload = ref(() => { })

const url = ref("")
const file = ref()
const props = defineProps({
    title: {
        type: String,
        // default: "Title"
    },
    uploadEnable: {
        type: Boolean,
        // default: falses
    },
    count: {
        type: Number,
        default: 0
    },
    recall: {
        type: Object
    }
})
function confirm() {
    destory()
    if (props.recall)
        props.recall.value = url.value;
}
function destory() {
    const notification = document.querySelector('.Input')?.parentElement
    if (notification) {
        notification.remove()
    }
}
const handleUploadImage = async (e: any) => {
    file.value = e.target.files
}
function checkFileName(fileName: string) {
    const reg = /^[\w\d_\-.]+$/;
    if (!reg.test(fileName)) {
        return false;
    }
    return true;
}
onMounted(() => {
    upload.value = () => {
        if (!file.value) {
            notify(
                1500,      // 0 means the notification will not be destoryed automatically after recall()
                'uploads',
                'No file selected !',
                'info',
                null);
            return
        }
        const fd = new FormData()
        if (!checkFileType(file.value[0].type)) return
        if (file.value[0].size > 10485760) {
            notify(
                1500,      // 0 means the notification will not be destoryed automatically after recall()
                'uploads',
                '[file-too-large] File is larger than 10485760 bytes !',
                'info',
                null);
            destory()
            return
        }
        if (!checkFileName(file.value[0].name)) {
            console.log("files name contain special character, rename file to timestamp")
            var renamedFile = new File([file.value[0]], `${new Date().getTime().toString()}.${file.value[0].name.split('.').pop()}`, { type: file.value[0].type });
            fd.append("file", renamedFile);
        } else {
            fd.append('file', file.value[0])
        }

        fd.append('directory', "WikiBreezeUpload/" + history.state.index + "/" + history.state.content)

        notify(
            0,      // 0 means the notification will not be destoryed automatically after recall()
            'uploads',
            'Uploading ' + file.value[0].name + '...',
            'info',
            new Promise((resolve) => {
                let { data: fileURL, error: uploadError } = useFetch(requestUrl.value + "/WikiBreezeUpload/", {
                    method: 'POST',
                    body: fd,
                    headers: {
                        'Content-Length': file.value[0].size.toString()
                    }
                }).get().json();
                watch(fileURL, (val) => {
                    if (val && val.location) {
                        console.log("upload " + file.value[0].name + " successful\nURL:" + val.location);
                        if (props.recall)
                            props.recall.value = val.location;
                        resolve({ success: true, notify: "upload successful !" })
                    } else {
                        console.log(val)
                        resolve({ success: false, notify: val })
                    }
                })
                watch(uploadError, (val) => {
                    console.log("upload failed :" + uploadError)
                    resolve({ success: false, notify: val })
                })
            })
        );
        destory()
        return
    }
})
const allowedTypes = ["image/png", "image/jpeg", "image/jpg", "image/svg+xml"];

function checkFileType(file: string) {
    if (!allowedTypes.includes(file)) {
        notify(
            1500,      // 0 means the notification will not be destoryed automatically after recall()
            'Tpye error',
            'Only PNG, JPG, JPEG and SVG images are allowed !',
            'info',
            null);
        return false;
    }

    return true;
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
                    top: 0;
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

            button:hover {
                transform: scale(1.2);
            }
        }
    }
}
</style>