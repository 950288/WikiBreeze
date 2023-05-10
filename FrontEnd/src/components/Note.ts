import Paragraph from '@tiptap/extension-paragraph'
import { findParentNodeClosestToPos } from '@tiptap/vue-3'
export const Note = Paragraph.extend({
    name: "note",
    content: "text*",
    addKeyboardShortcuts() {
        return {
            'Backspace': function ({ editor }) {
                let tr = editor.state.tr
                console.log("Backspace")
                const { $from } = tr.selection
                let imagePro = <any>findParentNodeClosestToPos($from, (node) => node.type === editor.schema.nodes.ImagePro)
                let tablePro = <any>findParentNodeClosestToPos($from, (node) => node.type === editor.schema.nodes.TablePro)
                if ((imagePro && typeof (imagePro.node.content.content[1]?.content.content[0]?.text) == "undefined")) {
                    console.log("delete ImagePro")
                    editor.chain().focus().deleteImagePro().run()
                    return true
                }
                if ((tablePro && typeof (tablePro.node.content.content[1]?.content.content[0]?.text) == "undefined")) {
                    console.log("delete TablePro")
                    editor.chain().focus().deleteTablePro().run()
                    return true
                }
                return false
            },
        }
    }
})
