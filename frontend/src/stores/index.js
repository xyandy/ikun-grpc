import { defineStore } from 'pinia';
import { reactive } from 'vue';

export const alertStore = defineStore('alert', () => {
  const state = reactive({
    show: false,
    type: '',
    title: '',
  });

  function set(type, title) {
    state.show = true;
    state.type = type;
    state.title = title;
  }

  function setSuccessTitle(title) {
    set('success', title);
  }

  function setErrorTitle(title) {
    set('error', title);
  }

  return { state, setSuccessTitle, setErrorTitle };
});

// const counter = useCounterStore();
// counter.increment()
// counter.count++
