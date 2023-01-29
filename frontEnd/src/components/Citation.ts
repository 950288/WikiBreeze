import { Node } from '@tiptap/core'
import { mergeAttributes, VueNodeViewRenderer ,NodeViewContent } from "@tiptap/vue-3"
import CitationComponent from './CitationComponent.vue'

declare module '@tiptap/core' {
    interface Commands<ReturnType> {
        citation: {
            addCitation: () => ReturnType,
        }
    }
}

export default Node.create({
    name: 'citation',

    group: 'inline-block',

    content: 'inline*',

    parseHTML() {
        return [
            {
                tag: 'citation',
            },
        ]
    },

    renderHTML({ HTMLAttributes }) {
        return ['citation', mergeAttributes(HTMLAttributes) , 0 ]
    },

    addCommands() {
        return {
            addCitation: () => ({ editor , tr }) => {
                const citation = this.editor.schema.nodes.citation.createChecked({
                    class: "citation"
                },[
                    editor.schema.text('add table note'),
                ])
                console.log(JSON.stringify(citation))
                tr.insert(tr.selection.from - 1, citation)
                return true
            }
        }
    },
    addNodeView () {
        return VueNodeViewRenderer(CitationComponent)
    },
})