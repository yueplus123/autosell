import { ref } from 'vue'

/**
 * 全站共用的「联系方式」弹窗开关。
 * 弹窗本体（XinhuoContactModal）只在 XinhuoLayout 里挂一次，
 * 任何地方（「快速响应」卡片、展示模式下的「联系渠道合作」按钮）调用 openContactModal() 即可打开。
 */
export const contactModalOpen = ref(false)

export const openContactModal = () => {
  contactModalOpen.value = true
}

export const closeContactModal = () => {
  contactModalOpen.value = false
}
