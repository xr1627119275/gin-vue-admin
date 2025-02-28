<script setup>

  import { nextTick, onMounted, ref, unref, watch } from 'vue'
  import CodeMirror from 'codemirror'
  import 'codemirror/addon/lint/lint.css'
  import 'codemirror/lib/codemirror.css'
  import 'codemirror/theme/monokai.css'
  import 'codemirror/mode/yaml/yaml'
  import 'codemirror/addon/lint/lint'
  import 'codemirror/addon/lint/yaml-lint'
  import YAML from 'js-yaml';
  // window.jsyaml = require("js-yaml");
  const visible =  defineModel("visible")
  const code =  defineModel("code")
  const yamlEditor = ref(null)
  const codeRef = ref(null)
  watch(visible, async (val)=> {
    if (!val) return
    await nextTick()
    if (!yamlEditor.value)
      yamlEditor.value = CodeMirror.fromTextArea(codeRef.value, {
        lineNumbers: true, // 显示行号
        height: '60vh',
        mode: 'text/x-yaml', // 语法 model
        gutters: ['CodeMirror-lint-markers'],  // 语法检查器
        theme: 'monokai', // 编辑器主题
        readOnly: true,
      });
    unref(yamlEditor).setValue(code.value);
  })
</script>

<template>
  <el-dialog :footer-class="'display-none none'"  width="70vw"  v-model="visible">
    <div id="CodeContainer" class="h-[60vh]">
      <textarea  ref="codeRef"></textarea>
    </div>
  </el-dialog>
</template>

<style lang="scss" >
  #CodeContainer {
    overflow-y: scroll ;
    .CodeMirror {
      height: auto !important;
    }
  }
</style>