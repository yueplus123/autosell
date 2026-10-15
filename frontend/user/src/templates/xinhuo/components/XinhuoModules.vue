<template>
  <div class="mx-auto max-w-[1200px] space-y-12 px-4 py-12 sm:space-y-16 sm:px-6 sm:py-16">
    <section v-for="mod in homeModules" :key="mod.key">
      <div class="mb-4 sm:mb-6">
        <h2 class="text-[26px] font-black leading-tight tracking-[-0.01em] sm:text-[36px]">{{ t(mod.titleKey) }}</h2>
        <p v-if="mod.descKey" class="mt-1.5 text-[14px] text-[color:var(--xh-muted)] sm:text-[16px]">{{ t(mod.descKey) }}</p>
      </div>

      <!-- 方块：手机 2 列，桌面 4 列 -->
      <div v-if="mod.layout === 'square'" class="grid grid-cols-2 gap-3 sm:gap-5 lg:grid-cols-4">
        <RouterLink
          v-for="tile in mod.tiles"
          :key="tile.key"
          :to="moduleTileLink(tile)"
          class="group relative isolate flex aspect-square flex-col justify-between overflow-hidden rounded-[22px] p-4 transition duration-300 hover:-translate-y-1 sm:rounded-[28px] sm:p-6"
          :class="tile.tone === 'dark' ? 'bg-[#0b0b0c] text-white' : 'bg-white text-[color:var(--xh-ink)]'"
        >
          <span class="pointer-events-none absolute -right-10 -top-10 -z-10 h-40 w-40 rounded-full opacity-70 blur-3xl transition-opacity duration-300 group-hover:opacity-100" :style="{ background: tile.glow }"></span>
          <span
            class="grid h-11 place-items-center rounded-[24%] sm:h-14"
            :class="[tile.wide ? 'w-20 sm:w-24' : 'w-11 sm:w-14', tile.tone === 'dark' ? 'border border-white/15 bg-black' : 'bg-[color:var(--xh-page)]']"
          >
            <img :src="iconSrc(tile.icon)" alt="" class="object-contain" :class="[tile.wide ? 'h-[46%] w-[80%]' : 'h-[58%] w-[58%]', tile.tone === 'dark' ? 'invert' : '']" />
          </span>
          <div>
            <div class="text-[30px] font-black leading-none tracking-[-0.02em] sm:text-[44px]">{{ tileLabel(tile) }}</div>
            <div class="mt-2 flex items-center justify-between gap-2 text-[12px] sm:text-[14px]" :class="tile.tone === 'dark' ? 'text-white/65' : 'text-[color:var(--xh-muted)]'">
              <span class="truncate">{{ tile.subKey ? t(tile.subKey) : '' }}</span>
              <span class="grid h-7 w-7 flex-none place-items-center rounded-full transition group-hover:translate-x-0.5 sm:h-8 sm:w-8" :class="tile.tone === 'dark' ? 'bg-white text-black' : 'bg-[color:var(--xh-ink)] text-white'">
                <ArrowRight class="h-3.5 w-3.5" aria-hidden="true" />
              </span>
            </div>
          </div>
        </RouterLink>
      </div>

      <!-- 横向长条 -->
      <div v-else class="flex flex-col gap-3 sm:gap-4">
        <RouterLink
          v-for="tile in mod.tiles"
          :key="tile.key"
          :to="moduleTileLink(tile)"
          class="group relative isolate flex min-h-[112px] items-center gap-4 overflow-hidden rounded-[22px] bg-white p-4 transition duration-300 hover:-translate-y-0.5 hover:shadow-[0_12px_40px_rgba(0,0,0,0.08)] sm:min-h-[136px] sm:gap-7 sm:rounded-[28px] sm:px-8"
        >
          <span class="pointer-events-none absolute -left-16 top-1/2 -z-10 h-48 w-72 -translate-y-1/2 rounded-full opacity-60 blur-3xl" :style="{ background: tile.glow }"></span>
          <span class="grid h-16 flex-none place-items-center rounded-[22%] bg-[color:var(--xh-page)] shadow-[inset_0_2px_0_#fff,0_10px_24px_rgba(0,0,0,0.08)] sm:h-20" :class="tile.wide ? 'w-28 sm:w-36' : 'w-16 sm:w-20'">
            <img :src="iconSrc(tile.icon)" alt="" class="object-contain" :class="tile.wide ? 'h-[46%] w-[78%]' : 'h-[58%] w-[58%]'" />
          </span>
          <div class="min-w-0 flex-1">
            <div class="text-[19px] font-black leading-snug sm:text-[26px]">{{ tileLabel(tile) }}</div>
            <div v-if="tile.subKey" class="mt-1 text-[13px] text-[color:var(--xh-muted)] sm:text-[15px]">{{ t(tile.subKey) }}</div>
          </div>
          <span class="inline-flex flex-none items-center gap-1.5 rounded-full bg-[color:var(--xh-ink)] px-4 py-2 text-[13px] font-medium text-white transition group-hover:gap-2.5 max-sm:hidden sm:text-[14px]">
            {{ t('xinhuo.modules.buy') }}
            <ArrowRight class="h-3.5 w-3.5" aria-hidden="true" />
          </span>
          <ArrowRight class="h-5 w-5 flex-none text-[color:var(--xh-muted)] sm:hidden" aria-hidden="true" />
        </RouterLink>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import { ArrowRight } from 'lucide-vue-next'
import { HERO_ICON_BASE } from '../heroTiles'
import { homeModules, moduleTileLink, type ModuleTile } from '../homeModules'

const { t } = useI18n()

const iconSrc = (file: string) => HERO_ICON_BASE + file
const tileLabel = (tile: ModuleTile) => tile.label ?? (tile.labelKey ? t(tile.labelKey) : '')
</script>
