<template>
  <div>
    <!--    <div class="gva-search-box"></div>-->
    <div class="gva-table-box">
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        :show-overflow-tooltip="true"
        row-key="ID"
      >
        <!-- <el-table-column type="selection" width="55"/> -->

        <!-- 新增展开列 -->
        <el-table-column type="expand">
          <template #default="{ row: { http_infos: row } }">
            <el-row :gutter="8" class="w-full p-2">
              <el-col :span="12" class="flex flex-col gap-[10px]">
                <h4>请求Request
                <ai-gva/></h4>
                <el-input
                  type="textarea"
                  autosize
                  :value="
                    Object.keys(row.headers)
                      .map((k) => {
                        return `${k}: ${row.headers[k]}`
                      })
                      .join('\n')
                  "
                ></el-input>
                <h5>
                  <el-tag type="info">请求体</el-tag>
                </h5>
                <el-input
                  type="textarea"
                  :autosize="{ maxRows: 50 }"
                  :value="row.reqContent"
                >
                </el-input>
              </el-col>
              <el-col :span="12" class="flex flex-col gap-[10px]">
                <h4>响应Response</h4>
                <el-input
                  type="textarea"
                  autosize
                  :value="
                    Object.keys(row.resHeader)
                      .map((k) => {
                        return `${k}: ${row.resHeader[k]}`
                      })
                      .join('\n')
                  "
                ></el-input>
                <template v-if="row.resContent">
                  <h5>
                    <el-tag type="info">响应体</el-tag>
                  </h5>
                  <el-input
                    type="textarea"
                    :autosize="{ maxRows: 10 }"
                    :value="row.resContent"
                  >
                  </el-input>
                </template>
              </el-col>
            </el-row>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="时间"
          prop="http_infos.createdAt"
          width="180"
        >
          <template #default="scope">{{
            formatDate(scope.row.http_infos.createdAt)
          }}</template>
        </el-table-column>
        <el-table-column align="left" label="网站" prop="http_infos.host">
          <template #default="{ row: { http_infos } }">
            <el-button type="text">
              http://{{ http_infos.host }}{{ http_infos.path }}
            </el-button>
          </template>
        </el-table-column>

        <el-table-column align="left" label="响应码" width="100">
          <template #default="{ row }">
            <el-tag>{{ row.http_infos.statusCode }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="弱密码" prop="password" />
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
  </div>
</template>

<script setup>
  import { getHttpWeakPassWordInfosList } from '@/api/httpInfos/httpInfos'

  // 全量引入格式化工具 请按需保留
  import { formatDate } from '@/utils/format'
  import { onMounted, ref } from 'vue'
  import useWebsocket from '@/hooks/use-websocket'

  // 控制更多查询条件显示/隐藏状态
  // =========== 表格控制部分 ===========
  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(30)
  const tableData = ref([])
  const searchInfo = ref({})

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
    const table = await getHttpWeakPassWordInfosList({
      page: page.value,
      pageSize: pageSize.value,
      weakPassword: true,
      ...searchInfo.value
    })
    if (table.code === 0) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
  }
  onMounted(() => {
    const { connectWebSocket } = useWebsocket({})
    connectWebSocket('liveWeakPasswordLog', (data) => {
      data = JSON.parse(JSON.parse(data))
      tableData.value.unshift(data)
    })
  })
  getTableData()

  // 获取流量数据的函数
  const fetchTrafficData = async () => {}

  // 定时器定期获取流量数据
  setInterval(fetchTrafficData, 5000) // 每5秒获取一次数据

  // 组件挂载时调用获取流���数据的函数
  fetchTrafficData()
</script>

<style>
  .el-table .warning-row {
    background: oldlace;
  }
</style>
