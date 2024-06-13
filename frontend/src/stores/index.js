import { computed, ref } from 'vue';
import { defineStore } from 'pinia';

export const useEndpointStore = defineStore('endpoint', () => {
  const endpoint = ref('hello');

  // function increment() {
  //   count.value++;
  // }

  return { endpoint };
});

// const counter = useCounterStore();
// counter.increment()
// counter.count++
