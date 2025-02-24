<script setup lang="ts">
  import { ref } from 'vue'
  import { format } from 'date-fns'

  interface Task {
    id: number
    title: string
    description: string
    status: 'pending' | 'in_progress' | 'completed'
    assignee: string
    created_at: Date
    started_at?: Date
    completed_at?: Date
    priority: 'high' | 'medium' | 'low'
    ai_summary?: string
  }

  const newTaskDescription = ref('')

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

  const getPriorityColor = (priority: string) => {
    switch (priority) {
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

  const startTask = (taskId: number) => {
    const task = tasks.value.find(t => t.id === taskId)
    if (task && task.status === 'pending') {
      task.status = 'in_progress'
      task.started_at = new Date()
    }
  }

  const completeTask = (taskId: number) => {
    const task = tasks.value.find(t => t.id === taskId)
    if (task && task.status === 'in_progress') {
      task.status = 'completed'
      task.completed_at = new Date()
    }
  }

  const addNewTask = () => {
    if (!newTaskDescription.value.trim()) return

    // 模拟AI提炼过程
    const aiSummary = `任务分析：${newTaskDescription.value}。建议分配给安全团队处理，预计耗时2小时。`

    const newTask: Task = {
      id: tasks.value.length + 1,
      title: newTaskDescription.value.split('.')[0], // 使用第一句作为标题
      description: newTaskDescription.value,
      status: 'pending',
      assignee: '待分配',
      created_at: new Date(),
      priority: 'medium',
      ai_summary: aiSummary
    }

    tasks.value.unshift(newTask)
    newTaskDescription.value = ''
  }
</script>

<template>
  <div class="space-y-6">
    <!-- 新建任务 -->
    <div class="card">
      <h2 class="text-xl font-semibold mb-4">创建新任务</h2>
      <div class="space-y-4">
        <div>
          <label for="taskDescription" class="block text-sm font-medium text-gray-700 mb-2">
            任务描述
          </label>
          <textarea
            id="taskDescription"
            v-model="newTaskDescription"
            rows="3"
            class="w-full px-3 py-2 border border-solid  border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500"
            placeholder="请输入任务描述，AI将自动分析并提炼关键信息..."
          ></textarea>
        </div>
        <button
          @click="addNewTask"
          class="px-4 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
        >
          创建任务
        </button>
      </div>
    </div>

    <!-- 任务列表 -->
    <div class="card">
      <h2 class="text-xl font-semibold mb-4">任务列表</h2>
      <div class="space-y-4">
        <div v-for="task in tasks" :key="task.id" class="border rounded-lg p-4">
          <div class="flex items-start justify-between">
            <div class="space-y-2">
              <div class="flex items-center gap-2">
                <h3 class="text-lg font-medium">{{ task.title }}</h3>
                <span :class="'px-2 py-1 rounded-full text-xs ' + getStatusColor(task.status)">
                  {{ getStatusText(task.status) }}
                </span>
                <span :class="'px-2 py-1 rounded-full text-xs ' + getPriorityColor(task.priority)">
                  优先级: {{ task.priority === 'high' ? '高' : task.priority === 'medium' ? '中' : '低' }}
                </span>
              </div>
              <p class="text-gray-600">{{ task.description }}</p>
              <div class="bg-gray-50 p-3 rounded-md">
                <p class="text-sm text-gray-700">
                  <span class="font-medium">AI 分析：</span>
                  {{ task.ai_summary }}
                </p>
              </div>
              <div class="text-sm text-gray-500 space-y-1">
                <p>负责人: {{ task.assignee }}</p>
                <p>创建时间: {{ format(task.created_at, 'yyyy-MM-dd HH:mm') }}</p>
                <p v-if="task.started_at">
                  开始时间: {{ format(task.started_at, 'yyyy-MM-dd HH:mm') }}
                </p>
                <p v-if="task.completed_at">
                  完成时间: {{ format(task.completed_at, 'yyyy-MM-dd HH:mm') }}
                </p>
              </div>
            </div>
            <div class="space-x-2">
              <button
                v-if="task.status === 'pending'"
                @click="startTask(task.id)"
                class="px-3 py-1 bg-blue-600 text-white rounded hover:bg-blue-700"
              >
                开始任务
              </button>
              <button
                v-if="task.status === 'in_progress'"
                @click="completeTask(task.id)"
                class="px-3 py-1 bg-green-600 text-white rounded hover:bg-green-700"
              >
                完成任务
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>