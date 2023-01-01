<template>
  <div class="main content">
    <h3>You are editing <u>{{ state.index }}</u> on page <u>{{ state.node }}.</u></h3>
    <div class="buttons">

    <button class="button is-success" v-if="mounted" @click="editor.chain().focus().toggleHeading({ level: 1 }).run()">
      H1
    </button>
    <button class="button is-success" v-if="mounted" @click="editor.chain().focus().toggleHeading({ level: 2 }).run()">
      H2
    </button>
    <button class="button is-success" v-if="mounted" @click="editor.chain().focus().toggleHeading({ level: 3 }).run()">
      H3
    </button>
    <button class="button is-success" v-if="mounted" @click="editor.chain().focus().toggleBold().run()">
      Bold
    </button>
    <button class="button is-success " v-if="mounted" @click="editor.chain().focus().toggleItalic().run()">
      <em>Italic</em>
    </button>
    <button class="button is-success " v-if="mounted" @click="editor.chain().focus().toggleSuperscript().run()">
      <sup>Sup</sup>
    </button>
    <button class="button is-success " v-if="mounted" @click="editor.chain().focus().toggleSubscript().run()">
      <sub>Sub</sub>
    </button>
    <button class="button is-success " v-if="mounted" @click="editor.chain().focus().toggleUnderline().run()">
      Underline
    </button>
    <button class="button is-success " v-if="mounted" @click="editor.chain().focus().toggleBulletList().run()">
      List
    </button>
    <button class="button is-success " v-if="mounted" @click="editor.chain().focus().toggleCode().run()">
      Code
    </button>
    <button class="button is-success " v-if="mounted" @click="editor.chain().focus().setLink({ href: 'https://example.com' }).run()">
      set Link
    </button>
    <button class="button is-success " v-if="mounted" @click="editor.chain().focus().unsetLink().run()">
      unset Link
    </button>
    <button class="button is-success " v-if="mounted" @click="editor.chain().focus().setImage({ src: 'https://example.com/foobar.png' }).run()">
      image
    </button>
    <button class="button is-warning " v-if="mounted" @click="editor.chain().focus().unsetAllMarks().run()">
      unset
    </button>
    <button class="button is-link" @click="save">save</button>
  </div>
  <editor-content :editor="editor" class="edit" />
  <br>
  <div class="block">click save after edited !</div>
  </div>

</template>
  
<script setup lang="ts">
import { useEditor, EditorContent } from '@tiptap/vue-3'
import Document from '@tiptap/extension-document'
import Paragraph from '@tiptap/extension-paragraph'
import Italic from '@tiptap/extension-italic'
import Subscript from '@tiptap/extension-subscript'
import Superscript from '@tiptap/extension-superscript'
import Code from '@tiptap/extension-code'
import Text from '@tiptap/extension-text'
import Bold from '@tiptap/extension-bold'
import Underline from '@tiptap/extension-underline'
import Heading from '@tiptap/extension-heading'
import Bulletlist from '@tiptap/extension-bullet-list'
import Listitem from '@tiptap/extension-list-item'
import Image from '@tiptap/extension-image'
import Link from '@tiptap/extension-link'

import { generateHTML } from '@tiptap/html'
import { onMounted, ref , defineEmits } from 'vue'
const emit = defineEmits(['save'])
const view = ref("click save!")

const state = ref({
  index: history.state.index,
  node: history.state.node})

const props = defineProps(['json'])

const mounted = ref(false);
Heading.configure({
  levels: [1, 2 , 3],
})
const CodePre = Code.extend({
  renderHTML({ HTMLAttributes }) {
    return ['pre', HTMLAttributes, 0]
  },
})
Link.configure({
  autolink: true,
})
const extensions = [
  Document,
  Paragraph,
  Text,
  Bold,
  Italic,
  Subscript,
  Superscript,
  CodePre,
  Underline,
  Heading,
  Bulletlist,
  Listitem,
  Image,
  Link
]
const editor: any = useEditor({
  extensions: extensions,
  content: props.json,
  autofocus: true,
  editable: true,
  injectCSS: false,
})

onMounted(() => {
  mounted.value = true;
  console.log(editor.value)
  console.log(editor.value.chain().focus().toggleBold())
})
function save(){
  console.log(view.value = editor.value.getJSON())
  console.log(view.value = encodeURI(JSON.stringify(editor.value.getJSON())))
  console.log(view.value = generateHTML(editor.value.getJSON(JSON.parse(decodeURI(view.value))),extensions))
  emit('save',editor.value.getJSON());
}
</script>
<style scoped lang="scss">
.main{
  background-color: var(--has-background-lightest-grey);
  padding: 0.3em;
  border-radius: 4px;
  margin: 0 0 0 0;
}
.bold {
  font-weight: bold;
}
.buttons{
  margin: 0 0 0 0;
}
.buttons{
  position: sticky;
  top: 60px;
  // opacity: 0.6;
  // background-color: var(--has-background-lightest-grey);
}
button {
  margin: 0 5px 5px 0;
  display: inline-block;
  opacity: 0.9;
}

button:hover{
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
  padding-left: 0.4em;
  min-height: 25ch;
  outline: auto;
  outline-color: var(--has-border-dark);
  outline-style: dashed;
}

.ProseMirror-focused {
  outline: auto;
  >*+* {
    margin-top: 0.75em;
  }
}
</style>