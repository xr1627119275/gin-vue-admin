<script setup lang="ts">
  import { ref, computed, unref, reactive, nextTick } from 'vue'

  const route = useRoute()
  console.log(route)
  const target = ref(route.query.target )
  const scanConfig = reactive({
    checkbox: [ ]
  })

  // import temp from './temp'
  const scanning = ref(false)
  const results = ref([])
  const showPocModal = ref(false)
  const searchQuery = ref('')
  const selectedPocs = ref<string[]>([])
  import nucleiInfo from '../nucleiInfo/nucleiInfo.vue'
  import {ElMessageBox } from  'element-plus'
  import {
    createNucleiScan,
    getNucleiPocData
  } from '@/api/nucleiInfo/nucleiInfo'
  import useWebsocket from '@/hooks/use-websocket'
  import { useRoute } from 'vue-router'
  import PocViewer from '@/view/nucleiInfo/pocViewer.vue'




  const checkDatas = ref([])

  const startScan = async () => {

    scanning.value = true
    results.value = []

    const queryData = {
      targets: unref(target).split('\n').map(item => {
        return item?.replaceAll('http://', '')?.replaceAll('https://', '')
      }),
      poc_filter: {
        IDs: unref(checkDatas).map(item => item.id),
        ExcludeSeverities: "info"
      }
    }

    if (scanConfig.checkbox.includes('excludeInfo')) {
      delete queryData.poc_filter.ExcludeSeverities
    }


    let res = await createNucleiScan(queryData)
    if (res.code === 0) {
      const id = res.data.id
      const { connectWebSocket } = useWebsocket({
        async handleClose() {
          scanning.value = false
          console.log("results:", results)
        }
      })

      connectWebSocket(id, (data) => {
        data = JSON.parse(JSON.parse(data))
        results.value.push(data)
      })
    }
    // 模拟扫描结果
    // setTimeout(() => {
    //   results.value = selectedPocs.value.map(pocId => {
    //     const poc = pocs.value.find(p => p.id === pocId)
    //     const isVulnerable = Math.random() > 0.5 // 随机模拟是否存在漏洞
    //     return {
    //       id: poc?.id,
    //       name: poc?.name,
    //       severity: poc?.severity,
    //       category: poc?.category,
    //       description: isVulnerable ? `发现${poc?.name}漏洞` : `未发现${poc?.name}漏洞`,
    //       vulnerable: isVulnerable,
    //       location: isVulnerable ? `${target.value}/api/vulnerable-endpoint` : '',
    //       details: poc?.description,
    //       recommendations: isVulnerable ? [
    //         '及时更新相关组件到最新版本',
    //         '检查并修复相关代码逻辑',
    //         '增加相应的安全防护措施'
    //       ] : [],
    //       timestamp: new Date().toISOString()
    //     }
    //   })
    //   scanning.value = false
    // }, 2000)
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

  const dialogData = ref("")
  const showDataDialog = ref(false)
  function showPoc(data) {
    dialogData.value = data
    showDataDialog.value = true
  }
  async function showPocData(id) {
    const res = await getNucleiPocData({id})
    if (res.code === 0) {
      console.log(res.data.data)
      showPoc(res.data.data)
    }
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
          <div class="flex gap-2 pb-2 flex-wrap" >
            <el-tag v-for="(data, index) in checkDatas" :key="data.id" :type="'success'"  closable @close="checkDatas.splice(index, 1)">{{ data.id }}</el-tag>
          </div>
          <button
            type="button"
            @click="showPocModal = true"
            class="w-full bg-gray-100 hover:bg-gray-200 text-gray-800 font-medium py-2 px-4 rounded-lg border border-gray-300"
          >
            选择 漏洞库 ({{ checkDatas.length ? `已选择${checkDatas.length}个` : '默认全选'}})
          </button>
        </div>
        <div v-if="!checkDatas.length">
          <label for="target" class="block text-sm font-medium mb-2">选项</label>
          <el-checkbox-group v-model="scanConfig.checkbox">
            <el-checkbox value="excludeInfo" label="指纹信息识别"></el-checkbox>
          </el-checkbox-group>
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
    <el-dialog v-model="showPocModal" title="选择 POC" width="90vw">
      <div class="bg-white dark:bg-gray-800    w-full  flex flex-col">
        <div class="overflow-y-auto flex-1 ">
          <nuclei-info type="select" v-model:multiple-selection="checkDatas"/>
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
    </el-dialog>

    <!-- 扫描结果 -->
    <div v-if="results.length > 0" class="bg-white dark:bg-gray-800 rounded-lg shadow">
      <div class="p-6">
        <h2 class="text-xl font-semibold mb-4">扫描结果</h2>
        <div class="space-y-6">
          <div
            v-for="result in results"
            :key="result.info.id"
            class="border rounded-lg overflow-hidden"
            :class="result.vulnerable ? 'border-red-200 bg-red-50' : 'border-green-200 bg-green-50'"
          >
            <div class="p-4">
              <div class="flex items-start justify-between">
                <div class="flex-1">
                  <div class="flex items-center space-x-2">
                    <h3 class="font-medium">{{ result.info.name }}</h3>
                    <span  :class="['text-xs font-medium px-2 py-1 rounded-full flex-shrink-0', getSeverityBgColor(result.info.severity.toUpperCase()), getSeverityColor(result.info.severity.toUpperCase())]">
                      {{ getSeverityText(result.info.severity.toUpperCase()) }}
                    </span>
                    <span class="text-xs text-gray-500">{{ result.url }}</span>
                  </div>
                  <p class="text-sm text-gray-600 mt-1"><strong>ID:</strong> {{ result['template-id'] }}</p>
                </div>
                <span class="text-sm text-gray-500">{{ formatDate(result.timestamp) }}</span>
              </div>
              <div class="mt-4">
<!--                result.vulnerable ? 'text-red-600' : 'text-green-600'-->
<!--                <p class="font-medium">-->
<!--                  {{ result.info.description }}-->
<!--                </p>-->

                <template v-if="result">
                  <p class=" text-sm text-gray-600">
                    <strong>漏洞位置:</strong> {{ result['matched-at'] }}
                  </p>
                  <p v-if="result['extracted-results']?.length" class="mt-2 text-sm text-gray-600">
                    <strong>漏洞详情:</strong> <pre>{{ result['extracted-results'] }}</pre>
                  </p>
                </template>
              </div>
              <div class="mt-4 flex justify-between gap-2">
                <el-button type="info" size="small" @click="showPocData(result['template-id'])">查看poc</el-button>
<!--                <el-button v-if="result['curl-command']" type="info" size="small"  @click="showDialogData(result['curl-command'])">查看CURL</el-button>-->
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
    <poc-viewer v-model:code="dialogData" v-model:visible="showDataDialog"></poc-viewer>
  </div>
</template>

