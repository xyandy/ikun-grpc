<template>
  <div ref="container"></div>
</template>

<script setup>
import JSONEditor from 'jsoneditor';
import 'jsoneditor/dist/jsoneditor.css';
import { nextTick, onBeforeUnmount, onMounted, reactive, ref, watch } from 'vue';

const props = defineProps({
  modelValue: Object,
  // https://github.com/josdejong/jsoneditor/blob/develop/docs/api.md#configuration-options
  options: Object,
});
const emits = defineEmits(['update:modelValue']);

const container = ref(null);
const state = reactive({
  editor: null,
  json: {},
  internalChange: false,
});

watch(
  () => props.modelValue,
  async (value) => {
    if (!state.internalChange) {
      state.json = value;
      state.editor && state.editor.set(value);
    }
  },
  { immediate: true }
);

onMounted(() => {
  const finalOptions = {
    language: 'en',
    // mode: 'code',
    // modes: ['view', 'code', 'text', 'preview'],
    enableSort: false,
    enableTransform: false,
    statusBar: false,
    ...props.options,
    onChange: async () => {
      try {
        if (!state.editor) {
          return;
        }
        const json = state.editor.get() || {};
        state.json = json;
        emits('update:modelValue', json);
        state.internalChange = true;

        await nextTick(() => {
          state.internalChange = false;
        });
      } catch (err) {
        console.error(err.message);
      }
    },
  };

  const json = props.modelValue || {};
  state.editor = new JSONEditor(container.value, finalOptions, json);
});

onBeforeUnmount(() => {
  if (state.editor) {
    state.editor.destroy();
  }
});
</script>

<style>
.jsoneditor-repair {
  display: none;
}

.jsoneditor-poweredBy {
  display: none;
}
</style>
