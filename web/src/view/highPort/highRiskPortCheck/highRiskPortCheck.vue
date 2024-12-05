<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list"></div>
      <!--      <el-tabs type="border-card" class="demo-tabs">-->
      <!--        <el-tab-pane>-->
      <!--          <template #label>-->
      <!--            <span class="custom-tabs-label">-->
      <!--              <el-badge :value="12" class="item">-->
      <!--                <span>Route</span>-->
      <!--              </el-badge>-->
      <!--            </span>-->
      <!--          </template>-->
      <!--          Route-->
      <!--        </el-tab-pane>-->
      <!--        <el-tab-pane label="Config">Config</el-tab-pane>-->
      <!--        <el-tab-pane label="Role">Role</el-tab-pane>-->
      <!--        <el-tab-pane label="Task">Task</el-tab-pane>-->
      <!--      </el-tabs>-->

      <div class="mb-4">
        <el-button
          text
          :bg="curr_port_number === 0"
          :type="curr_port_number === 0 && 'primary'"
          @click="curr_port_number = 0"
        >
          全部
        </el-button>
        <el-button
          v-for="port in portsList"
          :key="port.ID"
          :type="curr_port_number === port.portNumber && 'primary'"
          text
          :bg="curr_port_number === port.portNumber"
          @click="curr_port_number = port.portNumber"
        >
          {{ port.portNumber }}
        </el-button>
      </div>

      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
      >
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" label="日期" prop="createdAt">
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>
        <el-table-column label="最新时间" width="180">
          <template #default="scope">{{
            formatRelativeTime(scope.row.UpdatedAt)
          }}</template>
        </el-table-column>

        <el-table-column label="高危端口">
          <template #default="scope">
            <el-tag
              v-if="scope.row.PortConfig"
              :title="scope.row.PortConfig.portDescription"
              >{{ scope.row.PortConfig?.portName }}</el-tag
            >
          </template>
        </el-table-column>
        <el-table-column label="目标地址">
          <template #default="scope">
            <div class="flex">
              <span class="shrink-0 w-[3em]">IP:</span>{{ scope.row.Ip }}
            </div>
            <hr />
            <div class="flex">
              <span class="shrink-0 w-[3em]">MAC:</span>{{ scope.row.Mac }}
            </div>
          </template>
        </el-table-column>

        <el-table-column label="访问来源">
          <template #default="scope">
            <div class="flex">
              <span class="shrink-0 w-[3em]">IP:</span>{{ scope.row.FromIP }}
            </div>
            <hr />
            <div class="flex">
              <span class="shrink-0 w-[3em]">MAC:</span>{{ scope.row.FromMac }}
            </div>
          </template>
        </el-table-column>
        <!--        <el-table-column label="记录" prop="info"> </el-table-column>-->

        <!--        <el-table-column align="left" label="操作" fixed="right" width="240">-->
        <!--          <template #default="scope">-->
        <!--            <el-button-->
        <!--              type="primary"-->
        <!--              link-->
        <!--              class="table-button"-->
        <!--              @click="getDetails(scope.row)"-->
        <!--              ><el-icon style="margin-right: 5px"><InfoFilled /></el-icon-->
        <!--              >查看</el-button-->
        <!--            >-->
        <!--          </template>-->
        <!--        </el-table-column>-->
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <el-drawer
      destroy-on-close
      size="800"
      v-model="detailShow"
      :show-close="true"
      :before-close="closeDetailShow"
      title="查看"
    >
      <el-descriptions :column="1" border>
        <el-descriptions-item label="端口号">
          {{ detailFrom.portNumber }}
        </el-descriptions-item>
        <el-descriptions-item label="端口名称">
          {{ detailFrom.portName }}
        </el-descriptions-item>
        <el-descriptions-item label="风险等级">
          {{ detailFrom.riskLevel }}
        </el-descriptions-item>
        <el-descriptions-item label="描述">
          {{ detailFrom.portDescription }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
  import {
    findHighRiskPortConfig,
    getHighRiskPortConfigList,
    getHighRiskPortLogsList
  } from '@/api/highPort/highRiskPortConfig'

  // 全量引入格式化工具 请按需保留
  import { formatDate, WsPath } from '@/utils/format'
  import { ref, watch, onMounted, onBeforeUnmount } from 'vue'

  defineOptions({
    name: 'HighRiskPortConfig'
  })

  const ws = ref(null)
  const logs = ref([]) // 用于存储高危端口日志

  const connectWebSocket = () => {
    ws.value = new WebSocket(WsPath + 'system_ws/?messageId=livePortLog') // 替换为你的WebSocket地址
    ws.value.onopen = () => {
      console.log('WebSocket connected')
    }
    ws.value.onmessage = (event) => {
      const data = event.data
      if (data.includes('高危PORT')) {
        getTableData()
      }
      console.log(event.data)

      const newLogs = JSON.parse(event.data)
      logs.value = newLogs // 更新日志列表
    }

    ws.value.onclose = () => {
      console.log('WebSocket closed, attempting to reconnect...')
      setTimeout(connectWebSocket, 1000) // 重新连接
    }
  }

  onMounted(() => {
    connectWebSocket()
  })

  onBeforeUnmount(() => {
    if (ws.value) {
      ws.value.close()
    }
  })

  const portsList = ref([])
  const curr_port_number = ref(0)
  function findPorts() {
    getHighRiskPortConfigList({
      page: 1,
      pageSize: 1000,
      status: 1
    }).then((res) => {
      portsList.value = res.data.list
    })
  }
  findPorts()
  watch(curr_port_number, () => {
    searchInfo.value['portNumber'] = curr_port_number.value
    getTableData()
  })
  // =========== 表格控制部分 ===========
  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const tableData = ref([])
  const searchInfo = ref({})
  // 排序
  const sortChange = ({ prop, order }) => {
    const sortMap = {
      portNumber: 'port_number',
      riskLevel: 'risk_level'
    }

    let sort = sortMap[prop]
    if (!sort) {
      sort = prop.replace(/[A-Z]/g, (match) => `_${match.toLowerCase()}`)
    }

    searchInfo.value.sort = sort
    searchInfo.value.order = order
    getTableData()
  }

  // 分页
  const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
  }

  // 修改页面容量
  const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
  }

  // 查询
  const getTableData = async () => {
    const table = await getHighRiskPortLogsList({
      page: page.value,
      pageSize: pageSize.value,
      ...searchInfo.value
    })
    if (table.code === 0) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
  }

  getTableData()

  // ============== 表格控制部分结束 ===============

  // 获取需要的字典 可能为空 按需保留
  const setOptions = async () => {}

  // 获取需要的字典 可能为空 按需保留
  setOptions()

  // 多选数据
  const multipleSelection = ref([])
  // 多选
  const handleSelectionChange = (val) => {
    multipleSelection.value = val
  }

  const detailFrom = ref({})

  // 查看详情控制标记
  const detailShow = ref(false)

  // 打开详情弹窗
  const openDetailShow = () => {
    detailShow.value = true
  }

  // 打开详情
  const getDetails = async (row) => {
    // 打开弹窗
    const res = await findHighRiskPortConfig({ ID: row.ID })
    if (res.code === 0) {
      detailFrom.value = res.data
      openDetailShow()
    }
  }

  // 关闭详情弹窗
  const closeDetailShow = () => {
    detailShow.value = false
    detailFrom.value = {}
  }

  const formatRelativeTime = (dateString) => {
    const now = new Date()
    const date = new Date(dateString)
    const seconds = Math.floor((now - date) / 1000)

    if (seconds === 0) return '当前在线' // Change for 0 seconds

    let interval = Math.floor(seconds / 31536000)
    if (interval > 1) return `${interval}年之前`
    interval = Math.floor(seconds / 2592000)
    if (interval > 1) return `${interval}个月之前`
    interval = Math.floor(seconds / 86400)
    if (interval > 1) return `${interval}天之前`
    interval = Math.floor(seconds / 3600)
    if (interval > 1) return `${interval}小时之前`
    interval = Math.floor(seconds / 60)
    if (interval > 1) return `${interval}分钟之前`
    return `${seconds}秒之前`
  }
</script>

<style></style>
