<template>
  <node-view-wrapper class="citation-wrapper" contenteditable="false">
      [<a>
        <node-view-content contenteditable="true">1</node-view-content>
      </a>]
  </node-view-wrapper>
</template>

<script lang="ts">
import { NodeViewContent, nodeViewProps, NodeViewWrapper } from '@tiptap/vue-3'
export default {
  props: nodeViewProps,
  components: {
    NodeViewWrapper,
    NodeViewContent,
  },
}
</script>
<style lang="scss">
.citation-wrapper{
  display: inline-block;  
  vertical-align: top;
  div{
    display: inline;
  }
}
</style>