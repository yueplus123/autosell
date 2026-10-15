<template>
  <div class="xinhuo-scope flex min-h-screen flex-col">
    <!-- 公告条：读取后台首页公告（标题），可当次会话关闭 -->
    <div v-if="announcementText && !announcementClosed" class="bg-[color:var(--xh-accent)] text-[13px] text-white">
      <div class="relative mx-auto flex w-full max-w-[1200px] items-center justify-center gap-3 px-12 py-2 text-center">
        <span class="flex-none rounded-full bg-white/20 px-2 py-0.5 text-xs font-medium">{{ t('xinhuo.announcementTag') }}</span>
        <span class="truncate">{{ announcementText }}</span>
        <RouterLink v-if="noticeEnabled" to="/notice" class="flex-none font-medium underline underline-offset-2 max-sm:hidden">{{ t('xinhuo.viewDetails') }}</RouterLink>
        <button
          type="button"
          class="absolute right-2 top-1/2 grid h-8 w-8 -translate-y-1/2 place-items-center rounded-full hover:bg-white/15"
          :aria-label="t('xinhuo.close')"
          @click="closeAnnouncement"
        ><X class="h-4 w-4" /></button>
      </div>
    </div>

    <!-- 黑色顶栏 -->
    <header class="sticky top-0 z-50 border-b border-[#1c1c1f] bg-black text-[#f5f5f7]">
      <div class="mx-auto flex h-14 w-full max-w-[1200px] items-center gap-6 px-4 sm:px-6">
        <RouterLink to="/" class="flex min-w-0 items-center gap-2.5 text-[17px] font-bold text-[#f5f5f7]" :title="brandName">
          <img v-if="brandLogo" :src="brandLogo" :alt="brandName" class="h-7 max-w-[140px] object-contain" />
          <span v-else class="grid h-7 w-7 flex-none place-items-center rounded-lg bg-[color:var(--xh-accent)]">
            <Flame class="h-4 w-4 text-white" />
          </span>
          <span class="truncate">{{ brandName }}</span>
        </RouterLink>

        <!-- 顶栏只保留「订单查询」 -->
        <nav class="ml-4 flex items-center gap-7 text-sm max-[900px]:hidden">
          <RouterLink :to="lookupTarget" class="whitespace-nowrap text-[#b8b8be] transition-colors hover:text-white" active-class="!text-white">{{ t('xinhuo.nav.orderLookup') }}</RouterLink>
        </nav>

        <div class="ml-auto flex items-center gap-1 sm:gap-2">
          <RouterLink to="/cart" class="relative grid h-10 w-10 place-items-center rounded-full text-[#d2d2d7] transition-colors hover:bg-white/10 hover:text-white" :aria-label="t('navbar.cart')">
            <ShoppingBag class="h-[19px] w-[19px]" />
            <span v-if="cartCount > 0" class="absolute right-0.5 top-0.5 grid h-[18px] min-w-[18px] place-items-center rounded-full bg-[color:var(--xh-accent)] px-1 text-[11px] font-bold text-white">{{ cartCount }}</span>
          </RouterLink>

          <div ref="langEl" class="relative max-[900px]:hidden">
            <button type="button" class="grid h-10 w-10 place-items-center rounded-full text-[#d2d2d7] transition-colors hover:bg-white/10 hover:text-white" :aria-label="t('navbar.selectLanguage')" @click="toggleLang">
              <Globe class="h-[19px] w-[19px]" />
            </button>
            <div v-if="langOpen" class="absolute right-0 top-[calc(100%+8px)] z-[60] flex min-w-[160px] flex-col gap-0.5 rounded-xl border border-[#2a2a2e] bg-[#161618] p-1.5 shadow-2xl">
              <button v-for="lang in languages" :key="lang.code" type="button" class="flex items-center justify-between rounded-lg px-3 py-2.5 text-left text-sm transition-colors hover:bg-white/10" :class="appStore.locale === lang.code ? 'text-white' : 'text-[#b8b8be]'" @click="changeLanguage(lang.code)">
                {{ lang.name }}
                <Check v-if="appStore.locale === lang.code" class="h-4 w-4" />
              </button>
            </div>
          </div>

          <template v-if="userAuthStore.isAuthenticated">
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
              <RouterLink :to="lookupTarget" class="rounded-lg px-3 py-2.5 text-sm text-[#d2d2d7] hover:bg-white/10 hover:text-white" @click="moreOpen = false">{{ t('xinhuo.nav.orderLookup') }}</RouterLink>
              <div class="my-1 h-px bg-[#2a2a2e]"></div>
              <RouterLink v-if="userAuthStore.isAuthenticated" to="/me" class="rounded-lg px-3 py-2.5 text-sm text-[#d2d2d7] hover:bg-white/10 hover:text-white" @click="moreOpen = false">{{ t('navbar.personalCenter') }}</RouterLink>
              <button v-if="userAuthStore.isAuthenticated" type="button" class="rounded-lg px-3 py-2.5 text-left text-sm text-[#d2d2d7] hover:bg-white/10 hover:text-white" @click="userAuthStore.logout(); moreOpen = false">{{ t('navbar.logout') }}</button>
              <template v-else>
                <RouterLink to="/auth/login" class="rounded-lg px-3 py-2.5 text-sm text-[#d2d2d7] hover:bg-white/10 hover:text-white" @click="moreOpen = false">{{ t('navbar.login') }}</RouterLink>
                <RouterLink to="/auth/register" class="rounded-lg px-3 py-2.5 text-sm text-[#d2d2d7] hover:bg-white/10 hover:text-white" @click="moreOpen = false">{{ t('xinhuo.nav.register') }}</RouterLink>
              </template>
              <div class="my-1 h-px bg-[#2a2a2e]"></div>
              <button v-for="lang in languages" :key="`ml-${lang.code}`" type="button" class="flex items-center justify-between rounded-lg px-3 py-2.5 text-left text-sm hover:bg-white/10" :class="appStore.locale === lang.code ? 'text-white' : 'text-[#b8b8be]'" @click="changeLanguage(lang.code); moreOpen = false">
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
      <div class="mx-auto grid w-full max-w-[1200px] grid-cols-2 gap-8 px-4 pb-8 pt-12 sm:px-6 lg:grid-cols-[1.6fr_repeat(3,1fr)]">
        <div class="col-span-2 lg:col-span-1">
          <div class="text-base font-bold text-[#f5f5f7]">{{ brandName }}</div>
          <p class="mt-3 max-w-[36ch] leading-relaxed">{{ brandDescription || t('xinhuo.footer.tagline') }}</p>
        </div>
        <div class="flex flex-col gap-2.5">
          <div class="font-medium text-[#f5f5f7]">{{ t('xinhuo.footer.shop') }}</div>
          <RouterLink to="/products" class="hover:text-white">{{ t('xinhuo.footer.allProducts') }}</RouterLink>
          <RouterLink v-if="noticeEnabled" to="/notice" class="hover:text-white">{{ t('xinhuo.footer.notice') }}</RouterLink>
        </div>
        <div class="flex flex-col gap-2.5">
          <div class="font-medium text-[#f5f5f7]">{{ t('xinhuo.footer.service') }}</div>
          <RouterLink :to="lookupTarget" class="hover:text-white">{{ t('xinhuo.footer.orderLookup') }}</RouterLink>
          <a v-if="contact?.telegram" :href="contact.telegram" target="_blank" rel="noopener noreferrer" class="hover:text-white">Telegram</a>
          <a v-if="contact?.whatsapp" :href="contact.whatsapp" target="_blank" rel="noopener noreferrer" class="hover:text-white">WhatsApp</a>
        </div>
        <div class="flex flex-col gap-2.5">
          <div class="font-medium text-[#f5f5f7]">{{ t('xinhuo.footer.company') }}</div>
          <RouterLink v-if="aboutEnabled" to="/about" class="hover:text-white">{{ t('xinhuo.footer.about') }}</RouterLink>
          <RouterLink v-if="blogEnabled" to="/blog" class="hover:text-white">{{ t('xinhuo.footer.blog') }}</RouterLink>
          <a v-for="link in footerLinks" :key="link.name" :href="link.url || 'javascript:void(0)'" :target="link.url ? '_blank' : undefined" rel="noopener noreferrer" class="hover:text-white">{{ link.name }}</a>
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
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { Check, Flame, Globe, LogOut, Menu, ShoppingBag, X } from 'lucide-vue-next'
import { useAppStore } from '../../../stores/app'
import { useCartStore } from '../../../stores/cart'
import { useUserAuthStore } from '../../../stores/userAuth'
import { useNavConfig } from '../../../composables/useNavConfig'
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

const { blogEnabled, noticeEnabled, aboutEnabled } = useNavConfig()

// 订单查询是独立页面（游客凭邮箱 + 订单密码查询，登录用户可在个人中心查看）
const lookupTarget = '/guest/orders'

const footerLinks = computed(() => {
  const links = appStore.config?.footer_links
  if (!Array.isArray(links)) return []
  return links
    .map((item: { name?: unknown; url?: unknown }) => ({
      name: typeof item?.name === 'string' ? item.name.trim() : getLocalizedText(item?.name as Record<string, string>, appStore.locale),
      url: String(item?.url || '').trim(),
    }))
    .filter((item) => item.name)
})

const contact = computed(() => appStore.config?.contact as { telegram?: string; whatsapp?: string } | undefined)
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
const toggleLang = () => { langOpen.value = !langOpen.value; moreOpen.value = false }
const toggleMore = () => { moreOpen.value = !moreOpen.value; langOpen.value = false }

const onDocClick = (e: MouseEvent) => {
  const target = e.target as Node
  if (langOpen.value && langEl.value && !langEl.value.contains(target)) langOpen.value = false
  if (moreOpen.value && moreEl.value && !moreEl.value.contains(target)) moreOpen.value = false
}

onMounted(() => {
  document.addEventListener('click', onDocClick)
  document.body.classList.add('xinhuo-tokens')
})
onUnmounted(() => {
  document.removeEventListener('click', onDocClick)
  document.body.classList.remove('xinhuo-tokens')
})
</script>
