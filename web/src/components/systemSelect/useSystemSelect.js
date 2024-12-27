import { ref } from 'vue'

export const showSelectSystem = ref(false)
export const selectCallback = ref(() => {})
export const setSelectSystem = (target) => {
  showSelectSystem.value = target
  if (target) {
    return new Promise((resolve) => {
      selectCallback.value = (conf, index) => {
        resolve({ conf, index })
        showSelectSystem.value = false
      }
    })
  }
}
