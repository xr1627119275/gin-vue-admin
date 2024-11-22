<template>
  <el-tabs type="border-card" class="m-[20px]">
    <el-tab-pane>
      <template #label>
        <span class="custom-tabs-label items-center flex gap-1">
          <el-icon><calendar/></el-icon>
          <div>资产扫描</div>
        </span>
      </template>
      <div class="penetration-test-container">
        <el-card class="box-card">
          <template #header>
            <div class="card-header">
              <h2>扫描</h2>
            </div>
          </template>
          <el-form :model="formData" label-position="top">
            <el-row :gutter="20">
              <el-col :span="16">
                <el-form-item label="IP"
                              title="目标ip: 192.168.11.11 | 192.168.11.11-255 | 192.168.11.11,192.168.11.12">
                  <el-input v-model="formData.host"
                            placeholder="目标ip: 192.168.11.11 | 192.168.11.11-255 | 192.168.11.11,192.168.11.12"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="8">
                <el-form-item label="端口"
                              title='设置扫描的端口: 22 | 1-65535 | 22,80,3306 (default "21,22,80,81,135,139,443,445,1433,3306,5432,6379,7001,8000,8080,8089,9000,9200,11211,27017")'>
                  <el-input v-model="formData.port"
                            placeholder='设置扫描的端口: 22 | 1-65535 | 22,80,3306 (default "21,22,80,81,135,139,443,445,1433,3306,5432,6379,7001,8000,8080,8089,9000,9200,11211,27017")'></el-input>
                </el-form-item>
              </el-col>
            </el-row>
            <el-form-item>
              <el-button type="primary" @click="executeTest" :loading="isLoading" class="execute-btn">
                执行
              </el-button>
            </el-form-item>
          </el-form>
        </el-card>

        <el-card v-if="logs" class="box-card log-card">
          <template #header>
            <div class="card-header">
              <h3>日志</h3>
            </div>
          </template>
          <div class="log-content">
            <el-input type="textarea" :value="logs" :autosize="{ minRows: 10, maxRows: 30}">
            </el-input>
          </div>
        </el-card>
      </div>
    </el-tab-pane>
    <el-tab-pane label="扫描历史">
      <div class="gva-search-box">
        <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule"
                 @keyup.enter="onSubmit">
          <el-form-item label="创建日期" prop="createdAt">
            <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled/></el-icon>
          </el-tooltip>
        </span>
            </template>
            <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期"
                            :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期"
                            :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
          </el-form-item>


          <template v-if="showAllQuery">
            <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
            <el-form-item label="HOST" prop="host">
              <el-input v-model="searchInfo.host" placeholder="搜索条件"/>

            </el-form-item>

          </template>

          <el-form-item>
            <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
            <el-button icon="refresh" @click="onReset">重置</el-button>
            <!--            <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开-->
            <!--            </el-button>-->
            <!--            <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>-->
          </el-form-item>
        </el-form>
      </div>
      <div class="gva-table-box">
        <el-table
          ref="multipleTable"
          style="width: 100%"
          tooltip-effect="dark"
          :data="tableData"
          row-key="ID"
          @selection-change="handleSelectionChange"
        >
          <el-table-column type="selection" width="55"/>

          <el-table-column align="left" label="日期" prop="createdAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
          </el-table-column>

          <el-table-column align="left" label="IP" prop="host" />
          <el-table-column align="left" label="Port" prop="port" />
          <el-table-column align="left" label="链接" prop="args.u" />

          <el-table-column  label="操作" fixed="right" >
            <template #default="scope">
              <el-button type="primary" link @click="detailInfo(scope.row)">查看日志</el-button>
              <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
          </el-table-column>
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
    </el-tab-pane>
  </el-tabs>
  <el-dialog v-model="logsInfoVisible" title="日志信息" width="800">
    <el-input type="textarea" readonly :value="currInfo" :autosize="{ minRows: 10, maxRows: 25}"></el-input>
  </el-dialog>
</template>

<script setup>
import {reactive, ref} from 'vue'
import {createWebScan, deleteWebScan, deleteWebScanByIds, findWebScan, getWebScanList} from "@/api/webScan/webScan";
import {formatDate} from "@/utils/format";
import { ElInput, ElMessage, ElMessageBox} from "element-plus";

defineOptions({
  name: 'VulScan'
})
const logsInfoVisible = ref(false)
const currInfo = ref('')

const formData = reactive({
  host: '',
  port: '',
  args: {
    extra: []
  }
})

const logs = ref("")
const timeId = ref("")
const isLoading = ref(false)

const executeTest = async () => {
  clearInterval(timeId.value)
  // if (!formData.host) {
  //   ElMessage.error('请输入主机和端口')
  //   return
  // }
  const params = JSON.parse(JSON.stringify(formData))

  const {extra} = params.args
  if (extra && extra?.length) {
    extra.forEach((item) => {
      params.args[item] = ''
    })
  }
  delete params.args.extra
  let res = await createWebScan(params)
  console.log(res)
  const jobId = res.data
  if (!jobId) return

  isLoading.value = true

  logs.value = ""
  timeId.value = setInterval(async () => {
    try {
      res = await findWebScan({ID: String(jobId)})
      const result = res?.data?.result
      logs.value = result
      if (result.includes('扫描结束') || res?.data?.scan_over) {
        return clearInterval(timeId.value)
      }
      // 模拟渗透测试过程
    } catch (error) {
      clearInterval(timeId.value)
    } finally {
      isLoading.value = false
    }
  }, 800)

}


const rule = reactive({})

const searchRule = reactive({
  createdAt: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change'
    }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
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
  const table = await getWebScanList({page: page.value, pageSize: pageSize.value, ...searchInfo.value})
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// 更新行
const updateWebScanFunc = async (row) => {
  const res = await findWebScan({ID: row.ID})
  // type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    // dialogFormVisible.value = true
  }
}

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteWebScanFunc(row)
  })
}

// 删除行
const detailInfo = (row) => {
  currInfo.value = row.result
  logsInfoVisible.value = true
}

// 多选删除
const onDelete = async () => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const IDs = []
    if (multipleSelection.value.length === 0) {
      ElMessage({
        type: 'warning',
        message: '请选择要删除的数据'
      })
      return
    }
    multipleSelection.value &&
    multipleSelection.value.map(item => {
      IDs.push(item.ID)
    })
    const res = await deleteWebScanByIds({IDs})
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === IDs.length && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}


// 删除行
const deleteWebScanFunc = async (row) => {
  const res = await deleteWebScan({ID: row.ID})
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}
const delay = (ms) => new Promise(resolve => setTimeout(resolve, ms))
</script>

<style scoped>
.penetration-test-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

.box-card {
  margin-bottom: 20px;
  border-radius: 8px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-header h2, .card-header h3 {
  margin: 0;
  font-weight: 600;
}

.execute-btn {
  width: 100%;
  margin-top: 10px;
}

.log-card {
  max-height: 400px;
  overflow-y: auto;
}

.log-content {
  font-family: 'Courier New', Courier, monospace;
  white-space: pre-wrap;
}

.log-content p {
  margin: 5px 0;
  padding: 5px;
  border-radius: 4px;
}

.log-info {
  background-color: #f4f4f5;
  color: #909399;
}

.log-success {
  background-color: #f0f9eb;
  color: #67c23a;
}

.log-warning {
  background-color: #fdf6ec;
  color: #e6a23c;
}

.log-error {
  background-color: #fef0f0;
  color: #f56c6c;
}
</style>