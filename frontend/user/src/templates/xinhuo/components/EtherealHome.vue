<template>
  <!-- 官网首页（参考 Ethereal 博客主题布局）：全屏壁纸首屏 + 动态条 + 三栏（品牌卡片 / 文章列表 / 公告与统计） -->
  <div class="eth-page">
    <EtherealSakura />

    <!-- ==================== 全屏壁纸首屏 ==================== -->
    <section class="eth-hero">
      <div
        v-for="(bg, i) in heroBackgrounds"
        :key="bg.key"
        class="eth-hero-bg"
        :class="{ 'is-active': i === activeBg }"
        :style="{ '--bg': `url(${bg.image})`, '--bg-mobile': `url(${bg.mobile || bg.image})` }"
        aria-hidden="true"
      ></div>
      <div class="eth-hero-shade" aria-hidden="true"></div>

      <div class="relative z-[1] mx-auto flex h-full max-w-[1200px] flex-col items-center justify-center px-4 text-center text-white">
        <h1 class="eth-hero-title">{{ heroTitle }}</h1>
        <p class="eth-hero-sub">
          <span>{{ typedSubtitle }}</span><span class="eth-caret" aria-hidden="true">|</span>
        </p>
        <div class="mt-7 flex flex-wrap items-center justify-center gap-3">
          <RouterLink :to="aiZonePath" class="eth-hero-btn">
            <Sparkles class="h-4 w-4" aria-hidden="true" />
            {{ t('home.business.ai.cta') }}
          </RouterLink>
          <RouterLink :to="gameZonePath" class="eth-hero-btn">
            <Gamepad2 class="h-4 w-4" aria-hidden="true" />
            {{ t('home.business.game.cta') }}
          </RouterLink>
          <button type="button" class="eth-hero-icon" :aria-label="t('xinhuo.ethereal.contact')" @click="openContactModal()">
            <MessageCircle class="h-[18px] w-[18px]" aria-hidden="true" />
          </button>
          <button type="button" class="eth-hero-icon" :aria-label="t('xinhuo.ethereal.toArticles')" @click="scrollToArticles">
            <ChevronDown class="h-[18px] w-[18px]" aria-hidden="true" />
          </button>
        </div>
      </div>

      <!-- 底部波浪 -->
      <svg class="eth-wave" viewBox="0 0 1440 120" preserveAspectRatio="none" aria-hidden="true">
        <path d="M0 64 C 240 110 480 20 720 56 C 960 92 1200 30 1440 60 L1440 120 L0 120 Z" />
      </svg>
    </section>

    <div class="eth-container">
      <!-- ==================== 动态条（读取后台「公告」类文章） ==================== -->
      <div v-if="moments.length" class="eth-card eth-moments">
        <span class="eth-moments-tag"><Zap class="h-3.5 w-3.5" aria-hidden="true" />{{ t('xinhuo.ethereal.moments') }}</span>
        <div class="eth-marquee">
          <div class="eth-marquee-track" :style="{ animationDuration: `${Math.max(20, moments.length * 9)}s` }">
            <RouterLink v-for="(m, i) in [...moments, ...moments]" :key="`${m.id}-${i}`" to="/notice" class="eth-moment">
              <b>{{ brandName }}</b>：{{ getLocalizedText(m.title) }}
            </RouterLink>
          </div>
        </div>
        <RouterLink to="/notice" class="eth-moments-more">{{ t('xinhuo.ethereal.more') }}<ChevronRight class="h-3.5 w-3.5" aria-hidden="true" /></RouterLink>
      </div>

      <!-- ==================== 三栏 ==================== -->
      <div id="eth-articles" class="eth-grid">
        <!-- 左栏：品牌卡片 -->
        <aside class="eth-left">
          <div class="eth-card eth-profile">
            <div class="eth-profile-cover">
              <img v-if="brandLogo" :src="brandLogo" :alt="brandName" />
              <Flame v-else class="h-16 w-16 text-white" aria-hidden="true" />
            </div>
            <div class="px-5 pb-5 pt-4 text-center">
              <div class="text-lg font-bold">{{ brandName }}</div>
              <span class="eth-profile-bar" aria-hidden="true"></span>
              <p class="mt-2 text-[13px] leading-relaxed text-[color:var(--eth-muted)]">{{ brandDescription }}</p>
              <div class="mt-4 grid grid-cols-2 gap-2">
                <RouterLink :to="aiZonePath" class="eth-chip-btn">{{ t('zone.ai.badge') }}</RouterLink>
                <RouterLink :to="gameZonePath" class="eth-chip-btn">{{ t('zone.game.badge') }}</RouterLink>
              </div>
              <div v-if="contactChannels.length" class="mt-4 flex justify-center gap-2">
                <button
                  v-for="ch in contactChannels"
                  :key="ch.key"
                  type="button"
                  class="eth-round-btn"
                  :aria-label="ch.label"
                  :title="ch.label"
                  @click="openContactModal()"
                >
                  <component :is="ch.icon" class="h-4 w-4" aria-hidden="true" />
                </button>
              </div>
            </div>
          </div>
        </aside>

        <!-- 中栏：分类 + 文章 -->
        <main class="eth-main">
          <div class="eth-card eth-cats">
            <button type="button" class="eth-cat eth-cat--home" :class="{ 'is-active': activeCategory === 0 }" :aria-label="t('xinhuo.ethereal.all')" @click="selectCategory(0)">
              <Home class="h-4 w-4" aria-hidden="true" />
            </button>
            <button
              v-for="cat in categoryChips"
              :key="cat.id"
              type="button"
              class="eth-cat"
              :class="{ 'is-active': activeCategory === cat.id }"
              @click="selectCategory(cat.id)"
            >
              {{ cat.name }}<span class="eth-cat-count">{{ cat.count }}</span>
            </button>
          </div>

          <p v-if="keyword" class="px-1 text-sm text-[color:var(--eth-muted)]">
            {{ t('xinhuo.ethereal.searchResult', { q: keyword, n: filteredPosts.length }) }}
            <button type="button" class="ml-2 text-[color:var(--eth-accent)] hover:underline" @click="clearSearch">{{ t('xinhuo.ethereal.clearSearch') }}</button>
          </p>

          <template v-if="pagedPosts.length">
            <RouterLink
              v-for="(post, i) in pagedPosts"
              :key="post.id"
              :to="`/blog/${post.slug}`"
              class="eth-card eth-post"
              :style="{ animationDelay: `${i * 60}ms` }"
            >
              <div class="eth-post-body">
                <h2 class="eth-post-title">{{ getLocalizedText(post.title) }}</h2>
                <div class="eth-post-meta">
                  <span v-if="categoryName(post)" class="eth-meta-chip"><BookOpen class="h-3.5 w-3.5" aria-hidden="true" />{{ categoryName(post) }}</span>
                  <span class="eth-meta-item"><CalendarDays class="h-3.5 w-3.5" aria-hidden="true" />{{ formatDate(post.published_at) }}</span>
                  <span class="eth-meta-item"><FileText class="h-3.5 w-3.5" aria-hidden="true" />{{ t('xinhuo.ethereal.words', { n: wordCount(post) }) }}</span>
                </div>
                <p v-if="getLocalizedText(post.summary)" class="eth-post-summary">{{ getLocalizedText(post.summary) }}</p>
              </div>
              <div v-if="post.thumbnail" class="eth-post-cover">
                <img :src="getImageUrl(post.thumbnail)" :alt="getLocalizedText(post.title)" loading="lazy" />
              </div>
            </RouterLink>

            <nav v-if="totalPages > 1" class="eth-pager" :aria-label="t('xinhuo.ethereal.pager')">
              <button type="button" class="eth-page-btn" :disabled="page <= 1" :aria-label="t('xinhuo.ethereal.prev')" @click="goPage(page - 1)"><ChevronLeft class="h-4 w-4" /></button>
              <button v-for="n in totalPages" :key="n" type="button" class="eth-page-btn" :class="{ 'is-active': n === page }" @click="goPage(n)">{{ n }}</button>
              <button type="button" class="eth-page-btn" :disabled="page >= totalPages" :aria-label="t('xinhuo.ethereal.next')" @click="goPage(page + 1)"><ChevronRight class="h-4 w-4" /></button>
            </nav>
          </template>
          <div v-else-if="!loading" class="eth-card px-6 py-14 text-center text-sm text-[color:var(--eth-muted)]">
            {{ keyword || activeCategory ? t('xinhuo.ethereal.noMatch') : t('xinhuo.posts.empty') }}
          </div>
          <div v-else class="flex flex-col gap-4">
            <div v-for="n in 3" :key="n" class="eth-card h-[180px] animate-pulse"></div>
          </div>
        </main>

        <!-- 右栏：公告 + 站点统计 -->
        <aside class="eth-right">
          <div v-if="announcementTitle || announcementText" class="eth-card eth-widget">
            <h3 class="eth-widget-title">{{ t('xinhuo.ethereal.notice') }}</h3>
            <p v-if="announcementTitle" class="text-sm font-medium">{{ announcementTitle }}</p>
            <p v-if="announcementText" class="mt-1 line-clamp-4 text-[13px] leading-relaxed text-[color:var(--eth-muted)]">{{ announcementText }}</p>
            <RouterLink v-if="aboutEnabled" to="/about" class="eth-chip-btn mt-3 inline-flex">{{ t('xinhuo.ethereal.about') }}</RouterLink>
          </div>

          <div class="eth-card eth-widget">
            <h3 class="eth-widget-title">{{ t('xinhuo.ethereal.stats') }}</h3>
            <ul class="flex flex-col gap-3 text-sm">
              <li class="eth-stat"><span><FileText class="h-4 w-4" aria-hidden="true" />{{ t('xinhuo.ethereal.statPosts') }}</span><b>{{ posts.length }}</b></li>
              <li class="eth-stat"><span><Tags class="h-4 w-4" aria-hidden="true" />{{ t('xinhuo.ethereal.statCategories') }}</span><b>{{ categoryChips.length }}</b></li>
              <li class="eth-stat"><span><PenLine class="h-4 w-4" aria-hidden="true" />{{ t('xinhuo.ethereal.statWords') }}</span><b>{{ formatCompact(totalWords) }}</b></li>
              <li v-if="lastUpdated" class="eth-stat"><span><CalendarDays class="h-4 w-4" aria-hidden="true" />{{ t('xinhuo.ethereal.statUpdated') }}</span><b>{{ lastUpdated }}</b></li>
            </ul>
          </div>
        </aside>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref, watch, type Component } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import {
  BookOpen, CalendarDays, ChevronDown, ChevronLeft, ChevronRight, FileText, Flame, Gamepad2, Home,
  MessageCircle, PenLine, Send, Smartphone, Sparkles, Tags, Zap,
} from 'lucide-vue-next'
import { bannerAPI, postAPI, postCategoryAPI } from '../../../api'
import { getImageUrl } from '../../../utils/image'
import { businessZonePath } from '../../../utils/businessLines'
import { useLocalized } from '../../../composables/useProduct'
import { useNavConfig } from '../../../composables/useNavConfig'
import { useAppStore } from '../../../stores/app'
import { openContactModal } from '../contactModal'
import EtherealSakura from './EtherealSakura.vue'

const { t } = useI18n()
const route = useRoute()
const router = useRouter()
const appStore = useAppStore()
const { getLocalizedText } = useLocalized()
const { aboutEnabled } = useNavConfig()

const aiZonePath = businessZonePath('ai')
const gameZonePath = businessZonePath('game')

// ==================== 品牌 ====================
const brandName = computed(() => String(appStore.config?.brand?.site_name || '').trim() || '薪火数卡')
const brandLogo = computed(() => {
  const raw = String(appStore.config?.brand?.site_logo || '').trim()
  return raw ? getImageUrl(raw) : ''
})
const brandDescription = computed(() => {
  const desc = appStore.config?.brand?.site_description
  const text = desc && typeof desc === 'object' ? getLocalizedText(desc as Record<string, string>).trim() : ''
  return text || t('xinhuo.footer.tagline')
})

const contactChannels = computed(() => {
  const c = (appStore.config?.contact || {}) as Record<string, unknown>
  const list: { key: string; label: string; icon: Component }[] = []
  if (String(c.wechat || '').trim()) list.push({ key: 'wechat', label: t('xinhuo.contact.wechat'), icon: MessageCircle })
  if (String(c.qq || '').trim()) list.push({ key: 'qq', label: 'QQ', icon: Smartphone })
  if (String(c.telegram || '').trim()) list.push({ key: 'telegram', label: 'Telegram', icon: Send })
  return list
})

// ==================== 首屏：壁纸读取后台「轮播图」，没有时用默认壁纸 ====================
const DEFAULT_WALLPAPER = '/xinhuo/ethereal/sky.svg'
const banners = ref<any[]>([])
const heroBackgrounds = computed(() => {
  const list = banners.value
    .filter((b) => String(b?.image || '').trim())
    .map((b) => ({ key: `b-${b.id}`, image: getImageUrl(b.image), mobile: b.mobile_image ? getImageUrl(b.mobile_image) : '' }))
  return list.length ? list : [{ key: 'default', image: DEFAULT_WALLPAPER, mobile: '' }]
})
const firstBanner = computed(() => banners.value[0])
const heroTitle = computed(() => getLocalizedText(firstBanner.value?.title).trim() || t('xinhuo.hero.title'))
const heroSubtitleFull = computed(() => getLocalizedText(firstBanner.value?.subtitle).trim() || t('xinhuo.hero.subtitle'))

const activeBg = ref(0)
let bgTimer: ReturnType<typeof setInterval> | undefined

// 副标题打字机效果
const typedSubtitle = ref('')
let typeTimer: ReturnType<typeof setTimeout> | undefined
const reducedMotion = () => typeof window.matchMedia === 'function' && window.matchMedia('(prefers-reduced-motion: reduce)').matches
const startTyping = () => {
  clearTimeout(typeTimer)
  const full = heroSubtitleFull.value
  if (reducedMotion()) {
    typedSubtitle.value = full
    return
  }
  let i = 0
  const step = () => {
    typedSubtitle.value = full.slice(0, i)
    if (i < full.length) {
      i++
      typeTimer = setTimeout(step, 90)
    }
  }
  step()
}
watch(heroSubtitleFull, startTyping)

const scrollToArticles = () => {
  document.getElementById('eth-articles')?.scrollIntoView({ behavior: 'smooth', block: 'start' })
}

// ==================== 动态条 / 公告 ====================
const moments = ref<any[]>([])
const stripHtml = (html: string) => html.replace(/<[^>]*>/g, ' ').replace(/&nbsp;/g, ' ').replace(/\s+/g, ' ').trim()
const announcement = computed(() => appStore.config?.announcement as { title?: Record<string, string>; content?: Record<string, string> } | undefined)
const announcementTitle = computed(() => (announcement.value?.title ? getLocalizedText(announcement.value.title).trim() : ''))
const announcementText = computed(() => (announcement.value?.content ? stripHtml(getLocalizedText(announcement.value.content)) : ''))

// ==================== 文章 / 分类 ====================
const PAGE_SIZE = 8
const loading = ref(true)
const posts = ref<any[]>([])
const categories = ref<any[]>([])
const activeCategory = ref(0)
const page = ref(1)
const keyword = computed(() => (typeof route.query.q === 'string' ? route.query.q.trim() : ''))

const categoryMap = computed(() => {
  const map = new Map<number, string>()
  for (const c of categories.value) map.set(Number(c.id), getLocalizedText(c.name))
  return map
})
const categoryName = (post: any) => (post?.category_id ? categoryMap.value.get(Number(post.category_id)) || '' : '')
const categoryChips = computed(() => {
  const counts = new Map<number, number>()
  for (const p of posts.value) {
    const id = Number(p?.category_id || 0)
    if (id) counts.set(id, (counts.get(id) || 0) + 1)
  }
  return categories.value
    .map((c) => ({ id: Number(c.id), name: getLocalizedText(c.name), count: counts.get(Number(c.id)) || 0 }))
    .filter((c) => c.count > 0)
})

const filteredPosts = computed(() => {
  const q = keyword.value.toLowerCase()
  return posts.value.filter((p) => {
    if (activeCategory.value && Number(p?.category_id || 0) !== activeCategory.value) return false
    if (!q) return true
    const hay = `${getLocalizedText(p.title)} ${getLocalizedText(p.summary)}`.toLowerCase()
    return hay.includes(q)
  })
})
const totalPages = computed(() => Math.max(1, Math.ceil(filteredPosts.value.length / PAGE_SIZE)))
const pagedPosts = computed(() => filteredPosts.value.slice((page.value - 1) * PAGE_SIZE, page.value * PAGE_SIZE))

const selectCategory = (id: number) => {
  activeCategory.value = id
  page.value = 1
}
const goPage = (n: number) => {
  page.value = Math.min(Math.max(1, n), totalPages.value)
  scrollToArticles()
}
const clearSearch = () => {
  const query = { ...route.query }
  delete query.q
  void router.replace({ query })
}
watch(keyword, () => {
  page.value = 1
  if (keyword.value) scrollToArticles()
})

const wordCount = (post: any) => stripHtml(getLocalizedText(post?.content)).replace(/\s/g, '').length
const totalWords = computed(() => posts.value.reduce((sum, p) => sum + wordCount(p), 0))
const formatCompact = (n: number) => {
  try {
    return new Intl.NumberFormat(appStore.locale || 'zh-CN', { notation: 'compact', maximumFractionDigits: 1 }).format(n)
  } catch {
    return String(n)
  }
}
const formatDate = (value: string) => (value ? new Date(value).toLocaleDateString(appStore.locale || undefined) : '')
const lastUpdated = computed(() => formatDate(posts.value[0]?.published_at || ''))

// ==================== 加载 ====================
onMounted(async () => {
  startTyping()
  const [postRes, catRes, noticeRes, bannerRes] = await Promise.allSettled([
    postAPI.list({ page: 1, page_size: 200, type: 'blog' }),
    postCategoryAPI.list(),
    postAPI.list({ page: 1, page_size: 10, type: 'notice' }),
    bannerAPI.list({ position: 'home_hero', limit: 5 }),
  ])
  if (postRes.status === 'fulfilled') posts.value = postRes.value.data.data || []
  if (catRes.status === 'fulfilled') categories.value = catRes.value.data.data || []
  if (noticeRes.status === 'fulfilled') moments.value = noticeRes.value.data.data || []
  if (bannerRes.status === 'fulfilled') banners.value = bannerRes.value.data.data || []
  loading.value = false

  if (heroBackgrounds.value.length > 1 && !reducedMotion()) {
    bgTimer = setInterval(() => {
      activeBg.value = (activeBg.value + 1) % heroBackgrounds.value.length
    }, 7000)
  }
  if (keyword.value) scrollToArticles()
})

onUnmounted(() => {
  clearTimeout(typeTimer)
  clearInterval(bgTimer)
})
</script>

<style scoped>
.eth-page {
  --eth-bg: #eef1f8;
  --eth-card: #ffffff;
  --eth-ink: #1f2433;
  --eth-muted: #6b7288;
  --eth-line: #e4e8f2;
  --eth-accent: #4f7cf0;
  --eth-accent-soft: #e8efff;
  background: var(--eth-bg);
  color: var(--eth-ink);
  padding-bottom: 56px;
}

/* ---------- 首屏 ---------- */
.eth-hero {
  position: relative;
  height: min(100svh, 920px);
  min-height: 560px;
  overflow: hidden;
  background: #2e2753;
}
.eth-hero-bg {
  position: absolute;
  inset: 0;
  background-image: var(--bg);
  background-size: cover;
  background-position: center;
  opacity: 0;
  transform: scale(1.04);
  transition: opacity 1.2s ease, transform 8s ease;
}
.eth-hero-bg.is-active {
  opacity: 1;
  transform: scale(1);
}
@media (max-width: 640px) {
  .eth-hero-bg { background-image: var(--bg-mobile); }
}
.eth-hero-shade {
  position: absolute;
  inset: 0;
  background: linear-gradient(180deg, rgba(10, 12, 30, 0.25) 0%, rgba(10, 12, 30, 0.15) 45%, rgba(10, 12, 30, 0.35) 100%);
}
.eth-hero-title {
  font-size: clamp(34px, 6vw, 72px);
  font-weight: 900;
  line-height: 1.12;
  letter-spacing: -0.01em;
  text-shadow: 0 4px 24px rgba(0, 0, 0, 0.35);
  text-wrap: balance;
}
.eth-hero-sub {
  margin-top: 14px;
  min-height: 1.6em;
  font-size: clamp(15px, 1.8vw, 22px);
  color: rgba(255, 255, 255, 0.92);
  text-shadow: 0 2px 12px rgba(0, 0, 0, 0.35);
}
.eth-caret {
  margin-left: 2px;
  animation: eth-blink 1s steps(1) infinite;
}
@keyframes eth-blink { 50% { opacity: 0; } }
.eth-hero-btn,
.eth-hero-icon {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  height: 44px;
  border-radius: 999px;
  border: 1px solid rgba(255, 255, 255, 0.35);
  background: rgba(255, 255, 255, 0.18);
  backdrop-filter: blur(12px);
  -webkit-backdrop-filter: blur(12px);
  color: #fff;
  font-size: 15px;
  font-weight: 600;
  transition: background 0.2s ease, transform 0.2s ease;
}
.eth-hero-btn { padding: 0 20px; }
.eth-hero-icon { width: 44px; }
.eth-hero-btn:hover,
.eth-hero-icon:hover {
  background: rgba(255, 255, 255, 0.3);
  transform: translateY(-2px);
}
.eth-wave {
  position: absolute;
  left: 0;
  right: 0;
  bottom: -1px;
  width: 100%;
  height: 70px;
  fill: var(--eth-bg);
}

/* ---------- 布局 ---------- */
.eth-container {
  position: relative;
  z-index: 2;
  max-width: 1240px;
  margin: -28px auto 0;
  padding: 0 16px;
  display: flex;
  flex-direction: column;
  gap: 16px;
}
.eth-card {
  background: var(--eth-card);
  border: 1px solid var(--eth-line);
  border-radius: 16px;
  box-shadow: 0 6px 24px rgba(31, 36, 51, 0.05);
}
.eth-grid {
  display: grid;
  grid-template-columns: 260px minmax(0, 1fr) 280px;
  gap: 16px;
  align-items: start;
  scroll-margin-top: 80px;
}
.eth-left,
.eth-right {
  position: sticky;
  top: 80px;
  display: flex;
  flex-direction: column;
  gap: 16px;
  min-width: 0;
}
.eth-main {
  display: flex;
  flex-direction: column;
  gap: 16px;
  min-width: 0;
}
@media (max-width: 1100px) {
  .eth-grid { grid-template-columns: minmax(0, 1fr) 280px; }
  .eth-left { display: none; }
}
@media (max-width: 820px) {
  .eth-grid { grid-template-columns: minmax(0, 1fr); }
  .eth-right { position: static; }
}

/* ---------- 动态条 ---------- */
.eth-moments {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 10px 12px;
  min-width: 0;
}
.eth-moments-tag {
  flex: none;
  display: inline-flex;
  align-items: center;
  gap: 4px;
  padding: 4px 10px;
  border-radius: 8px;
  background: var(--eth-accent-soft);
  color: var(--eth-accent);
  font-size: 13px;
  font-weight: 600;
}
.eth-marquee {
  flex: 1;
  min-width: 0;
  overflow: hidden;
  mask-image: linear-gradient(90deg, transparent, #000 6%, #000 94%, transparent);
}
.eth-marquee-track {
  display: inline-flex;
  gap: 32px;
  white-space: nowrap;
  animation: eth-marquee linear infinite;
}
.eth-marquee:hover .eth-marquee-track { animation-play-state: paused; }
@keyframes eth-marquee { from { transform: translateX(0); } to { transform: translateX(-50%); } }
.eth-moment {
  font-size: 13px;
  color: var(--eth-muted);
}
.eth-moment b { color: var(--eth-accent); font-weight: 600; }
.eth-moment:hover { color: var(--eth-ink); }
.eth-moments-more {
  flex: none;
  display: inline-flex;
  align-items: center;
  gap: 2px;
  padding: 4px 10px;
  border-radius: 8px;
  background: var(--eth-accent-soft);
  color: var(--eth-accent);
  font-size: 13px;
}

/* ---------- 左栏 ---------- */
.eth-profile { overflow: hidden; }
.eth-profile-cover {
  margin: 10px 10px 0;
  aspect-ratio: 1 / 1;
  border-radius: 12px;
  display: grid;
  place-items: center;
  background: linear-gradient(135deg, #6d8cf5, #a78bfa 55%, #f9a8d4);
  overflow: hidden;
}
.eth-profile-cover img {
  width: 62%;
  height: 62%;
  object-fit: contain;
  filter: drop-shadow(0 8px 18px rgba(0, 0, 0, 0.18));
}
.eth-profile-bar {
  display: block;
  width: 28px;
  height: 4px;
  margin: 8px auto 0;
  border-radius: 999px;
  background: var(--eth-accent);
}
.eth-chip-btn {
  justify-content: center;
  align-items: center;
  height: 34px;
  padding: 0 12px;
  border-radius: 10px;
  background: var(--eth-accent-soft);
  color: var(--eth-accent);
  font-size: 13px;
  font-weight: 600;
  display: inline-flex;
  transition: background 0.2s ease, color 0.2s ease;
}
.eth-chip-btn:hover { background: var(--eth-accent); color: #fff; }
.eth-round-btn {
  display: grid;
  place-items: center;
  width: 36px;
  height: 36px;
  border-radius: 10px;
  background: var(--eth-bg);
  color: var(--eth-muted);
  transition: background 0.2s ease, color 0.2s ease;
}
.eth-round-btn:hover { background: var(--eth-accent-soft); color: var(--eth-accent); }

/* ---------- 分类条 ---------- */
.eth-cats {
  display: flex;
  gap: 8px;
  padding: 10px;
  overflow-x: auto;
  scrollbar-width: none;
}
.eth-cats::-webkit-scrollbar { display: none; }
.eth-cat {
  flex: none;
  display: inline-flex;
  align-items: center;
  gap: 8px;
  height: 36px;
  padding: 0 14px;
  border-radius: 10px;
  background: var(--eth-bg);
  color: var(--eth-ink);
  font-size: 14px;
  transition: background 0.2s ease, color 0.2s ease;
}
.eth-cat:hover { background: var(--eth-accent-soft); color: var(--eth-accent); }
.eth-cat.is-active { background: var(--eth-accent); color: #fff; }
.eth-cat--home { padding: 0 12px; }
.eth-cat-count {
  font-size: 12px;
  opacity: 0.6;
  font-variant-numeric: tabular-nums;
}

/* ---------- 文章卡片 ---------- */
.eth-post {
  display: flex;
  gap: 20px;
  padding: 14px;
  min-width: 0;
  transition: transform 0.25s ease, box-shadow 0.25s ease, border-color 0.25s ease;
  animation: eth-rise 0.5s ease both;
}
.eth-post:hover {
  transform: translateY(-3px);
  border-color: #cdd8f6;
  box-shadow: 0 14px 34px rgba(79, 124, 240, 0.14);
}
@keyframes eth-rise { from { opacity: 0; transform: translateY(14px); } to { opacity: 1; transform: none; } }
.eth-post-body {
  flex: 1;
  min-width: 0;
  padding: 10px 4px 10px 8px;
  display: flex;
  flex-direction: column;
}
.eth-post-title {
  position: relative;
  padding-left: 14px;
  font-size: 22px;
  font-weight: 800;
  line-height: 1.35;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
.eth-post-title::before {
  content: '';
  position: absolute;
  left: 0;
  top: 0.35em;
  width: 4px;
  height: 0.9em;
  border-radius: 4px;
  background: var(--eth-accent);
}
.eth-post:hover .eth-post-title { color: var(--eth-accent); }
.eth-post-meta {
  margin-top: 12px;
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 8px 14px;
  font-size: 13px;
  color: var(--eth-muted);
}
.eth-meta-chip {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 3px 10px;
  border-radius: 8px;
  background: var(--eth-accent-soft);
  color: var(--eth-accent);
}
.eth-meta-item {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  font-variant-numeric: tabular-nums;
}
.eth-post-summary {
  margin-top: 12px;
  font-size: 14px;
  line-height: 1.75;
  color: var(--eth-muted);
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
.eth-post-cover {
  flex: none;
  width: 36%;
  max-width: 340px;
  aspect-ratio: 16 / 10;
  border-radius: 12px;
  overflow: hidden;
  background: var(--eth-bg);
}
.eth-post-cover img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.5s ease;
}
.eth-post:hover .eth-post-cover img { transform: scale(1.05); }
@media (max-width: 640px) {
  .eth-post { flex-direction: column-reverse; gap: 4px; padding: 10px; }
  .eth-post-cover { width: 100%; max-width: none; }
  .eth-post-title { font-size: 18px; }
  .eth-post-body { padding: 8px 4px; }
}

/* ---------- 分页 ---------- */
.eth-pager {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 8px;
  padding-top: 4px;
}
.eth-page-btn {
  min-width: 38px;
  height: 38px;
  padding: 0 10px;
  border-radius: 10px;
  background: var(--eth-card);
  border: 1px solid var(--eth-line);
  display: inline-flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  font-variant-numeric: tabular-nums;
}
.eth-page-btn:hover:not(:disabled) { border-color: var(--eth-accent); color: var(--eth-accent); }
.eth-page-btn.is-active { background: var(--eth-accent); border-color: var(--eth-accent); color: #fff; }
.eth-page-btn:disabled { opacity: 0.4; cursor: not-allowed; }

/* ---------- 右栏小组件 ---------- */
.eth-widget { padding: 18px 18px 20px; }
.eth-widget-title {
  position: relative;
  margin-bottom: 12px;
  padding-left: 12px;
  font-size: 17px;
  font-weight: 700;
}
.eth-widget-title::before {
  content: '';
  position: absolute;
  left: 0;
  top: 0.25em;
  width: 4px;
  height: 1em;
  border-radius: 4px;
  background: var(--eth-accent);
}
.eth-stat {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
}
.eth-stat span {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  color: var(--eth-muted);
}
.eth-stat span :deep(svg) { color: var(--eth-accent); }
.eth-stat b { font-variant-numeric: tabular-nums; }

@media (prefers-reduced-motion: reduce) {
  .eth-post, .eth-marquee-track, .eth-caret { animation: none; }
  .eth-hero-bg { transition: none; }
}
</style>
