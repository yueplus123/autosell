<template>
  <div class="zone-page min-h-screen bg-background text-foreground" :class="isAi ? 'zone-ai' : 'zone-game'">

    <!-- ==================== AI 专区：深色科技风 ==================== -->
    <section v-if="isAi" class="relative overflow-hidden bg-slate-950 pb-14 text-white md:pb-20" :class="heroTopClass">
      <div class="pointer-events-none absolute inset-0 opacity-[0.15]" :style="aiGridStyle"></div>
      <span class="pointer-events-none absolute -left-24 top-10 h-72 w-72 rounded-full bg-violet-600/40 blur-3xl"></span>
      <span class="pointer-events-none absolute -right-20 -top-10 h-80 w-80 rounded-full bg-cyan-500/30 blur-3xl"></span>
      <span class="pointer-events-none absolute bottom-0 left-1/2 h-40 w-[36rem] -translate-x-1/2 rounded-full bg-fuchsia-600/20 blur-3xl"></span>

      <div class="container relative mx-auto px-4">
        <BreadcrumbLite :label="t('home.business.ai.eyebrow')" dark />
        <div class="mt-6 max-w-3xl space-y-5">
          <span class="inline-flex items-center gap-2 rounded-full border border-cyan-400/30 bg-cyan-400/10 px-3 py-1 font-mono text-xs uppercase tracking-[0.2em] text-cyan-300">
            <Sparkles class="h-3.5 w-3.5" aria-hidden="true" />
            {{ t('zone.ai.badge') }}
          </span>
          <h1 class="text-4xl font-black tracking-tight md:text-6xl">
            <span class="bg-linear-to-r from-cyan-300 via-violet-300 to-fuchsia-300 bg-clip-text text-transparent">{{ t('home.business.ai.title') }}</span>
          </h1>
          <p class="max-w-2xl text-base leading-relaxed text-slate-300 md:text-lg">{{ t('home.business.ai.description') }}</p>
          <div class="flex flex-wrap gap-2 pt-1">
            <span v-for="i in 3" :key="i" class="inline-flex items-center gap-1.5 rounded-md border border-white/10 bg-white/5 px-3 py-1.5 text-sm text-slate-200">
              <Check class="h-3.5 w-3.5 text-cyan-300" aria-hidden="true" />
              {{ t(`home.business.ai.highlight${i}`) }}
            </span>
          </div>
        </div>

        <!-- AI 子分类：终端风标签 -->
        <div v-if="children.length" class="mt-10">
          <div class="mb-3 font-mono text-xs uppercase tracking-[0.2em] text-slate-400">{{ t('zone.ai.pickTitle') }}</div>
          <div class="flex flex-wrap gap-2">
            <button
              v-for="opt in filterOptions"
              :key="opt.key"
              type="button"
              class="rounded-lg border px-4 py-2 text-sm font-medium transition"
              :class="opt.active
                ? 'border-cyan-300 bg-cyan-300 text-slate-950 shadow-[0_0_24px_rgba(103,232,249,0.45)]'
                : 'border-white/15 bg-white/5 text-slate-200 hover:border-cyan-300/60 hover:text-white'"
              @click="selectChild(opt.slug)"
            >
              {{ opt.label }}
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- ==================== 二次元专区：明亮活泼风 ==================== -->
    <section v-else class="relative overflow-hidden bg-linear-to-br from-pink-100 via-rose-50 to-amber-100 pb-14 text-slate-900 md:pb-20 dark:from-pink-950/60 dark:via-slate-950 dark:to-amber-950/40 dark:text-white" :class="heroTopClass">
      <span class="pointer-events-none absolute -right-24 -top-24 h-80 w-80 rounded-full bg-pink-300/50 blur-3xl dark:bg-pink-600/30"></span>
      <span class="pointer-events-none absolute -bottom-24 -left-16 h-72 w-72 rounded-full bg-amber-300/50 blur-3xl dark:bg-amber-600/20"></span>
      <Star class="pointer-events-none absolute right-[12%] top-28 h-8 w-8 rotate-12 fill-yellow-300 text-yellow-400" aria-hidden="true" />
      <Star class="pointer-events-none absolute right-[22%] top-48 h-5 w-5 -rotate-12 fill-pink-300 text-pink-400" aria-hidden="true" />
      <Heart class="pointer-events-none absolute right-[8%] top-60 h-6 w-6 fill-rose-300 text-rose-400" aria-hidden="true" />
      <Gamepad2 class="pointer-events-none absolute -bottom-10 right-4 h-56 w-56 -rotate-12 text-pink-400/15" :stroke-width="1.25" aria-hidden="true" />

      <div class="container relative mx-auto px-4">
        <BreadcrumbLite :label="t('zone.game.badge')" />
        <div class="mt-6 max-w-3xl space-y-5">
          <span class="inline-flex items-center gap-2 rounded-full bg-white/80 px-4 py-1.5 text-sm font-bold text-pink-600 shadow-sm dark:bg-white/10 dark:text-pink-300">
            <Gamepad2 class="h-4 w-4" aria-hidden="true" />
            {{ t('zone.game.badge') }}
          </span>
          <h1 class="text-4xl font-black tracking-tight md:text-6xl">
            <span class="bg-linear-to-r from-pink-500 via-rose-500 to-amber-500 bg-clip-text text-transparent">{{ t('home.business.game.title') }}</span>
          </h1>
          <p class="max-w-2xl text-base leading-relaxed text-slate-600 md:text-lg dark:text-slate-300">{{ t('home.business.game.description') }}</p>
          <div class="flex flex-wrap gap-2 pt-1">
            <span v-for="i in 3" :key="i" class="inline-flex items-center gap-1.5 rounded-full bg-white px-3.5 py-1.5 text-sm font-semibold text-rose-600 shadow-sm dark:bg-white/10 dark:text-rose-300">
              <Heart class="h-3.5 w-3.5 fill-current" aria-hidden="true" />
              {{ t(`home.business.game.highlight${i}`) }}
            </span>
          </div>
        </div>

        <!-- 游戏图标墙 -->
        <div v-if="children.length" class="mt-10">
          <div class="mb-3 text-sm font-bold text-slate-700 dark:text-slate-200">{{ t('zone.game.pickTitle') }}</div>
          <div class="grid grid-cols-3 gap-3 sm:grid-cols-4 md:grid-cols-6 lg:grid-cols-8">
            <button
              v-for="opt in filterOptions"
              :key="opt.key"
              type="button"
              class="group flex flex-col items-center gap-2 rounded-2xl p-3 transition"
              :class="opt.active
                ? 'bg-white shadow-lg ring-2 ring-pink-400 dark:bg-white/10'
                : 'bg-white/60 hover:-translate-y-0.5 hover:bg-white hover:shadow-md dark:bg-white/5 dark:hover:bg-white/10'"
              @click="selectChild(opt.slug)"
            >
              <span class="grid h-14 w-14 place-items-center overflow-hidden rounded-2xl bg-linear-to-br from-pink-400 to-amber-300 text-white shadow-inner">
                <img v-if="opt.icon" :src="getImageUrl(opt.icon)" :alt="opt.label" loading="lazy" class="h-full w-full object-cover" />
                <LayoutGrid v-else-if="!opt.slug" class="h-6 w-6" aria-hidden="true" />
                <span v-else class="text-lg font-black">{{ opt.label.slice(0, 1) }}</span>
              </span>
              <span class="line-clamp-1 text-xs font-semibold text-slate-700 dark:text-slate-200">{{ opt.label }}</span>
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- ==================== 商品区 ==================== -->
    <section class="relative z-10 py-10 md:py-14">
      <div class="container mx-auto px-4">
        <div class="mb-6 flex flex-col gap-4 md:flex-row md:items-center md:justify-between">
          <h2 class="flex items-center gap-3 text-2xl font-bold md:text-3xl">
            <span class="h-7 w-1.5 rounded-full" :class="isAi ? 'bg-linear-to-b from-cyan-400 to-violet-500' : 'bg-linear-to-b from-pink-400 to-amber-400'"></span>
            {{ selectedChild ? catName(selectedChild) : t(`home.business.${zoneKey}.hotTitle`) }}
          </h2>
          <div class="relative w-full md:w-72">
            <Search class="pointer-events-none absolute left-3.5 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" aria-hidden="true" />
            <input
              v-model="searchQuery"
              type="text"
              class="h-10 w-full rounded-xl border bg-card pl-10 pr-9 text-sm text-foreground transition-shadow placeholder:text-muted-foreground focus:outline-none focus:ring-2"
              :class="isAi ? 'focus:ring-violet-400/40' : 'focus:ring-pink-400/40'"
              :placeholder="t('products.searchBoxPlaceholder')"
              @keydown.enter="applySearch"
            />
            <button v-if="searchQuery" type="button" class="absolute inset-y-0 right-3 flex items-center text-muted-foreground hover:text-foreground" @click="clearSearch">
              <X class="h-4 w-4" />
            </button>
          </div>
        </div>

        <div v-if="loading" class="grid grid-cols-2 gap-3 md:grid-cols-3 md:gap-4 lg:grid-cols-4">
          <div v-for="i in 8" :key="i" class="h-64 rounded-2xl border bg-card theme-skeleton"></div>
        </div>
        <template v-else-if="products.length">
          <div class="grid grid-cols-2 gap-3 md:grid-cols-3 md:gap-4 lg:grid-cols-4">
            <ProductCard
              v-for="(product, idx) in products"
              :key="product.id"
              :product="product"
              :index="idx"
              :animation-step="50"
              @click="goToProduct"
              @quick-buy="openQuickBuy"
            />
          </div>
          <div class="mt-8">
            <PaginationNav :current-page="currentPage" :total-pages="totalPages" :loading="loading" @change-page="changePage" />
          </div>
        </template>
        <EmptyState
          v-else
          variant="soft"
          icon="package"
          :title="t(searchActive ? 'products.empty' : 'zone.empty')"
          :description="searchActive ? '' : t('zone.emptyDesc')"
        />
      </div>
    </section>

    <!-- ==================== 服务亮点 ==================== -->
    <section class="pb-16">
      <div class="container mx-auto px-4">
        <div class="grid grid-cols-1 gap-4 md:grid-cols-3">
          <div
            v-for="i in 3"
            :key="i"
            class="rounded-2xl border p-6"
            :class="isAi ? 'border-violet-500/20 bg-linear-to-br from-violet-500/5 to-cyan-500/5' : 'border-pink-300/40 bg-linear-to-br from-pink-50 to-amber-50 dark:from-pink-500/10 dark:to-amber-500/5'"
          >
            <span class="mb-4 grid h-11 w-11 place-items-center rounded-xl text-white" :class="isAi ? 'bg-linear-to-br from-violet-600 to-cyan-500' : 'bg-linear-to-br from-pink-500 to-amber-400'">
              <component :is="featureIcons[i - 1]" class="h-5 w-5" aria-hidden="true" />
            </span>
            <h3 class="text-base font-bold">{{ t(`zone.${zoneKey}.f${i}Title`) }}</h3>
            <p class="mt-1.5 text-sm leading-relaxed text-muted-foreground">{{ t(`zone.${zoneKey}.f${i}Desc`) }}</p>
          </div>
        </div>
      </div>
    </section>

    <ProductQuickBuy
      v-if="quickBuyProduct"
      :product="quickBuyProduct"
      :visible="quickBuyVisible"
      @update:visible="quickBuyVisible = $event"
    />
  </div>
</template>

<script setup lang="ts">
import { computed, defineComponent, h, onMounted, ref, watch } from 'vue'
import { RouterLink, useRoute, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import {
  BadgeCheck, Check, ChevronRight, Gamepad2, Headphones, Heart, LayoutGrid, Layers, Search, Sparkles, Star, Tag, X, Zap,
} from 'lucide-vue-next'
import { categoryAPI, productAPI } from '../api'
import { getImageUrl } from '../utils/image'
import { useLocalized } from '../composables/useProduct'
import { usePageSeo } from '../composables/usePageSeo'
import { getActiveTemplate } from '../templates/registry'
import type { PublicCategory } from '../utils/category'
import { businessZonePath, isBusinessLineKey, resolveBusinessLines, type BusinessLineKey } from '../utils/businessLines'
import ProductCard from '../components/ProductCard.vue'
import ProductQuickBuy from '../components/ProductQuickBuy.vue'
import PaginationNav from '../components/PaginationNav.vue'
import EmptyState from '../components/EmptyState.vue'

const route = useRoute()
const router = useRouter()
const { t } = useI18n()
const { getLocalizedText } = useLocalized()

const PAGE_SIZE = 20

const zoneKey = computed<BusinessLineKey>(() => {
  const key = String(route.params.key || '')
  return isBusinessLineKey(key) ? key : 'ai'
})
const isAi = computed(() => zoneKey.value === 'ai')
// classic 模板导航栏是 fixed，需要让出高度；vault / xinhuo 的导航栏是 sticky，不需要
const heroTopClass = getActiveTemplate() === 'classic' ? 'pt-24' : 'pt-10 md:pt-14'

// 简易面包屑：首页 › 专区名
const BreadcrumbLite = defineComponent({
  props: { label: { type: String, required: true }, dark: Boolean },
  setup(props) {
    return () => h('nav', { class: ['flex items-center gap-1.5 text-sm', props.dark ? 'text-slate-400' : 'text-slate-500 dark:text-slate-400'] }, [
      h(RouterLink, { to: '/', class: props.dark ? 'hover:text-white' : 'hover:text-slate-900 dark:hover:text-white' }, () => t('nav.home')),
      h(ChevronRight, { class: 'h-3.5 w-3.5', 'aria-hidden': 'true' }),
      h('span', { class: props.dark ? 'text-slate-200' : 'text-slate-700 dark:text-slate-200' }, props.label),
    ])
  },
})

const aiGridStyle = {
  backgroundImage:
    'linear-gradient(rgba(148,163,184,0.5) 1px, transparent 1px), linear-gradient(90deg, rgba(148,163,184,0.5) 1px, transparent 1px)',
  backgroundSize: '44px 44px',
  maskImage: 'radial-gradient(ellipse at 50% 30%, black 30%, transparent 75%)',
  WebkitMaskImage: 'radial-gradient(ellipse at 50% 30%, black 30%, transparent 75%)',
}

const featureIcons = computed(() => (isAi.value ? [Zap, Layers, Headphones] : [Zap, Tag, BadgeCheck]))

// ==================== 分类 ====================
const categories = ref<PublicCategory[]>([])
const categoriesLoaded = ref(false)
const line = computed(() => resolveBusinessLines(categories.value).find((l) => l.key === zoneKey.value) ?? null)
const children = computed(() => line.value?.children ?? [])

const catName = (cat: PublicCategory) => getLocalizedText(cat.name) || String(cat.slug || '')

const selectedSlug = computed(() => {
  const c = route.query.c
  return typeof c === 'string' ? c : ''
})
const selectedChild = computed(() => children.value.find((c) => String(c.slug) === selectedSlug.value) ?? null)

const filterOptions = computed(() => [
  { key: 'all', slug: '', label: t('zone.all'), icon: '', active: !selectedChild.value },
  ...children.value.map((c) => ({
    key: String(c.id),
    slug: String(c.slug),
    label: catName(c),
    icon: c.icon || '',
    active: selectedChild.value?.id === c.id,
  })),
])

const selectChild = (slug: string) => {
  router.replace(businessZonePath(zoneKey.value, slug || undefined))
}

// ==================== 商品 ====================
const products = ref<any[]>([])
const loading = ref(true)
const currentPage = ref(1)
const totalPages = ref(0)
const searchQuery = ref('')
const appliedSearch = ref('')
const searchActive = computed(() => appliedSearch.value !== '')

const loadProducts = async () => {
  const categoryId = selectedChild.value?.id ?? line.value?.category?.id
  if (!categoryId) {
    products.value = []
    totalPages.value = 0
    loading.value = false
    return
  }
  loading.value = true
  try {
    const params: Record<string, unknown> = { page: currentPage.value, page_size: PAGE_SIZE, category_id: categoryId }
    if (appliedSearch.value) params.search = appliedSearch.value
    const res = await productAPI.list(params)
    products.value = res.data.data || []
    totalPages.value = res.data.pagination?.total_page || 0
  } catch (error) {
    console.error('Failed to load zone products:', error)
    products.value = []
  } finally {
    loading.value = false
  }
}

const changePage = (page: number) => {
  currentPage.value = page
  loadProducts()
}

const applySearch = () => {
  appliedSearch.value = searchQuery.value.trim()
  currentPage.value = 1
  loadProducts()
}

const clearSearch = () => {
  searchQuery.value = ''
  applySearch()
}

watch([zoneKey, selectedSlug], () => {
  if (!categoriesLoaded.value) return
  currentPage.value = 1
  loadProducts()
})

const quickBuyProduct = ref<any>(null)
const quickBuyVisible = ref(false)
const openQuickBuy = (product: any) => {
  quickBuyProduct.value = product
  quickBuyVisible.value = true
}
const goToProduct = (slug: string) => router.push(`/products/${slug}`)

usePageSeo({
  canonicalPath: () => route.path,
  title: () => t(`zone.${zoneKey.value}.badge`),
  description: () => t(`home.business.${zoneKey.value}.description`),
})

onMounted(async () => {
  try {
    const res = await categoryAPI.list()
    categories.value = res.data.data || []
  } catch (error) {
    console.error('Failed to load categories:', error)
  } finally {
    categoriesLoaded.value = true
  }
  await loadProducts()
})
</script>
