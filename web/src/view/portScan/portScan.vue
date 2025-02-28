<script setup>
import {ref} from 'vue'
import {getPortScan, getPortScanList, portScan} from '@/api/highPort/highRiskPortConfig'
import {formatDate} from '@/utils/format'
import ScanPortInfo from '@/components/port/scanPortInfo.vue'
import {portConfig} from '@/utils/port'
import useWebsocket from "@/hooks/use-websocket";

defineOptions({
  name: 'PortScan'
})
const portRadio = ref('top100')
const formData = ref({
  target: '',
  port: portConfig['top100'],
  taskName: ''
})
const visible = ref(false)
const cardLogVisible = ref(false)
const visibleList = ref(false)
const currLogs = ref('')
const currScanInfo = ref([])

const logs = ref([])

const scanData = ref([])
const isLoading = ref(false)

function getScanInfoById(id) {
  scanData.value = []
  getPortScan({id}).then((res) => {
    scanData.value = JSON.parse(res.data?.jsonResult || '[]')
  })
}

function handleInfoList(info) {
  visibleList.value = true
  currScanInfo.value = info.jsonResult
}

const cardListData = ref([])

async function getData() {
  const res = await getPortScanList({pageNo: 0, pageSize: 1000})
  cardListData.value = (res.data?.list || []).map((item) => {
    try {
      const jsonResult = JSON.parse(item.jsonResult || '{}')
      item.jsonResult = jsonResult
      item.filterResult = jsonResult.reduce((acc, curr) => {
        let find = acc.find((item) => item.ip === curr.IP)
        if (find) {
          find.list.push(curr)
          return acc
        }
        acc.push({
          ip: curr.IP,
          list: [curr]
        })
        return acc
      }, [])
    } catch (e) {
      console.log(e)
    }
    return item
  })
}

// getData()

const executeTest = () => {
  // 这里添加执行扫描的逻辑
  isLoading.value = true
  logs.value = []
  portScan(formData.value).then((res) => {
    isLoading.value = false
    console.log('扫描执行完成:', res)
    const id = res.data.id
    getPortScan({id}).then(() => {
      visible.value = true
      const {connectWebSocket} = useWebsocket({
        handleClose() {
          getScanInfoById(id)
        }
      })
      connectWebSocket(id, (data) => {
        data = String(data).replaceAll('"\\r', '').replaceAll('\\n"', '')
        logs.value.push(data)
      })
    })
  }) // 模拟异步操作
}
</script>

<template>
  <div>
    <!--  <el-tabs type="border-card" class="m-[20px]">-->
    <!--    <el-tab-pane>-->
    <!--      <template #label>-->
    <!--        <span class="custom-tabs-label items-center flex gap-1">-->
    <!--          <el-icon><calendar /></el-icon>-->
    <!--          <div>端口扫描</div>-->
    <!--        </span>-->
    <!--      </template>-->
    <div class="penetration-test-container">
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <h2>资产扫描</h2>
          </div>
        </template>
        <el-form :model="formData" label-width="auto">
          <el-row :gutter="20">
            <el-col :span="12" class="flex flex-col">
              <el-form-item label="任务名称" inline>
                <el-input
                    v-model="formData.taskName"
                    placeholder="请输入任务名称"
                ></el-input>
              </el-form-item>
              <el-form-item class="h-full">
                <el-input
                    class="h-full"
                    type="textarea"
                    input-style="height: 100%"
                    v-model="formData.target"
                    placeholder="请输入目标 "
                ></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="端口配置">
                <el-radio-group
                    v-model="portRadio"
                    style="width: 100%"
                    @change="formData.port = portConfig[portRadio]"
                >
                  <el-radio label="web">Web端口</el-radio>
                  <el-radio label="top100">Top100端口</el-radio>
                  <el-radio label="top1000">Top1000端口</el-radio>
                  <el-radio label="rce">常见RCE端口</el-radio>
                  <el-radio label="lite">精简端口</el-radio>
                  <el-radio label="all">全部端口</el-radio>
                </el-radio-group>
              </el-form-item>
              <el-form-item>
                <el-input
                    type="textarea"
                    v-model="formData.port"
                    :autosize="{
                    minRows: 3,
                    maxRows: 3
                  }"
                    placeholder="扫描的端口"
                ></el-input>
              </el-form-item>
              <el-form-item>
                <div class="flex">
                  <el-button
                      type="primary"
                      @click="executeTest"
                      :loading="isLoading"
                      class="execute-btn"
                  >
                    执行扫描
                  </el-button>
                  <el-button v-if="logs.length" type="text" @click="visible = true"
                  >日志
                  </el-button
                  >
                </div>
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
        </el-form>
      </el-card>
      <scan-port-info class="mt-2" :data="scanData"/>
    </div>
    <!--    </el-tab-pane>-->
    <!--  </el-tabs>-->
    <el-dialog v-model="visible" title="运行日志" width="800">
      <el-scrollbar class="bg-black">
        <div class="text-white bg-black p-4">
          <pre v-for="(log, index) in logs" :key="index">{{ log }}</pre>
        </div>
      </el-scrollbar>
      <template #footer></template>
    </el-dialog>

    <el-dialog v-model="cardLogVisible" title="运行日志" width="800">
      <el-scrollbar class="bg-black">
        <div class="text-white bg-black p-4">
          <pre>{{ currLogs }}</pre>
        </div>
      </el-scrollbar>
      <template #footer></template>
    </el-dialog>

    <el-dialog v-model="visibleList" title="详细信息" width="1100">
      <scan-port-info :data="currScanInfo"/>
      <template #footer></template>
    </el-dialog>

    <el-row v-if="cardListData.length && false">
      <el-col :span="24" v-for="card in cardListData" :key="card.id">
        <el-card class="mt-[20px]">
          <template #header>
            <div class="card-header">
              <h2 class="flex gap-[10px]">
                {{ card.task_name }}
                <el-tag>{{ card.target }}</el-tag>
                <el-tag v-if="card.port" type="primary">{{ card.port }}</el-tag>
                <el-tag class="ml-auto" type="info">{{
                    formatDate(card.CreatedAt)
                  }}
                </el-tag>
              </h2>
            </div>
          </template>
          <div class="flex gap-2 flex-wrap">
            <template v-if="card.filterResult.length > 1">
              <div class="w-full">
                <template v-for="res in card.filterResult" :key="res.ip">
                  <div class="font-bold my-1">
                    {{ res.ip }}
                  </div>
                  <div class="flex gap-2 flex-wrap">
                    <el-tag
                        :title="data.URL"
                        v-for="(data, index) in res.list"
                        :key="index"
                    >
                      {{ data.Port }}{{ data.Service && ' / ' + data.Service }}
                    </el-tag>
                  </div>
                </template>
              </div>
            </template>
            <el-tag
                v-else
                :title="data.URL"
                v-for="(data, index) in card.jsonResult"
                :key="index"
            >
              {{ data.Port }}{{ data.Service && ' / ' + data.Service }}
            </el-tag
            >
            <!--        <pre>{{ JSON.stringify(JSON.parse(card.jsonResult), null, 2) }}</pre>-->
          </div>
          <template #footer>
            <div class="flex">
              <el-button
                  size="small"
                  type="info"
                  @click="
                  () => {
                    currLogs = card.info
                    cardLogVisible = true
                  }
                "
              >查看日志
              </el-button
              >
              <el-button
                  size="small"
                  type="success"
                  class="ml-auto"
                  @click="handleInfoList(card)"
              >详细信息
              </el-button
              >
            </div>
          </template>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<style scoped lang="scss"></style>
