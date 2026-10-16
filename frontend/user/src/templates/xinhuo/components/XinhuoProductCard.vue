<template>
  <!-- 主站展示模式：卡片不可点进详情，不显示价格、库存和购买按钮，缺货也按正常样式展示 -->
  <component
    :is="showcase ? 'div' : RouterLink"
    :to="showcase ? undefined : `/products/${product.slug}`"
    class="group flex rounded-[20px] bg-white text-[color:var(--xh-ink)] transition-shadow hover:shadow-[0_8px_30px_rgba(0,0,0,0.08)] max-sm:items-center max-sm:gap-3 max-sm:rounded-2xl max-sm:p-3.5 sm:min-h-[236px] sm:flex-col sm:gap-3.5 sm:p-[22px]"
  >
    <!-- 图标 + 标签 -->
    <div class="flex items-start justify-between max-sm:contents">
      <div class="grid h-[52px] w-[52px] flex-none place-items-center overflow-hidden rounded-[14px] text-xl font-bold text-white max-sm:h-12 max-sm:w-12 max-sm:rounded-xl" :style="iconImage ? undefined : { background: tileColor }">
        <img v-if="iconImage" :src="iconImage" :alt="title" class="h-full w-full object-cover" loading="lazy" @error="imageErrored = true" />
        <span v-else>{{ initial }}</span>
      </div>
      <div class="flex flex-wrap justify-end gap-1.5 max-sm:hidden">
        <span v-for="tag in badges" :key="tag.label" class="rounded-full px-2.5 py-1 text-xs font-medium" :class="tag.tone">{{ tag.label }}</span>
      </div>
    </div>

    <!-- 标题 / 描述 / 元信息 -->
    <div class="min-w-0 max-sm:flex-1">
      <h3 class="truncate text-lg font-bold max-sm:text-[15px]">{{ title }}</h3>
      <p v-if="summary" class="mt-1.5 line-clamp-2 text-sm text-[color:var(--xh-muted)] max-sm:mt-0.5 max-sm:line-clamp-1 max-sm:text-xs">{{ summary }}</p>
      <p class="mt-2 flex flex-wrap gap-x-3 gap-y-1 text-xs text-[color:var(--xh-muted)] max-sm:mt-1 max-sm:text-[11px]">
        <span class="sm:hidden">{{ fulfillmentLabel }}</span>
        <span v-if="!showcase">{{ stockLabel }}</span>
        <span v-if="categoryName">{{ categoryName }}</span>
      </p>
    </div>

    <!-- 价格 + 购买 -->
    <div v-if="!showcase" class="flex items-center justify-between max-sm:flex-none max-sm:flex-col max-sm:items-end max-sm:gap-1.5 sm:mt-auto">
      <div class="flex items-baseline gap-2 max-sm:flex-col max-sm:items-end max-sm:gap-0">
        <span class="text-[22px] font-semibold tabular-nums max-sm:text-base">{{ displayPrice }}</span>
        <span v-if="promo" class="text-[13px] text-[#86868b] line-through max-sm:text-[11px]">{{ originalPrice }}</span>
      </div>
      <span
        class="rounded-full px-[18px] py-2 text-sm font-medium max-sm:px-3.5 max-sm:py-1.5 max-sm:text-[13px]"
        :class="soldOut ? 'bg-[#e8e8ed] text-[color:var(--xh-muted)]' : buyTone"
      >{{ soldOut ? t('xinhuo.shop.soldOut') : t('xinhuo.shop.buy') }}</span>
    </div>
  </component>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { RouterLink } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { useAppStore } from '../../../stores/app'
import { isMainShowcase } from '../../../utils/storefrontZone'
import { getFirstImageUrl, getImageUrl } from '../../../utils/image'
import { useLocalized, useProductLabels } from '../../../composables/useProduct'

// tone：专区配色（ai 黑色按钮、game 粉紫渐变按钮），默认薪火蓝
const props = withDefaults(defineProps<{ product: any; index?: number; tone?: 'default' | 'ai' | 'game' }>(), { index: 0, tone: 'default' })

const BUY_TONES = {
  default: 'bg-[color:var(--xh-accent)] text-white group-hover:bg-[color:var(--xh-accent-hover)]',
  ai: 'bg-[#0b0b0c] text-white group-hover:bg-black',
  game: 'bg-linear-to-r from-violet-600 via-fuchsia-500 to-orange-400 text-white group-hover:brightness-110',
} as const
const buyTone = computed(() => BUY_TONES[props.tone ?? 'default'])

// 主站展示模式（或接口已隐藏价格）时只展示，不可购买
const appStore = useAppStore()
const showcase = computed(() => isMainShowcase(appStore.config as any) || props.product?.price_hidden === true)

const { t } = useI18n()
const { getLocalizedText, formatPrice } = useLocalized()
const { getStockStatusLabel, getFulfillmentTypeLabel, isSoldOut, hasPromotionPrice, getPromotionPriceAmount } = useProductLabels()

const tileColors = ['#c2410c', '#047857', '#1d1d1f', '#4338ca', '#be185d', '#334155', '#9a3412', '#0369a1']
const tileColor = computed(() => tileColors[(props.index ?? 0) % tileColors.length])

const title = computed(() => getLocalizedText(props.product?.title))
const summary = computed(() => getLocalizedText(props.product?.description))
const categoryName = computed(() => getLocalizedText(props.product?.category?.name))
const initial = computed(() => (title.value || '?').trim().charAt(0).toUpperCase())

const imageErrored = ref(false)
const iconImage = computed(() => {
  if (imageErrored.value) return ''
  const primary = getFirstImageUrl(props.product?.images)
  if (primary) return primary
  const icon = props.product?.category?.icon
  return icon ? getImageUrl(icon) : ''
})

const soldOut = computed(() => isSoldOut(props.product))
const promo = computed(() => hasPromotionPrice(props.product))
const displayPrice = computed(() => formatPrice(promo.value ? getPromotionPriceAmount(props.product) : props.product?.price_amount))
const originalPrice = computed(() => formatPrice(props.product?.price_amount))
const stockLabel = computed(() => getStockStatusLabel(props.product))
const fulfillmentLabel = computed(() => getFulfillmentTypeLabel(props.product?.fulfillment_type))

// 发货方式 + 后台给商品打的前两个标签（可用来写“质保 30 天”“独享账号”等）
const badges = computed(() => {
  const list: { label: string; tone: string }[] = []
  list.push({
    label: fulfillmentLabel.value,
    tone: props.product?.fulfillment_type === 'auto' ? 'bg-[#e3f5ea] text-[#0f6b3a]' : 'bg-[#fdf1dc] text-[#8a4b08]',
  })
  const tags: unknown[] = Array.isArray(props.product?.tags) ? props.product.tags : []
  tags
    .map((tag) => String(tag || '').trim())
    .filter(Boolean)
    .slice(0, 2)
    .forEach((label) => list.push({ label, tone: 'bg-[#e5edff] text-[#1e3a8a]' }))
  return list
})
</script>
