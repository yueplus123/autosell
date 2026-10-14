<template>
  <!-- 二次元赞助专区入口：紧接在 AI 首屏下方的渐变色横幅 -->
  <section class="relative isolate overflow-hidden bg-linear-to-br from-violet-600 via-fuchsia-500 to-orange-400 text-white">
    <span class="pointer-events-none absolute -left-24 -top-24 -z-10 h-96 w-96 rounded-full bg-white/20 blur-3xl"></span>
    <span class="pointer-events-none absolute -bottom-32 right-0 -z-10 h-[28rem] w-[28rem] rounded-full bg-amber-300/40 blur-3xl"></span>
    <div class="pointer-events-none absolute inset-0 -z-10 opacity-25" :style="dotPatternStyle"></div>
    <Star class="pointer-events-none absolute left-[8%] top-12 h-6 w-6 rotate-12 fill-white/80 text-white/80" aria-hidden="true" />
    <Sparkle class="pointer-events-none absolute left-[46%] top-8 h-8 w-8 fill-yellow-200 text-yellow-200 max-md:hidden" aria-hidden="true" />
    <Heart class="pointer-events-none absolute right-[6%] top-16 h-6 w-6 fill-rose-200 text-rose-200" aria-hidden="true" />

    <div class="mx-auto grid max-w-[1200px] items-center gap-12 px-4 py-14 sm:px-6 md:grid-cols-2 md:gap-10 md:py-20">
      <div class="space-y-5 text-center md:text-left">
        <span class="inline-flex items-center gap-2 rounded-full bg-white/20 px-4 py-1.5 text-sm font-bold backdrop-blur-sm">
          <Gamepad2 class="h-4 w-4" aria-hidden="true" />
          {{ t('home.business.game.eyebrow') }}
        </span>
        <h2 class="text-[clamp(32px,5vw,56px)] font-black leading-[1.1] tracking-[-0.02em] drop-shadow-sm">{{ t('home.business.game.title') }}</h2>
        <p class="mx-auto max-w-lg text-[15px] leading-relaxed text-white/90 md:mx-0 md:text-[17px]">{{ t('home.business.game.description') }}</p>
        <div class="flex flex-wrap justify-center gap-2 md:justify-start">
          <span v-for="i in 3" :key="i" class="inline-flex items-center gap-1.5 rounded-full bg-white/20 px-3.5 py-1.5 text-sm font-semibold backdrop-blur-sm">
            <Heart class="h-3.5 w-3.5 fill-current" aria-hidden="true" />
            {{ t(`home.business.game.highlight${i}`) }}
          </span>
        </div>
        <div class="pt-1">
          <RouterLink
            :to="zoneLink"
            class="group inline-flex items-center gap-2 rounded-full bg-white px-7 py-3 text-[15px] font-bold text-fuchsia-600 shadow-xl shadow-fuchsia-900/20 transition hover:scale-105"
          >
            {{ t('home.business.game.cta') }}
            <ArrowRight class="h-4 w-4 transition-transform group-hover:translate-x-1" aria-hidden="true" />
          </RouterLink>
        </div>
      </div>

      <!-- 游戏卡片扇形展开：图片取自后台「二次元赞助」下各子分类的图标 -->
      <div class="flex justify-center">
        <div class="relative flex h-56 w-full max-w-md items-center justify-center sm:h-72">
          <RouterLink
            v-for="card in cards"
            :key="card.key"
            :to="card.link"
            :aria-label="card.label || t('home.business.game.eyebrow')"
            class="absolute flex h-44 w-32 flex-col items-center justify-end overflow-hidden rounded-3xl border-4 border-white/80 bg-white shadow-2xl shadow-fuchsia-950/30 transition duration-300 hover:z-20 hover:-translate-y-4 sm:h-60 sm:w-44"
            :style="{ transform: `translateX(${card.fan.x}) rotate(${card.fan.r}deg)`, zIndex: card.fan.z }"
          >
            <img v-if="card.icon" :src="getImageUrl(card.icon)" :alt="card.label" loading="lazy" class="absolute inset-0 h-full w-full object-cover" />
            <div v-else class="absolute inset-0 grid place-items-center" :class="card.fan.bg">
              <span v-if="card.initial" class="text-5xl font-black text-white/95 drop-shadow sm:text-6xl">{{ card.initial }}</span>
              <component :is="card.fan.icon" v-else class="h-12 w-12 text-white/95" aria-hidden="true" />
            </div>
            <span v-if="card.label" class="relative z-10 w-full bg-linear-to-t from-black/70 to-transparent px-2 pb-3 pt-8 text-center text-sm font-bold text-white">{{ card.label }}</span>
          </RouterLink>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { ArrowRight, Gamepad2, Heart, Sparkle, Sparkles, Star } from 'lucide-vue-next'
import { categoryAPI } from '../../../api'
import { getImageUrl } from '../../../utils/image'
import { useLocalized } from '../../../composables/useProduct'
import type { PublicCategory } from '../../../utils/category'
import { businessZonePath, resolveBusinessLines } from '../../../utils/businessLines'

const { t } = useI18n()
const { getLocalizedText } = useLocalized()

const categories = ref<PublicCategory[]>([])
const gameLine = computed(() => resolveBusinessLines(categories.value).find((line) => line.key === 'game'))
const zoneLink = computed(() => gameLine.value?.link ?? businessZonePath('game'))

const FAN = [
  { x: '-46%', r: -14, z: 1, bg: 'bg-linear-to-br from-sky-400 to-violet-500', icon: Star },
  { x: '-15%', r: -5, z: 3, bg: 'bg-linear-to-br from-pink-400 to-rose-500', icon: Heart },
  { x: '15%', r: 5, z: 4, bg: 'bg-linear-to-br from-amber-300 to-orange-500', icon: Sparkles },
  { x: '46%', r: 14, z: 2, bg: 'bg-linear-to-br from-emerald-300 to-teal-500', icon: Gamepad2 },
]

const cards = computed(() => {
  const children = gameLine.value?.children.slice(0, FAN.length) ?? []
  return FAN.map((fan, i) => {
    const cat = children[i]
    if (!cat) return { key: `placeholder-${i}`, fan, link: zoneLink.value, label: '', icon: '', initial: '' }
    const label = getLocalizedText(cat.name) || String(cat.slug || '')
    return {
      key: `cat-${cat.id}`,
      fan,
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

onMounted(async () => {
  try {
    const res = await categoryAPI.list()
    categories.value = res.data.data || []
  } catch (error) {
    console.error('Failed to load categories:', error)
  }
})
</script>
