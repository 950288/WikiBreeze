import { mergeAttributes, Node, findParentNodeClosestToPos,NodeView } from '@tiptap/core'

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
      deleteImagePro: () => ReturnType,
    }
  }
}

export default Node.create({
  name: 'ImagePro',

  group: 'block',

  content: 'image paragraph',

  allowGapCursor: true,

  parseHTML() {
    return [
      { tag: 'image-pro' },
    ]
  },

  renderHTML() {
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
          { src: url,
            allowGapCursor: false,
          }
        )
        console.log(image)
        const note = editor.schema.nodes.paragraph.createChecked(
          {} 
          , 
          [
            editor.schema.text('edit your image note here')
          ]
        )
        console.log(JSON.stringify(note))
        const imagePro = editor.schema.nodes.ImagePro.createChecked(
          { class: "imagePro" }
          , 
          [
            image,
            note
          ]
        )
        tr.replaceSelectionWith(imagePro)
        editor.view.dispatch(tr)
        return true
      },
      deleteImagePro: () => ({ commands, editor, tr }) => {
        const { $from } = tr.selection
        const imagePro = findParentNodeClosestToPos($from, (node) => node.type === editor.schema.nodes.ImagePro)
        if (imagePro) {
          tr.deleteRange(imagePro.pos, imagePro.pos + imagePro.node.nodeSize)
          return true
        }
        return false
      }
    }
  },
  addNodeView(): NodeViewRenderer<ImageProNodeType, Editor> {
    return {
      name: 'ImagePro',
      // 监听删除事件，触发删除命令
      destroy: (view: EditorView, node: ImageProNodeType) => {
        const { deleteImagePro } = view.state.commands;
        try {
          deleteImagePro();
        } catch (err) {
          // 如果删除失败，说明这个节点已经被删除了，不需要再做任何事情
        }
      },
    };
  },
})
