import Image from '@tiptap/extension-image'
import { findParentNodeClosestToPos } from '@tiptap/vue-3'

export const ImageX = Image.extend({
    name: "imageX",

    draggable: false,

    addKeyboardShortcuts() {
        return {
            'Enter': () => {
                // console.log("Enter")
                if (this.editor.isActive('imageX')) {
                    console.log("Enter")
                    return true
                }
                return false
            },
        }
    }
})

console.log(Image)
console.log(ImageX)