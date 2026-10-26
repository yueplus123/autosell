<template>
  <div class="bg-[color:var(--xh-page)] text-[color:var(--xh-ink)]">
    <!-- ==================== 第一屏：科技风首屏（AI + Visa + 二次元） ==================== -->
    <section class="xh-hero overflow-hidden text-[#f5f5f7]">
      <div class="xh-scan" aria-hidden="true"></div>
      <div class="mx-auto max-w-[1240px] px-4 pt-12 text-center sm:px-5 sm:pt-14">
        <span class="inline-flex items-center gap-2 rounded-full border border-cyan-400/30 bg-cyan-400/10 px-3 py-1 font-mono text-[11px] uppercase tracking-[0.24em] text-cyan-300 sm:text-xs">
          <span class="h-1.5 w-1.5 animate-pulse rounded-full bg-cyan-300"></span>
          AI · VISA · ACG
        </span>
        <h1 class="xh-hero-title mt-4 text-[clamp(40px,7vw,88px)] font-black leading-[1.1] tracking-[-0.02em]">{{ heroTitle }}</h1>
        <p class="mt-3 text-[clamp(15px,1.6vw,21px)] text-[#a1a1a6] sm:mt-4">{{ heroSubtitle }}</p>

        <div class="xh-wall mt-5 sm:mt-6">
          <div
            v-for="tile in tiles"
            :key="tile.icon"
            role="img"
            :aria-label="tile.label"
            class="xh-tile"
            :class="[`xh-tile--${tile.tone}`, `xh-tile--${tile.row}`, { 'xh-tile--wide': tile.wide }]"
            :style="tile.style"
          >
            <img :src="tile.src" alt="" />
          </div>
        </div>

        <!-- 官网展示模式：首屏直接给两个专区入口 -->
        <div v-if="showcase" class="flex flex-wrap justify-center gap-3 pb-8 pt-2 sm:pb-10">
          <RouterLink :to="aiZonePath" class="inline-flex items-center gap-2 rounded-full bg-white px-6 py-3 text-[15px] font-semibold text-black shadow-[0_0_36px_rgba(34,211,238,0.35)] transition hover:scale-105">
            {{ t('home.business.ai.cta') }}
            <ArrowRight class="h-4 w-4" aria-hidden="true" />
          </RouterLink>
          <RouterLink :to="gameZonePath" class="inline-flex items-center gap-2 rounded-full bg-linear-to-r from-violet-600 via-fuchsia-500 to-orange-400 px-6 py-3 text-[15px] font-semibold text-white shadow-[0_0_36px_rgba(217,70,239,0.35)] transition hover:scale-105">
            {{ t('home.business.game.cta') }}
            <ArrowRight class="h-4 w-4" aria-hidden="true" />
          </RouterLink>
        </div>
        <a v-else href="#zones" class="mt-2 inline-flex flex-col items-center gap-1.5 pb-6 pt-1 text-[13px] text-[#a1a1a6] transition-colors hover:text-white sm:pb-7" @click.prevent="scrollToId('zones')">
          <span>{{ t('xinhuo.hero.scroll') }}</span>
          <ChevronDown class="h-5 w-5" />
        </a>
      </div>
    </section>

    <!-- ==================== 官网展示模式：文章列表（首页即博客） ==================== -->
    <section v-if="showcase" id="articles" class="mx-auto max-w-[1200px] px-4 py-10 sm:px-6 sm:py-16">
      <h2 class="mb-4 text-2xl font-black sm:mb-6 sm:text-[32px]">{{ t('xinhuo.posts.title') }}</h2>
      <div v-if="articles.length" class="grid grid-cols-1 gap-2.5 sm:grid-cols-[repeat(auto-fill,minmax(300px,1fr))] sm:gap-5">
        <RouterLink v-for="post in articles" :key="post.id" :to="`/blog/${post.slug}`" class="flex min-w-0 gap-3 overflow-hidden rounded-2xl bg-white p-3 text-[color:var(--xh-ink)] transition-shadow hover:shadow-[0_8px_30px_rgba(0,0,0,0.08)] sm:flex-col sm:gap-0 sm:rounded-[20px] sm:p-0">
          <div class="h-[68px] w-[88px] flex-none overflow-hidden rounded-[10px] bg-[#e8e8ed] sm:h-[180px] sm:w-full sm:rounded-none">
            <img v-if="post.thumbnail" :src="getImageUrl(post.thumbnail)" :alt="getLocalizedText(post.title)" class="h-full w-full object-cover" loading="lazy" />
          </div>
          <div class="min-w-0 sm:px-[22px] sm:pb-6 sm:pt-5">
            <span class="text-[11px] text-[color:var(--xh-muted)] sm:text-[13px]">{{ formatDate(post.published_at) }}</span>
            <h3 class="mt-0.5 line-clamp-2 text-sm font-bold leading-snug sm:mt-2 sm:text-lg">{{ getLocalizedText(post.title) }}</h3>
            <p v-if="getLocalizedText(post.summary)" class="mt-2 line-clamp-2 text-sm text-[color:var(--xh-muted)] max-sm:hidden">{{ getLocalizedText(post.summary) }}</p>
          </div>
        </RouterLink>
      </div>
      <p v-else-if="!articlesLoading" class="rounded-[20px] bg-white px-6 py-14 text-center text-sm text-[color:var(--xh-muted)]">{{ t('xinhuo.posts.empty') }}</p>
      <div v-if="articlesHasMore" class="mt-6 flex justify-center">
        <button type="button" class="rounded-full border border-[color:var(--xh-line)] bg-white px-6 py-2.5 text-sm font-medium transition hover:border-[color:var(--xh-accent)] hover:text-[color:var(--xh-accent)] disabled:opacity-50" :disabled="articlesLoading" @click="loadArticles(articlesPage + 1)">
          {{ articlesLoading ? t('xinhuo.posts.loading') : t('xinhuo.posts.more') }}
        </button>
      </div>
    </section>

    <template v-if="!showcase">
    <!-- ==================== 第二屏：左 AI 专区 / 右二次元赞助专区 ==================== -->
    <XinhuoZoneSplit />

    <!-- ==================== 商品展示模块（ChatGPT / 海外账号与礼品卡 / 更多会员） ==================== -->
    <XinhuoModules />
    </template>

    <!-- ==================== 为什么在薪火数卡购买 ==================== -->
    <XinhuoWhy />

    <template v-if="!showcase">
    <!-- ==================== 三步购买 ==================== -->
    <section class="bg-white">
      <div class="mx-auto max-w-[1200px] px-4 py-10 sm:px-6 sm:py-16">
        <h2 class="text-2xl font-black sm:text-center sm:text-[32px]">{{ t('xinhuo.steps.title') }}</h2>
        <ol class="mt-5 flex flex-col gap-4 sm:mt-10 sm:grid sm:grid-cols-3 sm:gap-8 sm:text-center">
          <li v-for="(step, idx) in stepItems" :key="idx" class="flex items-center gap-3.5 sm:block">
            <span class="w-11 flex-none text-[28px] font-bold tabular-nums text-[color:var(--xh-accent)] sm:block sm:w-auto sm:text-[44px]">{{ String(idx + 1).padStart(2, '0') }}</span>
            <div>
              <h3 class="text-[15px] font-bold sm:mt-2 sm:text-lg">{{ step.title }}</h3>
              <p class="mt-0.5 text-[13px] text-[color:var(--xh-muted)] sm:mt-1.5 sm:text-sm">{{ step.desc }}</p>
            </div>
          </li>
        </ol>
      </div>
    </section>

    <!-- ==================== 关于我们 ==================== -->
    <section class="mx-auto max-w-[1200px] px-4 py-10 sm:px-6 sm:py-16">
      <div class="flex flex-wrap items-center gap-6 sm:gap-12">
        <div class="min-w-0 flex-[1_1_420px]">
          <p class="text-[13px] font-medium text-[color:var(--xh-accent)] sm:text-sm">{{ t('xinhuo.about.eyebrow') }}</p>
          <h2 class="mb-3 mt-2 text-2xl font-black leading-snug sm:mb-4 sm:text-[32px]">{{ aboutTitle }}</h2>
          <p class="text-sm leading-[1.8] text-[color:var(--xh-muted)] sm:text-base">{{ aboutText }}</p>
          <RouterLink v-if="aboutEnabled" to="/about" class="mt-4 inline-block text-sm font-medium text-[color:var(--xh-accent)] hover:text-[color:var(--xh-accent-hover)] sm:mt-5 sm:text-[15px]">{{ t('xinhuo.about.more') }} ›</RouterLink>
        </div>
        <ul v-if="aboutServices.length" class="flex min-w-0 flex-[1_1_420px] flex-col gap-3 rounded-[18px] bg-white p-5 sm:rounded-3xl sm:p-8">
          <li v-for="(item, idx) in aboutServices" :key="idx" class="flex items-start gap-3 text-sm sm:text-[15px]">
            <span class="mt-0.5 grid h-5 w-5 flex-none place-items-center rounded-full bg-[color:var(--xh-accent)] text-white"><Check class="h-3 w-3" /></span>
            <span>{{ item }}</span>
          </li>
        </ul>
      </div>
    </section>

    <!-- ==================== 最新文章（分站不显示官网文章） ==================== -->
    <section v-if="!isReseller && blogEnabled && posts.length" class="mx-auto max-w-[1200px] px-4 pb-10 sm:px-6 sm:pb-16">
      <div class="mb-3.5 flex items-baseline justify-between gap-4 sm:mb-6">
        <h2 class="text-2xl font-black sm:text-[32px]">{{ t('xinhuo.posts.title') }}</h2>
        <RouterLink to="/blog" class="text-sm font-medium text-[color:var(--xh-accent)] hover:text-[color:var(--xh-accent-hover)] sm:text-[15px]">{{ t('xinhuo.posts.all') }} ›</RouterLink>
      </div>
      <div class="flex flex-col gap-2.5 sm:grid sm:grid-cols-[repeat(auto-fit,minmax(300px,1fr))] sm:gap-5">
        <RouterLink v-for="post in posts" :key="post.id" :to="`/blog/${post.slug}`" class="flex gap-3 overflow-hidden rounded-2xl bg-white p-3 text-[color:var(--xh-ink)] transition-shadow hover:shadow-[0_8px_30px_rgba(0,0,0,0.08)] sm:flex-col sm:gap-0 sm:rounded-[20px] sm:p-0">
          <div class="h-[68px] w-[88px] flex-none overflow-hidden rounded-[10px] bg-[#e8e8ed] sm:h-[170px] sm:w-full sm:rounded-none">
            <img v-if="post.thumbnail" :src="getImageUrl(post.thumbnail)" :alt="getLocalizedText(post.title)" class="h-full w-full object-cover" loading="lazy" />
          </div>
          <div class="min-w-0 sm:px-[22px] sm:pb-6 sm:pt-5">
            <span class="text-[11px] text-[color:var(--xh-muted)] sm:text-[13px]">{{ formatDate(post.published_at) }}</span>
            <h3 class="mt-0.5 line-clamp-2 text-sm font-bold leading-snug sm:mt-2 sm:text-lg">{{ getLocalizedText(post.title) }}</h3>
            <p v-if="getLocalizedText(post.summary)" class="mt-2 line-clamp-2 text-sm text-[color:var(--xh-muted)] max-sm:hidden">{{ getLocalizedText(post.summary) }}</p>
          </div>
        </RouterLink>
      </div>
    </section>

    <!-- ==================== 常见问题 ==================== -->
    <section class="bg-white">
      <div class="mx-auto max-w-[820px] px-4 py-10 sm:px-6 sm:py-16">
        <h2 class="mb-2 text-2xl font-black sm:mb-6 sm:text-center sm:text-[32px]">{{ t('xinhuo.faq.title') }}</h2>
        <details v-for="(item, idx) in faqItems" :key="idx" class="group border-b border-[color:var(--xh-line)] py-3.5 sm:py-[18px]">
          <summary class="flex cursor-pointer list-none items-center justify-between gap-4 text-[15px] font-medium sm:text-base">
            {{ item.q }}
            <Plus class="h-4 w-4 flex-none text-[color:var(--xh-muted)] transition-transform group-open:rotate-45" />
          </summary>
          <p class="mt-2 text-[13px] leading-[1.8] text-[color:var(--xh-muted)] sm:mt-2.5 sm:text-sm">{{ item.a }}</p>
        </details>
      </div>
    </section>

    <!-- ==================== 购买须知 ==================== -->
    <section id="notice" class="mx-auto max-w-[1200px] px-4 py-8 sm:px-6 sm:pb-14 sm:pt-12">
      <h2 class="mb-3 text-[17px] font-bold sm:mb-4 sm:text-xl">{{ t('xinhuo.notice.title') }}</h2>
      <ol class="flex list-decimal flex-col gap-2 pl-[18px] text-xs leading-[1.75] text-[color:var(--xh-muted)] sm:gap-2.5 sm:pl-5 sm:text-[13px] sm:leading-[1.8]">
        <li v-for="(line, idx) in noticeItems" :key="idx">{{ line }}</li>
      </ol>
      <p class="mt-4 text-xs text-[color:var(--xh-muted)] sm:text-[13px]">
        {{ t('xinhuo.notice.more') }}
        <RouterLink to="/terms" class="text-[color:var(--xh-accent)] hover:underline">{{ t('xinhuo.notice.terms') }}</RouterLink>
        {{ t('xinhuo.notice.and') }}
        <RouterLink to="/privacy" class="text-[color:var(--xh-accent)] hover:underline">{{ t('xinhuo.notice.privacy') }}</RouterLink>
      </p>
    </section>
    </template>
  </div>
</template>

<script setup lang="ts">
import { computed, nextTick, onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { ArrowRight, Check, ChevronDown, Plus } from 'lucide-vue-next'
import { postAPI } from '../../api'
import { getImageUrl } from '../../utils/image'
import { useLocalized } from '../../composables/useProduct'
import { useNavConfig } from '../../composables/useNavConfig'
import { usePageSeo } from '../../composables/usePageSeo'
import { useAppStore } from '../../stores/app'
import { isMainShowcase } from '../../utils/storefrontZone'
import { businessZonePath } from '../../utils/businessLines'
import XinhuoZoneSplit from './components/XinhuoZoneSplit.vue'
import XinhuoModules from './components/XinhuoModules.vue'
import XinhuoWhy from './components/XinhuoWhy.vue'
import { HERO_ICON_BASE, heroTiles } from './heroTiles'

const route = useRoute()
const { t, te } = useI18n()
const { getLocalizedText } = useLocalized()
const appStore = useAppStore()
const { blogEnabled, aboutEnabled } = useNavConfig()

usePageSeo({ canonicalPath: () => route.path })

// 官网展示模式：首页 = 首屏（两个专区入口）+ 文章列表 + 为什么选择薪火；分站保持商城首页
const showcase = computed(() => isMainShowcase(appStore.config as any))
const isReseller = computed(() => String((appStore.config as any)?.tenant?.mode || '') === 'reseller')
const aiZonePath = businessZonePath('ai')
const gameZonePath = businessZonePath('game')

// ==================== 首屏 ====================
const heroTitle = computed(() => t('xinhuo.hero.title'))
const heroSubtitle = computed(() => t('xinhuo.hero.subtitle'))

const tiles = heroTiles.map((tile) => ({
  ...tile,
  src: HERO_ICON_BASE + tile.icon,
  style: {
    '--x': `${tile.x}%`,
    '--y': `${tile.y}%`,
    '--s': `${tile.s}%`,
    '--mx': `${tile.mx}%`,
    '--my': `${tile.my}%`,
    '--ms': `${tile.ms}%`,
    ...(tile.glow ? { '--glow': tile.glow } : {}),
    zIndex: tile.z ?? 1,
  } as Record<string, string | number>,
}))

const scrollToId = (id: string) => {
  document.getElementById(id)?.scrollIntoView({ behavior: 'smooth', block: 'start' })
}

// ==================== 文案（i18n 里的列表，按下标逐条读取，增删条目无需改代码） ====================
const MAX_ITEMS = 20
const readList = <T>(base: string, pick: (key: string) => T | null): T[] => {
  void appStore.locale // 切换语言时重新计算
  const list: T[] = []
  for (let i = 0; i < MAX_ITEMS; i++) {
    const item = pick(`${base}.${i}`)
    if (item === null) break
    list.push(item)
  }
  return list
}
const stepItems = computed(() => readList('xinhuo.steps.items', (k) => (te(`${k}.title`) ? { title: t(`${k}.title`), desc: t(`${k}.desc`) } : null)))
const faqItems = computed(() => readList('xinhuo.faq.items', (k) => (te(`${k}.q`) ? { q: t(`${k}.q`), a: t(`${k}.a`) } : null)))
const noticeItems = computed(() => readList('xinhuo.notice.items', (k) => (te(k) ? t(k) : null)))

// ==================== 关于我们 ====================
// 优先用后台“关于我们”页的配置（标题 / 简介 / 服务项目），没配置时用模板默认文案
interface AboutConfig {
  hero?: { title?: Record<string, string> }
  introduction?: Record<string, string>
  services?: { items?: Record<string, string>[] }
}
const aboutConfig = computed(() => (appStore.config?.about || {}) as AboutConfig)
const aboutTitle = computed(() => getLocalizedText(aboutConfig.value.hero?.title).trim() || t('xinhuo.about.title'))
const aboutText = computed(() => getLocalizedText(aboutConfig.value.introduction).trim() || t('xinhuo.about.desc'))
const aboutServices = computed(() => {
  const items = aboutConfig.value.services?.items
  if (!Array.isArray(items)) return []
  return items.map((item) => getLocalizedText(item).trim()).filter(Boolean).slice(0, 6)
})

// ==================== 文章 ====================
const posts = ref<any[]>([])
const loadPosts = async () => {
  if (!blogEnabled.value) return
  try {
    const res = await postAPI.list({ page: 1, page_size: 3, type: 'blog' })
    posts.value = res.data.data || []
  } catch (err) {
    console.error('Failed to load posts:', err)
  }
}

// 官网文章列表（分页加载）
const ARTICLE_PAGE_SIZE = 9
const articles = ref<any[]>([])
const articlesPage = ref(0)
const articlesTotal = ref(0)
const articlesLoading = ref(false)
const articlesHasMore = computed(() => articles.value.length < articlesTotal.value)
const loadArticles = async (page: number) => {
  if (articlesLoading.value) return
  articlesLoading.value = true
  try {
    const res = await postAPI.list({ page, page_size: ARTICLE_PAGE_SIZE, type: 'blog' })
    const list = res.data.data || []
    articles.value = page === 1 ? list : [...articles.value, ...list]
    articlesTotal.value = Number(res.data.pagination?.total ?? articles.value.length)
    articlesPage.value = page
  } catch (err) {
    console.error('Failed to load articles:', err)
  } finally {
    articlesLoading.value = false
  }
}
const formatDate = (value: string) => (value ? new Date(value).toLocaleDateString() : '')

onMounted(async () => {
  await appStore.loadConfig()
  if (showcase.value) {
    await loadArticles(1)
  } else if (!isReseller.value) {
    await loadPosts()
  }
  // 带锚点进入首页时（如 /#zones），数据加载完再定位
  if (route.hash) {
    await nextTick()
    scrollToId(route.hash.slice(1))
  }
})
</script>
