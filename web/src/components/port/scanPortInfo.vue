<script setup>
  import { createAFrogWebScan, findWebScan } from '@/api/webScan/webScan'
  import useWebsocket from '@/hooks/use-websocket'
  defineProps({
    data: {
      type: Array,
      default: () => []
    }
  })
  async function handleClick(row) {
    let res = await createAFrogWebScan({ target: row.URL })
    if (res.code === 0) {
      const id = res.data.id
      const { connectWebSocket } = useWebsocket({
        async handleClose() {
          let res = await findWebScan({ ID: id })
          console.log(res.data)
        }
      })

      connectWebSocket(id, (data) => {
        console.log(data)
      })
    }
  }
</script>

<template>
  <el-table :data="data" border>
    <el-table-column label="IP" prop="IP" width="150" />
    <el-table-column label="端口" prop="Port" width="100" />
    <el-table-column label="Service" prop="Service" width="100" />
    <el-table-column label="Keyword" prop="Keyword" />
    <el-table-column label="Response" prop="Response">
      <template #default="{ row }">
        <div v-if="row.Response" class="line-clamp-2" :title="row.Response">
          {{ row.Response }}
        </div>
      </template>
    </el-table-column>
    <el-table-column label="操作" width="150">
      <template #default="{ row }">
        <el-button v-if="row?.URL?.startsWith('http')" @click="handleClick(row)"
          >web漏洞扫描</el-button
        >
      </template>
    </el-table-column>
  </el-table>
</template>

<style scoped lang="scss"></style>
