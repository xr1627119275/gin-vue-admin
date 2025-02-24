<template>
  <div class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
    <h2 class="text-2xl font-bold mb-4">Scan Results</h2>
    <div v-if="scanning" class="text-center">
      <p class="text-lg">Scanning in progress...</p>
      <div class="loader ease-linear rounded-full border-4 border-t-4 border-gray-200 h-12 w-12 mb-4 mx-auto"></div>
    </div>
    <div v-else-if="results.length === 0" class="text-center">
      <p class="text-lg">No results yet. Start a scan to see vulnerabilities.</p>
    </div>
    <div v-else>
      <table class="min-w-full">
        <thead>
        <tr>
          <th class="px-4 py-2 text-left">ID</th>
          <th class="px-4 py-2 text-left">Severity</th>
          <th class="px-4 py-2 text-left">Name</th>
          <th class="px-4 py-2 text-left">Target</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="result in results" :key="result.id">
          <td class="border px-4 py-2">{{ result.id }}</td>
          <td class="border px-4 py-2">
              <span :class="severityClass(result.severity)">
                {{ result.severity }}
              </span>
          </td>
          <td class="border px-4 py-2">{{ result.name }}</td>
          <td class="border px-4 py-2">{{ result.target }}</td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
  import { defineProps } from 'vue';

  const props = defineProps({
    results: {
      type: Array,
      default: () => []
    },
    scanning: {
      type: Boolean,
      default: false
    }
  });

  const severityClass = (severity) => {
    switch (severity.toLowerCase()) {
      case 'high':
        return 'bg-red-500 text-white px-2 py-1 rounded';
      case 'medium':
        return 'bg-yellow-500 text-white px-2 py-1 rounded';
      case 'low':
        return 'bg-green-500 text-white px-2 py-1 rounded';
      default:
        return 'bg-gray-500 text-white px-2 py-1 rounded';
    }
  };
</script>

<style scoped>
  .loader {
    border-top-color: #3498db;
    -webkit-animation: spinner 1.5s linear infinite;
    animation: spinner 1.5s linear infinite;
  }

  @-webkit-keyframes spinner {
    0% { -webkit-transform: rotate(0deg); }
    100% { -webkit-transform: rotate(360deg); }
  }

  @keyframes spinner {
    0% { transform: rotate(0deg); }
    100% { transform: rotate(360deg); }
  }
</style>