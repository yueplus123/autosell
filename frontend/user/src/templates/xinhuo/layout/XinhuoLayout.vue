<template>
  <div class="xinhuo-scope flex min-h-screen flex-col overflow-x-clip">
    <!-- 公告条：读取后台首页公告（标题），可当次会话关闭 -->
    <div v-if="announcementText && !announcementClosed && !immersive" class="bg-[color:var(--xh-accent)] text-[13px] text-white">
      <div class="relative mx-auto flex w-full max-w-[1200px] items-center justify-center gap-3 px-12 py-2 text-center">
        <span class="flex-none rounded-full bg-white/20 px-2 py-0.5 text-xs font-medium">{{ t('xinhuo.announcementTag') }}</span>
        <span class="truncate">{{ announcementText }}</span>
        <button
          type="button"
          class="absolute right-2 top-1/2 grid h-8 w-8 -translate-y-1/2 place-items-center rounded-full hover:bg-white/15"
          :aria-label="t('xinhuo.close')"
          @click="closeAnnouncement"
        ><X class="h-4 w-4" /></button>
      </div>
    </div>

    <!-- 黑色顶栏；官网首页（Ethereal 风格）时变成浮在壁纸上的毛玻璃顶栏 -->
    <header
      class="z-50 text-[#f5f5f7]"
      :class="immersive ? 'fixed inset-x-0 top-0 px-3 sm:px-4' : 'sticky top-0 border-b border-[#1c1c1f] bg-black'"
    >
      <div
        class="mx-auto flex h-14 w-full max-w-[1200px] items-center gap-6 px-4 sm:px-6"
        :class="immersive ? 'rounded-b-2xl border border-t-0 border-white/15 bg-black/30 shadow-[0_8px_30px_rgba(0,0,0,0.18)] backdrop-blur-xl' : ''"
      >
        <RouterLink :to="showcase ? '/' : homePath" class="flex min-w-0 items-center gap-2.5 text-[17px] font-bold text-[#f5f5f7]" :title="brandName">
          <img v-if="brandLogo" :src="brandLogo" :alt="brandName" class="h-7 max-w-[140px] object-contain" />
          <span v-else class="grid h-7 w-7 flex-none place-items-center rounded-lg bg-[color:var(--xh-accent)]">
            <Flame class="h-4 w-4 text-white" />
          </span>
          <span class="truncate">{{ brandName }}</span>
        </RouterLink>

        <!-- 官网：首页 / AI 专区 / 二次元专区（两个专区是两个独立页面）；首页另有文章搜索 -->
        <nav v-if="showcase" class="ml-2 flex items-center gap-6 text-sm max-[900px]:hidden">
          <RouterLink v-for="item in showcaseNav" :key="item.to" :to="item.to" class="whitespace-nowrap transition-colors hover:text-white" :class="item.active ? 'font-medium text-white' : 'text-[#d2d2d7]'">{{ item.label }}</RouterLink>
        </nav>
        <form v-if="immersive" class="relative ml-auto max-[900px]:hidden" role="search" @submit.prevent="submitSearch">
          <Search class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-white/60" aria-hidden="true" />
          <input
            id="eth-search"
            v-model.trim="searchText"
            type="search"
            :placeholder="t('xinhuo.ethereal.search')"
            class="h-9 w-56 rounded-full border border-white/15 bg-white/10 pl-9 pr-3 text-sm text-white outline-none placeholder:text-white/50 focus:border-white/40"
          />
        </form>

        <!-- 顶栏只保留「订单查询」；官网展示模式不能下单，不显示 -->
        <nav v-if="!showcase" class="ml-4 flex items-center gap-7 text-sm max-[900px]:hidden">
          <RouterLink :to="lookupTarget" class="whitespace-nowrap text-[#b8b8be] transition-colors hover:text-white" active-class="!text-white">{{ t('xinhuo.nav.orderLookup') }}</RouterLink>
        </nav>

        <div class="flex items-center gap-1 sm:gap-2" :class="immersive ? 'max-[900px]:ml-auto' : 'ml-auto'">
          <!-- 官网首页：显示设置（主题色相 / 文章布局）和深浅色切换 -->
          <div v-if="immersive" ref="paletteEl" class="relative">
            <button type="button" class="grid h-10 w-10 place-items-center rounded-full text-[#d2d2d7] transition-colors hover:bg-white/10 hover:text-white" :class="{ 'bg-white/15 text-white': paletteOpen }" :aria-label="t('xinhuo.ethereal.display')" :aria-expanded="paletteOpen" @click="togglePalette">
              <Palette class="h-[19px] w-[19px]" />
            </button>
            <div v-if="paletteOpen" class="absolute right-0 top-[calc(100%+8px)] z-[60] w-[280px] rounded-2xl border border-white/15 bg-[#161618]/95 p-4 text-[#f5f5f7] shadow-2xl backdrop-blur-xl max-sm:-right-12">
              <div class="flex items-center justify-between">
                <span class="text-sm font-semibold">{{ t('xinhuo.ethereal.hue') }}</span>
                <span class="flex items-center gap-2">
                  <span class="rounded-md bg-white/10 px-2 py-0.5 text-xs tabular-nums">{{ ethHue }}</span>
                  <button type="button" class="grid h-7 w-7 place-items-center rounded-md text-[#b8b8be] hover:bg-white/10 hover:text-white" :aria-label="t('xinhuo.ethereal.reset')" :title="t('xinhuo.ethereal.reset')" @click="resetEtherealTheme()"><RotateCcw class="h-3.5 w-3.5" /></button>
                </span>
              </div>
              <input id="eth-hue" :value="ethHue" type="range" min="0" max="360" step="1" class="eth-hue-range mt-3 w-full" :aria-label="t('xinhuo.ethereal.hue')" @input="setHue" />
              <div class="mt-4 text-sm font-semibold">{{ t('xinhuo.ethereal.layout') }}</div>
              <div class="mt-2 grid grid-cols-2 gap-2">
                <button type="button" class="flex h-9 items-center justify-center gap-1.5 rounded-lg text-sm transition-colors" :class="ethLayout === 'list' ? 'bg-white text-black' : 'bg-white/10 text-[#d2d2d7] hover:bg-white/15'" @click="setLayout('list')"><List class="h-4 w-4" />{{ t('xinhuo.ethereal.layoutList') }}</button>
                <button type="button" class="flex h-9 items-center justify-center gap-1.5 rounded-lg text-sm transition-colors" :class="ethLayout === 'grid' ? 'bg-white text-black' : 'bg-white/10 text-[#d2d2d7] hover:bg-white/15'" @click="setLayout('grid')"><LayoutGrid class="h-4 w-4" />{{ t('xinhuo.ethereal.layoutGrid') }}</button>
              </div>
            </div>
          </div>
          <button v-if="immersive" type="button" class="grid h-10 w-10 place-items-center rounded-full text-[#d2d2d7] transition-colors hover:bg-white/10 hover:text-white" :aria-label="ethDark ? t('xinhuo.ethereal.light') : t('xinhuo.ethereal.dark')" :title="ethDark ? t('xinhuo.ethereal.light') : t('xinhuo.ethereal.dark')" @click="toggleDark">
            <Contrast class="h-[19px] w-[19px]" />
          </button>
          <!-- 主站展示模式不能下单，不显示购物车 -->
          <RouterLink v-if="!showcase" to="/cart" class="relative grid h-10 w-10 place-items-center rounded-full text-[#d2d2d7] transition-colors hover:bg-white/10 hover:text-white" :aria-label="t('navbar.cart')">
            <ShoppingBag class="h-[19px] w-[19px]" />
            <span v-if="cartCount > 0" class="absolute right-0.5 top-0.5 grid h-[18px] min-w-[18px] place-items-center rounded-full bg-[color:var(--xh-accent)] px-1 text-[11px] font-bold text-white">{{ cartCount }}</span>
          </RouterLink>

          <div ref="langEl" class="relative" :class="showcase ? '' : 'max-[900px]:hidden'">
            <button type="button" class="grid h-10 w-10 place-items-center rounded-full text-[#d2d2d7] transition-colors hover:bg-white/10 hover:text-white" :class="{ 'bg-white/15 text-white': langOpen }" :aria-label="t('navbar.selectLanguage')" @click="toggleLang">
              <Languages v-if="showcase" class="h-[19px] w-[19px]" />
              <Globe v-else class="h-[19px] w-[19px]" />
            </button>
            <div v-if="langOpen" class="absolute right-0 top-[calc(100%+8px)] z-[60] flex min-w-[160px] flex-col gap-0.5 rounded-xl border border-[#2a2a2e] bg-[#161618] p-1.5 shadow-2xl">
              <button v-for="lang in languages" :key="lang.code" type="button" class="flex items-center justify-between rounded-lg px-3 py-2.5 text-left text-sm transition-colors hover:bg-white/10" :class="appStore.locale === lang.code ? 'text-white' : 'text-[#b8b8be]'" @click="changeLanguage(lang.code)">
                {{ lang.name }}
                <Check v-if="appStore.locale === lang.code" class="h-4 w-4" />
              </button>
            </div>
          </div>

          <!-- 官网不需要登录注册 -->
          <template v-if="showcase"></template>
          <template v-else-if="userAuthStore.isAuthenticated">
            <RouterLink to="/me" class="rounded-full px-3 py-1.5 text-sm text-[#d2d2d7] transition-colors hover:text-white max-[900px]:hidden">{{ t('navbar.personalCenter') }}</RouterLink>
            <button type="button" class="grid h-10 w-10 place-items-center rounded-full text-[#d2d2d7] transition-colors hover:bg-white/10 hover:text-white max-[900px]:hidden" :aria-label="t('navbar.logout')" @click="userAuthStore.logout()"><LogOut class="h-[18px] w-[18px]" /></button>
          </template>
          <template v-else>
            <RouterLink to="/auth/login" class="px-2 text-sm text-[#d2d2d7] transition-colors hover:text-white max-[900px]:hidden">{{ t('navbar.login') }}</RouterLink>
            <RouterLink to="/auth/register" class="rounded-full bg-[#f5f5f7] px-4 py-1.5 text-sm font-medium text-black transition-colors hover:bg-white max-[900px]:hidden">{{ t('xinhuo.nav.register') }}</RouterLink>
          </template>

          <!-- 移动端菜单 -->
          <div ref="moreEl" class="relative hidden max-[900px]:block">
            <button type="button" class="grid h-11 w-11 place-items-center rounded-full text-[#f5f5f7]" :aria-label="moreOpen ? t('xinhuo.nav.close') : t('xinhuo.nav.menu')" :aria-expanded="moreOpen" @click="toggleMore">
              <X v-if="moreOpen" class="pointer-events-none h-[22px] w-[22px]" />
              <Menu v-else class="pointer-events-none h-[22px] w-[22px]" />
            </button>
            <div v-if="moreOpen" class="absolute right-0 top-[calc(100%+8px)] z-[60] flex w-[220px] flex-col gap-0.5 rounded-xl border border-[#2a2a2e] bg-[#161618] p-1.5 shadow-2xl">
              <template v-if="showcase">
                <RouterLink v-for="item in showcaseNav" :key="`m-${item.to}`" :to="item.to" class="rounded-lg px-3 py-2.5 text-sm hover:bg-white/10 hover:text-white" :class="item.active ? 'text-white' : 'text-[#d2d2d7]'" @click="moreOpen = false">{{ item.label }}</RouterLink>
              </template>
              <template v-if="!showcase">
                <RouterLink :to="lookupTarget" class="rounded-lg px-3 py-2.5 text-sm text-[#d2d2d7] hover:bg-white/10 hover:text-white" @click="moreOpen = false">{{ t('xinhuo.nav.orderLookup') }}</RouterLink>
                <div class="my-1 h-px bg-[#2a2a2e]"></div>
              </template>
              <template v-if="!showcase">
                <RouterLink v-if="userAuthStore.isAuthenticated" to="/me" class="rounded-lg px-3 py-2.5 text-sm text-[#d2d2d7] hover:bg-white/10 hover:text-white" @click="moreOpen = false">{{ t('navbar.personalCenter') }}</RouterLink>
                <button v-if="userAuthStore.isAuthenticated" type="button" class="rounded-lg px-3 py-2.5 text-left text-sm text-[#d2d2d7] hover:bg-white/10 hover:text-white" @click="userAuthStore.logout(); moreOpen = false">{{ t('navbar.logout') }}</button>
                <template v-else>
                  <RouterLink to="/auth/login" class="rounded-lg px-3 py-2.5 text-sm text-[#d2d2d7] hover:bg-white/10 hover:text-white" @click="moreOpen = false">{{ t('navbar.login') }}</RouterLink>
                  <RouterLink to="/auth/register" class="rounded-lg px-3 py-2.5 text-sm text-[#d2d2d7] hover:bg-white/10 hover:text-white" @click="moreOpen = false">{{ t('xinhuo.nav.register') }}</RouterLink>
                </template>
                <div class="my-1 h-px bg-[#2a2a2e]"></div>
              </template>
              <button v-for="lang in (showcase ? [] : languages)" :key="`ml-${lang.code}`" type="button" class="flex items-center justify-between rounded-lg px-3 py-2.5 text-left text-sm hover:bg-white/10" :class="appStore.locale === lang.code ? 'text-white' : 'text-[#b8b8be]'" @click="changeLanguage(lang.code); moreOpen = false">
                {{ lang.name }}
                <Check v-if="appStore.locale === lang.code" class="h-4 w-4" />
              </button>
            </div>
          </div>
        </div>
      </div>
    </header>

    <main class="flex-1">
      <slot />
    </main>

    <!-- 页脚 -->
    <footer class="bg-[#0b0b0c] text-[13px] text-[#b8b8be]">
      <!-- 页脚只保留品牌、订单查询与联系方式：不放首页 / 全部商品 / 公告 / 博客等入口，顾客留在当前专区 -->
      <div class="mx-auto grid w-full max-w-[1200px] gap-8 px-4 pb-8 pt-12 sm:grid-cols-[1.6fr_1fr] sm:px-6">
        <div>
          <div class="text-base font-bold text-[#f5f5f7]">{{ brandName }}</div>
          <p class="mt-3 max-w-[36ch] leading-relaxed">{{ brandDescription || t('xinhuo.footer.tagline') }}</p>
        </div>
        <div class="flex flex-col gap-2.5">
          <div class="font-medium text-[#f5f5f7]">{{ t('xinhuo.footer.service') }}</div>
          <RouterLink v-if="!showcase" :to="lookupTarget" class="hover:text-white">{{ t('xinhuo.footer.orderLookup') }}</RouterLink>
          <span v-if="contact?.wechat">{{ t('xinhuo.contact.wechat') }}：{{ contact.wechat }}</span>
          <span v-if="contact?.qq">QQ：{{ contact.qq }}</span>
          <a v-if="contact?.telegram" :href="telegramHref" target="_blank" rel="noopener noreferrer" class="hover:text-white">Telegram</a>
          <a v-if="contact?.whatsapp" :href="contact.whatsapp" target="_blank" rel="noopener noreferrer" class="hover:text-white">WhatsApp</a>
        </div>
      </div>
      <div class="mx-auto flex w-full max-w-[1200px] flex-wrap justify-between gap-3 border-t border-[#2a2a2e] px-4 pb-8 pt-5 sm:px-6">
        <span>© {{ year }} {{ brandName }}</span>
        <span class="flex flex-wrap gap-x-2">
          <RouterLink to="/terms" class="hover:text-white">{{ t('xinhuo.footer.terms') }}</RouterLink>
          <span aria-hidden="true">·</span>
          <RouterLink to="/privacy" class="hover:text-white">{{ t('xinhuo.footer.privacy') }}</RouterLink>
        </span>
      </div>
    </footer>

    <!-- 全站共用的联系方式弹窗（「快速响应」「联系渠道合作」都打开它） -->
    <XinhuoContactModal />
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { Check, Contrast, Flame, Globe, Languages, LayoutGrid, List, LogOut, Menu, Palette, RotateCcw, Search, ShoppingBag, X } from 'lucide-vue-next'
import { ethDark, ethHue, ethLayout, resetEtherealTheme, type EtherealLayout } from '../etherealTheme'
import { useRoute, useRouter } from 'vue-router'
import { businessZonePath } from '../../../utils/businessLines'
import { useAppStore } from '../../../stores/app'
import { useCartStore } from '../../../stores/cart'
import { useUserAuthStore } from '../../../stores/userAuth'
import { useStorefrontLinks } from '../../../composables/useStorefrontLinks'
import { isMainShowcase } from '../../../utils/storefrontZone'
import XinhuoContactModal from '../components/XinhuoContactModal.vue'
import { getImageUrl } from '../../../utils/image'
import { getLocalizedText } from '../../../utils/resellerSiteConfig'
import '../styles/xinhuo.css'

const { t } = useI18n()
const appStore = useAppStore()
const cartStore = useCartStore()
const userAuthStore = useUserAuthStore()

const year = new Date().getFullYear()
const langOpen = ref(false)
const moreOpen = ref(false)
const langEl = ref<HTMLElement | null>(null)
const paletteOpen = ref(false)
const paletteEl = ref<HTMLElement | null>(null)
const moreEl = ref<HTMLElement | null>(null)

const brandName = computed(() => String(appStore.config?.brand?.site_name || '').trim() || '薪火数卡')
const brandLogo = computed(() => {
  const raw = String(appStore.config?.brand?.site_logo || '').trim()
  return raw ? getImageUrl(raw) : ''
})
const brandDescription = computed(() => {
  const desc = appStore.config?.brand?.site_description
  if (desc && typeof desc === 'object') return getLocalizedText(desc as Record<string, string>, appStore.locale).trim()
  return ''
})

// 顶栏 logo：在专区里（或分站只代理一个专区）时回到专区，而不是总首页
const { homePath } = useStorefrontLinks()
const showcase = computed(() => isMainShowcase(appStore.config as any))

// 官网首页使用 Ethereal 风格：毛玻璃顶栏浮在壁纸上，带专区入口和文章搜索
const route = useRoute()
const router = useRouter()
const immersive = computed(() => showcase.value && route.name === 'home')
const aiZonePath = businessZonePath('ai')
const gameZonePath = businessZonePath('game')
const showcaseNav = computed(() => {
  const key = route.name === 'business-zone' ? String(route.params.key || '') : ''
  return [
    { to: '/', label: t('xinhuo.ethereal.navHome'), active: route.name === 'home' },
    { to: aiZonePath, label: t('zone.ai.badge'), active: key === 'ai' },
    { to: gameZonePath, label: t('zone.game.badge'), active: key === 'game' },
  ]
})
const searchText = ref(typeof route.query.q === 'string' ? route.query.q : '')
const submitSearch = () => {
  const q = searchText.value.trim()
  void router.push({ path: '/', query: q ? { q } : {} })
}

// 订单查询是独立页面（游客凭邮箱 + 订单密码查询，登录用户可在个人中心查看）
const lookupTarget = '/guest/orders'

const contact = computed(() => appStore.config?.contact as { telegram?: string; whatsapp?: string; wechat?: string; qq?: string } | undefined)
const telegramHref = computed(() => {
  const raw = String(contact.value?.telegram || '').trim()
  return /^https?:\/\//i.test(raw) ? raw : `https://t.me/${raw.replace(/^@/, '')}`
})
const cartCount = computed(() => cartStore.totalItems)

// ==================== 公告条 ====================
const ANNOUNCEMENT_BAR_KEY = 'xinhuo_announcement_bar_closed'
const announcement = computed(() => appStore.config?.announcement as { title?: Record<string, string>; version?: string } | undefined)
const announcementText = computed(() => {
  const title = announcement.value?.title
  return title ? getLocalizedText(title, appStore.locale).trim() : ''
})
const readClosedVersion = () => {
  try { return sessionStorage.getItem(ANNOUNCEMENT_BAR_KEY) || '' } catch { return '' }
}
const closedVersion = ref(readClosedVersion())
const announcementClosed = computed(() => closedVersion.value !== '' && closedVersion.value === String(announcement.value?.version || ''))
const closeAnnouncement = () => {
  const version = String(announcement.value?.version || 'closed')
  closedVersion.value = version
  try { sessionStorage.setItem(ANNOUNCEMENT_BAR_KEY, version) } catch { /* 存储不可用时仅本次隐藏 */ }
}

// ==================== 语言 / 菜单 ====================
const languages = [
  { code: 'zh-CN', name: '简体中文' },
  { code: 'zh-TW', name: '繁體中文' },
  { code: 'en-US', name: 'English' },
]
const changeLanguage = (code: string) => {
  appStore.setLocale(code)
  langOpen.value = false
}
const toggleLang = () => { langOpen.value = !langOpen.value; moreOpen.value = false; paletteOpen.value = false }
const toggleMore = () => { moreOpen.value = !moreOpen.value; langOpen.value = false; paletteOpen.value = false }
const setHue = (e: Event) => { ethHue.value = Number((e.target as HTMLInputElement).value) || 0 }
const setLayout = (layout: EtherealLayout) => { ethLayout.value = layout }
const toggleDark = () => { ethDark.value = !ethDark.value }
const togglePalette = () => { paletteOpen.value = !paletteOpen.value; langOpen.value = false; moreOpen.value = false }

const onDocClick = (e: MouseEvent) => {
  const target = e.target as Node
  if (langOpen.value && langEl.value && !langEl.value.contains(target)) langOpen.value = false
  if (moreOpen.value && moreEl.value && !moreEl.value.contains(target)) moreOpen.value = false
  if (paletteOpen.value && paletteEl.value && !paletteEl.value.contains(target)) paletteOpen.value = false
}

onMounted(() => {
  document.addEventListener('click', onDocClick)
  document.body.classList.add('xinhuo-tokens')
  // 薪火模板只有浅色设计：告诉手机浏览器不要强制深色（否则白卡片被涂黑、标题看不清）
  document.documentElement.style.colorScheme = 'only light'
})
onUnmounted(() => {
  document.removeEventListener('click', onDocClick)
  document.body.classList.remove('xinhuo-tokens')
  document.documentElement.style.colorScheme = ''
})
</script>

<style scoped>
/* 主题色相滑块：彩虹色轨道 */
.eth-hue-range {
  -webkit-appearance: none;
  appearance: none;
  height: 14px;
  border-radius: 999px;
  background: linear-gradient(90deg, hsl(0 85% 65%), hsl(60 85% 60%), hsl(120 70% 55%), hsl(180 75% 55%), hsl(240 80% 65%), hsl(300 75% 65%), hsl(360 85% 65%));
  outline: none;
}
.eth-hue-range::-webkit-slider-thumb {
  -webkit-appearance: none;
  width: 10px;
  height: 22px;
  border-radius: 4px;
  background: #fff;
  box-shadow: 0 1px 6px rgba(0, 0, 0, 0.35);
  cursor: pointer;
}
.eth-hue-range::-moz-range-thumb {
  width: 10px;
  height: 22px;
  border: 0;
  border-radius: 4px;
  background: #fff;
  box-shadow: 0 1px 6px rgba(0, 0, 0, 0.35);
  cursor: pointer;
}
</style>
