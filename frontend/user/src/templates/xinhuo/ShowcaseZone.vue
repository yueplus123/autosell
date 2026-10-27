<template>
  <div class="bg-[color:var(--xh-page)] text-[color:var(--xh-ink)]">
    <!-- 官网专区页（只展示业务，不显示价格、不能下单；与代理分站的专区页 ResellerZone 完全分开） -->
    <!-- ==================== 专区首屏 ==================== -->
    <!-- AI：黑底科技风，复用首页首屏的网格 / 扫描光 -->
    <section v-if="isAi" class="xh-hero overflow-hidden text-[#f5f5f7]">
      <div class="xh-scan" aria-hidden="true"></div>
      <div class="mx-auto grid max-w-[1200px] items-center gap-10 px-4 pb-14 pt-12 sm:px-6 md:grid-cols-[1.15fr_1fr] md:pb-20 md:pt-16">
        <div class="space-y-5">
          <span class="inline-flex items-center gap-2 rounded-full border border-cyan-400/30 bg-cyan-400/10 px-3 py-1 font-mono text-[11px] uppercase tracking-[0.24em] text-cyan-300 sm:text-xs">
            <span class="h-1.5 w-1.5 animate-pulse rounded-full bg-cyan-300"></span>
            AI ZONE
          </span>
          <h1 class="xh-hero-title text-[clamp(36px,5.6vw,68px)] font-black leading-[1.08] tracking-[-0.02em]">{{ t('home.business.ai.title') }}</h1>
          <p class="max-w-xl text-[15px] leading-relaxed text-[#a1a1a6] sm:text-[17px]">{{ t('zonePage.showcase.ai.description') }}</p>
          <div class="flex flex-wrap gap-2">
            <span v-for="i in 3" :key="i" class="inline-flex items-center gap-1.5 rounded-md border border-white/10 bg-white/5 px-3 py-1.5 text-[13px] text-white/85">
              <Check class="h-3.5 w-3.5 text-cyan-300" aria-hidden="true" />
              {{ t(`zonePage.showcase.ai.highlight${i}`) }}
            </span>
          </div>
          <a href="#zone-shop" class="inline-flex items-center gap-2 rounded-full bg-white px-6 py-3 text-[15px] font-semibold text-black shadow-[0_0_36px_rgba(34,211,238,0.35)] transition hover:scale-105" @click.prevent="scrollToShop">
            {{ t('zonePage.browse') }}
            <ArrowDown class="h-4 w-4" aria-hidden="true" />
          </a>
        </div>
        <!-- AI 业务图标 -->
        <BrandIconWall :tiles="ZONE_AI_TILES" variant="zone" />
      </div>
    </section>

    <!-- 二次元：粉紫橙渐变 + 游戏卡片 -->
    <section v-else class="relative isolate overflow-hidden bg-linear-to-br from-violet-600 via-fuchsia-500 to-orange-400 text-white">
      <div class="pointer-events-none absolute inset-0 -z-10 opacity-25" :style="dotStyle"></div>
      <span class="pointer-events-none absolute -left-24 -top-24 -z-10 h-96 w-96 rounded-full bg-white/20 blur-3xl"></span>
      <span class="pointer-events-none absolute -bottom-32 right-0 -z-10 h-[28rem] w-[28rem] rounded-full bg-amber-300/40 blur-3xl"></span>
      <Star class="pointer-events-none absolute left-[7%] top-10 h-6 w-6 rotate-12 fill-white/80 text-white/80" aria-hidden="true" />
      <Sparkle class="pointer-events-none absolute left-[44%] top-8 h-8 w-8 fill-yellow-200 text-yellow-200 max-md:hidden" aria-hidden="true" />
      <Heart class="pointer-events-none absolute right-[6%] top-14 h-6 w-6 fill-rose-200 text-rose-200" aria-hidden="true" />

      <div class="mx-auto grid max-w-[1200px] items-center gap-12 px-4 pb-14 pt-12 sm:px-6 md:grid-cols-2 md:pb-20 md:pt-16">
        <div class="space-y-5 text-center md:text-left">
          <span class="inline-flex items-center gap-2 rounded-full bg-white/20 px-4 py-1.5 text-sm font-bold backdrop-blur-sm">
            <Gamepad2 class="h-4 w-4" aria-hidden="true" />
            {{ t('zone.game.badge') }}
          </span>
          <h1 class="text-[clamp(34px,5.4vw,64px)] font-black leading-[1.08] tracking-[-0.02em] drop-shadow-sm">{{ t('home.business.game.title') }}</h1>
          <p class="mx-auto max-w-xl text-[15px] leading-relaxed text-white/90 sm:text-[17px] md:mx-0">{{ t('zonePage.showcase.game.description') }}</p>
          <div class="flex flex-wrap justify-center gap-2 md:justify-start">
            <span v-for="i in 3" :key="i" class="inline-flex items-center gap-1.5 rounded-full bg-white/20 px-3.5 py-1.5 text-sm font-semibold backdrop-blur-sm">
              <Heart class="h-3.5 w-3.5 fill-current" aria-hidden="true" />
              {{ t(`zonePage.showcase.game.highlight${i}`) }}
            </span>
          </div>
          <a href="#zone-shop" class="inline-flex items-center gap-2 rounded-full bg-white px-6 py-3 text-[15px] font-bold text-fuchsia-600 shadow-xl shadow-fuchsia-900/20 transition hover:scale-105" @click.prevent="scrollToShop">
            {{ t('zonePage.browse') }}
            <ArrowDown class="h-4 w-4" aria-hidden="true" />
          </a>
        </div>
        <!-- 二次元业务图标 -->
        <BrandIconWall :tiles="ZONE_ACG_TILES" variant="zone" />
      </div>
    </section>

    <!-- ==================== 筛选条（吸顶）：子分类 + 搜索 ==================== -->
    <div id="zone-shop" class="sticky top-14 z-40 scroll-mt-14 border-b backdrop-blur-xl" :class="isAi ? 'border-white/10 bg-black/85' : 'border-fuchsia-200/60 bg-white/85'">
      <div class="mx-auto flex max-w-[1200px] flex-col gap-3 px-4 py-3 sm:px-6 md:flex-row md:items-center">
        <div class="xh-scroll-x -mx-4 flex flex-1 gap-2 overflow-x-auto px-4 md:mx-0 md:flex-wrap md:overflow-visible md:px-0">
          <button
            v-for="opt in filterOptions"
            :key="opt.key"
            type="button"
            class="h-9 flex-none whitespace-nowrap rounded-full border px-4 text-[13px] font-medium transition-colors sm:text-sm"
            :class="chipClass(opt.active)"
            :aria-pressed="opt.active"
            @click="selectChild(opt.slug)"
          >{{ opt.label }}</button>
        </div>
      </div>
    </div>

    <!-- ==================== 商品 ==================== -->
    <section class="mx-auto max-w-[1200px] px-4 py-8 sm:px-6 sm:py-12">
      <div class="mb-4 flex flex-wrap items-end justify-between gap-3 sm:mb-6">
        <h2 class="text-[24px] font-black sm:text-[32px]">{{ listTitle }}</h2>
        <span v-if="!loading && total > 0" class="text-sm text-[color:var(--xh-muted)]">{{ t('zonePage.count', { n: total }) }}</span>
      </div>

      <div v-if="loading" class="grid grid-cols-1 gap-2.5 sm:grid-cols-[repeat(auto-fill,minmax(260px,1fr))] sm:gap-5">
        <div v-for="n in 8" :key="`sk-${n}`" class="h-[86px] animate-pulse rounded-2xl bg-white sm:h-[236px] sm:rounded-[20px]"></div>
      </div>
      <div v-else-if="products.length" class="grid grid-cols-1 gap-2.5 sm:grid-cols-[repeat(auto-fill,minmax(260px,1fr))] sm:gap-5">
        <XinhuoProductCard v-for="(product, idx) in products" :key="product.id" :product="product" :index="idx" :tone="zoneKey" />
      </div>
      <div v-else class="flex flex-col items-center gap-3 rounded-[24px] bg-white px-6 py-16 text-center">
        <span class="grid h-14 w-14 place-items-center rounded-2xl text-white" :class="isAi ? 'bg-black' : 'bg-linear-to-br from-violet-600 via-fuchsia-500 to-orange-400'">
          <component :is="isAi ? Sparkles : Gamepad2" class="h-6 w-6" aria-hidden="true" />
        </span>
        <p class="text-lg font-bold">{{ appliedSearch ? t('products.empty') : t('zone.empty') }}</p>
        <p class="text-sm text-[color:var(--xh-muted)]">{{ appliedSearch ? t('zonePage.emptySearch') : t('zone.emptyDesc') }}</p>
        <button v-if="appliedSearch || selectedChild" type="button" class="mt-1 rounded-full border px-4 py-1.5 text-sm hover:bg-black/5" @click="resetFilters">{{ t('zonePage.reset') }}</button>
      </div>

      <!-- 分页 -->
      <nav v-if="totalPages > 1" class="mt-8 flex items-center justify-center gap-2" :aria-label="t('zonePage.pagination')">
        <button type="button" class="grid h-10 w-10 place-items-center rounded-full border bg-white disabled:opacity-40" :disabled="page <= 1" :aria-label="t('common.previousBanner')" @click="goPage(page - 1)">
          <ChevronLeft class="h-4 w-4" />
        </button>
        <button
          v-for="p in pageWindow"
          :key="p"
          type="button"
          class="h-10 min-w-10 rounded-full px-3 text-sm font-semibold"
          :class="p === page ? (isAi ? 'bg-black text-white' : 'bg-fuchsia-500 text-white') : 'border bg-white'"
          @click="goPage(p)"
        >{{ p }}</button>
        <button type="button" class="grid h-10 w-10 place-items-center rounded-full border bg-white disabled:opacity-40" :disabled="page >= totalPages" :aria-label="t('common.nextBanner')" @click="goPage(page + 1)">
          <ChevronRight class="h-4 w-4" />
        </button>
      </nav>
    </section>

  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { ArrowDown, Check, ChevronLeft, ChevronRight, Gamepad2, Heart, Sparkle, Sparkles, Star } from 'lucide-vue-next'
import { categoryAPI, productAPI } from '../../api'
import { useLocalized } from '../../composables/useProduct'
import { usePageSeo } from '../../composables/usePageSeo'
import type { PublicCategory } from '../../utils/category'
import { businessZonePath, isBusinessLineKey, resolveBusinessLines, type BusinessLineKey } from '../../utils/businessLines'
import XinhuoProductCard from './components/XinhuoProductCard.vue'
import { ZONE_ACG_TILES, ZONE_AI_TILES } from './brandIcons'
import BrandIconWall from './components/BrandIconWall.vue'

const route = useRoute()
const router = useRouter()
const { t } = useI18n()
const { getLocalizedText } = useLocalized()

const PAGE_SIZE = 24

const zoneKey = computed<BusinessLineKey>(() => {
  const key = String(route.params.key || '')
  return isBusinessLineKey(key) ? key : 'ai'
})
const isAi = computed(() => zoneKey.value === 'ai')

usePageSeo({
  canonicalPath: () => route.path,
  title: () => t(`zone.${zoneKey.value}.badge`),
  description: () => t(`zonePage.showcase.${zoneKey.value}.description`),
})

// ==================== 分类 ====================
const categories = ref<PublicCategory[]>([])
const categoriesLoaded = ref(false)
const line = computed(() => resolveBusinessLines(categories.value).find((l) => l.key === zoneKey.value) ?? null)
const children = computed(() => line.value?.children ?? [])
const catName = (cat: PublicCategory) => getLocalizedText(cat.name) || String(cat.slug || '')

const selectedSlug = computed(() => (typeof route.query.c === 'string' ? route.query.c : ''))
const selectedChild = computed(() => children.value.find((c) => String(c.slug) === selectedSlug.value) ?? null)

const filterOptions = computed(() => [
  { key: 'all', slug: '', label: t('zone.all'), active: !selectedChild.value },
  ...children.value.map((c) => ({ key: String(c.id), slug: String(c.slug), label: catName(c), active: selectedChild.value?.id === c.id })),
])

const chipClass = (active: boolean) => {
  if (isAi.value) return active ? 'border-cyan-300 bg-cyan-300 text-black' : 'border-white/15 text-white/80 hover:border-white/40 hover:text-white'
  return active ? 'border-transparent bg-linear-to-r from-violet-600 via-fuchsia-500 to-orange-400 text-white' : 'border-fuchsia-200 bg-white text-[color:var(--xh-ink)] hover:border-fuchsia-400'
}

const listTitle = computed(() => (selectedChild.value ? catName(selectedChild.value) : t(`zonePage.${zoneKey.value}.listTitle`)))

// ==================== 首屏图标（官网专区：AI / 二次元业务图标，见 brandIcons.ts） ====================
const dotStyle = {
  backgroundImage: 'radial-gradient(rgba(255,255,255,0.5) 1.2px, transparent 1.2px)',
  backgroundSize: '22px 22px',
}

// ==================== 商品 ====================
const products = ref<any[]>([])
const loading = ref(true)
const page = ref(1)
const total = ref(0)
const totalPages = ref(0)
const appliedSearch = computed(() => (typeof route.query.q === 'string' ? route.query.q.trim() : ''))

const pageWindow = computed(() => {
  const end = Math.min(totalPages.value, Math.max(page.value + 2, 5))
  const start = Math.max(1, end - 4)
  return Array.from({ length: end - start + 1 }, (_, i) => start + i)
})

const loadProducts = async () => {
  const categoryId = selectedChild.value?.id ?? line.value?.category?.id
  if (!categoryId) {
    products.value = []
    total.value = 0
    totalPages.value = 0
    loading.value = false
    return
  }
  loading.value = true
  try {
    const params: Record<string, unknown> = { page: page.value, page_size: PAGE_SIZE, category_id: categoryId }
    if (appliedSearch.value) params.search = appliedSearch.value
    const res = await productAPI.list(params)
    products.value = res.data.data || []
    total.value = Number(res.data.pagination?.total || products.value.length)
    totalPages.value = Number(res.data.pagination?.total_page || 0)
  } catch (error) {
    console.error('Failed to load zone products:', error)
    products.value = []
  } finally {
    loading.value = false
  }
}

const updateQuery = (next: { c?: string; q?: string }) => {
  const query: Record<string, string> = {}
  const c = next.c ?? selectedSlug.value
  const q = next.q ?? appliedSearch.value
  if (c) query.c = c
  if (q) query.q = q
  router.replace({ path: businessZonePath(zoneKey.value), query })
}

const selectChild = (slug: string) => updateQuery({ c: slug })
const resetFilters = () => {
  router.replace({ path: businessZonePath(zoneKey.value) })
}
const goPage = (p: number) => {
  if (p < 1 || p > totalPages.value || p === page.value) return
  page.value = p
  loadProducts()
  scrollToShop()
}
const scrollToShop = () => {
  document.getElementById('zone-shop')?.scrollIntoView({ behavior: 'smooth', block: 'start' })
}

watch([zoneKey, selectedSlug, appliedSearch], () => {
  if (!categoriesLoaded.value) return
  page.value = 1
  loadProducts()
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
