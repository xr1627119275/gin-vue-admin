import { WsPath } from '@/utils/format'
import { onMounted, onUnmounted, ref } from 'vue'

export default function useWebsocket({ handleClose = () => {} }) {
  const ws = ref(null)
  const connectWebSocket = (messageId, cb) => {
    ws.value = new WebSocket(WsPath + `system_ws?messageId=${messageId}`) // 替换为你的WebSocket地址
    ws.value.onopen = () => {
      console.log('WebSocket connected')
    }
    ws.value.onmessage = (event) => {
      let data = event.data
      if (data?.includes(`{{over_end}}`)) {
        return wsClose()
      }
      cb && cb(data)
    }

    ws.value.onclose = () => {
      if (ws.value.hasClose) return
      console.log('WebSocket closed, attempting to reconnect...')
      setTimeout(connectWebSocket, 1000) // 重新连接
    }
  }
  function wsClose() {
    try {
      ws.value.hasClose = true
      ws.value.close()
    } catch (e) {
      console.log(e)
    } finally {
      handleClose && handleClose()
    }
  }

  onUnmounted(() => {
    if (ws.value) wsClose()
  })

  return {
    connectWebSocket,
    ws,
    wsClose
  }
}
