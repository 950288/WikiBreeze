<template>
  <div class="main content">
    <h3 class="reminder">You are editing <u>{{ state.content }}</u> on <u>{{ state.index }}</u> page.</h3>
    <div class="buttons">
      <div class="button Heading" id="buttons-toggle">
        <div class="buttons-wrap">
          <button v-for="(item) in costum.otherConfigurations ? costum.otherConfigurations.headingLevels : [1, 2, 3]"
            id="button" class="button is-success" v-if="mounted"
            @click="editor.chain().focus().toggleHeading({ level: item }).run()">
            H{{ item }}
          </button>
        </div>
      </div>
      <button id="button" class="button is-success" v-if="mounted" @click="editor.chain().focus().toggleBold().run()">
        Bold
      </button>
      <button id="button" class="button is-success " v-if="mounted" @click="editor.chain().focus().toggleItalic().run()">
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
      <button id="button" class="button is-success " v-if="mounted" @click="editor.chain().focus().toggleStrike().run()">
        <s>Strike</s>
      </button>
      <button id="button" class="button is-success " v-if="mounted"
        @click="editor.chain().focus().toggleBulletList().run()">
        List
      </button>
      <button id="button" class="button is-success " v-if="mounted" @click="editor.chain().focus().toggleCode().run()">
        Code
      </button>
      <div class="button CodeBlock" id="buttons-toggle">
        <div class="buttons-wrap">
          <button id="button" class="button is-success " v-if="mounted"
            @click="editor.chain().focus().toggleCodeBlock().run()">
            CodeBlock
          </button>
          <button id="button" class="button is-success " v-if="mounted"
            @click="editor.chain().focus().toggleCodeBlock().updateAttributes('codeBlock', { language: 'python' }).run()">
            python
          </button>
          <button id="button" class="button is-success " v-if="mounted"
            @click="editor.chain().focus().toggleCodeBlock().updateAttributes('codeBlock', { language: 'css' }).run()">
            css
          </button>
          <button id="button" class="button is-success " v-if="mounted"
            @click="editor.chain().focus().toggleCodeBlock().updateAttributes('codeBlock', { language: 'html' }).run()">
            html
          </button>
          <button id="button" class="button is-success " v-if="mounted"
            @click="editor.chain().focus().toggleCodeBlock().updateAttributes('codeBlock', { language: 'js' }).run()">
            js
          </button>
        </div>
      </div>
      <button id="button" class="button is-success " v-if="mounted" @click="editor.chain().focus().exitCode().run()">
        ExitCodeBlock
      </button>
      <div class="button setLink" id="buttons-toggle">
        <div class="buttons-wrap">
          <button id="button" class="button is-success " v-if="mounted" @click="setLink()">
            SetLink
          </button>
          <button id="button" class="button is-success " v-if="mounted" @click="editor.chain().focus().unsetLink().run()">
            Unset
          </button>
        </div>
      </div>
      <button id="button" class="button is-success " v-if="mounted" @click="editor.chain().focus().setHardBreak().run()">
        Break
      </button>
      <button id="button" class="button is-success "
        v-if="mounted && !(costum.otherConfigurations.imagePro == 'true' && costum.otherConfigurations.imageMustContainNote == 'true')"
        @click="setImageURL()">
        Image
      </button>
      <button id="button" class="button is-success " v-if="mounted && costum.otherConfigurations.imagePro == 'true'"
        @click="(_$event: any) => setImageProURL()">
        ImagePro
      </button>
      <div class="button textAlign" id="buttons-toggle">
        <div class="buttons-wrap">
          <button id="button" class="button is-success " v-if="mounted">
            Align
          </button>
          <button id="button" class="button is-success " v-if="mounted"
            @click="editor.chain().focus().setTextAlign('left').run()">
            Left
          </button>
          <button id="button" class="button is-success " v-if="mounted"
            @click="editor.chain().focus().setTextAlign('right').run()">
            Right
          </button>
          <button id="button" class="button is-success " v-if="mounted"
            @click="editor.chain().focus().setTextAlign('center').run()">
            Center
          </button>
          <!-- <button id="button" class="button is-success " v-if="mounted"
            @click="editor.chain().focus().setTextAlign('justify').run()">
            Justify
          </button> -->
        </div>
      </div>
      <button id="button" class="button is-info Table" v-if="mounted" @click="TableToogle = TableToogle ? false : true">
        <p>Table</p><img :class="{ imghover: TableToogle }" src="@/assets/angle.svg" />
      </button>
      <button id="button" class="button is-info "
        v-if="mounted && TableToogle && !(costum.otherConfigurations.tablePro == 'true' && costum.otherConfigurations.tableMustContainNote == 'true')"
        @click="editor.chain().focus().insertTable({ rows: 3, cols: 3, withHeaderRow: true }).run()">
        addTable
      </button>
      <button id="button" class="button is-info "
        v-if="mounted && TableToogle && costum.otherConfigurations.tablePro == 'true'"
        @click="editor.chain().focus().insertTablePro({ rows: 3, cols: 3, withHeaderRow: true }).run()">
        TablePro
      </button>
      <button id="button" class="button is-info  " v-if="mounted && TableToogle"
        @click="editor.chain().focus().addColumnBefore().run()">
        addColBefore
      </button>
      <button id="button" class="button is-info " v-if="mounted && TableToogle"
        @click="editor.chain().focus().addColumnAfter().run()">
        addColAfter
      </button>
      <button id="button" class="button is-info  " v-if="mounted && TableToogle"
        @click="editor.chain().focus().deleteColumn().run()">
        delCol
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
        delRow
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
        setHeader
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
      <button id="button" class="button is-warning " v-if="mounted" @click="editor.chain().focus().unsetAllMarks().run()">
        RemoveMark
      </button>
      <button id="button" class="button is-warning " v-if="mounted" @click="editor.chain().focus().clearContent().run()">
        RemoveAll
      </button>
      <button id="button" class="button is-link" @click="save">save</button>
    </div>
    <editor-content :editor="editor" class="edit" />
    <br>
    <div class="block">Press CTRL+S to save !</div>
  </div>
</template>
  
<script setup lang="ts">
import { useEditor, EditorContent, type Extensions } from '@tiptap/vue-3'
import Document from '@tiptap/extension-document'
import Paragraphs from './Note'
import Italic from '@tiptap/extension-italic'
import Subscript from '@tiptap/extension-subscript'
import Superscript from '@tiptap/extension-superscript'
import Code from '@tiptap/extension-code'
import codeBlock from '@tiptap/extension-code-block-lowlight'
import Text from '@tiptap/extension-text'
import Bold from '@tiptap/extension-bold'
import Underline from '@tiptap/extension-underline'
import Heading from '@tiptap/extension-heading'
import Bulletlist from '@tiptap/extension-bullet-list'
import Listitem from '@tiptap/extension-list-item'
import Image from '@tiptap/extension-image'
import Link from '@tiptap/extension-link'
import Strike from '@tiptap/extension-strike'
import HardBreak from '@tiptap/extension-hard-break'
import TablePro from './TablePro'
import ImagePro from './ImagePro'
import Table from '@tiptap/extension-table'
import TextAlign from '@tiptap/extension-text-align'
import TableRow from '@tiptap/extension-table-row'
import tableHeader from '@tiptap/extension-table-header'
import tableCell from '@tiptap/extension-table-cell'
import History from '@tiptap/extension-history'
import css from 'highlight.js/lib/languages/css'
import js from 'highlight.js/lib/languages/javascript'
import html from 'highlight.js/lib/languages/xml'
import python from 'highlight.js/lib/languages/python'
import { lowlight } from 'lowlight'
import { generateHTML } from '@tiptap/html'
import Gapcursor from '@tiptap/extension-gapcursor'
import { onMounted, ref, watch, type Ref } from 'vue'
import { useInputStore } from "@/main";

const inputStore = useInputStore();

lowlight.registerLanguage('html', html)
lowlight.registerLanguage('css', css)
lowlight.registerLanguage('js', js)
lowlight.registerLanguage('python', python)


const emit = defineEmits(['save'])
const TableToogle = ref(false)
const state = ref({
  index: history.state.index,
  content: history.state.content
})

const props = defineProps(['contentJson', "renderConfigJson", "uploadEnable"])

const mounted = ref(false);
let costum = props.renderConfigJson;
Heading.configure({
  levels: costum.otherConfigurations ? costum.otherConfigurations.headingLevels : [1, 2, 3],
})
Link.configure({
  autolink: true,
})
History.configure({
  depth: 100,
  newGroupDelay: 500,
})
const TableHeader = tableHeader.extend({
  content: 'paragraph+',
})
const TableCell = tableCell.extend({
  content: 'paragraph+',
})
let CodeBlock = codeBlock.extend({
  isolating: true,
})

const extensions = [
  Document,
  Paragraphs.configure({
    HTMLAttributes: {
      class: 'my-custom-class',
    },
  }),
  Text,
  Bold,
  Italic,
  Subscript,
  Superscript,
  Code,
  CodeBlock.configure({
    defaultLanguage: null,
    lowlight,
  }),
  Underline,
  Heading,
  Bulletlist,
  Listitem,
  Image,
  ImagePro,
  Link,
  HardBreak,
  TextAlign.configure({
    types: ['heading', 'paragraph'],
  }),
  Table,
  TablePro,
  TableRow,
  TableHeader,
  TableCell,
  Strike,
  History,
  Gapcursor
]

let extensions_costum: Extensions = [];

extensions.forEach((extension, index) => {
  if (costum[extension.name] && extension.type == "node") {
    let newExtension = costum[extension.name].tag ? extension.extend({
      renderHTML({ HTMLAttributes }: any) {
        const val:any = (this as any).parent?.({ HTMLAttributes });
        if (val && val[0]) {
          val[0] = costum[extension.name].tag
        }
        return val
      }
    }) : extension.extend()

    if (costum[extension.name].HTMLAttributes && newExtension.options) {
      newExtension.configure({
        HTMLAttributes: costum[extension.name].HTMLAttributes,
      })

      newExtension.options.HTMLAttributes = JSON.parse(JSON.stringify(costum[extension.name].HTMLAttributes))
    }
    extensions_costum[index] = newExtension
  } else if (costum[extension.name] && extension.type == "Mark") {
    extensions_costum[index] = extension.extend()
  } else {
    extensions_costum[index] = extension.extend()
  }
})

const editor: Ref = useEditor({
  extensions: extensions,
  content: props.contentJson,
  autofocus: true,
  editable: true,
  injectCSS: false,
})

onMounted(() => {
  mounted.value = true;
  document.addEventListener('keydown', (e) => {
    if (e.ctrlKey && (e.key == 's' || e.key == 'S')) {
      e.preventDefault();
      save();
    }
    if (e.ctrlKey && (e.key == 'Z')) {
      e.preventDefault();
      editor.value.chain().focus().undo().run()
    }
    if (e.ctrlKey && (e.key == 'Y')) {
      e.preventDefault();
      editor.value.chain().focus().redo().run()
    }
  });
})
function save() {
  let html = generateHTML(editor.value.getJSON(), extensions_costum)
  emit('save', editor.value.getJSON(), html);
}
function setImageURL() {
  let recall = inputStore.input(
    "Insert image",
    props.uploadEnable,

  );
  watch(recall, (val) => {
    editor.value.chain().focus().setImage({ src: val }).run()
  })
}
function setImageProURL() {
  let recall = inputStore.input(
    "Insert image",
    props.uploadEnable
  );
  watch(recall, (val) => {
    editor.value.chain().focus().insertImagePro(val).run()
  })
}
function setLink() {
  let recall = inputStore.input(
    "Insert Link",
    false
  );
  watch(recall, (val) => {
    editor.value.chain().focus().setLink({ href: val, target: '_blank' }).run()
  })
}
</script>
<style scoped lang="scss">
html #button,
html #buttons-toggle {
  --buttons-height: 2.5rem;

  @media (max-width: 1350px) {
    --buttons-height: 1.7rem;
  }

}

.main {
  background-color: var(--has-background-lightest-grey);
  padding: 0.3em;
  border-radius: 4px;
  margin: 0 0 0 0;

  .reminder {
    margin-bottom: .3em;
  }
}

.bold {
  font-weight: bold;
}

.buttons {
  position: sticky;
  top: 60px;
  margin: 0 0 0 0;
  z-index: 1;

  #button {
    margin-right: 0rem;
    margin: 0 4px 4px 0;
    height: var(--buttons-height);
    display: inline-block;
    opacity: 0.9;
    transition: ease 0.2s;
  }

  @media (max-width: 1350px) {
    #button {
      padding: 0.2rem;
      line-height: 1.3rem;
    }
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

    @media (max-width: 1350px) {
      p {
        padding-right: .2rem;
      }
    }

    img {
      height: 1.4em;
      width: 1.4em;
      transform: rotate(-90deg);
      transition: transform 0.5s ease;
    }

    .imghover {
      transform: rotate(90deg);
    }
  }
}

#buttons-toggle {
  position: relative;
  margin: 0 4px 4px 0;
  padding: 0;
  flex-direction: column;
  justify-content: flex-start;
  border-width: 0px;
  height: var(--buttons-height);
  opacity: 0.9;
  background-color: hsl(153, 53%, 53%);

  .buttons-wrap {
    position: absolute;
    margin: 0;
    top: 0;
    left: 0;
    padding: 0;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    height: var(--buttons-height);
    overflow: hidden;
    transition: ease 0.2s;
    opacity: 1;
    width: 100%;


    #button {
      margin: 0;
      transition: ease 0.2s;
      transform: scale(1);
      opacity: 1;
      display: inline;
      height: var(--buttons-height);
      padding-bottom: calc(0.5em - 1px);
      padding-top: calc(0.5em - 1px);
      width: 100%;
    }

    #button:first-child {
      padding-right: 1.9em;

    }

    .button:first-child::after {
      content: '';
      height: 1.5em;
      width: 1.5em;
      transform: rotate(0deg);
      transition: transform 0.5s ease;
      margin-left: .5em;
      display: block;
      position: absolute;
      top: 50%;
      transform: translateY(-50%);
      right: .5rem;
      background: url(@/assets/angle.svg);
    }

    @media (max-width: 1350px) {
      #button {
        padding: 0.2rem;
      }

      .button:first-child::after {
        right: .2rem;
      }
    }

    .button:first-child:hover::after {
      transform: translateY(-50%) rotate(180deg);
    }
  }

  .buttons-wrap:hover {
    height: auto;
    transform: scale(1.1);
    opacity: 1;
    // z-index: 2;

    .button:hover {
      opacity: 1;
    }
  }

  .button:first-child {
    img {
      display: block;
    }
  }
}

#buttons-toggle.Heading {
  width: 5em;
}

#buttons-toggle.textAlign {
  width: 6em;
}

#buttons-toggle.CodeBlock {
  width: 8.5em;
}

#buttons-toggle.setLink {
  width: 7em;
}

#buttons-toggle.CodeBlock:hover,
#buttons-toggle.textAlign:hover,
#buttons-toggle.Heading:hover,
#buttons-toggle.setLink:hover {
  z-index: 2;
}


@media (max-width: 1350px) {
  #buttons-toggle.Heading {
    width: 4em;
  }

  #buttons-toggle.textAlign {
    width: 4.5em;
  }

  #buttons-toggle.CodeBlock {
    width: 7em;
  }

  #buttons-toggle.setLink {
    width: 5.5em;
  }
}

button:hover {
  opacity: 1;
}

button * {
  color: #fff;
}
</style>
<style lang="scss">
code {
  font-size: 1.1rem;
  padding: 0.15em;
  border-radius: 0.15em;
  background-color: rgba(#616161, 0.1);
  color: var(--has-text-dark-grey);
  box-decoration-break: clone;
}

.ProseMirror pre code {
  font-size: 1.1rem;
}

sup {
  vertical-align: top;
}

sub {
  vertical-align: sub;
}

a {
  color: var(--has-text-link);
}
</style>
