<template>
  <!-- 业务图标墙：错落排列（后排小、带模糊；前排大、带倒影），复用首屏图标墙样式 xh-wall / xh-tile -->
  <div class="xh-wall" :class="{ 'xh-wall--zone': variant === 'zone' }">
    <div
      v-for="tile in styled"
      :key="tile.key"
      role="img"
      :aria-label="tile.label"
      class="xh-tile"
      :class="[`xh-tile--${tile.tone}`, `xh-tile--${tile.row}`, { 'xh-tile--wide': tile.wide }]"
      :style="tile.style"
    >
      <img :src="tile.src" alt="" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { BrandTile } from '../brandIcons'

const props = withDefaults(defineProps<{ tiles: BrandTile[]; variant?: 'hero' | 'zone' }>(), { variant: 'hero' })

const styled = computed(() =>
  props.tiles.map((tile) => ({
    ...tile,
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
  })),
)
</script>
