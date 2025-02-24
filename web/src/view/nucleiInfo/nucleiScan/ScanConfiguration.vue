<template>
  <div class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
    <h2 class="text-2xl font-bold mb-4">Scan Configuration</h2>
    <form @submit.prevent="startScan">
      <div class="mb-4">
        <label class="block text-gray-700 text-sm font-bold mb-2" for="target">
          Target URL
        </label>
        <input
          class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          id="target"
          v-model="config.target"
          type="url"
          placeholder="https://example.com"
          required
        >
      </div>
      <div class="mb-4">
        <label class="block text-gray-700 text-sm font-bold mb-2" for="templates">
          Templates
        </label>
        <input
          class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          id="templates"
          v-model="config.templates"
          type="text"
          placeholder="cves,vulnerabilities"
        >
      </div>
      <div class="mb-4">
        <label class="flex items-center">
          <input type="checkbox" v-model="config.verbose" class="form-checkbox">
          <span class="ml-2">Verbose Output</span>
        </label>
      </div>
      <div class="flex items-center justify-between">
        <button
          class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
          type="submit"
        >
          Start Scan
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
  import { reactive } from 'vue';

  const config = reactive({
    target: '',
    templates: '',
    verbose: false
  });

  const emit = defineEmits(['start-scan']);

  const startScan = () => {
    emit('start-scan', { ...config });
  };
</script>