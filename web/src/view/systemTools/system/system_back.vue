<template>
  <div class="system">
    <el-form
      ref="form"
      :model="config"
      label-width="240px"
    >
      <!--  System start  -->
      <el-collapse v-model="activeNames">
        <el-collapse-item
          title="系统配置"
          name="1"
        >
          <el-form-item label="抓包网卡">
            <el-select
                v-model="config.scan['device']"
                style="width:100%"
            >
              <el-option v-for="device in deviceInfos" :key="device.Name" :value="device.Name" :label="device.Description" :title="device.Addresses && device.Addresses.map(item => item.IP).join('\n')" />
              <el-option value="pgsql" />
            </el-select>
          </el-form-item>
          <el-form-item label="端口值">
            <el-input v-model.number="config.system.addr" />
          </el-form-item>
          <el-form-item label="数据库类型">
            <el-select
              v-model="config.system['db-type']"
              style="width:100%"
            >
              <el-option value="mysql" />
              <el-option value="pgsql" />
            </el-select>
          </el-form-item>
          <el-form-item label="Oss类型">
            <el-select
              v-model="config.system['oss-type']"
              style="width:100%"
            >
              <el-option value="local" />
              <el-option value="qiniu" />
              <el-option value="tencent-cos" />
              <el-option value="aliyun-oss" />
              <el-option value="huawei-obs" />
            </el-select>
          </el-form-item>
          <el-form-item label="多点登录拦截">
            <el-checkbox v-model="config.system['use-multipoint']">开启</el-checkbox>
          </el-form-item>
          <el-form-item label="开启redis">
            <el-checkbox v-model="config.system['use-redis']">开启</el-checkbox>
          </el-form-item>
          <el-form-item label="限流次数">
            <el-input-number v-model.number="config.system['iplimit-count']" />
          </el-form-item>
          <el-form-item label="限流时间">
            <el-input-number v-model.number="config.system['iplimit-time']" />
          </el-form-item>
          <el-tooltip
            content="请修改完成后，注意一并修改前端env环境下的VITE_BASE_PATH"
            placement="top-start"
          >
            <el-form-item label="全局路由前缀">
              <el-input v-model="config.system['router-prefix']" />
            </el-form-item>
          </el-tooltip>
        </el-collapse-item>
      </el-collapse>
    </el-form>
    <div class="mt-4">
      <el-button
        type="primary"
        @click="update"
      >立即更新</el-button>
      <el-button
        type="primary"
        @click="reload"
      >重启服务（开发中）</el-button>
    </div>
  </div>
</template>

<script setup>
import {getNetDrivers, getSystemConfig, setSystemConfig} from '@/api/system'
import { emailTest } from '@/api/email'
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'

defineOptions({
  name: 'Config'
})

const activeNames = reactive([])
const config = ref({
  system: {
    'iplimit-count': 0,
    'iplimit-time': 0
  },
  scan: {
    device: ''
  },
  jwt: {},
  mysql: {},
  pgsql: {},
  excel: {},
  autocode: {},
  redis: {},
  mongo: {
    coll: '',
    options: '',
    database: '',
    username: '',
    password: '',
    'min-pool-size': '',
    'max-pool-size': '',
    'socket-timeout-ms': '',
    'connect-timeout-ms': '',
    'is-zap': '',
    hosts: [
      {
        host: '',
        port: ''
      }
    ]
  },
  qiniu: {},
  'tencent-cos': {},
  'aliyun-oss': {},
  'hua-wei-obs': {},
  captcha: {},
  zap: {},
  local: {},
  email: {},
  timer: {
    detail: {}
  }
})
const deviceInfos = ref([])

const initForm = async() => {
  const res = await getSystemConfig()
  if (res.code === 0) {
    config.value = res.data.config
  }
  const deviceRes = await getNetDrivers()
  deviceInfos.value = deviceRes.data['devices']
  // console.log("device: ", deviceRes.data['devices'])
}
initForm()
const reload = () => {}
const update = async() => {
  const res = await setSystemConfig({ config: config.value })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '配置文件设置成功'
    })
    await initForm()
  }
}
const email = async() => {
  const res = await emailTest()
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '邮件发送成功'
    })
    await initForm()
  } else {
    ElMessage({
      type: 'error',
      message: '邮件发送失败'
    })
  }
}

</script>

<style lang="scss">
.system {
  @apply bg-white p-9 rounded dark:bg-slate-900;
  h2 {
    @apply p-2.5 my-2.5 text-lg shadow;
  }
}
</style>
