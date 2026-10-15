<template>
  <section class="mx-auto max-w-[1200px] px-4 py-10 sm:px-6 sm:py-14">
    <h2 class="text-2xl font-black sm:text-[32px]">{{ t('xinhuo.why.title', { brand: brandName }) }}</h2>
    <p class="mb-4 mt-1.5 text-[15px] text-[color:var(--xh-muted)] sm:mb-6 sm:text-[17px]">{{ t('xinhuo.why.subtitle') }}</p>

    <!-- 手机和电脑都是一排 2 个，4 项分 2 行 -->
    <div class="grid grid-cols-2 gap-2.5 sm:gap-5">
      <component
        :is="item.contact ? 'button' : 'div'"
        v-for="(item, idx) in items"
        :key="idx"
        :type="item.contact ? 'button' : undefined"
        class="rounded-2xl bg-white p-4 text-left sm:rounded-[20px] sm:p-7"
        :class="item.contact ? 'group cursor-pointer ring-[color:var(--xh-accent)] transition hover:-translate-y-0.5 hover:shadow-[0_10px_30px_rgba(0,0,0,0.08)] hover:ring-2' : ''"
        @click="item.contact ? openContactModal() : undefined"
      >
        <component :is="ICONS[idx % ICONS.length]" class="h-[22px] w-[22px] text-[color:var(--xh-accent)] sm:h-7 sm:w-7" aria-hidden="true" />
        <h3 class="mb-1 mt-2.5 text-[15px] font-bold sm:mb-1.5 sm:mt-4 sm:text-lg">{{ item.title }}</h3>
        <p class="text-xs leading-relaxed text-[color:var(--xh-muted)] sm:text-sm sm:leading-[1.7]">{{ item.desc }}</p>
        <span v-if="item.contact" class="mt-3 inline-flex items-center gap-1 text-xs font-medium text-[color:var(--xh-accent)] sm:text-sm">
          {{ t('xinhuo.contact.open') }}
          <ChevronRight class="h-3.5 w-3.5 transition-transform group-hover:translate-x-0.5" aria-hidden="true" />
        </span>
      </component>
    </div>
  </section>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { ChevronRight, Headphones, ShieldCheck, Zap, BadgeCheck } from 'lucide-vue-next'
import { useAppStore } from '../../../stores/app'
import { openContactModal } from '../contactModal'

const { t, te } = useI18n()
const appStore = useAppStore()

const brandName = computed(() => String(appStore.config?.brand?.site_name || '').trim() || '薪火数卡')

// 条目来自语言包 xinhuo.why.items.<n>，带 contact 的条目点击后弹出联系方式（弹窗在 XinhuoLayout 中）
const ICONS = [BadgeCheck, Zap, ShieldCheck, Headphones]
const items = computed(() => {
  void appStore.locale
  const list: { title: string; desc: string; contact: boolean }[] = []
  for (let i = 0; i < 12; i++) {
    const k = `xinhuo.why.items.${i}`
    if (!te(`${k}.title`)) break
    list.push({ title: t(`${k}.title`), desc: t(`${k}.desc`), contact: te(`${k}.contact`) })
  }
  return list
})
</script>
