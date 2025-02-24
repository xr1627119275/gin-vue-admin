<script setup lang="ts">
  import { ref, computed } from 'vue'

  const target = ref('')
  const scanning = ref(false)
  const results = ref<any[]>([])
  const showPocModal = ref(false)
  const searchQuery = ref('')
  const selectedPocs = ref<string[]>([])
  import nucleiInfo from '../nucleiInfo/nucleiInfo.vue'
  interface Poc {
    id: string
    name: string
    description: string
    severity: string
    category: string
  }



  const pocs = ref<Poc[]>([
    {
      id: 'CVE-2021-1234',
      name: 'SQL注入漏洞检测',
      description: '检测常见的SQL注入漏洞',
      severity: 'HIGH',
      category: '注入漏洞'
    },
    {
      id: 'CVE-2021-5678',
      name: 'XSS跨站脚本检测',
      description: '检测跨站脚本攻击漏洞',
      severity: 'MEDIUM',
      category: 'XSS'
    },
    {
      id: 'CVE-2021-9012',
      name: '目录遍历漏洞检测',
      description: '检测目录遍历漏洞',
      severity: 'HIGH',
      category: '文件系统'
    },
    {
      id: 'CVE-2022-1111',
      name: '命令注入漏洞检测',
      description: '检测命令注入漏洞',
      severity: 'CRITICAL',
      category: '注入漏洞'
    }
  ])

  const filteredPocs = computed(() => {
    return pocs.value.filter(poc =>
      poc.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      poc.description.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      poc.category.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  })

  const togglePoc = (pocId: string) => {
    const index = selectedPocs.value.indexOf(pocId)
    if (index === -1) {
      selectedPocs.value.push(pocId)
    } else {
      selectedPocs.value.splice(index, 1)
    }
  }

  const startScan = async () => {

    scanning.value = true
    results.value = []

    // 模拟扫描结果
    setTimeout(() => {
      results.value = selectedPocs.value.map(pocId => {
        const poc = pocs.value.find(p => p.id === pocId)
        const isVulnerable = Math.random() > 0.5 // 随机模拟是否存在漏洞
        return {
          id: poc?.id,
          name: poc?.name,
          severity: poc?.severity,
          category: poc?.category,
          description: isVulnerable ? `发现${poc?.name}漏洞` : `未发现${poc?.name}漏洞`,
          vulnerable: isVulnerable,
          location: isVulnerable ? `${target.value}/api/vulnerable-endpoint` : '',
          details: poc?.description,
          recommendations: isVulnerable ? [
            '及时更新相关组件到最新版本',
            '检查并修复相关代码逻辑',
            '增加相应的安全防护措施'
          ] : [],
          timestamp: new Date().toISOString()
        }
      })
      scanning.value = false
    }, 2000)
  }

  const getSeverityColor = (severity: string) => {
    const colors = {
      'CRITICAL': 'text-purple-600',
      'HIGH': 'text-red-600',
      'MEDIUM': 'text-orange-500',
      'LOW': 'text-yellow-500',
      'INFO': 'text-blue-500'
    }
    return colors[severity as keyof typeof colors] || 'text-gray-600'
  }

  const getSeverityBgColor = (severity: string) => {
    const colors = {
      'CRITICAL': 'bg-purple-100',
      'HIGH': 'bg-red-100',
      'MEDIUM': 'bg-orange-100',
      'LOW': 'bg-yellow-100',
      'INFO': 'bg-blue-100'
    }
    return colors[severity as keyof typeof colors] || 'bg-gray-100'
  }

  const getSeverityText = (severity: string) => {
    const texts = {
      'CRITICAL': '严重',
      'HIGH': '高危',
      'MEDIUM': '中危',
      'LOW': '低危',
      'INFO': '信息'
    }
    return texts[severity as keyof typeof texts] || severity
  }

  const formatDate = (dateString: string) => {
    const date = new Date(dateString)
    return date.toLocaleString('zh-CN', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
      second: '2-digit'
    })
  }
</script>

<template>
  <div class="max-w-4xl mx-auto p-6">
<!--    <h1 class="text-2xl font-bold mb-6">Nuclei 漏洞扫描器</h1>-->

    <div class="bg-white dark:bg-gray-800 rounded-lg shadow p-6 mb-6">
      <form @submit.prevent="startScan" class="space-y-4">
        <div>
          <label for="target" class="block text-sm font-medium mb-2">目标 URL</label>
          <textarea
            id="target"
            v-model="target"
            type="url"
            placeholder="https://example.com"
            required
            class="w-full px-4 py-2 rounded-lg border border-gray-300 dark:border-gray-600 focus:ring-2 focus:ring-blue-500 dark:bg-gray-700"
          />
        </div>

        <div>
          <button
            type="button"
            @click="showPocModal = true"
            class="w-full bg-gray-100 hover:bg-gray-200 text-gray-800 font-medium py-2 px-4 rounded-lg border border-gray-300"
          >
            选择 漏洞库 ({{ selectedPocs.length ? `已选择${selectedPocs.length}个` : '默认全选'}})
          </button>
        </div>

        <button
          type="submit"
          :disabled="scanning"
          class="w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {{ scanning ? '扫描中...' : '开始扫描' }}
        </button>
      </form>
    </div>

    <!-- POC 选择弹窗 -->
    <div v-if="showPocModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4">
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow-xl max-w-[800px] w-full max-h-[80vh] flex flex-col">
        <div class="p-4 border-b border-gray-200 dark:border-gray-700">
          <div class="flex justify-between items-center">
            <h2 class="text-xl font-semibold">选择 POC</h2>
            <button
              @click="showPocModal = false"
              class="text-gray-500 hover:text-gray-700"
            >
              ✕
            </button>
          </div>
        </div>

        <div class="overflow-y-auto flex-1 p-4">
          <nuclei-info type="select"/>
        </div>

        <div class="p-4 border-t border-gray-200 dark:border-gray-700">
          <button
            @click="showPocModal = false"
            class="w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg"
          >
            确认选择
          </button>
        </div>
      </div>
    </div>

    <!-- 扫描结果 -->
    <div v-if="results.length > 0" class="bg-white dark:bg-gray-800 rounded-lg shadow">
      <div class="p-6">
        <h2 class="text-xl font-semibold mb-4">扫描结果</h2>
        <div class="space-y-6">
          <div
            v-for="result in results"
            :key="result.id"
            class="border rounded-lg overflow-hidden"
            :class="result.vulnerable ? 'border-red-200 bg-red-50' : 'border-green-200 bg-green-50'"
          >
            <div class="p-4">
              <div class="flex items-start justify-between">
                <div class="flex-1">
                  <div class="flex items-center space-x-2">
                    <h3 class="font-medium">{{ result.name }}</h3>
                    <span :class="['text-xs font-medium px-2 py-1 rounded-full', getSeverityBgColor(result.severity), getSeverityColor(result.severity)]">
                      {{ getSeverityText(result.severity) }}
                    </span>
                    <span class="text-xs text-gray-500">{{ result.category }}</span>
                  </div>
                  <p class="text-sm text-gray-600 mt-1">{{ result.id }}</p>
                </div>
                <span class="text-sm text-gray-500">{{ formatDate(result.timestamp) }}</span>
              </div>

              <div class="mt-4">
                <p :class="result.vulnerable ? 'text-red-600' : 'text-green-600'" class="font-medium">
                  {{ result.description }}
                </p>

                <template v-if="result.vulnerable">
                  <p class="mt-2 text-sm text-gray-600">
                    <strong>漏洞位置:</strong> {{ result.location }}
                  </p>
                  <p class="mt-2 text-sm text-gray-600">
                    <strong>漏洞详情:</strong> {{ result.details }}
                  </p>

                  <div class="mt-4">
                    <h4 class="font-medium text-gray-700 mb-2">修复建议:</h4>
                    <ul class="list-disc list-inside space-y-1">
                      <li
                        v-for="(recommendation, index) in result.recommendations"
                        :key="index"
                        class="text-sm text-gray-600"
                      >
                        {{ recommendation }}
                      </li>
                    </ul>
                  </div>
                </template>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div v-else-if="scanning" class="text-center p-6">
      <div class="animate-spin rounded-full h-12 w-12 border-4 border-blue-500 border-t-transparent mx-auto"></div>
      <p class="mt-4 text-gray-600">正在扫描目标漏洞...</p>
    </div>
  </div>
</template>

<style scoped>
  /* Add any additional component-specific styles here */
</style>