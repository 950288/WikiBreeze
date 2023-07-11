import { Node, findParentNodeClosestToPos } from '@tiptap/core'

/**
 * @name imagePro
  * @description The imagePro extension allows you to insert an image with a
  * note into the editor. The note is displayed below the image and can be
  * used to provide context or additional information about the image.
  * @example
  * insertimagePro()
 */

declare module '@tiptap/core' {
  interface Commands<ReturnType> {
    imagePro: {
      insertImagePro: (url: String) => ReturnType,
      deleteImagePro: () => ReturnType,
    }
  }
}

export default Node.create({
  name: 'imagePro',

  group: 'block',

  content: 'image paragraph',

  // allowGapCursor: true,

  isolating: true,

  parseHTML() {
    return [
      { tag: 'image-pro' },
    ]
  },

  renderHTML() {
    // return ['image-pro', { 'data-gapcursor': 'ImageProCursor' }, 0]
    return ['image-pro', 0]
  },

  addAttributes() {
    return {
      class: {
        default: 'imagePro',
      },
    }
  },

  addCommands() {
    return {
      insertImagePro: (url) => ({ commands, editor, tr }) => {
        console.log(url)
        if (!url) {
          return false
        }
        const image = editor.schema.nodes.image.createChecked(
          {
            src: url
          }
        )
        console.log(image)
        const tableNote = editor.schema.nodes.paragraph.createChecked(
          {
            class: "paragraph"
          }
          ,
          [
            editor.schema.text('Edit your image note here')
          ]
        )
        console.log(JSON.stringify(image))
        const imagePro = editor.schema.nodes.imagePro.createChecked(
          { class: "imagePro" }
          ,
          [
            image,
            tableNote
          ]
        )
        tr.replaceSelectionWith(imagePro)
        editor.view.dispatch(tr)
        return true
      },
      deleteImagePro: () => ({ commands, editor, tr }) => {
        const { $from } = tr.selection
        const imagePro = findParentNodeClosestToPos($from, (node) => node.type === editor.schema.nodes.imagePro)
        if (imagePro) {
          tr.deleteRange(imagePro.pos, imagePro.pos + imagePro.node.nodeSize)
          return true
        }
        return false
      },

    }
  },
})
