<script setup>
import { setPassDict, getPassDict } from '@/api/highPort/passwordrule'
import { onMounted, ref } from 'vue'
import { ElMessage } from 'element-plus'

const dicts = ref([])
const fileContent = ref('')
const loading = ref(false)

function getData() {
  loading.value = true
  getPassDict().then((res) => {
    dicts.value = res.data || []
    fileContent.value = dicts.value.join('\n')
    loading.value = false
  })
}

function handleSave() {
  if (!fileContent.value.trim()) {
    ElMessage.warning('请输入密码字典内容')
    return
  }
  
  loading.value = true
  setPassDict({ content: fileContent.value }).then(res => {
    if (res.code === 0) {
      ElMessage.success('保存成功')
      getData()
    } else {
      ElMessage.error('保存失败')
    }
    loading.value = false
  })
}

onMounted(() => {
  getData()
})
</script>

<template>
  <div class="password-dict">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>密码字典配置</span>
          <el-button type="primary" @click="handleSave" :loading="loading">
            保存配置
          </el-button>
        </div>
      </template>
      
      <el-input
        v-model="fileContent"
        type="textarea"
        :rows="15"
        placeholder="请输入密码字典内容，每行一个密码"
      />
      
      <div class="tips">
        <el-alert
          title="提示：每行输入一个密码，系统会自动过滤空行"
          type="info"
          :closable="false"
        />
      </div>
    </el-card>
  </div>
</template>

<style scoped lang="scss">
.password-dict {
  padding: 20px;
  
  .card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  
  .tips {
    margin-top: 16px;
  }
}
</style>
