<template>
  <div>
    <div class="gva-form-box">
      <el-form
        :model="formData"
        ref="elFormRef"
        label-position="right"
        :rules="rule"
        label-width="80px"
      >
        <el-form-item label="端口号:" prop="portNumber">
          <el-input
            v-model.number="formData.portNumber"
            :clearable="false"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="端口名称:" prop="portName">
          <el-input
            v-model="formData.portName"
            :clearable="false"
            placeholder="请输入端口名称"
          />
        </el-form-item>
        <el-form-item label="风险等级:" prop="riskLevel">
          <el-input
            v-model="formData.riskLevel"
            :clearable="false"
            placeholder="请输入风险等级"
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
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save"
            >保存</el-button
          >
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
  import {
    createHighRiskPortConfig,
    updateHighRiskPortConfig,
    findHighRiskPortConfig
  } from '@/api/highPort/highRiskPortConfig'

  defineOptions({
    name: 'HighRiskPortConfigForm'
  })

  // 自动获取字典
  import { getDictFunc } from '@/utils/format'
  import { useRoute, useRouter } from 'vue-router'
  import { ElMessage } from 'element-plus'
  import { ref, reactive } from 'vue'
  // 富文本组件
  import RichEdit from '@/components/richtext/rich-edit.vue'

  const route = useRoute()
  const router = useRouter()

  // 提交按钮loading
  const btnLoading = ref(false)

  const type = ref('')
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
      }
    ],
    riskLevel: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      }
    ]
  })

  const elFormRef = ref()

  // 初始化方法
  const init = async () => {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findHighRiskPortConfig({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
  }

  init()
  // 保存按钮
  const save = async () => {
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
      }
    })
  }

  // 返回按钮
  const back = () => {
    router.go(-1)
  }
</script>

<style></style>
