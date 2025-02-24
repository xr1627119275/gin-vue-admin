<script setup lang="ts">
  import { ref, computed } from 'vue'
  import { format } from 'date-fns'
  import type { Task } from '../types'

  const tasks = ref<Task[]>([
    {
      id: 1,
      title: '防火墙安全加固',
      description: '需要对防火墙配置进行全面检查，特别关注最近出现的异常登录尝试',
      status: 'in_progress',
      assignee: '张三',
      created_at: new Date(Date.now() - 86400000),
      started_at: new Date(Date.now() - 43200000),
      priority: 'high',
      ai_summary: '防火墙安全加固任务：重点关注异常登录，建议检查ACL规则，更新安全策略，预计需要4小时完成'
    },
    {
      id: 2,
      title: '漏洞扫描报告分析',
      description: '对最新的漏洞扫描报告进行分析，标注出误报项，确定修复优先级',
      status: 'pending',
      assignee: '李四',
      created_at: new Date(),
      priority: 'medium',
      ai_summary: '漏洞扫描分析任务：建议先处理高危漏洞，对误报进行标注，预计需要2小时完成分析'
    }
  ])

  const recentTasks = computed(() => {
    return tasks.value
      .sort((a, b) => b.created_at.getTime() - a.created_at.getTime())
      .slice(0, 3)
  })

  const getStatusText = (status: string) => {
    switch (status) {
      case 'pending':
        return '待处理'
      case 'in_progress':
        return '进行中'
      case 'completed':
        return '已完成'
      default:
        return '未知'
    }
  }

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'pending':
        return 'bg-gray-100 text-gray-800'
      case 'in_progress':
        return 'bg-blue-100 text-blue-800'
      case 'completed':
        return 'bg-green-100 text-green-800'
      default:
        return 'bg-gray-100 text-gray-800'
    }
  }
</script>

<template>
  <div class="card">
    <div class="flex items-center justify-between mb-4">
      <h2 class="text-xl font-semibold">最新工作任务</h2>
      <router-link
        :to="{name :'AiTaskList'}"
        class="text-indigo-600 hover:text-indigo-800 text-sm font-medium"
      >
        查看全部
      </router-link>
    </div>
    <div class="space-y-3">
      <div
        v-for="task in recentTasks"
        :key="task.id"
        class="flex items-center justify-between p-3 bg-gray-50 rounded-lg"
      >
        <div class="flex items-center space-x-3">
          <span :class="'px-2 py-1 rounded-full text-xs ' + getStatusColor(task.status)">
            {{ getStatusText(task.status) }}
          </span>
          <span class="font-medium text-gray-900">{{ task.title }}</span>
        </div>
        <span class="text-sm text-gray-500">{{ format(task.created_at, 'MM-dd HH:mm') }}</span>
      </div>
    </div>
  </div>
</template>