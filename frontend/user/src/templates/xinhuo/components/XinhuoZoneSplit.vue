<template>
  <!-- 第二屏：左右各一半，左 AI 专区、右二次元赞助专区 -->
  <section id="zones" class="grid scroll-mt-14 md:grid-cols-2">
    <!-- 左：AI 专区（深色科技风） -->
    <RouterLink
      :to="aiLink"
      class="group relative isolate flex min-h-[380px] flex-col justify-between overflow-hidden bg-[#05070d] p-7 text-white sm:p-10 md:min-h-[460px] lg:p-14"
    >
      <div class="pointer-events-none absolute inset-0 -z-10 opacity-60" :style="gridStyle"></div>
      <span class="pointer-events-none absolute -left-24 -top-20 -z-10 h-80 w-80 rounded-full bg-cyan-500/25 blur-3xl transition-opacity duration-500 group-hover:opacity-100 md:opacity-70"></span>
      <span class="pointer-events-none absolute -bottom-24 right-0 -z-10 h-80 w-80 rounded-full bg-violet-600/30 blur-3xl"></span>

      <div class="space-y-4">
        <span class="inline-flex items-center gap-2 rounded-full border border-cyan-400/30 bg-cyan-400/10 px-3 py-1 font-mono text-xs uppercase tracking-[0.2em] text-cyan-300">
          <Sparkles class="h-3.5 w-3.5" aria-hidden="true" />
          {{ t('home.business.ai.eyebrow') }}
        </span>
        <h2 class="text-[clamp(30px,3.6vw,48px)] font-black leading-[1.12] tracking-[-0.02em]">{{ t('home.business.ai.title') }}</h2>
        <p class="max-w-md text-[15px] leading-relaxed text-white/70">{{ t('home.business.ai.description') }}</p>
        <div class="flex flex-wrap gap-2">
          <span v-for="i in 3" :key="i" class="inline-flex items-center gap-1.5 rounded-md border border-white/10 bg-white/5 px-3 py-1.5 text-[13px] text-white/85">
            <Check class="h-3.5 w-3.5 text-cyan-300" aria-hidden="true" />
            {{ t(`home.business.ai.highlight${i}`) }}
          </span>
        </div>
      </div>

      <div class="mt-8 flex flex-wrap items-end justify-between gap-6">
        <span class="inline-flex items-center gap-2 rounded-full bg-white px-6 py-3 text-[15px] font-semibold text-black shadow-[0_0_36px_rgba(34,211,238,0.35)] transition group-hover:gap-3">
          {{ t('home.business.ai.cta') }}
          <ArrowRight class="h-4 w-4" aria-hidden="true" />
        </span>
        <div class="flex items-end gap-3" aria-hidden="true">
          <span v-for="(icon, i) in AI_ICONS" :key="icon.src" class="grid place-items-center rounded-[24%] transition-transform duration-500 group-hover:-translate-y-1.5" :class="[icon.tone === 'dark' ? 'border border-white/20 bg-black' : 'bg-[#f2f2f5]', i === 1 ? 'h-16 w-16 sm:h-20 sm:w-20' : 'h-12 w-12 sm:h-14 sm:w-14']" :style="{ transitionDelay: `${i * 60}ms` }">
            <img :src="icon.src" alt="" class="h-[56%] w-[56%] object-contain" :class="icon.tone === 'dark' ? 'invert' : ''" />
          </span>
        </div>
      </div>
    </RouterLink>

    <!-- 右：二次元赞助专区（渐变色） -->
    <RouterLink
      :to="gameLink"
      class="group relative isolate flex min-h-[380px] flex-col justify-between overflow-hidden bg-linear-to-br from-violet-600 via-fuchsia-500 to-orange-400 p-7 text-white sm:p-10 md:min-h-[460px] lg:p-14"
    >
      <div class="pointer-events-none absolute inset-0 -z-10 opacity-25" :style="dotStyle"></div>
      <span class="pointer-events-none absolute -right-20 -top-24 -z-10 h-80 w-80 rounded-full bg-white/20 blur-3xl"></span>
      <span class="pointer-events-none absolute -bottom-28 -left-10 -z-10 h-80 w-80 rounded-full bg-amber-300/40 blur-3xl"></span>
      <Star class="pointer-events-none absolute right-[10%] top-10 h-6 w-6 rotate-12 fill-white/80 text-white/80" aria-hidden="true" />
      <Sparkle class="pointer-events-none absolute right-[26%] top-20 h-5 w-5 fill-yellow-200 text-yellow-200" aria-hidden="true" />
      <Heart class="pointer-events-none absolute right-[16%] top-32 h-5 w-5 fill-rose-200 text-rose-200" aria-hidden="true" />

      <div class="space-y-4">
        <span class="inline-flex items-center gap-2 rounded-full bg-white/20 px-3.5 py-1 text-xs font-bold tracking-wide backdrop-blur-sm">
          <Gamepad2 class="h-3.5 w-3.5" aria-hidden="true" />
          {{ t('home.business.game.eyebrow') }}
        </span>
        <h2 class="text-[clamp(30px,3.6vw,48px)] font-black leading-[1.12] tracking-[-0.02em] drop-shadow-sm">{{ t('home.business.game.title') }}</h2>
        <p class="max-w-md text-[15px] leading-relaxed text-white/90">{{ t('home.business.game.description') }}</p>
        <div class="flex flex-wrap gap-2">
          <span v-for="i in 3" :key="i" class="inline-flex items-center gap-1.5 rounded-full bg-white/20 px-3 py-1.5 text-[13px] font-semibold backdrop-blur-sm">
            <Heart class="h-3 w-3 fill-current" aria-hidden="true" />
            {{ t(`home.business.game.highlight${i}`) }}
          </span>
        </div>
      </div>

      <div class="mt-8 flex flex-wrap items-end justify-between gap-6">
        <span class="inline-flex items-center gap-2 rounded-full bg-white px-6 py-3 text-[15px] font-bold text-fuchsia-600 shadow-xl shadow-fuchsia-900/20 transition group-hover:gap-3">
          {{ t('home.business.game.cta') }}
          <ArrowRight class="h-4 w-4" aria-hidden="true" />
        </span>
        <!-- 游戏子分类图标（后台上传），没上传时显示游戏名首字 -->
        <div class="flex items-end gap-3" aria-hidden="true">
          <span v-for="(game, i) in gameTiles" :key="game.key" class="grid place-items-center overflow-hidden rounded-[24%] border-2 border-white/70 bg-white/20 font-black backdrop-blur-sm transition-transform duration-500 group-hover:-translate-y-1.5" :class="i === 1 ? 'h-16 w-16 text-2xl sm:h-20 sm:w-20 sm:text-3xl' : 'h-12 w-12 text-lg sm:h-14 sm:w-14 sm:text-xl'" :style="{ transitionDelay: `${i * 60}ms` }">
            <img v-if="game.icon" :src="getImageUrl(game.icon)" alt="" class="h-full w-full object-cover" />
            <template v-else-if="game.initial">{{ game.initial }}</template>
            <Gamepad2 v-else class="h-1/2 w-1/2" />
          </span>
        </div>
      </div>
    </RouterLink>
  </section>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { ArrowRight, Check, Gamepad2, Heart, Sparkle, Sparkles, Star } from 'lucide-vue-next'
import { categoryAPI } from '../../../api'
import { getImageUrl } from '../../../utils/image'
import { useLocalized } from '../../../composables/useProduct'
import type { PublicCategory } from '../../../utils/category'
import { businessZonePath, resolveBusinessLines } from '../../../utils/businessLines'
import { HERO_ICON_BASE } from '../heroTiles'

const { t } = useI18n()
const { getLocalizedText } = useLocalized()

const AI_ICONS = [
  { src: `${HERO_ICON_BASE}gemini.png`, tone: 'light' },
  { src: `${HERO_ICON_BASE}chatgpt.png`, tone: 'dark' },
  { src: `${HERO_ICON_BASE}grok.png`, tone: 'dark' },
] as const

const categories = ref<PublicCategory[]>([])
const lines = computed(() => resolveBusinessLines(categories.value))
const aiLink = computed(() => lines.value.find((l) => l.key === 'ai')?.link ?? businessZonePath('ai'))
const gameLink = computed(() => lines.value.find((l) => l.key === 'game')?.link ?? businessZonePath('game'))

const gameTiles = computed(() => {
  const children = lines.value.find((l) => l.key === 'game')?.children.slice(0, 3) ?? []
  return [0, 1, 2].map((i) => {
    const cat = children[i]
    if (!cat) return { key: `placeholder-${i}`, icon: '', initial: '' }
    const name = getLocalizedText(cat.name) || String(cat.slug || '')
    return { key: `cat-${cat.id}`, icon: cat.icon ? String(cat.icon) : '', initial: name.slice(0, 1) }
  })
})

const gridStyle = {
  backgroundImage:
    'linear-gradient(rgba(34,211,238,0.12) 1px, transparent 1px), linear-gradient(90deg, rgba(34,211,238,0.12) 1px, transparent 1px)',
  backgroundSize: '40px 40px',
  maskImage: 'radial-gradient(ellipse at 30% 40%, black 20%, transparent 75%)',
  WebkitMaskImage: 'radial-gradient(ellipse at 30% 40%, black 20%, transparent 75%)',
}
const dotStyle = {
  backgroundImage: 'radial-gradient(rgba(255,255,255,0.5) 1.2px, transparent 1.2px)',
  backgroundSize: '22px 22px',
}

onMounted(async () => {
  try {
    const res = await categoryAPI.list()
    categories.value = res.data.data || []
  } catch (error) {
    console.error('Failed to load categories:', error)
  }
})
</script>
