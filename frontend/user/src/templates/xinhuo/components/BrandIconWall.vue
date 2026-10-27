<template>
  <!-- 业务图标墙：砖块轻微浮动，首屏用 hero 尺寸，专区首屏用 zone 尺寸 -->
  <ul class="bw" :class="[`bw--${size}`]" :style="{ '--cols': String(columns), '--cols-m': String(mobileColumns) }">
    <li
      v-for="(icon, i) in icons"
      :key="icon.key"
      class="bw-tile"
      :class="[icon.dark ? 'bw-tile--dark' : 'bw-tile--light', { 'bw-tile--wide': icon.wide }]"
      :style="{ animationDelay: `${(i % 4) * -1.5}s` }"
      :title="icon.label"
    >
      <img :src="icon.src" :alt="icon.label" loading="lazy" />
    </li>
  </ul>
</template>

<script setup lang="ts">
import type { BrandIcon } from '../brandIcons'

withDefaults(defineProps<{ icons: BrandIcon[]; size?: 'hero' | 'zone'; columns?: number; mobileColumns?: number }>(), {
  size: 'hero',
  columns: 6,
  mobileColumns: 4,
})
</script>

<style scoped>
.bw {
  list-style: none;
  margin: 0 auto;
  padding: 0;
  display: grid;
  grid-template-columns: repeat(var(--cols), minmax(0, 1fr));
  gap: 22px;
  width: 100%;
}
.bw--hero { max-width: 760px; }
.bw--zone { max-width: 460px; gap: 16px; }
.bw-tile {
  aspect-ratio: 1;
  border-radius: 24%;
  display: grid;
  place-items: center;
  animation: bw-bob 6s ease-in-out infinite;
  transition: transform 0.25s ease;
}
.bw-tile:hover { transform: translateY(-4px) scale(1.04); }
.bw-tile img {
  width: 56%;
  height: 56%;
  object-fit: contain;
  pointer-events: none;
  user-select: none;
}
.bw-tile--wide img { width: 76%; }
.bw-tile--light {
  background: linear-gradient(180deg, #ffffff, #e9ebf1);
  box-shadow:
    0 18px 34px rgba(0, 0, 0, 0.45),
    inset 0 2px 0 #fff,
    inset 0 -5px 12px rgba(0, 0, 0, 0.1);
}
.bw-tile--dark {
  background: linear-gradient(180deg, #2a2a2e, #0b0b0c);
  border: 1.5px solid #3a3a40;
  box-shadow:
    0 20px 40px rgba(0, 0, 0, 0.6),
    inset 0 2px 0 rgba(255, 255, 255, 0.18);
}
.bw-tile--dark img { filter: invert(1); }
.bw--zone .bw-tile--light { box-shadow: 0 14px 28px rgba(0, 0, 0, 0.28), inset 0 2px 0 #fff, inset 0 -4px 10px rgba(0, 0, 0, 0.08); }
@keyframes bw-bob {
  0%, 100% { translate: 0 0; }
  50% { translate: 0 -6px; }
}
@media (max-width: 640px) {
  .bw { grid-template-columns: repeat(var(--cols-m), minmax(0, 1fr)); gap: 14px; }
  .bw--hero { max-width: 340px; }
  .bw--zone { max-width: 320px; gap: 12px; }
}
@media (prefers-reduced-motion: reduce) {
  .bw-tile { animation: none; }
}
</style>
