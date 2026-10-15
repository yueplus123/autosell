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
        @click="item.contact ? (contactOpen = true) : undefined"
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

    <!-- 联系方式弹窗：读取后台「设置 → 联系方式」 -->
    <Teleport to="body">
      <Transition name="xh-fade">
        <div v-if="contactOpen" class="xinhuo-scope fixed inset-0 z-[80] grid place-items-center bg-black/50 p-4 backdrop-blur-sm" @click.self="contactOpen = false">
          <div class="w-full max-w-sm rounded-3xl bg-white p-6 text-[#1d1d1f] shadow-2xl" role="dialog" aria-modal="true" :aria-label="t('xinhuo.contact.title')">
            <div class="flex items-start justify-between gap-4">
              <div>
                <h3 class="text-lg font-bold">{{ t('xinhuo.contact.title') }}</h3>
                <p class="mt-1 text-sm text-[#6e6e73]">{{ t('xinhuo.contact.desc') }}</p>
              </div>
              <button type="button" class="grid h-9 w-9 flex-none place-items-center rounded-full hover:bg-black/5" :aria-label="t('xinhuo.close')" @click="contactOpen = false">
                <X class="h-4 w-4" />
              </button>
            </div>

            <ul v-if="channels.length" class="mt-5 flex flex-col gap-2.5">
              <li v-for="ch in channels" :key="ch.key" class="flex items-center gap-3 rounded-2xl bg-[#f5f5f7] px-4 py-3">
                <span class="grid h-9 w-9 flex-none place-items-center rounded-xl text-white" :class="ch.color">
                  <component :is="ch.icon" class="h-[18px] w-[18px]" aria-hidden="true" />
                </span>
                <div class="min-w-0 flex-1">
                  <div class="text-xs text-[#6e6e73]">{{ ch.label }}</div>
                  <div class="truncate text-[15px] font-semibold">{{ ch.display }}</div>
                </div>
                <a v-if="ch.href" :href="ch.href" target="_blank" rel="noopener noreferrer" class="flex-none rounded-full bg-[#1d1d1f] px-3.5 py-1.5 text-xs font-medium text-white hover:bg-black">{{ t('xinhuo.contact.go') }}</a>
                <button v-else type="button" class="flex-none rounded-full bg-[#1d1d1f] px-3.5 py-1.5 text-xs font-medium text-white hover:bg-black" @click="copy(ch)">
                  {{ copiedKey === ch.key ? t('xinhuo.contact.copied') : t('xinhuo.contact.copy') }}
                </button>
              </li>
            </ul>
            <p v-else class="mt-5 rounded-2xl bg-[#f5f5f7] px-4 py-6 text-center text-sm text-[#6e6e73]">{{ t('xinhuo.contact.empty') }}</p>
          </div>
        </div>
      </Transition>
    </Teleport>
  </section>
</template>

<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref, type Component } from 'vue'
import { useI18n } from 'vue-i18n'
import { ChevronRight, Headphones, MessageCircle, Send, ShieldCheck, Smartphone, Zap, BadgeCheck, X } from 'lucide-vue-next'
import { useAppStore } from '../../../stores/app'

const { t, te } = useI18n()
const appStore = useAppStore()

const brandName = computed(() => String(appStore.config?.brand?.site_name || '').trim() || '薪火数卡')

// 条目来自语言包 xinhuo.why.items.<n>，带 contact: true 的条目点击后弹出联系方式
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

interface Channel {
  key: string
  label: string
  display: string
  href: string
  icon: Component
  color: string
}

const isUrl = (v: string) => /^https?:\/\//i.test(v)

const channels = computed<Channel[]>(() => {
  const c = (appStore.config?.contact || {}) as Record<string, unknown>
  const val = (k: string) => String(c[k] ?? '').trim()
  const list: Channel[] = []
  const wechat = val('wechat')
  if (wechat) list.push({ key: 'wechat', label: t('xinhuo.contact.wechat'), display: wechat, href: isUrl(wechat) ? wechat : '', icon: MessageCircle, color: 'bg-[#07c160]' })
  const qq = val('qq')
  if (qq) list.push({ key: 'qq', label: 'QQ', display: qq, href: isUrl(qq) ? qq : '', icon: Smartphone, color: 'bg-[#12b7f5]' })
  const telegram = val('telegram')
  if (telegram) {
    const href = isUrl(telegram) ? telegram : `https://t.me/${telegram.replace(/^@/, '')}`
    list.push({ key: 'telegram', label: 'Telegram', display: telegram.replace(/^https?:\/\/(t|telegram)\.me\//i, '@'), href, icon: Send, color: 'bg-[#229ed9]' })
  }
  const whatsapp = val('whatsapp')
  if (whatsapp) list.push({ key: 'whatsapp', label: 'WhatsApp', display: whatsapp.replace(/^https?:\/\/wa\.me\//i, '+'), href: isUrl(whatsapp) ? whatsapp : '', icon: MessageCircle, color: 'bg-[#25d366]' })
  return list
})

const contactOpen = ref(false)
const copiedKey = ref('')
let copiedTimer: ReturnType<typeof setTimeout> | undefined
const copy = async (ch: Channel) => {
  try {
    await navigator.clipboard.writeText(ch.display)
  } catch {
    // 剪贴板不可用时退回到选中文本的旧方法
    const el = document.createElement('textarea')
    el.value = ch.display
    document.body.appendChild(el)
    el.select()
    try { document.execCommand('copy') } catch { /* 忽略 */ }
    el.remove()
  }
  copiedKey.value = ch.key
  clearTimeout(copiedTimer)
  copiedTimer = setTimeout(() => { copiedKey.value = '' }, 1800)
}

const onKey = (e: KeyboardEvent) => {
  if (e.key === 'Escape') contactOpen.value = false
}
onMounted(() => document.addEventListener('keydown', onKey))
onUnmounted(() => {
  document.removeEventListener('keydown', onKey)
  clearTimeout(copiedTimer)
})
</script>

<style scoped>
.xh-fade-enter-active,
.xh-fade-leave-active {
  transition: opacity 180ms ease;
}
.xh-fade-enter-from,
.xh-fade-leave-to {
  opacity: 0;
}
</style>
