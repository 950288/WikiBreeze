import { mergeAttributes, Node, findParentNodeClosestToPos } from '@tiptap/core'
import { createTable } from '@tiptap/extension-table'


/**
 * @name TablePro
 * @description The TablePro extension allows you to insert a table with a 
 * note into the editor. The note is displayed above the table and can be 
 * used to provide context or additional information about the table.
 * @example
 * insertTablePro() // inserts a table with a default size of 2 rows and 2 
 * columns and a header row
 * insertTablePro({ rows: 3, cols: 3, withHeaderRow: false }) // inserts a 
 * table with 3 rows, 3 columns, and no header row
 * @renderHTML
 * <table-pro>
 *   <p>edit your table note</p>
 *   <table>
 *      ****
 *   </table>
 * </table-pro>
 */
interface TableProOptions {
  HTMLAttributes: Record<string, any>
}

declare module '@tiptap/core' {
  interface Commands<ReturnType> {
    tablePro: {
      insertTablePro: () => ReturnType,
      deleteTablePro: () => ReturnType,
    }
  }
}

export default Node.create({
  name: 'TablePro',

  group: 'block',

  content: 'paragraph table',

  parseHTML() {
    console.log('parseHTML')
    return [
      { tag: 'table-pro' },
    ]
  },

  renderHTML({ HTMLAttributes = {} }) {
    // console.log('renderHTML')
    return ['table-pro', 0]
  },

  addCommands() {
    return {
      insertTablePro: ({ rows = 2, cols = 2, withHeaderRow = true } = {}) => ({ commands, editor, tr }) => {
        const table = createTable(editor.schema, rows, cols, withHeaderRow)
        // use createChecked to create a table node without checking the schema
        // const noteNode = editor.schema.nodes.paragraph.createChecked(null, [
        console.log(editor.schema.marks)
        const noteNode = editor.schema.nodes.paragraph.createChecked(
          { class: "tableNote" }
          , [
            editor.schema.text('add table note'),
          ])

        console.log("'''''''''''''''''''''''''")
        console.log(JSON.stringify(noteNode))
        //add class "note" to noteNode 
        if (!noteNode) {
          return false
        }

        const tableWithNoteNode = editor.schema.nodes.TablePro.createChecked({}, [noteNode, table])
        tr.insert(tr.selection.from - 1, tableWithNoteNode)
        console.log(JSON.stringify(tableWithNoteNode))
        return true
      },
      deleteTablePro: () => ({ commands, editor, tr }) => {
        const { $from } = tr.selection;
        const tablePro = findParentNodeClosestToPos($from, (node) => node.type === editor.schema.nodes.TablePro);
        if (tablePro) {
          tr.delete(tablePro.pos, tablePro.pos + tablePro.node.nodeSize)
          return true;
        }
        return false;
      }
    }
  }
})