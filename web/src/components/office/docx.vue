<template>
  <vue-office-docx
    :src="docx"
    @rendered="rendered"
  />
</template>
<script>
export default {
  name: 'Docx'
}
</script>
<script setup>
import { ref, watch } from 'vue'

//Introduce VueOfficeDocx component
import VueOfficeDocx from '@vue-office/docx'
//Introduce related styles
import '@vue-office/docx/lib/index.css'

const model = defineModel({
  type: String,
})

const docx = ref(null)
watch(
  () => model,
  value => { docx.value = value },
  { immediate: true }
)
const rendered = () => {
}
</script>
<style lang="scss" scoped>

</style>
