import Paragraph from "@tiptap/extension-paragraph";
import { findParentNodeClosestToPos } from "@tiptap/vue-3";
/**
 * this extended Paragraph allows you to delete an imagePro or tablePro node by pressing the backspace key
 **/
export default Paragraph.extend({
  addKeyboardShortcuts() {
    return {
      Backspace: function ({ editor }) {
        const tr = editor.state.tr;
        const { $from } = tr.selection;
        const imagePro = <any>(
          findParentNodeClosestToPos(
            $from,
            (node) => node.type === editor.schema.nodes.imagePro,
          )
        );
        const tablePro = <any>(
          findParentNodeClosestToPos(
            $from,
            (node) => node.type === editor.schema.nodes.tablePro,
          )
        );
        if (
          imagePro &&
          typeof imagePro.node.content.content[1]?.content.content[0]?.text ==
            "undefined"
        ) {
          console.log("delete ImagePro");
          editor.chain().focus().deleteImagePro().run();
          return true;
        }
        if (
          tablePro &&
          typeof tablePro.node.content.content[0]?.content.content[0]?.text ==
            "undefined"
        ) {
          console.log("delete TablePro");
          editor.chain().focus().deleteTablePro().run();
          return true;
        }
        return false;
      },
    };
  },
});
