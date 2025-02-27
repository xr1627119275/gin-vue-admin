
<template>
  <div>
    <div class="gva-search-box" >
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="POC搜索" prop="createdAt">
        <el-input v-model="searchInfo.input" placeholder="ID/名称/描述"></el-input>
      </el-form-item>
      <el-form-item label="类型" >
        <el-select
          v-model="searchInfo.poc_filter.Severity"
          placeholder="类型"
          style="width: 100px"
          :clearable="true"
        >
          <el-option
            v-for="(item, key) in Object.keys(severityType)"
            :key="key"
            :label="severityType[item]"
            :value="item"
          />
        </el-select>
      </el-form-item>
        <el-form-item label="标签" >
          <el-input-tag v-model="searchInfo.poc_filter.Tags" placeholder="标签搜索"></el-input-tag>
        </el-form-item>
        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
<!--          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>-->
<!--          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>-->
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
<!--        <div class="gva-btn-list">-->
<!--            <el-button  type="primary" icon="plus" @click="openDialog">新增</el-button>-->
<!--            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>-->
<!--        </div>-->
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        v-loading="tableLoading"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55"  :reserve-selection="true"/>
          <el-table-column align="left" label="模板ID" prop="id" />
          <el-table-column align="left" label="名称" >
            <template #default="{ row }">
              <div v-if="row.info.name" class="line-clamp-2" :title="row.info.name">
                {{ row.info.name }}
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left" label="描述" >
            <template #default="{ row }">
              <div v-if="row.info.description" class="line-clamp-2" :title="row.info.description">
                {{ row.info.description }}
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left" label="标签" >
            <template #default="{ row }">
              <div v-if="row.info" class="gap-1 flex flex-wrap">
                <el-tag v-for="tag in row.info.tags" :key="tag"> {{ tag }}</el-tag>
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left" label="类型" prop="info.severity" >
            <template  #default="{ row }">
              <template v-if="row.info?.severity">
                <el-tag
                  effect="dark"
                  :type="severityLevel(row.info.severity)"
                >{{severityTypeStr(row.info.severity)}}</el-tag>
              </template>
            </template>
          </el-table-column>
<!--        <el-table-column align="left" label="日期" prop="createdAt" width="180">-->
<!--            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>-->
<!--        </el-table-column>-->
        
<!--        <el-table-column align="left" label="操作" fixed="right" min-width="240">-->
<!--            <template #default="scope">-->
<!--            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>-->
<!--            <el-button  type="primary" link icon="edit" class="table-button" @click="updateNucleiFunc(scope.row)">编辑</el-button>-->
<!--            <el-button  type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>-->
<!--            </template>-->
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
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createNuclei,
  deleteNuclei,
  deleteNucleiByIds,
  updateNuclei,
  findNuclei,
  getNucleiList, getNucleiTemplateList
} from '@/api/nucleiInfo/nucleiInfo'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { severityType } from '@/utils/poc'



defineOptions({
    name: 'Nuclei'
})
const multipleSelection = defineModel('multipleSelection')


// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        })



// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const tableLoading = ref(false)
const searchInfo = ref({
  poc_filter: {
    Severity: "critical"
  }
})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
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
const getTableData = async() => {
  // const table = await getNucleiList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  // if (table.code === 0) {
  //   tableData.value = table.data.list
  //   total.value = table.data.total
  //   page.value = table.data.page
  //   pageSize.value = table.data.pageSize
  // }
  tableLoading.value = true
  const table = await getNucleiTemplateList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value})
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
  tableLoading.value = false
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
// 多选
const handleSelectionChange = (val) => {
  console.log(val)
  multipleSelection.value = val
}

function select (selection, row) {
  if (selection && selection.find(item => item && (item.id === row.id))) {
    addRows([row])
  } else {
    deleteRows([row])
  }
}
function selectAll (selection) {
  if (selection.length > 0) {
    addRows(tableData.value)
  } else {
    deleteRows(tableData.value)
  }
}

const selectedItems = ref([])
function addRows (rows) {
  rows.forEach(row => {
    if (selectedItems.value.find(item => item.id === row.id)) { return }
    selectedItems.value.push(row)
  });
}
function deleteRows (rows) {
  if (selectedItems.value.length === 0) { return }
  rows.forEach(row => {
    selectedItems.value = selectedItems.value.filter(item => item.id === row.id)
  })
}



function severityLevel(type) {
  switch (type) {
    case 'critical':
      return 'danger'
    case 'high':
      return 'warning'
    case 'info':
      return 'info'
    case 'medium':
      return 'primary'
    default:
      return 'info'
  }
}
function severityTypeStr(type) {
  switch (type) {
    case 'critical':
      return '严重'
    case 'high':
      return '高危'
    case 'info':
      return '信息'
    case 'medium':
      return '中危'
    case 'unknown':
      return '未知'
    default:
      return type
  }
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteNucleiFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
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
      const res = await deleteNucleiByIds({ IDs })
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

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateNucleiFunc = async(row) => {
    const res = await findNuclei({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteNucleiFunc = async (row) => {
    const res = await deleteNuclei({ ID: row.ID })
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
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'create':
                  res = await createNuclei(formData.value)
                  break
                case 'update':
                  res = await updateNuclei(formData.value)
                  break
                default:
                  res = await createNuclei(formData.value)
                  break
              }
              btnLoading.value = false
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
  const res = await findNuclei({ ID: row.ID })
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


</script>

<style>

</style>
