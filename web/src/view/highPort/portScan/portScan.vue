<script setup>
  import { ref } from 'vue'
  import { getPortScan, portScan } from '@/api/highPort/highRiskPortConfig'
  import { WsPath } from '@/utils/format'
  import ScanPortInfo from '@/components/port/scanPortInfo.vue'

  const formData = ref({
    target: '',
    port: '',
    taskName: ''
  })
  const visible = ref(false)

  const logs = ref([
    '[+]2024/12/06 14:11:17 当前环境为：windows, 输出编码为：utf-8\n',
    '[+]2024/12/06 14:11:18 成功加载HTTP指纹:[24758]条\n',
    '[+]2024/12/06 14:11:18 成功加载NMAP探针:[150]个,指纹[11916]条\n',
    '[+]2024/12/06 14:11:19 Domain、IP、Port、URL、Hydra引擎已准备就绪\n',
    '[+]2024/12/06 14:11:19 所有扫描任务已下发完毕\n',
    'ssh://6.6.6.3:22               ssh                        Version:8.2,Info:protocol2.0,Digest:\\"SSH-2.0-OpenSSH_8.2\\\\r\\\\n,Length:21,ProductName:OpenSSH,Port:22\n',
    'telnet://6.6.6.3:23            response is empty          Port:23,Length:0\n',
    'netbios-ns://6.6.6.3:137       netbios-ns                 Length:193,Hostname:XR_DISK,Port:137,Digest:CKAAAAAAAAAAAAAAAAAAAAA,ProductName:Sambanmbdnetbios-ns,Info:workgroup:WORKGROUP\n',
    'https://6.6.6.3:443            Hello!WelcometoSynologyWebStation! Digest:size:13px;\\\\ncolor:,Length:1949,FingerPrint:HTML5;nginx,Port:443\n',
    'mysql://6.6.6.3:3306           mysql                      ProductName:MySQL,Digest:5.7.44(p\\\\\\"vO_P=).(\\\\\\\\9_2f,Length:78,Port:3306,Version:5.7.44\n',
    'http://6.6.6.3:80              Hello!WelcometoSynologyWebStation! Port:80,FingerPrint:HTML5;nginx,Digest:min-height:580px;\\\\n,Length:1930\n',
    'netbios://6.6.6.3:139          netbios                    Version:4.6.2,DeviceType:v,Port:139,Digest:0\\\\x00\\\\x00\\\\x00\\\\x00\\\\x00@\\\\x,Length:41,ProductName:Sambasmbd\n',
    'rsync://6.6.6.3:873            rsync                      Digest:@ERROR:protocolstartup,Length:31,Port:873\n',
    'https://6.6.6.3:5001           xr_disk - Synology DiskStation FingerPrint:群晖-diskstation;Synology-DiskStation;HTML5;nginx,Digest:提供功能完整的网络存储解决方案，可让您轻松管理及,Port:5001,Length:14048\n',
    'http://6.6.6.3:5000            xr_disk - Synology DiskStation Port:5000,FingerPrint:群晖-diskstation;Synology-DiskStation;HTML5;nginx,Digest:提供功能完整的网络存储解决方案，可让您轻松管理及,Length:14071\n',
    'http://6.6.6.3:8000                                       Digest:\\"\\\\r\\\\n\\\\r\\\\nNotfound\\",Length:132,Port:8000\n',
    'http://6.6.6.3:8080            qBittorrent                Length:1605,FingerPrint:Perl,Port:8080,Digest:ption\\\\\\"content=\\\\\\"VueTor\n',
    'smb://6.6.6.3:445              smb                        Port:445,Digest:\\"SMB@Axr_diskGDyGJ`H+\\u003e0\\u003c,Length:258\n',
    'http://6.6.6.3:9000            Portainer                  Port:9000,Length:17718,FingerPrint:Portainer;Meta-Author;HTML5;portainer,FoundDomain:.googlecode.com、state.cu,Digest:\\"正在加载OS42aC0xMS4xVjE4Ni4\n',
    'http://6.6.6.3:8880            404NotFound                FingerPrint:nginx,Digest:\\"NotFound\\u003c/h1\\u003e\\u003c/center,Port:8880,Length:270\n',
    'http://6.6.6.3:8787            Bililive-go                Length:2260,Port:8787,Digest:6220e936.chunk.css\\\\\\"rel\n',
    'http://6.6.6.3:9876                                       Digest:\\"\\\\r\\\\n\\\\r\\\\n\\",Length:82,Port:9876\n',
    "http://6.6.6.3:8092            百度网盘｜Synology         Port:8092,Digest:百度网盘｜oscript\\u003e\\u003cstrong\\u003eWe',Length:883,FingerPrint:Perl;nginx/1.18.0;nginx\n",
    'http://6.6.6.3:4001            404NotFound                FingerPrint:Werkzeug/2.2.3Python/3.7.17,Length:369,Digest:\\"ouenteredtheURLmanu,Port:4001\n',
    'http://6.6.6.3:8096            Jellyfin                   Digest:\\"n:portrait)\\\\\\"rel=\\\\\\"ap,Port:8096,FingerPrint:JQuery;jellyfin;Kestrel,Length:10212\n',
    'http://6.6.6.3:7500            404NotFound                Length:369,Port:7500,FingerPrint:Werkzeug/3.1.3Python/3.9.20,Digest:eredtheURLmanuallypl\n',
    '[+]2024/12/06 14:11:34 程序执行总时长为：[16.2769902s]\n'
  ])

  const scanData = ref([])
  const isLoading = ref(false)
  function getScanInfoById(id) {
    scanData.value = []
    getPortScan({ id }).then((res) => {
      scanData.value = JSON.parse(res.data?.jsonResult || '[]')
    })
  }
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
        visible.value = true
        const ws = ref()
        const connectWebSocket = () => {
          ws.value = new WebSocket(WsPath + `system_ws?messageId=${id}`) // 替换为你的WebSocket地址
          ws.value.onopen = () => {
            console.log('WebSocket connected')
          }
          ws.value.onmessage = (event) => {
            let data = event.data
            data = String(data).replaceAll('"\\r', '').replaceAll('\\n"', '')
            if (data === '"{{over}}"\n') {
              ws.value.hasClose = true
              getScanInfoById(id)
              return ws.value.close()
            }
            logs.value.push(data)
          }

          ws.value.onclose = () => {
            if (ws.value.hasClose) return
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
              <h2>新建扫描任务</h2>
            </div>
          </template>
          <el-form :model="formData" label-position="top">
            <el-row :gutter="20">
              <el-col :span="24">
                <el-form-item label="任务名称">
                  <el-input
                    v-model="formData.taskName"
                    placeholder="请输入任务名称"
                  ></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="16">
                <el-form-item label="目标">
                  <el-input
                    v-model="formData.target"
                    placeholder="请输入目标 "
                  ></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="8">
                <el-form-item label="端口">
                  <el-input
                    v-model="formData.port"
                    placeholder="请输入端口"
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
        <el-card v-if="scanData.length" class="box-card mt-2">
          <template #header>
            <div class="card-header">
              <h2>结果</h2>
            </div>
          </template>
          <div class="-m-2">
            <scan-port-info :data="scanData" />
          </div>
        </el-card>
        <el-dialog v-model="visible" title="运行日志" width="800">
          <el-scrollbar class="bg-black">
            <div class="text-white bg-black p-4">
              <pre v-for="(log, index) in logs" :key="index">{{ log }}</pre>
            </div>
          </el-scrollbar>
          <template #footer> </template>
        </el-dialog>
      </div>
    </el-tab-pane>
  </el-tabs>
</template>

<style scoped lang="scss"></style>
