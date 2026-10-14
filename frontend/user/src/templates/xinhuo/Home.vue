<template>
  <div class="bg-[color:var(--xh-page)] text-[color:var(--xh-ink)]">
    <!-- ==================== 首屏 ==================== -->
    <section class="xh-hero overflow-hidden text-[#f5f5f7]">
      <div class="mx-auto max-w-[1240px] px-4 pt-12 text-center sm:px-5 sm:pt-14">
        <h1 class="text-[clamp(40px,7vw,88px)] font-black leading-[1.1] tracking-[-0.02em]">{{ heroTitle }}</h1>
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

        <a href="#shop" class="inline-flex flex-col items-center gap-1.5 pb-6 pt-1 text-[13px] text-[#a1a1a6] transition-colors hover:text-white sm:pb-7" @click.prevent="scrollToId('shop')">
          <span>{{ t('xinhuo.hero.scroll') }}</span>
          <ChevronDown class="h-5 w-5" />
        </a>
      </div>
    </section>

    <!-- ==================== 订单查询 ==================== -->
    <section id="lookup" class="mx-auto max-w-[1200px] scroll-mt-20 px-4 pt-6 sm:px-6 sm:pt-10">
      <div class="flex flex-wrap items-center gap-4 rounded-[18px] bg-white p-[18px] sm:gap-6 sm:rounded-[20px] sm:px-7 sm:py-6">
        <div class="min-w-0 flex-[1_1_240px]">
          <h2 class="text-[17px] font-bold sm:text-xl">{{ t('xinhuo.lookup.title') }}</h2>
          <p class="mt-1 text-[13px] text-[color:var(--xh-muted)] sm:text-sm">{{ t('xinhuo.lookup.desc') }}</p>
        </div>
        <form class="flex min-w-0 flex-[1_1_520px] flex-wrap gap-2.5 sm:gap-3" novalidate @submit.prevent="submitLookup">
          <label class="min-w-0 flex-[1_1_200px]">
            <span class="sr-only">{{ t('xinhuo.lookup.email') }}</span>
            <input v-model.trim="lookupEmail" type="email" autocomplete="email" :placeholder="t('xinhuo.lookup.email')" class="h-11 w-full rounded-xl border border-[color:var(--xh-line)] bg-[color:var(--xh-page)] px-4 text-[15px] outline-none transition-colors focus:border-[color:var(--xh-accent)] sm:h-12" @input="lookupError = ''" />
          </label>
          <label class="min-w-0 flex-[1_1_160px]">
            <span class="sr-only">{{ t('xinhuo.lookup.password') }}</span>
            <input v-model="lookupPassword" type="password" autocomplete="current-password" :placeholder="t('xinhuo.lookup.password')" class="h-11 w-full rounded-xl border border-[color:var(--xh-line)] bg-[color:var(--xh-page)] px-4 text-[15px] outline-none transition-colors focus:border-[color:var(--xh-accent)] sm:h-12" @input="lookupError = ''" />
          </label>
          <button type="submit" class="h-11 flex-none rounded-xl bg-[color:var(--xh-ink)] px-7 text-[15px] font-medium text-white transition-colors hover:bg-black sm:h-12">{{ t('xinhuo.lookup.submit') }}</button>
          <p v-if="lookupError" class="w-full text-[13px] text-[#c62828]" role="alert">{{ lookupError }}</p>
        </form>
      </div>
    </section>

    <!-- ==================== 热门商品 ==================== -->
    <section id="shop" class="mx-auto max-w-[1200px] scroll-mt-20 pb-6 pt-10 sm:px-6 sm:pt-14">
      <div class="flex flex-wrap items-end justify-between gap-4 px-4 sm:gap-6 sm:px-0">
        <div>
          <h2 class="text-[28px] font-black leading-tight tracking-[-0.01em] sm:text-[40px]">{{ t('xinhuo.shop.title') }}</h2>
          <p class="mt-2 text-[15px] text-[color:var(--xh-muted)] max-sm:hidden sm:text-[17px]">{{ t('xinhuo.shop.desc') }}</p>
        </div>
        <label class="flex h-11 w-full items-center gap-2.5 rounded-full border border-[color:var(--xh-line)] bg-white px-4 sm:h-12 sm:w-[340px] sm:px-[18px]">
          <Search class="h-[18px] w-[18px] flex-none text-[color:var(--xh-muted)]" />
          <span class="sr-only">{{ t('xinhuo.shop.search') }}</span>
          <input v-model.trim="keyword" type="search" :placeholder="t('xinhuo.shop.search')" class="min-w-0 flex-1 bg-transparent text-[15px] outline-none" />
        </label>
      </div>

      <div v-if="tabs.length > 1" class="xh-scroll-x mt-4 flex gap-2 overflow-x-auto px-4 pb-1 sm:mt-7 sm:flex-wrap sm:gap-2.5 sm:overflow-visible sm:px-0">
        <button
          v-for="tab in tabs"
          :key="tab.id"
          type="button"
          class="h-9 flex-none whitespace-nowrap rounded-full border px-4 text-[13px] transition-colors sm:h-10 sm:px-[18px] sm:text-sm"
          :class="activeTab === tab.id ? 'border-[color:var(--xh-ink)] bg-[color:var(--xh-ink)] font-medium text-white' : 'border-[color:var(--xh-line)] bg-white text-[color:var(--xh-ink)] hover:border-[color:var(--xh-muted)]'"
          :aria-pressed="activeTab === tab.id"
          @click="activeTab = tab.id"
        >{{ tab.label }}</button>
      </div>

      <div class="mt-3 grid gap-2.5 px-4 sm:mt-6 sm:grid-cols-[repeat(auto-fill,minmax(260px,1fr))] sm:gap-5 sm:px-0">
        <template v-if="productsLoading">
          <div v-for="n in 8" :key="`sk-${n}`" class="h-[86px] animate-pulse rounded-2xl bg-white sm:h-[236px] sm:rounded-[20px]"></div>
        </template>
        <XinhuoProductCard v-for="(product, idx) in visibleProducts" v-else :key="product.id" :product="product" :index="idx" />
      </div>
      <p v-if="!productsLoading && !visibleProducts.length" class="px-4 py-12 text-center text-sm text-[color:var(--xh-muted)]">
        {{ keyword ? t('xinhuo.shop.emptySearch') : t('xinhuo.shop.empty') }}
      </p>
      <div class="mt-5 text-center sm:mt-7">
        <RouterLink to="/products" class="text-sm font-medium text-[color:var(--xh-accent)] hover:text-[color:var(--xh-accent-hover)] sm:text-[15px]">{{ t('xinhuo.shop.viewAll') }} ›</RouterLink>
      </div>
    </section>

    <!-- ==================== 为什么选我们 ==================== -->
    <section class="mx-auto max-w-[1200px] px-4 py-10 sm:px-6 sm:py-14">
      <h2 class="mb-4 text-2xl font-black sm:mb-6 sm:text-[32px]">{{ t('xinhuo.why.title') }}</h2>
      <div class="grid grid-cols-2 gap-2.5 sm:grid-cols-[repeat(auto-fit,minmax(240px,1fr))] sm:gap-5">
        <div v-for="(item, idx) in whyItems" :key="idx" class="rounded-2xl bg-white p-4 sm:rounded-[20px] sm:p-7">
          <component :is="whyIcons[idx % whyIcons.length]" class="h-[22px] w-[22px] text-[color:var(--xh-accent)] sm:h-7 sm:w-7" />
          <h3 class="mb-1 mt-2.5 text-[15px] font-bold sm:mb-1.5 sm:mt-4 sm:text-lg">{{ item.title }}</h3>
          <p class="text-xs leading-relaxed text-[color:var(--xh-muted)] sm:text-sm sm:leading-[1.7]">{{ item.desc }}</p>
        </div>
      </div>
    </section>

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

    <!-- ==================== 最新文章 ==================== -->
    <section v-if="blogEnabled && posts.length" class="mx-auto max-w-[1200px] px-4 pb-10 sm:px-6 sm:pb-16">
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
  </div>
</template>

<script setup lang="ts">
import { computed, nextTick, onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { Check, ChevronDown, CreditCard, Headphones, Plus, Search, ShieldCheck, Zap } from 'lucide-vue-next'
import { categoryAPI, postAPI, productAPI } from '../../api'
import { buildCategoryGroups, type CategoryGroup } from '../../utils/category'
import { getImageUrl } from '../../utils/image'
import { saveGuestOrderAuth } from '../../utils/guestOrderAuth'
import { useLocalized } from '../../composables/useProduct'
import { useNavConfig } from '../../composables/useNavConfig'
import { usePageSeo } from '../../composables/usePageSeo'
import { useAppStore } from '../../stores/app'
import XinhuoProductCard from './components/XinhuoProductCard.vue'
import { HERO_ICON_BASE, heroTiles } from './heroTiles'

const route = useRoute()
const router = useRouter()
const { t, te } = useI18n()
const { getLocalizedText } = useLocalized()
const appStore = useAppStore()
const { blogEnabled, aboutEnabled } = useNavConfig()

usePageSeo({ canonicalPath: () => route.path })

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

// ==================== 订单查询 ====================
// 与游客订单页共用同一份凭据存储，跳转过去后会自动查询
const lookupEmail = ref('')
const lookupPassword = ref('')
const lookupError = ref('')
const submitLookup = () => {
  if (!lookupEmail.value || !lookupPassword.value) {
    lookupError.value = t('xinhuo.lookup.missing')
    return
  }
  saveGuestOrderAuth({ email: lookupEmail.value, order_password: lookupPassword.value })
  router.push('/guest/orders')
}

// ==================== 商品 ====================
const products = ref<any[]>([])
const productsLoading = ref(true)
const categoryGroups = ref<CategoryGroup[]>([])
const activeTab = ref<number>(0)
const keyword = ref('')
const HOME_PRODUCT_LIMIT = 8

const tabs = computed(() => [
  { id: 0, label: t('xinhuo.shop.all') },
  ...categoryGroups.value.map((group) => ({ id: group.id, label: getLocalizedText(group.name) || group.slug || '' })),
])

// 一级分类 → 自身及子分类 id，商品挂在子分类下时也能归到对应标签
const groupCategoryIds = computed(() => {
  const map = new Map<number, Set<number>>()
  categoryGroups.value.forEach((group) => {
    map.set(group.id, new Set([group.id, ...group.children.map((child) => child.id)]))
  })
  return map
})

const visibleProducts = computed(() => {
  let list = products.value
  if (activeTab.value) {
    const ids = groupCategoryIds.value.get(activeTab.value)
    list = ids ? list.filter((product) => ids.has(Number(product.category_id))) : list
  }
  const kw = keyword.value.toLowerCase()
  if (kw) {
    list = list.filter((product) => getLocalizedText(product.title).toLowerCase().includes(kw))
  }
  return list.slice(0, HOME_PRODUCT_LIMIT)
})

const loadProducts = async () => {
  productsLoading.value = true
  try {
    const res = await productAPI.list({ page: 1, page_size: 60 })
    products.value = res.data.data || []
  } catch (err) {
    console.error('Failed to load products:', err)
  } finally {
    productsLoading.value = false
  }
}

const loadCategories = async () => {
  try {
    const res = await categoryAPI.list()
    categoryGroups.value = buildCategoryGroups(res.data.data || []).slice(0, 8)
  } catch (err) {
    console.error('Failed to load categories:', err)
  }
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
const whyItems = computed(() => readList('xinhuo.why.items', (k) => (te(`${k}.title`) ? { title: t(`${k}.title`), desc: t(`${k}.desc`) } : null)))
const stepItems = computed(() => readList('xinhuo.steps.items', (k) => (te(`${k}.title`) ? { title: t(`${k}.title`), desc: t(`${k}.desc`) } : null)))
const faqItems = computed(() => readList('xinhuo.faq.items', (k) => (te(`${k}.q`) ? { q: t(`${k}.q`), a: t(`${k}.a`) } : null)))
const noticeItems = computed(() => readList('xinhuo.notice.items', (k) => (te(k) ? t(k) : null)))
const whyIcons = [Zap, ShieldCheck, CreditCard, Headphones]

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
const formatDate = (value: string) => (value ? new Date(value).toLocaleDateString() : '')

onMounted(async () => {
  await appStore.loadConfig()
  await Promise.all([loadProducts(), loadCategories(), loadPosts()])
  // 从其他页面点“订单查询”跳回首页时，数据加载完再定位
  if (route.hash) {
    await nextTick()
    scrollToId(route.hash.slice(1))
  }
})
</script>
