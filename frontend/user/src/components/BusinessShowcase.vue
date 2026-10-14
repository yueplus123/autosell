<template>
  <div class="business-showcase">

    <!-- ==================== AI 主视觉：黑底大标题 + 悬浮图标 ==================== -->
    <section class="relative isolate overflow-hidden bg-black pb-16 text-white md:pb-20" :class="heroTopClass">
      <!-- 底部彩色光晕 -->
      <span class="pointer-events-none absolute -bottom-40 -left-32 -z-10 h-[30rem] w-[40rem] rounded-full bg-blue-800/45 blur-[120px]"></span>
      <span class="pointer-events-none absolute -bottom-48 left-1/2 -z-10 h-[28rem] w-[36rem] -translate-x-1/2 rounded-full bg-orange-700/35 blur-[120px]"></span>
      <span class="pointer-events-none absolute -bottom-40 -right-32 -z-10 h-[30rem] w-[36rem] rounded-full bg-emerald-800/35 blur-[120px]"></span>

      <div :class="containerClass">
        <div class="text-center">
          <h2 class="text-5xl font-black tracking-tight sm:text-7xl md:text-8xl">{{ t('home.business.ai.heroTitle') }}</h2>
          <p class="mt-5 text-base text-white/65 sm:text-xl md:text-2xl">{{ aiHighlightsLine }}</p>
        </div>

        <!-- 图标方块 -->
        <div class="relative mx-auto mt-12 max-w-5xl md:mt-16">
          <!-- 上方两个悬浮小方块（仅桌面） -->
          <RouterLink
            v-for="slot in FLOAT_SLOTS"
            :key="`float-${slot.index}`"
            :to="tileLink(aiTiles[slot.index])"
            :aria-label="tileLabel(aiTiles[slot.index])"
            class="hero-tile hero-float absolute top-0 hidden size-[120px] md:grid"
            :class="[slot.position, tileSkin('light')]"
            :style="{ animationDelay: slot.delay, boxShadow: slot.glow }"
          >
            <TileFace :tile="aiTiles[slot.index]" variant="light" />
          </RouterLink>

          <!-- 底部一排 -->
          <div class="flex items-end justify-center gap-3 sm:gap-4 md:gap-6 md:pt-[104px]">
            <RouterLink
              v-for="slot in ROW_SLOTS"
              :key="`row-${slot.index}`"
              :to="tileLink(aiTiles[slot.index])"
              :aria-label="tileLabel(aiTiles[slot.index])"
              class="hero-tile hero-float shrink-0"
              :class="[slot.size, tileSkin(slot.variant)]"
              :style="{ animationDelay: slot.delay, boxShadow: slot.glow }"
            >
              <TileFace :tile="aiTiles[slot.index]" :variant="slot.variant" :large="slot.index === 0" />
            </RouterLink>
          </div>
        </div>

        <div class="mt-16 flex justify-center md:mt-20">
          <RouterLink
            :to="aiLine.link"
            class="group inline-flex items-center gap-2 rounded-full bg-white px-7 py-3.5 text-base font-semibold text-black shadow-[0_0_40px_rgba(255,255,255,0.25)] transition hover:scale-105"
          >
            {{ t('home.business.ai.cta') }}
            <ArrowRight class="h-4 w-4 transition-transform group-hover:translate-x-1" aria-hidden="true" />
          </RouterLink>
        </div>
      </div>
    </section>

    <!-- ==================== 二次元专区：渐变色 ==================== -->
    <section class="relative isolate overflow-hidden bg-linear-to-br from-violet-600 via-fuchsia-500 to-orange-400 py-20 text-white md:py-28">
      <span class="pointer-events-none absolute -left-24 -top-24 -z-10 h-96 w-96 rounded-full bg-white/20 blur-3xl"></span>
      <span class="pointer-events-none absolute -bottom-32 right-0 -z-10 h-[28rem] w-[28rem] rounded-full bg-amber-300/40 blur-3xl"></span>
      <div class="pointer-events-none absolute inset-0 -z-10 opacity-25" :style="dotPatternStyle"></div>
      <Star class="pointer-events-none absolute left-[8%] top-16 h-6 w-6 rotate-12 fill-white/80 text-white/80" aria-hidden="true" />
      <Sparkle class="pointer-events-none absolute left-[46%] top-10 h-8 w-8 fill-yellow-200 text-yellow-200" aria-hidden="true" />
      <Star class="pointer-events-none absolute bottom-14 left-[38%] h-4 w-4 -rotate-12 fill-pink-100 text-pink-100" aria-hidden="true" />
      <Heart class="pointer-events-none absolute right-[6%] top-20 h-6 w-6 fill-rose-200 text-rose-200" aria-hidden="true" />

      <div :class="containerClass">
        <div class="grid items-center gap-14 md:grid-cols-2 md:gap-10">
          <div class="space-y-6 text-center md:text-left">
            <span class="inline-flex items-center gap-2 rounded-full bg-white/20 px-4 py-1.5 text-sm font-bold backdrop-blur-sm">
              <Gamepad2 class="h-4 w-4" aria-hidden="true" />
              {{ t('home.business.game.eyebrow') }}
            </span>
            <h2 class="text-4xl font-black tracking-tight drop-shadow-sm sm:text-5xl md:text-6xl">{{ t('home.business.game.title') }}</h2>
            <p class="mx-auto max-w-lg text-base leading-relaxed text-white/90 md:mx-0 md:text-lg">{{ t('home.business.game.description') }}</p>
            <div class="flex flex-wrap justify-center gap-2 md:justify-start">
              <span v-for="i in 3" :key="i" class="inline-flex items-center gap-1.5 rounded-full bg-white/20 px-3.5 py-1.5 text-sm font-semibold backdrop-blur-sm">
                <Heart class="h-3.5 w-3.5 fill-current" aria-hidden="true" />
                {{ t(`home.business.game.highlight${i}`) }}
              </span>
            </div>
            <div class="pt-2">
              <RouterLink
                :to="gameLine.link"
                class="group inline-flex items-center gap-2 rounded-full bg-white px-7 py-3.5 text-base font-bold text-fuchsia-600 shadow-xl shadow-fuchsia-900/20 transition hover:scale-105"
              >
                {{ t('home.business.game.cta') }}
                <ArrowRight class="h-4 w-4 transition-transform group-hover:translate-x-1" aria-hidden="true" />
              </RouterLink>
            </div>
          </div>

          <!-- 游戏卡片扇形展开 -->
          <div class="flex justify-center">
            <div class="relative flex h-60 w-full max-w-md items-center justify-center sm:h-72">
              <RouterLink
                v-for="(card, i) in gameCards"
                :key="card.key"
                :to="card.link"
                :aria-label="card.label"
                class="game-card absolute flex h-48 w-36 flex-col items-center justify-end overflow-hidden rounded-3xl border-4 border-white/80 bg-white shadow-2xl shadow-fuchsia-950/30 transition duration-300 hover:z-20 hover:-translate-y-4 sm:h-60 sm:w-44"
                :style="{ transform: `translateX(${FAN[i].x}) rotate(${FAN[i].r}deg)`, zIndex: FAN[i].z }"
              >
                <img v-if="card.icon" :src="getImageUrl(card.icon)" :alt="card.label" loading="lazy" class="absolute inset-0 h-full w-full object-cover" />
                <div v-else class="absolute inset-0 grid place-items-center" :class="FAN[i].bg">
                  <span v-if="card.initial" class="text-6xl font-black text-white/95 drop-shadow">{{ card.initial }}</span>
                  <component :is="FAN[i].icon" v-else class="h-14 w-14 text-white/95" aria-hidden="true" />
                </div>
                <span v-if="card.label" class="relative z-10 w-full bg-linear-to-t from-black/70 to-transparent px-2 pb-3 pt-8 text-center text-sm font-bold text-white">{{ card.label }}</span>
              </RouterLink>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- ==================== 各业务线热门商品 ==================== -->
    <template v-if="showProducts">
      <section
        v-for="line in linesWithProducts"
        :key="`products-${line.key}`"
        class="relative z-10 py-8"
      >
        <div :class="containerClass">
          <div class="mb-5 flex items-end justify-between gap-4">
            <div class="flex items-center gap-3">
              <span class="grid h-9 w-9 place-items-center rounded-xl text-white" :class="line.key === 'ai' ? 'bg-black' : 'bg-linear-to-br from-violet-600 via-fuchsia-500 to-orange-400'">
                <component :is="line.key === 'ai' ? Sparkles : Gamepad2" class="h-[18px] w-[18px]" aria-hidden="true" />
              </span>
              <h2 :class="headingClass">{{ t(`home.business.${line.key}.hotTitle`) }}</h2>
            </div>
            <RouterLink :to="line.link" class="text-sm font-semibold text-muted-foreground transition-colors hover:text-foreground">
              {{ t('home.featured.viewAll') }}
            </RouterLink>
          </div>
          <div :class="productGridClass">
            <template v-for="(product, idx) in productsByKey[line.key]" :key="product.id">
              <slot name="product" :product="product" :index="idx" />
            </template>
          </div>
        </div>
      </section>
    </template>
  </div>
</template>

<script setup lang="ts">
import { computed, defineComponent, h, onMounted, reactive, ref, type PropType } from 'vue'
import { useI18n } from 'vue-i18n'
import {
  ArrowRight, Bot, Code, Gamepad2, Heart, Image as ImageIcon, MessageSquare, Mic, Sparkle, Sparkles, Star, Wand2,
} from 'lucide-vue-next'
import { categoryAPI, productAPI } from '../api'
import { getImageUrl } from '../utils/image'
import { useLocalized } from '../composables/useProduct'
import type { PublicCategory } from '../utils/category'
import {
  buildHeroTiles, businessZonePath, resolveBusinessLines, type BusinessLineKey, type HeroTile, type ResolvedBusinessLine,
} from '../utils/businessLines'

const props = withDefaults(defineProps<{
  showProducts?: boolean
  productLimit?: number
  containerClass?: string
  headingClass?: string
  productGridClass?: string
  /** AI 主视觉顶部留白：classic 导航栏是 fixed，需要更大的留白 */
  heroTopClass?: string
}>(), {
  showProducts: false,
  productLimit: 8,
  containerClass: 'container mx-auto px-4',
  headingClass: 'theme-section-heading text-2xl md:text-3xl',
  productGridClass: 'grid grid-cols-2 gap-3 md:grid-cols-3 md:gap-4 lg:grid-cols-4',
  heroTopClass: 'pt-28 md:pt-32',
})

const { t } = useI18n()
const { getLocalizedText } = useLocalized()

// ==================== 数据 ====================
const categories = ref<PublicCategory[]>([])
const lines = computed(() => resolveBusinessLines(categories.value))
const emptyLine = (key: BusinessLineKey): ResolvedBusinessLine => ({ key, category: null, children: [], link: businessZonePath(key) })
const aiLine = computed(() => lines.value.find((l) => l.key === 'ai') ?? emptyLine('ai'))
const gameLine = computed(() => lines.value.find((l) => l.key === 'game') ?? emptyLine('game'))

const catName = (cat: PublicCategory) => getLocalizedText(cat.name) || String(cat.slug || '')

// ==================== AI 主视觉 ====================
const aiHighlightsLine = computed(() => [1, 2, 3].map((i) => t(`home.business.ai.highlight${i}`)).filter(Boolean).join(' · '))
const aiTiles = computed(() => buildHeroTiles(aiLine.value.children, 7))

// 0 = 正中最大；1/2 = 两侧；3/4 = 最外侧；5/6 = 上方悬浮
const ROW_SLOTS = [
  { index: 3, variant: 'light', size: 'hidden sm:grid size-[72px] md:size-[164px]', delay: '0.6s', glow: '0 0 60px rgba(239,68,68,0.25)' },
  { index: 1, variant: 'light', size: 'grid size-[92px] md:size-[194px]', delay: '0.2s', glow: '0 0 70px rgba(59,130,246,0.45)' },
  { index: 0, variant: 'dark', size: 'grid size-[120px] md:size-[252px]', delay: '0s', glow: '0 0 80px rgba(255,255,255,0.08)' },
  { index: 2, variant: 'dark', size: 'grid size-[92px] md:size-[194px]', delay: '0.4s', glow: '0 0 70px rgba(249,115,22,0.3)' },
  { index: 4, variant: 'light', size: 'hidden sm:grid size-[72px] md:size-[164px]', delay: '0.8s', glow: '0 0 60px rgba(148,163,184,0.25)' },
] as const
const FLOAT_SLOTS = [
  { index: 5, position: 'left-[calc(50%-360px)]', delay: '1s', glow: '0 0 50px rgba(59,130,246,0.5)' },
  { index: 6, position: 'left-[calc(50%+240px)]', delay: '1.3s', glow: '0 0 50px rgba(59,130,246,0.45)' },
] as const

const FALLBACK_ICONS = [Sparkles, Bot, MessageSquare, ImageIcon, Code, Mic, Wand2]
const FALLBACK_COLORS = ['text-white', 'text-blue-500', 'text-white', 'text-violet-500', 'text-emerald-500', 'text-rose-500', 'text-amber-500']

type TileVariant = 'light' | 'dark'
const tileSkin = (variant: TileVariant) => variant === 'dark'
  ? 'place-items-center rounded-[24%] border border-white/15 bg-linear-to-b from-neutral-800 to-black'
  : 'place-items-center rounded-[24%] bg-linear-to-b from-white to-slate-200'

const tileLink = (tile: HeroTile | undefined) =>
  tile?.category?.slug ? businessZonePath('ai', String(tile.category.slug)) : aiLine.value.link
const tileLabel = (tile: HeroTile | undefined) => (tile?.category ? catName(tile.category) : t('home.business.ai.eyebrow'))

const TileFace = defineComponent({
  props: {
    tile: { type: Object as PropType<HeroTile | undefined>, default: undefined },
    variant: { type: String as PropType<TileVariant>, default: 'light' },
    large: Boolean,
  },
  setup(p) {
    return () => {
      const tile = p.tile
      const icon = tile?.category?.icon ? String(tile.category.icon) : ''
      if (icon) {
        return h('img', {
          src: getImageUrl(icon),
          alt: tile?.category ? catName(tile.category) : '',
          loading: 'lazy',
          class: 'h-[58%] w-[58%] object-contain',
        })
      }
      const idx = tile?.fallbackIndex ?? 0
      const color = p.variant === 'dark' ? 'text-white' : FALLBACK_COLORS[idx % FALLBACK_COLORS.length]
      if (tile?.category) {
        return h('span', { class: ['font-black leading-none', p.large ? 'text-5xl md:text-8xl' : 'text-3xl md:text-6xl', color] }, catName(tile.category).slice(0, 1))
      }
      return h(FALLBACK_ICONS[idx % FALLBACK_ICONS.length], { class: ['h-1/2 w-1/2', color], strokeWidth: 1.6, 'aria-hidden': 'true' })
    }
  },
})

// ==================== 二次元专区 ====================
const FAN = [
  { x: '-46%', r: -14, z: 1, bg: 'bg-linear-to-br from-sky-400 to-violet-500', icon: Star },
  { x: '-15%', r: -5, z: 3, bg: 'bg-linear-to-br from-pink-400 to-rose-500', icon: Heart },
  { x: '15%', r: 5, z: 4, bg: 'bg-linear-to-br from-amber-300 to-orange-500', icon: Sparkles },
  { x: '46%', r: 14, z: 2, bg: 'bg-linear-to-br from-emerald-300 to-teal-500', icon: Gamepad2 },
]

const gameCards = computed(() => {
  const children = gameLine.value.children.slice(0, FAN.length)
  return FAN.map((_, i) => {
    const cat = children[i]
    if (!cat) return { key: `placeholder-${i}`, link: gameLine.value.link, label: '', icon: '', initial: '' }
    const label = catName(cat)
    return {
      key: `cat-${cat.id}`,
      link: businessZonePath('game', String(cat.slug)),
      label,
      icon: cat.icon ? String(cat.icon) : '',
      initial: label.slice(0, 1),
    }
  })
})

const dotPatternStyle = {
  backgroundImage: 'radial-gradient(rgba(255,255,255,0.5) 1.2px, transparent 1.2px)',
  backgroundSize: '22px 22px',
}

// ==================== 热门商品 ====================
const productsByKey = reactive<Record<BusinessLineKey, any[]>>({ ai: [], game: [] })
const linesWithProducts = computed(() => lines.value.filter((line) => line.category && productsByKey[line.key].length > 0))

const loadProducts = async () => {
  await Promise.all(
    lines.value.map(async (line) => {
      if (!line.category) return
      try {
        const res = await productAPI.list({ page: 1, page_size: props.productLimit, category_id: line.category.id })
        productsByKey[line.key] = res.data.data || []
      } catch (error) {
        console.error('Failed to load business line products:', error)
      }
    }),
  )
}

onMounted(async () => {
  try {
    const res = await categoryAPI.list()
    categories.value = res.data.data || []
  } catch (error) {
    console.error('Failed to load categories:', error)
  }
  if (props.showProducts) await loadProducts()
})
</script>

<style scoped>
.hero-tile {
  transition: transform 300ms ease;
  -webkit-box-reflect: below 14px linear-gradient(transparent 62%, rgba(255, 255, 255, 0.16));
}
.hero-tile:hover {
  transform: translateY(-6px) scale(1.04);
}
.hero-float {
  animation: hero-bob 6s ease-in-out infinite;
}
.hero-float:hover {
  animation-play-state: paused;
}
@keyframes hero-bob {
  0%, 100% { translate: 0 0; }
  50% { translate: 0 -8px; }
}
@media (prefers-reduced-motion: reduce) {
  .hero-float { animation: none; }
  .hero-tile, .game-card { transition: none; }
}
</style>
