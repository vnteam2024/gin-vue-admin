// Listen to the resize event of window and return the width and height of the current window
import { shallowRef } from 'vue'
import { tryOnMounted, useEventListener } from '@vueuse/core'

const width = shallowRef(0)
const height = shallowRef(0)

export const useWindowResize = (cb) => {
  const onResize = () => {
    width.value = window.innerWidth
    height.value = window.innerHeight
    if (cb && typeof cb === 'function') {
      cb(width.value, height.value)
    }
  }

  tryOnMounted(onResize)
  useEventListener('resize', onResize, { passive: true })
  return {
    width,
    height,
  }
}
