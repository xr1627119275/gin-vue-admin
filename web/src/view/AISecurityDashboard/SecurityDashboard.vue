<script setup lang="ts">
  import { ref } from 'vue'
  import { format } from 'date-fns'
  import RecentTasks from './RecentTasks.vue'
  import type { Alert, Annotation } from '../types'

  const currentDate = format(new Date(), 'yyyy-MM-dd HH:mm:ss')

  const securityStatus = ref({
    threatLevel: '中等',
    summary: '今日网络安全状况显示活动程度适中。检测并阻止了多次未经授权的访问尝试。整体系统稳定性良好。',
    recommendations: [
      '加强外部访问点的监控',
      '根据已识别的威胁更新防火墙规则',
      '审查用户访问模式'
    ]
  })

  const alerts = ref<Alert[]>([
    {
      id: 1,
      severity: 'high',
      device: '防火墙-01',
      description: '检测到多次登录失败尝试',
      timestamp: new Date(),
      annotations: [
        {
          id: 1,
          user: '张三',
          content: '这个告警需要立即处理，可能是暴力破解攻击',
          timestamp: new Date(Date.now() - 3600000)
        }
      ]
    },
    {
      id: 2,
      severity: 'medium',
      device: '入侵检测系统',
      description: '检测到异常网络流量模式',
      timestamp: new Date(),
      annotations: [
        {
          id: 2,
          user: '李四',
          content: '已确认是测试团队的压力测试导致，可以忽略',
          timestamp: new Date(Date.now() - 7200000)
        }
      ]
    },
    {
      id: 3,
      severity: 'low',
      device: '数据库服务器-01',
      description: '资源使用率突增',
      timestamp: new Date(),
      annotations: [
        {
          id: 3,
          user: '张三',
          content: '这个扫描结果不是很准确，需要手动验证',
          timestamp: new Date(Date.now() - 1800000)
        }
      ]
    }
  ])

  const getSeverityColor = (severity: string) => {
    switch (severity) {
      case 'high':
        return 'bg-red-100 text-red-800'
      case 'medium':
        return 'bg-yellow-100 text-yellow-800'
      case 'low':
        return 'bg-green-100 text-green-800'
      default:
        return 'bg-gray-100 text-gray-800'
    }
  }

  const getSeverityText = (severity: string) => {
    switch (severity) {
      case 'high':
        return '高'
      case 'medium':
        return '中'
      case 'low':
        return '低'
      default:
        return '未知'
    }
  }
</script>

<template>
  <div class="min-h-screen p-6">
    <header class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900 flex justify-between">AI 网络安全监控仪表盘

        <el-button type="primary" class="fr">创建任务</el-button>
      </h1>
      <p class="text-gray-600">最后更新时间: {{ currentDate }}</p>
    </header>

    <div class="grid grid-cols-1 gap-6">
      <!-- Security Status -->
      <div class="card">
        <h2 class="text-xl font-semibold mb-4">今日安全状况</h2>
        <div class="space-y-4">
          <div class="flex items-center">
            <span class="font-medium mr-2">威胁等级:</span>
            <span :class="'px-2 py-1 rounded-full text-sm ' +
              (securityStatus.threatLevel === '高' ? 'bg-red-100 text-red-800' :
               securityStatus.threatLevel === '中等' ? 'bg-yellow-100 text-yellow-800' :
               'bg-green-100 text-green-800')">
              {{ securityStatus.threatLevel }}
            </span>
          </div>
          <p class="text-gray-700">{{ securityStatus.summary }}</p>
          <div>
            <h3 class="font-medium mb-2">AI 建议:</h3>
            <ul class="list-disc list-inside space-y-1 text-gray-700">
              <li v-for="(rec, index) in securityStatus.recommendations" :key="index">
                {{ rec }}
              </li>
            </ul>
          </div>
        </div>
      </div>

      <!-- Recent Tasks -->
      <RecentTasks />

      <!-- Device Alerts -->
      <div class="card">
        <h2 class="text-xl font-semibold mb-4">设备告警</h2>
        <div class="overflow-x-auto">
          <table class="min-w-full">
            <thead>
            <tr class="bg-gray-50">
              <th class="px-4 py-2 text-left text-sm font-medium text-gray-500">严重程度</th>
              <th class="px-4 py-2 text-left text-sm font-medium text-gray-500">设备</th>
              <th class="px-4 py-2 text-left text-sm font-medium text-gray-500">描述</th>
              <th class="px-4 py-2 text-left text-sm font-medium text-gray-500">时间</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="alert in alerts" :key="alert.id" class="border-t border border-solid  border-gray-300 group">
              <td class="px-4 py-2">
                  <span :class="'px-2 py-1 rounded-full text-sm ' + getSeverityColor(alert.severity)">
                    {{ getSeverityText(alert.severity) }}
                  </span>
              </td>
              <td class="px-4 py-2 text-gray-700">{{ alert.device }}</td>
              <td class="px-4 py-2 text-gray-700">
                <div>
                  {{ alert.description }}
                  <!-- Annotations -->
                  <div v-if="alert.annotations && alert.annotations.length > 0"
                       class="mt-2 space-y-2">
                    <div v-for="annotation in alert.annotations"
                         :key="annotation.id"
                         class="text-sm bg-gray-50 p-2 rounded-lg border border-solid  border-gray-200">
                      <div class="flex items-center gap-2 text-gray-600 mb-1">
                        <span class="font-medium">{{ annotation.user }}</span>
                        <span class="text-gray-400">{{ format(annotation.timestamp, 'MM-dd HH:mm') }}</span>
                      </div>
                      <p class="text-gray-700">{{ annotation.content }}</p>
                    </div>
                  </div>
                </div>
              </td>
              <td class="px-4 py-2 text-gray-600">
                {{ format(alert.timestamp, 'HH:mm:ss') }}
              </td>
            </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>
