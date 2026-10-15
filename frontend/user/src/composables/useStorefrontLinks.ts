import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { useAppStore } from '../stores/app'
import { currentZone, storefrontHomePath, storefrontShopPath } from '../utils/storefrontZone'

/**
 * 站内「首页 / 继续购物」链接的统一出口。
 * 顾客在专区里（或分站只代理一个专区）时，这些链接都回到专区，而不是总首页或全部商品页。
 */
export function useStorefrontLinks() {
  const appStore = useAppStore()
  const { t } = useI18n()

  const zone = computed(() => currentZone(appStore.config as any))
  const homePath = computed(() => storefrontHomePath(appStore.config as any))
  const shopPath = computed(() => storefrontShopPath(appStore.config as any))
  /** 面包屑第一级的文字：在专区时显示专区名 */
  const homeLabel = computed(() => (zone.value ? t(`zone.${zone.value}.badge`) : t('nav.home')))

  return { zone, homePath, shopPath, homeLabel }
}
