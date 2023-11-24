import { Node, findParentNodeClosestToPos } from "@tiptap/core";
import { createTable } from "@tiptap/extension-table";

/**
 * @name tablePro
 * @description The tablePro extension allows you to insert a table with a
 * note into the editor. The note is displayed above the table and can be
 * used to provide context or additional information about the table.
 * @example
 * insertTablePro() // inserts a table with a default size of 2 rows and 2
 * columns and a header row
 * insertTablePro({ rows: 3, cols: 3, withHeaderRow: false }) // inserts a
 * table with 3 rows, 3 columns, and no header row
 * @renderHTML
 * <table-pro>
 *   <p>add table note here</p>
 *   <table>
 *      ****
 *   </table>
 * </table-pro>
 */

declare module "@tiptap/core" {
  interface Commands<ReturnType> {
    tablePro: {
      insertTablePro: () => ReturnType;
      deleteTablePro: () => ReturnType;
    };
  }
}

export default Node.create({
  name: "tablePro",

  group: "block",

  content: "paragraph",

  defining: false,

  isolating: true,

  parseHTML() {
    return [{ tag: "table-pro" }];
  },

  renderHTML() {
    return ["table-pro", 0];
  },

  addAttributes() {
    return {
      allowGapCursor: true,
    };
  },

  addCommands() {
    return {
      insertTablePro:
        ({ rows = 2, cols = 2, withHeaderRow = true } = {}) =>
        ({ editor, tr }) => {
          const table = createTable(editor.schema, rows, cols, withHeaderRow);

          const noteNode = editor.schema.nodes.paragraph.createChecked(
            { class: "tableNote" },
            [editor.schema.text("add table note here")],
          );

          if (!noteNode) {
            return false;
          }

          const tableWithNoteNode = editor.schema.nodes.tablePro.createChecked(
            {},
            [noteNode, table],
          );
          tr.insert(tr.selection.from - 1, tableWithNoteNode);
          return true;
        },
      deleteTablePro:
        () =>
        ({ editor, tr }) => {
          const { $from } = tr.selection;
          const tablePro = findParentNodeClosestToPos(
            $from,
            (node) => node.type === editor.schema.nodes.tablePro,
          );
          if (tablePro) {
            tr.delete(tablePro.pos, tablePro.pos + tablePro.node.nodeSize);
            return true;
          }
          return false;
        },
    };
  },
});
