<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog"
          >新增</el-button
        >
        <el-button
          icon="delete"
          style="margin-left: 10px"
          :disabled="!multipleSelection.length"
          @click="onDelete"
          >删除</el-button
        >
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

        <el-table-column align="left" label="日期" prop="createdAt" width="180">
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>

        <el-table-column
          sortable
          align="left"
          label="端口号"
          prop="portNumber"
          width="120"
        />
        <el-table-column
          align="left"
          label="端口名称"
          prop="portName"
          width="120"
        />
        <!--        <el-table-column-->
        <!--          sortable-->
        <!--          align="left"-->
        <!--          label="风险等级"-->
        <!--          prop="riskLevel"-->
        <!--          width="120"-->
        <!--        />-->

        <el-table-column label="描述" prop="portDescription" width="200">
          <template #default="scope">
            {{ scope.row.portDescription }}
          </template>
        </el-table-column>
        <el-table-column label="启用检测" prop="status" width="200">
          <template #default="{ row }">
            <el-switch
              v-model="row.status"
              :active-value="1"
              :inactive-value="0"
              :before-change="
                () =>
                  updateHighRiskPortConfig({
                    ...row,
                    status: Number(!row.status)
                  })
              "
            />
          </template>
        </el-table-column>
        <el-table-column
          sortable
          align="left"
          label="排序"
          prop="order"
          width="120"
        />
        <el-table-column
          align="left"
          label="操作"
          fixed="right"
          min-width="240"
        >
          <template #default="scope">
            <!--            <el-button-->
            <!--              type="primary"-->
            <!--              link-->
            <!--              class="table-button"-->
            <!--              @click="getDetails(scope.row)"-->
            <!--              ><el-icon style="margin-right: 5px"><InfoFilled /></el-icon-->
            <!--              >查看</el-button-->
            <!--            >-->
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateHighRiskPortConfigFunc(scope.row)"
              >编辑</el-button
            >
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteRow(scope.row)"
              >删除</el-button
            >
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
    <el-drawer
      destroy-on-close
      size="800"
      v-model="dialogFormVisible"
      :show-close="false"
      :before-close="closeDialog"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '新增' : '编辑' }}</span>
          <div>
            <el-button :loading="btnLoading" type="primary" @click="enterDialog"
              >确 定</el-button
            >
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form
        :model="formData"
        label-position="top"
        ref="elFormRef"
        :rules="rule"
        label-width="80px"
      >
        <el-form-item label="端口号:" prop="portNumber">
          <el-input
            v-model.number="formData.portNumber"
            :clearable="false"
            placeholder="请输入端口号"
          />
        </el-form-item>
        <el-form-item label="端口名称:" prop="portName">
          <el-input
            v-model="formData.portName"
            :clearable="false"
            placeholder="请输入端口名称"
          />
        </el-form-item>
        <!--        <el-form-item label="风险等级:" prop="riskLevel">-->
        <!--          <el-input-->
        <!--            v-model="formData.riskLevel"-->
        <!--            :clearable="false"-->
        <!--            placeholder="请输入风险等级"-->
        <!--          />-->
        <!--        </el-form-item>-->

        <el-form-item label="排序:" prop="riskLevel">
          <el-input
            v-model="formData.order"
            :clearable="false"
            placeholder="请输入排序"
          />
        </el-form-item>

        <el-form-item label="描述:" prop="portDescription">
          <el-input
            v-model="formData.portDescription"
            :clearable="false"
            placeholder="请输入描述"
            type="textarea"
            :rows="5"
          />
        </el-form-item>
      </el-form>
    </el-drawer>

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
        <!--        <el-descriptions-item label="风险等级">-->
        <!--          {{ detailFrom.riskLevel }}-->
        <!--        </el-descriptions-item>-->
        <el-descriptions-item label="描述">
          {{ detailFrom.portDescription }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
  import {
    createHighRiskPortConfig,
    deleteHighRiskPortConfig,
    deleteHighRiskPortConfigByIds,
    updateHighRiskPortConfig,
    findHighRiskPortConfig,
    getHighRiskPortConfigList
  } from '@/api/highPort/highRiskPortConfig'
  // 富文本组件
  import RichEdit from '@/components/richtext/rich-edit.vue'

  // 全量引入格式化工具 请按需保留
  import {
    getDictFunc,
    formatDate,
    formatBoolean,
    filterDict,
    filterDataSource,
    returnArrImg,
    onDownloadFile
  } from '@/utils/format'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, reactive } from 'vue'

  defineOptions({
    name: 'HighRiskPortConfig'
  })
  // 提交按钮loading
  const btnLoading = ref(false)

  // 控制更多查询条件显示/隐藏状态
  const showAllQuery = ref(false)

  // 自动化生成的字典（可能为空）以及字段
  const formData = ref({
    portNumber: 0,
    portName: '',
    riskLevel: '',
    portDescription: ''
  })

  // 验证规则
  const rule = reactive({
    portNumber: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      }
    ],
    portName: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      },
      {
        whitespace: true,
        message: '不能只输入空格',
        trigger: ['input', 'blur']
      }
    ],
    riskLevel: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      },
      {
        whitespace: true,
        message: '不能只输入空格',
        trigger: ['input', 'blur']
      }
    ]
  })

  const searchRule = reactive({
    createdAt: [
      {
        validator: (rule, value, callback) => {
          if (
            searchInfo.value.startCreatedAt &&
            !searchInfo.value.endCreatedAt
          ) {
            callback(new Error('请填写结束日期'))
          } else if (
            !searchInfo.value.startCreatedAt &&
            searchInfo.value.endCreatedAt
          ) {
            callback(new Error('请填写开始日期'))
          } else if (
            searchInfo.value.startCreatedAt &&
            searchInfo.value.endCreatedAt &&
            (searchInfo.value.startCreatedAt.getTime() ===
              searchInfo.value.endCreatedAt.getTime() ||
              searchInfo.value.startCreatedAt.getTime() >
                searchInfo.value.endCreatedAt.getTime())
          ) {
            callback(new Error('开始日期应当早于结束日期'))
          } else {
            callback()
          }
        },
        trigger: 'change'
      }
    ]
  })

  const elFormRef = ref()
  const elSearchFormRef = ref()

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
    const table = await getHighRiskPortConfigList({
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

  // 删除行
  const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      deleteHighRiskPortConfigFunc(row)
    })
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
        multipleSelection.value.map((item) => {
          IDs.push(item.ID)
        })
      const res = await deleteHighRiskPortConfigByIds({ IDs })
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
  const updateHighRiskPortConfigFunc = async (row) => {
    const res = await findHighRiskPortConfig({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data
      dialogFormVisible.value = true
    }
  }

  // 删除行
  const deleteHighRiskPortConfigFunc = async (row) => {
    const res = await deleteHighRiskPortConfig({ ID: row.ID })
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
      portNumber: 0,
      portName: '',
      riskLevel: '',
      portDescription: ''
    }
  }
  // 弹窗确定
  const enterDialog = async () => {
    btnLoading.value = true
    elFormRef.value?.validate(async (valid) => {
      if (!valid) return (btnLoading.value = false)
      let res
      switch (type.value) {
        case 'create':
          res = await createHighRiskPortConfig(formData.value)
          break
        case 'update':
          res = await updateHighRiskPortConfig(formData.value)
          break
        default:
          res = await createHighRiskPortConfig(formData.value)
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
</script>

<style></style>
