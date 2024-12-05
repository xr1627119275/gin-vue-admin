<script setup>
  import { ref } from 'vue'
  import { getPortScan, portScan } from '@/api/highPort/highRiskPortConfig'
  import { WsPath } from '@/utils/format'

  const formData = ref({
    target: '',
    port: ''
  })
  const logs = ref([])
  const isLoading = ref(false)

  const executeTest = () => {
    // 这里添加执行扫描的逻辑
    isLoading.value = true
    logs.value = []
    portScan(formData.value).then((res) => {
      isLoading.value = false
      console.log('扫描执行完成:', res)
      const id = res.data.id
      getPortScan({ id }).then((res) => {
        console.log(res)
        const ws = ref()
        const connectWebSocket = () => {
          ws.value = new WebSocket(WsPath + `system_ws?messageId=${id}`) // 替换为你的WebSocket地址
          ws.value.onopen = () => {
            console.log('WebSocket connected')
          }
          ws.value.onmessage = (event) => {
            console.log(event.data)
            logs.value.push(
              String(event.data).replaceAll('"\\r', '').replaceAll('\\n"', '')
            )
          }

          ws.value.onclose = () => {
            console.log('WebSocket closed, attempting to reconnect...')
            setTimeout(connectWebSocket, 1000) // 重新连接
          }
        }
        connectWebSocket()
      })
    }) // 模拟异步操作
  }
</script>

<template>
  <el-tabs type="border-card" class="m-[20px]">
    <el-tab-pane>
      <template #label>
        <span class="custom-tabs-label items-center flex gap-1">
          <el-icon><calendar /></el-icon>
          <div>端口扫描</div>
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
                <el-form-item label="目标">
                  <el-input
                    v-model="formData.target"
                    placeholder="请输入目标 （url/ip）"
                  ></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="8">
                <el-form-item label="端口">
                  <el-input
                    v-model="formData.port"
                    placeholder="set port ranges to scan，default is top100"
                  ></el-input>
                </el-form-item>
              </el-col>
              <!--              <el-col :span="24">-->
              <!--                <el-form-item label="其他">-->
              <!--                  <el-checkbox-group v-model="formData.args.extra">-->
              <!--                    <el-checkbox label="t1000" :value="'t1000'">scan top1000 ports</el-checkbox>-->
              <!--                  </el-checkbox-group>-->
              <!--                </el-form-item>-->
              <!--              </el-col>-->
            </el-row>
            <el-form-item>
              <el-button
                type="primary"
                @click="executeTest"
                :loading="isLoading"
                class="execute-btn"
              >
                执行
              </el-button>
            </el-form-item>
          </el-form>
        </el-card>

        <el-card v-if="logs.length" class="box-card mt-2">
          <template #header>
            <div class="card-header">
              <h2>日志</h2>
            </div>
          </template>
          <div class="text-white bg-black">
            <pre v-for="(log, index) in logs" :key="index">{{ log }}</pre>
          </div>
        </el-card>
      </div>
    </el-tab-pane>
  </el-tabs>
</template>

<style scoped lang="scss"></style>
