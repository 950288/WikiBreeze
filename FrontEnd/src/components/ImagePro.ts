import { mergeAttributes, Node, findParentNodeClosestToPos } from '@tiptap/core'
import { createTable } from '@tiptap/extension-table'
import { Image } from '@tiptap/extension-image'


/**
 * @name ImagePro
  * @description The ImagePro extension allows you to insert an image with a
  * note into the editor. The note is displayed below the image and can be
  * used to provide context or additional information about the image.
  * @example
  * insertImagePro()
 */

declare module '@tiptap/core' {
  interface Commands<ReturnType> {
    ImagePro: {
      insertImagePro: (url: String) => ReturnType,
      // deleteImagePro: () => ReturnType,
    }
  }
}

export default Node.create({
  name: 'ImagePro',

  group: 'block',

  content: 'paragraph image',

  parseHTML() {
    return [
      { tag: 'image-pro' },
    ]
  },

  renderHTML() {
    return ['image-pro', 0]
  },

  addCommands() {
    return {
      insertImagePro: (url) => ({ commands, editor, tr }) => {
        console.log(url)
        if (!url) {
          return false
        }
        const image = editor.schema.nodes.image.createChecked(
          { src: url }
        )
        console.log(JSON.stringify(image))
        const note = editor.schema.nodes.paragraph.createChecked(
          { class: "imageNote" }
          , 
          [
            editor.schema.text('edit your image note')
          ]
        )
        console.log(JSON.stringify(note))
        const imagePro = editor.schema.nodes.ImagePro.createChecked(
          { class: "imagePro" }
          , 
          [
            note,
            image
          ]
        )
        tr.replaceSelectionWith(imagePro)
        editor.view.dispatch(tr)
        return true
      },
      // deleteTablePro: () => ({ commands, editor, tr }) => {
      // }
    }
  }
})