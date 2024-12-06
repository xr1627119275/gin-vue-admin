<script setup>
  defineProps({
    data: {
      type: Array,
      default: () => []
    }
  })
  function handleClick(scope) {
    console.log(scope)
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
        <span v-if="row.Response" :title="row.Response">{{
          row.Response.slice(0, 60)
        }}</span>
      </template>
    </el-table-column>
    <el-table-column label="操作" width="150">
      <template #default="scope">
        <el-button @click="handleClick(scope)">漏洞扫描</el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<style scoped lang="scss"></style>
