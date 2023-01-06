<template>
  <div class="main content">
    <h3>You are editing <u>{{ state.index }}</u> on page <u>{{ state.content }}.</u></h3>
    <div class="buttons">
      <button id="button" class="button is-success" v-if="mounted"
        @click="editor.chain().focus().toggleHeading({ level: 1 }).run()">
        H1
      </button>
      <button id="button" class="button is-success" v-if="mounted"
        @click="editor.chain().focus().toggleHeading({ level: 2 }).run()">
        H2
      </button>
      <button id="button" class="button is-success" v-if="mounted"
        @click="editor.chain().focus().toggleHeading({ level: 3 }).run()">
        H3
      </button>
      <button id="button" class="button is-success" v-if="mounted" @click="editor.chain().focus().toggleBold().run()">
        Bold
      </button>
      <button id="button" class="button is-success " v-if="mounted"
        @click="editor.chain().focus().toggleItalic().run()">
        <em>Italic</em>
      </button>
      <button id="button" class="button is-success " v-if="mounted"
        @click="editor.chain().focus().toggleSuperscript().run()">
        <sup>Sup</sup>
      </button>
      <button id="button" class="button is-success " v-if="mounted"
        @click="editor.chain().focus().toggleSubscript().run()">
        <sub>Sub</sub>
      </button>
      <button id="button" class="button is-success " v-if="mounted"
        @click="editor.chain().focus().toggleUnderline().run()">
        <u>Underline</u>
      </button>
      <button id="button" class="button is-success " v-if="mounted"
        @click="editor.chain().focus().toggleStrike().run()">
        <s>Strike</s>
      </button>
      <button id="button" class="button is-success " v-if="mounted"
        @click="editor.chain().focus().toggleBulletList().run()">
        List
      </button>
      <button id="button" class="button is-success " v-if="mounted" @click="editor.chain().focus().toggleCode().run()">
        Code
      </button>
      <button id="button" class="button is-success " v-if="mounted"
        @click="editor.chain().focus().toggleCodeBlock().run()">
        CodeBlock
      </button>
      <button id="button" class="button is-success " v-if="mounted" @click="editor.chain().focus().exitCode().run()">
        ExitCode
      </button>
      <button id="button" class="button is-success " v-if="mounted"
        @click="editor.chain().focus().setLink({ href: 'https://example.com' }).run()">
        set Link
      </button>
      <button id="button" class="button is-success " v-if="mounted"
        @click="editor.chain().focus().setHardBreak().run()">
        hardBreak
      </button>
      <button id="button" class="button is-success " v-if="mounted" @click="editor.chain().focus().unsetLink().run()">
        unset Link
      </button>
      <button id="button" class="button is-success " v-if="mounted"
        @click="editor.chain().focus().setImage({ src: 'https://example.com/foobar.png' }).run()">
        image
      </button>
      <button id="button" class="button is-info Table" v-if="mounted" @click="TableToogle = TableToogle ? false : true">
        <p>Table</p><img :class="{ imghover: TableToogle }" src="@/assets/angle.svg" />
      </button>
      <button id="button" class="button is-info " v-if="mounted && TableToogle"
        @click="editor.chain().focus().insertTable({ rows: 3, cols: 3, withHeaderRow: true }).run()">
        insertTable
      </button>
      <button id="button" class="button is-info  " v-if="mounted && TableToogle"
        @click="editor.chain().focus().addColumnBefore().run()">
        addColumnBefore
      </button>
      <button id="button" class="button is-info " v-if="mounted && TableToogle"
        @click="editor.chain().focus().addColumnAfter().run()">
        addColumnAfter
      </button>
      <button id="button" class="button is-info  " v-if="mounted && TableToogle"
        @click="editor.chain().focus().deleteColumn().run()">
        deleteColumn
      </button>
      <button id="button" class="button is-info  " v-if="mounted && TableToogle"
        @click="editor.chain().focus().addRowBefore().run()">
        addRowBefore
      </button>
      <button id="button" class="button is-info  " v-if="mounted && TableToogle"
        @click="editor.chain().focus().addRowAfter().run()">
        addRowAfter
      </button>
      <button id="button" class="button is-info  " v-if="mounted && TableToogle"
        @click="editor.chain().focus().deleteRow().run()">
        deleteRow
      </button>
      <button id="button" class="button is-info  " v-if="mounted && TableToogle"
        @click="editor.chain().focus().deleteTable().run()">
        deleteTable
      </button>
      <button id="button" class="button is-info  " v-if="mounted && TableToogle"
        @click="editor.chain().focus().mergeCells().run()">
        mergeCells
      </button>
      <button id="button" class="button is-info  " v-if="mounted && TableToogle"
        @click="editor.chain().focus().splitCell().run()">
        splitCell
      </button>
      <button id="button" class="button is-info  " v-if="mounted && TableToogle"
        @click="editor.chain().focus().toggleHeaderCell().run()">
        toggleHeaderCell
      </button>
      <button id="button" class="button is-info  " v-if="mounted && TableToogle"
        @click="editor.chain().focus().fixTables().run()">
        fixTables
      </button>
      <button id="button" class="button is-warning " v-if="mounted" @click="editor.chain().focus().undo().run()">
        undo⬅️
      </button>
      <button id="button" class="button is-warning " v-if="mounted" @click="editor.chain().focus().redo().run()">
        redo➡️
      </button>
      <button id="button" class="button is-warning " v-if="mounted"
        @click="editor.chain().focus().unsetAllMarks().run()">
        unset
      </button>
      <button id="button" class="button is-link" @click="save">save</button>
    </div>
    <editor-content :editor="editor" class="edit" />
    <br>
    <div class="block">click save after edited !</div>
  </div>

</template>
  
<script setup lang="ts">
import { mergeAttributes, Node } from '@tiptap/core'
import { useEditor, EditorContent } from '@tiptap/vue-3'
import Document from '@tiptap/extension-document'
import Paragraph from '@tiptap/extension-paragraph'
import Italic from '@tiptap/extension-italic'
import Subscript from '@tiptap/extension-subscript'
import Superscript from '@tiptap/extension-superscript'
import Code from '@tiptap/extension-code'
import CodeBlock from '@tiptap/extension-code-block'
import Text from '@tiptap/extension-text'
import Bold from '@tiptap/extension-bold'
import Underline from '@tiptap/extension-underline'
import Heading from '@tiptap/extension-heading'
import Bulletlist from '@tiptap/extension-bullet-list'
import Listitem from '@tiptap/extension-list-item'
import Image from '@tiptap/extension-image'
import Link from '@tiptap/extension-link'
import strike from '@tiptap/extension-strike'
import hardBreak from '@tiptap/extension-hard-break'
import Table from '@tiptap/extension-table'
import TableRow from '@tiptap/extension-table-row'
import TableHeader from '@tiptap/extension-table-header'
import hardCell from '@tiptap/extension-table-cell'
import History from '@tiptap/extension-history'


import { generateHTML } from '@tiptap/html'
import { onMounted, ref, defineEmits } from 'vue'
const emit = defineEmits(['save'])
const view = ref("click save!")
const TableToogle = ref(false)
const state = ref({
  index: history.state.index,
  content: history.state.content
})

const props = defineProps(['contenetjson' , "renderConfigJson"])

const mounted = ref(false);
Heading.configure({
  levels: [1, 2, 3],
})
const CodePre = Code.extend({
  // renderHTML({ HTMLAttributes }) {
  //   // return ['pre', HTMLAttributes, 0]
  // },
})
Link.configure({
  autolink: true,
})
const TablePre = Table.extend({
  addAttributes() {
    return {
      class: {
        default: "table"
      },
    }
  },
})
const Paragraphs = Paragraph.extend({
  renderHTML({ HTMLAttributes }) {
    return ['p', mergeAttributes(this.options.HTMLAttributes, HTMLAttributes), 0]
  },
  addAttributes() {
    return {
      // class: {
      //   default: "paragraph"
      // },
    }
  },
})
History.configure({
  depth: 100,
  newGroupDelay: 500,
})
const extensions = [
  Document,
  Paragraphs,
  Text,
  Bold,
  Italic,
  Subscript,
  Superscript,
  CodePre,
  CodeBlock,
  Underline,
  Heading,
  Bulletlist,
  Listitem,
  Image,
  Link,
  hardBreak,
  TablePre,
  TableRow,
  TableHeader,
  hardCell,
  strike,
  History,
]
let extensions_costum: any = extensions.slice();
let costum = props.renderConfigJson;
console.log(extensions_costum)
extensions_costum.forEach((extension: any, index: number) => {
  if (costum[extension.name]) {
    if (costum[extension.name].tag || costum[extension.name].HTMLAttributes) {
      extension = extension.extend({
        addAttributes() {
          return costum[extension.name].HTMLAttributes
        },
        renderHTML({ HTMLAttributes } : any) {
          return [costum[extension.name].tag, mergeAttributes(this.options.HTMLAttributes, HTMLAttributes), 0]
        },
      })
    }
    extensions_costum[index] = extension
  }
})
console.log(extensions_costum)
const editor: any = useEditor({
  extensions: extensions,
  content: props.contenetjson,
  autofocus: true,
  editable: true,
  injectCSS: false,
})

onMounted(() => {
  mounted.value = true;
  console.log(editor.value)
  console.log(editor.value.chain().focus().toggleBold())
})
function save() {
  console.log(view.value = editor.value.getJSON())
  // console.log(view.value = encodeURI(JSON.stringify(editor.value.getJSON())))
  console.log(generateHTML(editor.value.getJSON(view.value), extensions))
  console.log(generateHTML(editor.value.getJSON(view.value), extensions_costum))
  emit('save', editor.value.getJSON(), generateHTML(editor.value.getJSON(view.value), extensions_costum));    
}
</script>
<style scoped lang="scss">
.main {
  background-color: var(--has-background-lightest-grey);
  padding: 0.3em;
  border-radius: 4px;
  margin: 0 0 0 0;
}



.bold {
  font-weight: bold;
}


.buttons {
  position: sticky;
  top: 60px;
  margin: 0 0 0 0;

  #button {
    margin-right: 0rem;
    margin: 0 4px 4px 0;
    display: inline-block;
    opacity: 0.9;
    transition: ease 0.2s;

  }

  #button:hover {
    transform: scale(1.2);
  }

  #button.Table {
    display: flex;
    flex-direction: row;

    p {
      padding: 0 0 0 0;
      margin: 0 0 0 0;
      line-height: 1.5em;
      width: auto;
      display: inline-block;
      padding-right: 1em;
    }

    img {
      height: 1.5em;
      width: 1.5em;
      transform: rotate(-90deg);
      transition: transform 0.5s ease;
    }

    .imghover {
      transform: rotate(90deg);
    }
  }
}



button:hover {
  opacity: 1;
}

button * {
  color: #fff;
}

sup {
  vertical-align: top;
}

sub {
  vertical-align: sub;
}
</style>
<style lang="scss">
.ProseMirror {
  padding: 0.4em;
  min-height: 25ch;
  outline: auto;
  outline-color: var(--has-border-dark);
  outline-style: dashed;
  // height: calc(100% + 50px);
}

.ProseMirror-focused {
  outline: auto;

  >*+* {
    margin-top: 0.75em;
  }
}
</style>