<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule"
               @keyup.enter="onSubmit">
        <el-form-item label="创建日期" prop="createdAt">
          <template #label>
            <span>
              创建日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false" />
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false" />
        </el-form-item>
        <template v-if="showAllQuery">
          <!-- 将需制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开
          </el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
<!--      <div class="gva-btn-list">-->
<!--        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>-->
<!--        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">-->
<!--          删除-->
<!--        </el-button>-->
<!--      </div>-->
<el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        :show-overflow-tooltip="true"
        row-key="id"
        @selection-change="handleSelectionChange"
      >
        <!-- <el-table-column type="selection" width="55"/> -->

            <!-- 新增展开列 -->
            <el-table-column type="expand">
          <template #default="{ row }">
            <div style="padding: 0 20px">

              <div>

                <h4>请求Request  <ai-chat /></h4>

              <el-input type="textarea" autosize :value="Object.keys(row.headers).map(k => {
                    return `${k}: ${row.headers[k]}`
                  }).join('\n')"></el-input>
                <h5>
                  <el-tag type="info">Payload请求载荷</el-tag>
                </h5>
                <el-input type="textarea" :autosize="{maxRows: 50}" :value="row.reqContent">

                </el-input>

              </div>
              <div>
                <h4>响应Response</h4>
                <el-input type="textarea" autosize :value="Object.keys(row.resHeader).map(k => {
                    return `${k}: ${row.resHeader[k]}`
                  }).join('\n')"></el-input>
                <h5>
                  <el-tag type="info">响应体</el-tag>
                </h5>
                <el-input type="textarea" :autosize="{maxRows: 50}" :value="row.resContent">

                </el-input>
              </div>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="时间" prop="createdAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.createdAt) }}</template>
        </el-table-column>
  <el-table-column align="left" label="协议" prop="proto"  width="100"/>
  <el-table-column align="left" label="HTTP方法" prop="method"  width="100"/>
  <el-table-column align="left" label="响应码" prop="statusCode"  width="100">

    <template #default="{ row }">
      <el-tag :type="row.statusCode == 200 ? 'success': 'danger'">{{  row.statusCode }}</el-tag>
    </template>
  </el-table-column>
        <!-- <el-table-column label="headers字段" prop="headers" >
          <template #default="{ row }">
            {{
              Object.keys(row.headers).map(k => {
                return `${k}: ${row.headers[k]}`
              }).join("\n")
            }}
          </template>
        </el-table-column> -->
        <el-table-column align="left" label="访问HOST" prop="host" />
        <el-table-column align="left" label="路径" prop="path" />
        <el-table-column align="left" label="来源" prop="remoteAddr" width="330"/>
  <!-- <el-table-column align="left" label="响应内容" prop="resContent" />
  <el-table-column label="resHeader字段" prop="resHeader" >
    <template #default="{ row }">
      {{
        Object.keys(row.resHeader).map(k => {
          return `${k}: ${row.resHeader[k]}`
        }).join("\n")
      }}
    </template>
  </el-table-column> -->

        
    

        <!-- <el-table-column align="left" label="操作" fixed="right" width="240">
          <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateHttpInfosFunc(scope.row)">
              变更
            </el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column> -->
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
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '添加' : '修改' }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="createdAt字段:" prop="createdAt">
          <el-date-picker v-model="formData.createdAt" type="date" style="width:100%" placeholder="选择日期"
                          :clearable="true"/>
        </el-form-item>
        <el-form-item label="headers字段:" prop="headers">
          // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.headers 后端会按照json的类型进行存取
          {{ formData.headers }}
        </el-form-item>
        <el-form-item label="域名:" prop="host">
          <el-input v-model="formData.host" :clearable="true" placeholder="请输入域名"/>
        </el-form-item>
        <el-form-item label="路径:" prop="path">
          <el-input v-model="formData.path" :clearable="true" placeholder="请输入路径"/>
        </el-form-item>
        <el-form-item label="原始IP:" prop="remoteAddr">
          <el-input v-model="formData.remoteAddr" :clearable="true" placeholder="请输入原始IP"/>
        </el-form-item>
        <el-form-item label="响应内容字符串:" prop="resContent">
          <el-input v-model="formData.resContent" :clearable="true" placeholder="请输入响应内容字符串"/>
        </el-form-item>
        <el-form-item label="resHeader字段:" prop="resHeader">
          // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.resHeader 后端会按照json的类型进行存取
          {{ formData.resHeader }}
        </el-form-item>
        <el-form-item label="响应码:" prop="statusCode">
          <el-input v-model.number="formData.statusCode" :clearable="true" placeholder="请输入响应码"/>
        </el-form-item>
      </el-form>
    </el-drawer>

    <!-- New section for monitoring traffic -->
<!--    <div class="gva-traffic-monitor">-->
<!--      <h2>流量监控</h2>-->
<!--      <div v-for="(request, index) in trafficData" :key="index" class="traffic-request-card" @click="showDetails(request)">-->
<!--        <h3>请��� {{ index + 1 }}</h3>-->
<!--        <p><strong>时间:</strong> {{ request.time }}</p>-->
<!--        <p><strong>请求方法:</strong> {{ request.method }}</p>-->
<!--        <p><strong>请求URL:</strong> {{ request.url }}</p>-->
<!--        <p><strong>流量:</strong> {{ request.traffic }}</p>-->
<!--        <p><strong>请求数:</strong> {{ request.requestCount }}</p>-->
<!--        <p><strong>响应时间:</strong> {{ request.responseTime }}</p>-->
<!--        <el-tag>{{ request.method }}</el-tag>-->
<!--      </div>-->
<!--    </div>-->

    <!-- 请求详细信息弹窗 -->
    <el-dialog v-model:visible="detailsVisible" title="请求详细信息" width="600px">
      <div v-if="selectedRequest" class="request-details-card">
        <h3>请求信息</h3>
        <p><strong>时间:</strong> {{ selectedRequest.time }}</p>
        <p><strong>请求方法:</strong> {{ selectedRequest.method }}</p>
        <p><strong>请求URL:</strong> {{ selectedRequest.url }}</p>
        <p><strong>流量:</strong> {{ selectedRequest.traffic }}</p>
        <p><strong>请求数:</strong> {{ selectedRequest.requestCount }}</p>
        <p><strong>响应时间:</strong> {{ selectedRequest.responseTime }}</p>

        <h4>请求头</h4>
        <ul>
          <li v-for="(value, key) in selectedRequest.headers" :key="key">
            <strong>{{ key }}:</strong> {{ value }}
          </li>
        </ul>

        <h4>响应头</h4>
        <ul>
          <li v-for="(value, key) in selectedRequest.resHeaders" :key="key">
            <strong>{{ key }}:</strong> {{ value }}
          </li>
        </ul>

        <h4>响应内容</h4>
        <pre>{{ selectedRequest.resContent }}</pre>
      </div>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createHttpInfos,
  deleteHttpInfos,
  deleteHttpInfosByIds,
  findHttpInfos,
  getHttpInfosList,
  updateHttpInfos
} from '@/api/httpInfos/httpInfos'

// 全量引入格式化工具 请按需保留
import {formatDate} from '@/utils/format'
import {ElMessage, ElMessageBox, ElSpace } from 'element-plus'
import {reactive, ref} from 'vue'

defineOptions({
  name: 'HttpInfos'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  createdAt: new Date(),
  headers: {},
  host: '',
  path: '',
  remoteAddr: '',
  resContent: '',
  resHeader: {},
  statusCode: undefined,
})


// 验证规则
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
const pageSize = ref(30)
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
  const table = await getHttpInfosList({page: page.value, pageSize: pageSize.value, ...searchInfo.value})
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
const setOptions = async () => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()


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
    deleteHttpInfosFunc(row)
  })
}

// 多选删除
const onDelete = async () => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const ids = []
    if (multipleSelection.value.length === 0) {
      ElMessage({
        type: 'warning',
        message: '请选择要删除的数据'
      })
      return
    }
    multipleSelection.value &&
    multipleSelection.value.map(item => {
      ids.push(item.id)
    })
    const res = await deleteHttpInfosByIds({ids})
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === ids.length && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateHttpInfosFunc = async (row) => {
  const res = await findHttpInfos({id: row.id})
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteHttpInfosFunc = async (row) => {
  const res = await deleteHttpInfos({id: row.id})
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

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    createdAt: new Date(),
    headers: {},
    host: '',
    path: '',
    remoteAddr: '',
    resContent: '',
    resHeader: {},
    statusCode: undefined,
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createHttpInfos(formData.value)
        break
      case 'update':
        res = await updateHttpInfos(formData.value)
        break
      default:
        res = await createHttpInfos(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      getTableData()
    }
  })
}

// New reactive variable for traffic data
const trafficData = ref([])

// 请求详细信息的状态
const detailsVisible = ref(false)
const selectedRequest = ref(null)

// 获取流量数据的函数
const fetchTrafficData = async () => {
  // const response = await getTrafficData()
  // if (response.code === 0) {
  //   trafficData.value = response.data.map(item => ({
  //     time: item.time,
  //     method: item.method,
  //     url: item.url,
  //     traffic: item.traffic,
  //     requestCount: item.requestCount,
  //     responseTime: item.responseTime,
  //     headers: item.headers, // 假设API返回此字段
  //     resHeaders: item.resHeaders, // 假设API返回此字段
  //     resContent: item.resContent // 假设API返回此字段
  //   }))
  // }
}

// 显示请求详细信息
const showDetails = (request) => {
  selectedRequest.value = request
  detailsVisible.value = true
}

// 定时器定期获取流量数据
setInterval(fetchTrafficData, 5000) // 每5秒获取一次数据

// 组件挂载时调用获取流���数据的函数
fetchTrafficData()

</script>

<style>
.gva-traffic-monitor {
  margin-top: 20px;
}

.traffic-request-card {
  border: 1px solid #e0e0e0;
  border-radius: 5px;
  padding: 15px;
  margin-bottom: 10px;
  background-color: #f9f9f9;
  cursor: pointer; /* 鼠标悬停时显示为可点击 */
}

.request-details-card {
  padding: 20px;
  background-color: #ffffff;
  border-radius: 5px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.request-details-card h3,
.request-details-card h4 {
  margin: 10px 0;
}

.request-details-card ul {
  list-style-type: none;
  padding: 0;
}

.request-details-card pre {
  background-color: #f5f5f5;
  padding: 10px;
  border-radius: 5px;
  overflow: auto;
}
</style>
